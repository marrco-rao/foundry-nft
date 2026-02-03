// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable/contracts/proxy/utils/UUPSUpgradeable.sol";
import {Initializable} from "@openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {IERC2981} from "@openzeppelin/contracts/interfaces/IERC2981.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "../interfaces/IPaymentToken.sol";
import "../interfaces/IWETH.sol";

contract NFTMarketplaceUUPS is Initializable, UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuard, IPaymentToken {
    using SafeERC20 for IERC20;
    // 这里可以添加NFT市场的功能，例如列出NFT、购买NFT等
    // NFT挂单结构体
    struct Listing {
        address seller; // 卖家地址
        address nftContract;      // NFT合约地址
        uint256 tokenId; // NFT的Token ID   
        uint256 price;  // 售价（wei）
        PaymentMethod paymentMethod; // 支付方式
        bool isActive;  // 挂单是否激活 
    }

    // 拍卖单结构体
    struct Auction {
        address seller;      // 卖家地址
        address nftContract;      // NFT合约地址
        uint256 tokenId;     // NFT的Token ID
        uint256 startingBid; // 起始竞拍价（wei）
        uint256 highestBid;  // 最高竞拍价（wei）
        address highestBidder; // 最高竞拍者地址
        uint256 endTime;     // 拍卖结束时间（时间戳）
        PaymentMethod paymentMethod; // 支付方式
        bool isActive;       // 拍卖是否激活
    }

    // 映射：Token ID => 挂单信息
    mapping(uint256 => Listing) public listings;
    uint256 public listingCount; // 挂单计数器

    // 拍卖映射
    mapping(uint256 => Auction) public auctions;
    uint256 public auctionCounter;
    
    /// @notice 待退款映射（用于拍卖），记录每个NFT的待退回出价金额
    /// @dev 第一层mapping的key为拍卖ID，第二层mapping的key为出价者地址，value为待退回的金额
    /// @dev 当出价被更高的出价超越时，原出价金额会存储在此映射中，等待出价者主动提取
    mapping(uint256 => mapping(address => uint256)) public pendingReturns;
    
    // 平台手续费（基点，10000 = 100%）
    uint256 public platformFee = 250; // 2.5%
    
    // 手续费接收地址
    address public feeRecipient;

    // 是否启用版税
    bool public royaltyEnabled;

    // 支付代币地址
    address public wethAddress;

    // Chainlink 价格预言机
    AggregatorV3Interface public ethUsdPriceFeed;

    // 支持的支付方式映射
    mapping(PaymentMethod => bool) public supportedPaymentMethods;
    

    // 事件：NFT上架
    event NFTListed(address indexed seller, uint256 indexed tokenId, uint256 price, PaymentMethod paymentMethod);
    // 事件：NFT下架
    event NFTDelisted(address indexed seller, uint256 indexed tokenId);
    // 事件：NFT售出
    event NFTPurchased(address indexed buyer, uint256 indexed tokenId, uint256 price, PaymentMethod paymentMethod);
    // 事件：NFT价格更新
    event NFTListingPriceUpdated(address indexed seller, uint256 indexed tokenId, uint256 newPrice);
    // 事件：价格查询（记录 USD 价格）
    event PriceInUSD(uint256 indexed listingId, uint256 priceInWei, uint256 priceInUSD);

    // 事件：NFT拍卖上架
    event NFTAuctionListed(address indexed seller, uint256 indexed tokenId, uint256 startingBid, uint256 endTime, PaymentMethod paymentMethod);
    // 事件：NFT拍卖下架
    event NFTAuctionDelisted(address indexed seller, uint256 indexed tokenId);
    // 事件：竞拍出价事件
    event NFTAuctionBidPlaced(address indexed bidder, uint256 indexed tokenId, uint256 bidAmount);
    // 事件：当NFT拍卖结束时触发
    event NFTAuctionEnded(address indexed winner, uint256 indexed tokenId, uint256 finalBidAmount); 
    // 事件：当NFT拍卖被取消时触发
    event NFTAuctionCancelled(address indexed seller, uint256 indexed tokenId);

    /**
     * @notice 初始化函数
     * @dev 初始化合约所有者和手续费信息 
     */
    function initialize(
        uint256 _feeBps,
        address _feeRecipient,
        address _wethAddress,
        address _ethUsdPriceFeed
    ) public initializer {
        __Ownable_init(msg.sender);

        platformFee = _feeBps;
        feeRecipient = _feeRecipient;
        royaltyEnabled = true;
        wethAddress = _wethAddress;
        ethUsdPriceFeed = AggregatorV3Interface(_ethUsdPriceFeed);
        
        // 默认支持 ETH 和 WETH
        supportedPaymentMethods[PaymentMethod.ETH] = true;
        supportedPaymentMethods[PaymentMethod.WETH] = true;
    }

    /**
    * @notice nft挂单上架销售
    * @dev 创建新的挂单
    * @param nftContract NFT合约地址
    * @param tokenId NFT的Token ID
    * @param price 售价（wei）
    * @param paymentMethod 支付方式（ETH 或 WETH）
     */
    function listNFT(
        address nftContract,
        uint256 tokenId,
        uint256 price,
        PaymentMethod paymentMethod
    ) external returns (uint256) {
        // 挂单逻辑
        require(price > 0, "Price must be greater than zero");
        require(nftContract != address(0),"Invalid NFT contract address");
        require(isPaymentMethodSupported(paymentMethod), "Payment method not supported");

        IERC721 nft = IERC721(nftContract);
        require(nft.ownerOf(tokenId) == msg.sender, "Not the owner of the NFT");
        require(nft.getApproved(tokenId) == address(this) || nft.isApprovedForAll(msg.sender, address(this)), "Marketplace not approved");

        listingCount++ ;
        listings[listingCount] = Listing({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            price: price,
            paymentMethod: paymentMethod,
            isActive: true
        });
        emit NFTListed(msg.sender, tokenId, price, paymentMethod);
        return listingCount;
    }

    /** 
    * @notice 下架NFT挂单
    * @dev 取消已有的挂单
    * @param listingId 挂单ID    
     */
    function delistNFT(uint256 listingId) external {
        Listing storage listing = listings[listingId];
        require(listing.isActive, "Listing is not active");
        require(listing.seller == msg.sender, "Not the seller");

        listing.isActive = false;
        emit NFTDelisted(msg.sender, listing.tokenId);
    }

    /** 
    * @notice 修改NFT挂单价格
     */
    function updateListingPrice(uint256 listingId, uint256 newPrice) external {
        require(newPrice > 0, "Price must be greater than zero");

        Listing storage listing = listings[listingId];
        require(listing.isActive, "Listing is not active");
        require(listing.seller == msg.sender, "Not the seller");

        listing.price = newPrice;
        emit NFTListingPriceUpdated(msg.sender, listing.tokenId, newPrice);
    }

    /** 
    * @notice 购买已挂单的NFT
    * @dev 处理NFT购买逻辑，包括支付和转移NFT，支持 ETH 和 WETH
    * @param nftContract NFT合约地址
    * @param listingId 挂单ID
     */
    function purchaseNFT(
        address nftContract,
        uint256 listingId
    ) external payable nonReentrant{
        Listing storage listing = listings[listingId];
        require(listing.isActive, "Listing is not active");
        require(msg.sender != listing.seller, "Cannot buy your own NFT");
  
        IERC721 nft = IERC721(nftContract);
        // 确保卖家仍然拥有NFT
        require(nft.ownerOf(listing.tokenId) == listing.seller, "Seller no longer owns the NFT");

        // 根据支付方式验证支付
        if (listing.paymentMethod == PaymentMethod.ETH) {
            require(msg.value >= listing.price, "Insufficient ETH payment");
        } else if (listing.paymentMethod == PaymentMethod.WETH) {
            require(msg.value == 0, "Should not send ETH for WETH payment");
            require(IERC20(wethAddress).balanceOf(msg.sender) >= listing.price, "Insufficient WETH balance");
            require(IERC20(wethAddress).allowance(msg.sender, address(this)) >= listing.price, "Insufficient WETH allowance");
        }

        // 标记挂单为不活跃：先更新状态（CEI原则）
        listing.isActive = false;   

        // 计算平台手续费
        uint256 feeAmount = (listing.price * platformFee) / 10000;
        
        // 获取版税信息
        (address royaltyReceiver, uint256 royaltyAmount) = _getRoyaltyInfo(
            listing.nftContract,
            listing.tokenId,
            listing.price
        );

        // 计算卖家应得金额
        uint256 sellerAmount = listing.price - feeAmount - royaltyAmount;   

        // 转移NFT给买家
        nft.safeTransferFrom(listing.seller, msg.sender, listing.tokenId);

        // 资金分配：版税 -> 平台手续费 -> 卖家收益
        if (listing.paymentMethod == PaymentMethod.ETH) {
            // ETH 支付
            if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
                (bool success, ) = royaltyReceiver.call{value: royaltyAmount}("");
                require(success, "Royalty transfer failed");
            }
            (bool feeSuccess, ) = feeRecipient.call{value: feeAmount}("");
            require(feeSuccess, "Fee transfer failed");

            (bool sellerSuccess, ) = listing.seller.call{value: sellerAmount}("");
            require(sellerSuccess, "Seller transfer failed");

            // 退还多余的 ETH
            uint256 refund = msg.value - listing.price;
            if (refund > 0) {
                (bool refundSuccess, ) = msg.sender.call{value: refund}("");
                require(refundSuccess, "Refund failed");
            }
        } else if (listing.paymentMethod == PaymentMethod.WETH) {
            // WETH 支付
            IERC20 weth = IERC20(wethAddress);
            if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
                weth.safeTransferFrom(msg.sender, royaltyReceiver, royaltyAmount);
            }
            weth.safeTransferFrom(msg.sender, feeRecipient, feeAmount);
            weth.safeTransferFrom(msg.sender, listing.seller, sellerAmount);
        }

        emit NFTPurchased(msg.sender, listing.tokenId, listing.price, listing.paymentMethod);
    }

    /**
    * @notice nft拍卖上架销售
    * @dev 创建新的拍卖
    * @param nftContract NFT合约地址
    * @param tokenId NFT的Token ID
    * @param startingBid 起始竞拍价（wei）
    * @param durationHours 拍卖时长（小时）
    * @param paymentMethod 支付方式（ETH 或 WETH）
    */ 
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startingBid,
        uint256 durationHours,
        PaymentMethod paymentMethod
    ) external returns (uint256) {
        // 拍卖逻辑
        require(startingBid > 0, "Starting bid must be greater than zero");
        require(durationHours >= 1, "Duration must be at least 1 hour");
        require(nftContract != address(0),"Invalid NFT contract address");
        require(isPaymentMethodSupported(paymentMethod), "Payment method not supported");

        IERC721 nft = IERC721(nftContract);
        // 验证所有权
        require(nft.ownerOf(tokenId) == msg.sender, "Not the owner of the NFT");
        // 验证授权
        require(nft.getApproved(tokenId) == address(this) || nft.isApprovedForAll(msg.sender, address(this)), "Marketplace not approved");

        auctionCounter++ ;
        auctions[auctionCounter] = Auction({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            startingBid: startingBid,
            highestBid: 0,
            highestBidder: address(0),
            endTime: block.timestamp + (durationHours * 1 hours),
            paymentMethod: paymentMethod,
            isActive: true
        });
        emit NFTAuctionListed(msg.sender, tokenId, startingBid, auctions[auctionCounter].endTime, paymentMethod);
        return auctionCounter;
    }  

    /**
     * @dev 出价
     * @param auctionId 拍卖ID
     * @notice 支持 ETH 和 WETH 出价，出价必须高于当前最高出价的5%
     */
    function placeBid(uint256 auctionId) external payable {
        Auction storage auction = auctions[auctionId];
        require(auction.isActive, "Auction is not active");
        require(block.timestamp < auction.endTime, "Auction has ended");
        require(msg.sender != auction.seller, "Seller cannot bid");

        // 计算最低出价
        uint256 minBid;       
        if (auction.highestBid == 0) {
            minBid = auction.startingBid;
        } else {
            minBid = auction.highestBid + (auction.highestBid / 20); // 最低出价需高于当前最高出价的5%
        }

        uint256 bidAmount;
        if (auction.paymentMethod == PaymentMethod.ETH) {
            require(msg.value >= minBid, "ETH bid too low");
            bidAmount = msg.value;
        } else if (auction.paymentMethod == PaymentMethod.WETH) {
            require(msg.value == 0, "Should not send ETH for WETH auction");
            require(IERC20(wethAddress).balanceOf(msg.sender) >= minBid, "Insufficient WETH balance");
            require(IERC20(wethAddress).allowance(msg.sender, address(this)) >= minBid, "Insufficient WETH allowance");
            bidAmount = minBid;
            
            // 转移 WETH 到合约
            IERC20(wethAddress).safeTransferFrom(msg.sender, address(this), bidAmount);
        }

        // 如果有之前的最高出价者，退还资金
        if (auction.highestBidder != address(0)) {
            if (auction.paymentMethod == PaymentMethod.ETH) {
                pendingReturns[auctionId][auction.highestBidder] += auction.highestBid;
            } else if (auction.paymentMethod == PaymentMethod.WETH) {
                // 直接退还 WETH 给之前的最高出价者
                IERC20(wethAddress).safeTransfer(auction.highestBidder, auction.highestBid);
            }
        }

        // 更新最高出价和最高出价者
        auction.highestBid = bidAmount;
        auction.highestBidder = msg.sender;

        emit NFTAuctionBidPlaced(msg.sender, auction.tokenId, bidAmount);
    }

    // 提取出价退款
    function withdrawBidRefund(uint256 auctionId) external {
        uint256 amount = pendingReturns[auctionId][msg.sender];
        require(amount > 0, "No funds to withdraw");

        pendingReturns[auctionId][msg.sender] = 0;
        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "Withdrawal failed");
    }

    /** 
    * @notice 结束拍卖 
    * @dev 只有拍卖卖家或合约所有者可以结束拍卖，支持 ETH 和 WETH 支付
    * @param auctionId 拍卖ID
    * @param nftContract NFT合约地址
    */
    function endAuction(uint256 auctionId, address nftContract) external nonReentrant {
        Auction storage auction = auctions[auctionId];
        require(auction.isActive, "Auction is not active");
        require(msg.sender == auction.seller || msg.sender == owner(), "Only seller or owner can end auction");
        require(block.timestamp >= auction.endTime, "Auction has not ended yet");

        auction.isActive = false;

        if (auction.highestBidder != address(0)) {
            // 有人出价，转移NFT给最高竞拍者
            IERC721 nft = IERC721(nftContract);
            nft.safeTransferFrom(auction.seller, auction.highestBidder, auction.tokenId);

            // 计算平台手续费
            uint256 feeAmount = (auction.highestBid * platformFee) / 10000;

            // 获取版税信息
            (address royaltyReceiver, uint256 royaltyAmount) = _getRoyaltyInfo(
                nftContract,
                auction.tokenId,
                auction.highestBid
            );

            // 计算卖家应得金额
            uint256 sellerAmount = auction.highestBid - feeAmount - royaltyAmount;

            // 资金分配
            if (auction.paymentMethod == PaymentMethod.ETH) {
                // ETH 支付
                if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
                    (bool success, ) = royaltyReceiver.call{value: royaltyAmount}("");
                    require(success, "Royalty transfer failed");
                }
                (bool feeSuccess, ) = feeRecipient.call{value: feeAmount}("");
                require(feeSuccess, "Fee transfer failed");
                (bool sellerSuccess, ) = auction.seller.call{value: sellerAmount}("");
                require(sellerSuccess, "Seller transfer failed");
            } else if (auction.paymentMethod == PaymentMethod.WETH) {
                // WETH 支付
                IERC20 weth = IERC20(wethAddress);
                if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
                    weth.safeTransfer(royaltyReceiver, royaltyAmount);
                }
                weth.safeTransfer(feeRecipient, feeAmount);
                weth.safeTransfer(auction.seller, sellerAmount);
            }

            emit NFTAuctionEnded(auction.highestBidder, auction.tokenId, auction.highestBid);
        } else {
            // 如果没有竞拍，拍卖被取消
            emit NFTAuctionCancelled(auction.seller, auction.tokenId);
        }
    }

    /**
    * @notice 查询挂单列表，在售的NFT（简单查询，不考虑分页）
    * @return listingIds 挂单ID数组 
     */
    function getActiveListings() external view returns (uint256[] memory listingIds) {
        uint256 activeCount = 0;
        for (uint256 i = 1; i <= listingCount; i++) {
            if (listings[i].isActive) {
                activeCount++;
            }
        }

        listingIds = new uint256[](activeCount);
        uint256 index = 0;
        for (uint256 i = 1; i <= listingCount; i++) {
            if (listings[i].isActive) {
                listingIds[index] = i;
                index++;
            }
        }
    }

    /**
    * @notice 获取挂单信息
    * @param listingId 挂单ID
    * @return seller 卖家地址
    * @return nftContract NFT合约地址
    * @return tokenId NFT的Token ID
    * @return price 售价（wei）
    * @return paymentMethod 支付方式
    * @return isActive 挂单是否激活
     */
    function getListing(uint256 listingId) external view returns (
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 price,
        PaymentMethod paymentMethod,
        bool isActive
    ) {
        Listing storage listing = listings[listingId];
        return (
            listing.seller, 
            listing.nftContract,
            listing.tokenId,
            listing.price,
            listing.paymentMethod,
            listing.isActive
        );
    }

    /**
    * @notice 获取拍卖信息
    * @param auctionId 拍卖ID   
    * @return seller 卖家地址
    * @return nftContract NFT合约地址
    * @return tokenId NFT的Token ID
    * @return startingBid 起始竞拍价（wei） 
    * @return highestBid 最高竞拍价（wei）
    * @return highestBidder 最高竞拍者地址
    * @return endTime 拍卖结束时间（时间戳）
    * @return paymentMethod 支付方式
    * @return isActive 拍卖是否激活
     */
    function getAuction(uint256 auctionId) external view returns (
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 startingBid,
        uint256 highestBid,
        address highestBidder,
        uint256 endTime,
        PaymentMethod paymentMethod,
        bool isActive
    ) {
        Auction storage auction = auctions[auctionId];
        return (    
            auction.seller,
            auction.nftContract,
            auction.tokenId,
            auction.startingBid,
            auction.highestBid,
            auction.highestBidder,
            auction.endTime,
            auction.paymentMethod,
            auction.isActive
        );
    }   

    /**
     * @dev 获取版税信息
     * @param nftContract NFT合约地址
     * @param tokenId Token ID
     * @param salePrice 售价
     * @return receiver 版税接收地址
     * @return royaltyAmount 版税金额
     * @notice 内部函数，检查NFT合约是否支持ERC2981标准
     */
    function _getRoyaltyInfo( 
        address nftContract,
        uint256 tokenId,
        uint256 salePrice
    ) internal view returns (address receiver, uint256 royaltyAmount) {
        // 检查是否启用版税
        if (!royaltyEnabled) {
            return (address(0) , 0);
        }
        // 检查NFT合约是否支持ERC2981
        if (IERC165(nftContract).supportsInterface(type(IERC2981).interfaceId)) {
            (receiver, royaltyAmount) = IERC2981(nftContract).royaltyInfo(tokenId, salePrice);
            
        } else {
            receiver = address(0);
            royaltyAmount = 0;
        }
    }

    /**
    * @notice 启用或禁用版税
    * @dev 只有合约所有者可以调用
    * @param _enabled 是否启用版税  
     */
    function setRoyaltyEnabled(bool _enabled) external onlyOwner {
        royaltyEnabled = _enabled;
    }

    /** 
    * @notice 升级授权函数
    * @dev 只有合约所有者可以升级合约实现  
    */ 
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @notice 设置平台手续费
     * @dev 只有手续费接收地址可以调用
     * @param _feeBps 新的手续费，基点表示，最大1000（10%）
     */
    function setPlatformFee(uint256 _feeBps) external onlyOwner {

        require(_feeBps <= 1000, "Fee too high"); // 最大10%
        platformFee = _feeBps;  
    }

    /**
     * @notice 设置手续费接收地址
     * @dev 只有合约所有者可以调用
     * @param _feeRecipient 新的手续费接收地址
     */
    function setFeeRecipient(address _feeRecipient) external onlyOwner {
        require(_feeRecipient != address(0),"Invalid address");
        feeRecipient = _feeRecipient;
    }

    /**
     * @notice 获取 ETH/USD 价格
     * @dev 从 Chainlink 价格预言机获取最新价格
     * @return price ETH/USD 价格（8位小数）
     */
    function getETHPrice() public view returns (uint256 price) {
        (, int256 answer, , , ) = ethUsdPriceFeed.latestRoundData();
        require(answer > 0, "Invalid price feed");
        return uint256(answer);
    }

    /**
     * @notice 获取挂单的 USD 价格
     * @param listingId 挂单ID
     * @return priceInUSD 价格（以美分计，2位小数）
     */
    function getListingPriceInUSD(uint256 listingId) external view returns (uint256 priceInUSD) {
        Listing storage listing = listings[listingId];
        require(listing.isActive, "Listing is not active");
        
        uint256 ethPrice = getETHPrice(); // 8 decimals
        // listing.price 是 wei (18 decimals)
        // 转换为 USD cents (2 decimals): (price * ethPrice) / 10^(18 + 8 - 2)
        priceInUSD = (listing.price * ethPrice) / 1e24;
        
        return priceInUSD;
    }

    /**
     * @notice 获取拍卖的 USD 价格
     * @param auctionId 拍卖ID
     * @return startingBidUSD 起始价格（美分）
     * @return highestBidUSD 当前最高出价（美分）
     */
    function getAuctionPriceInUSD(uint256 auctionId) external view returns (
        uint256 startingBidUSD,
        uint256 highestBidUSD
    ) {
        Auction storage auction = auctions[auctionId];
        require(auction.isActive, "Auction is not active");
        
        uint256 ethPrice = getETHPrice();
        startingBidUSD = (auction.startingBid * ethPrice) / 1e24;
        highestBidUSD = auction.highestBid > 0 ? (auction.highestBid * ethPrice) / 1e24 : 0;
        
        return (startingBidUSD, highestBidUSD);
    }

    // ========== IPaymentToken 接口实现 ==========

    /**
     * @notice 获取支付方式对应的代币地址
     * @param method 支付方式
     * @return 代币地址，ETH 返回 address(0)
     */
    function getPaymentTokenAddress(PaymentMethod method) external view override returns (address) {
        if (method == PaymentMethod.ETH) {
            return address(0);
        } else if (method == PaymentMethod.WETH) {
            return wethAddress;
        }
        return address(0);
    }

    /**
     * @notice 检查支付方式是否支持
     * @param method 支付方式
     * @return 是否支持
     */
    function isPaymentMethodSupported(PaymentMethod method) public view override returns (bool) {
        return supportedPaymentMethods[method];
    }

    /**
     * @notice 设置支付方式支持状态（仅所有者）
     * @param method 支付方式
     * @param supported 是否支持
     */
    function setPaymentMethodSupported(PaymentMethod method, bool supported) external onlyOwner {
        supportedPaymentMethods[method] = supported;
    }

    /**
     * @notice 更新 WETH 地址（仅所有者）
     * @param _wethAddress 新的 WETH 地址
     */
    function setWETHAddress(address _wethAddress) external onlyOwner {
        require(_wethAddress != address(0), "Invalid WETH address");
        wethAddress = _wethAddress;
    }

    /**
     * @notice 更新价格预言机地址（仅所有者）
     * @param _priceFeed 新的价格预言机地址
     */
    function setPriceFeed(address _priceFeed) external onlyOwner {
        require(_priceFeed != address(0), "Invalid price feed address");
        ethUsdPriceFeed = AggregatorV3Interface(_priceFeed);
    }
}