// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable/contracts/proxy/utils/UUPSUpgradeable.sol";
import {Initializable} from "@openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {IERC2981} from "@openzeppelin/contracts/interfaces/IERC2981.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract NFTMarketplaceUUPS is Initializable, UUPSUpgradeable, OwnableUpgradeable, ReentrancyGuard {
    // 这里可以添加NFT市场的功能，例如列出NFT、购买NFT等
    // NFT挂单结构体
    struct Listing {
        address seller; // 卖家地址
        address nftContract;      // NFT合约地址
        uint256 tokenId; // NFT的Token ID   
        uint256 price;  // 售价（wei）
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
    

    // 事件：NFT上架
    event NFTListed(address indexed seller, uint256 indexed tokenId, uint256 price);
    // 事件：NFT下架
    event NFTDelisted(address indexed seller, uint256 indexed tokenId);
    // 事件：NFT售出
    event NFTPurchased(address indexed buyer, uint256 indexed tokenId, uint256 price);
    // 事件：NFT价格更新
    event NFTListingPriceUpdated(address indexed seller, uint256 indexed tokenId, uint256 newPrice);

    // 事件：NFT拍卖上架
    event NFTAuctionListed(address indexed seller, uint256 indexed tokenId, uint256 startingBid, uint256 endTime);
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
        address _feeRecipient
    ) public initializer {
        __Ownable_init(msg.sender);

        platformFee = _feeBps;
        feeRecipient = _feeRecipient;
        royaltyEnabled = true;
    }

    /**
    * @notice nft挂单上架销售
    * @dev 创建新的挂单
    * @param nftContract NFT合约地址
    * @param tokenId NFT的Token ID
    * @param price 售价（wei）
     */
    function listNFT(
        address nftContract,
        uint256 tokenId,
        uint256 price
    ) external returns (uint256) {
        // 挂单逻辑
        require(price > 0, "Price must be greater than zero");
        require(nftContract != address(0),"Invalid NFT contract address");

        IERC721 nft = IERC721(nftContract);
        require(nft.ownerOf(tokenId) == msg.sender, "Not the owner of the NFT");
        require(nft.getApproved(tokenId) == address(this) || nft.isApprovedForAll(msg.sender, address(this)), "Marketplace not approved");

        listingCount++ ;
        listings[listingCount] = Listing({
            seller: msg.sender,
            nftContract: nftContract,
            tokenId: tokenId,
            price: price,
            isActive: true
        });
        emit NFTListed(msg.sender, tokenId, price);
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
    * @dev 处理NFT购买逻辑，包括支付和转移NFT
    * @param nftContract NFT合约地址
    * @param listingId 挂单ID
     */
    function purchaseNFT(
        address nftContract,
        uint256 listingId
    ) external payable nonReentrant{
        Listing storage listing = listings[listingId];
        require(listing.isActive, "Listing is not active");
        require(msg.value >= listing.price, "Insufficient payment");  
        require(msg.sender != listing.seller, "Cannot buy your own NFT");
  
        IERC721 nft = IERC721(nftContract);
        // 确保卖家仍然拥有NFT
        require(nft.ownerOf(listing.tokenId) == listing.seller, "Seller no longer owns the NFT");

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

        // 如果有版税，转移版税给版税接收者（这里假设版税接收者是合约所有者）
        if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
            (bool success, ) = royaltyReceiver.call{value: royaltyAmount}("");
            require(success, "Royalty transfer failed");
        }
        // 转移平台手续费给手续费接收地址
        (bool feeSuccess, ) = feeRecipient.call{value: feeAmount}("");
        require(feeSuccess, "Fee transfer failed");

        // 转移以太币给卖家
        (bool sellerSuccess, ) = listing.seller.call{value: sellerAmount}("");
        require(sellerSuccess, "Seller transfer failed");

        // 计算剩余的以太币退款给买家（如果有多付）
        uint256 refund = msg.value - listing.price;
        if (refund > 0) {
            (bool refundSuccess, ) = msg.sender.call{value: refund}("");
            require(refundSuccess, "Refund failed");
        }   

        emit NFTPurchased(msg.sender, listing.tokenId, listing.price);
    }

    /**
    * @notice nft拍卖上架销售
    * @dev 创建新的拍卖
    * @param nftContract NFT合约地址
    * @param tokenId NFT的Token ID
    * @param startingBid 起始竞拍价（wei）
    * @param durationHours 拍卖时长（小时）
    */ 
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startingBid,
        uint256 durationHours
    ) external returns (uint256) {
        // 拍卖逻辑
        require(startingBid > 0, "Starting bid must be greater than zero");
        require(durationHours >= 1, "Duration must be at least 1 hour");
        require(nftContract != address(0),"Invalid NFT contract address");

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
            isActive: true
        });
        emit NFTAuctionListed(msg.sender, tokenId, startingBid,auctions[auctionCounter].endTime);
        return auctionCounter;
    }  

    /**
     * @dev 出价
     * @param auctionId 拍卖ID
     * @notice 需要支付足够的ETH，出价必须高于当前最高出价的5%
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
        require(msg.value >= minBid, "Bid too low");

        // 如果有之前的最高出价者，记录待退款金额
        if (auction.highestBidder != address(0)) {
            pendingReturns[auctionId][auction.highestBidder] += auction.highestBid;
        }

        // 更新最高出价和最高出价者
        auction.highestBid = msg.value;
        auction.highestBidder = msg.sender;

        emit NFTAuctionBidPlaced(msg.sender, auction.tokenId, msg.value);

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
    * @dev 只有拍卖卖家或合约所有者可以结束拍卖
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
            // 转移版税
            if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
                (bool success, ) = royaltyReceiver.call{value: royaltyAmount}("");
                require(success, "Royalty transfer failed");
            }
            // 转移平台手续费
            (bool feeSuccess, ) = feeRecipient.call{value: feeAmount}("");
            require(feeSuccess, "Fee transfer failed");
            // 转移卖家收益
            (bool sellerSuccess, ) = auction.seller.call{value: sellerAmount}("");
            require(sellerSuccess, "Seller transfer failed");

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
    * @return isActive 挂单是否激活
     */
    function getListing(uint256 listingId) external view returns (
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 price,
        bool isActive
    ) {
        Listing storage listing = listings[listingId];
        return (
            listing.seller, 
            listing.nftContract,
            listing.tokenId,
            listing.price,
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
}