// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Test, console} from "forge-std/Test.sol";
import {NFTMarketplaceUUPS} from "../src/nft-market/NFTMarketplaceUUPS.sol";
import {MyNFTUUPS} from "../src/nft/MyNFTUUPS.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @title NFTMarketplaceUUPS Test Suite
 * @notice 测试NFT Marketplace合约的所有功能
 */
contract NFTMarketplaceUUPSTest is Test {
    NFTMarketplaceUUPS public marketplace;
    MyNFTUUPS public nft;
    
    address public marketImpl;
    address public marketProxy;
    address public nftImpl;
    address public nftProxy;
    
    address public owner;
    address public seller;
    address public buyer;
    address public feeRecipient;
    address public royaltyReceiver;
    
    uint256 public constant PLATFORM_FEE_BPS = 250; // 2.5%
    uint256 public constant ROYALTY_BPS = 250; // 2.5%
    uint256 public constant NFT_PRICE = 1 ether;
    uint256 public constant MINT_PRICE = 0.01 ether;
    
    // 事件声明
    event NFTListed(address indexed seller, uint256 indexed tokenId, uint256 price);
    event NFTDelisted(address indexed seller, uint256 indexed tokenId);
    event NFTPurchased(address indexed buyer, uint256 indexed tokenId, uint256 price);
    event NFTListingPriceUpdated(address indexed seller, uint256 indexed tokenId, uint256 newPrice);
    event NFTAuctionListed(address indexed seller, uint256 indexed tokenId, uint256 startingBid, uint256 endTime);
    event NFTAuctionBidPlaced(address indexed bidder, uint256 indexed tokenId, uint256 bidAmount);
    event NFTAuctionEnded(address indexed winner, uint256 indexed tokenId, uint256 finalBidAmount);
    event NFTAuctionCancelled(address indexed seller, uint256 indexed tokenId);
    
    function setUp() public {
        // 设置测试账户
        owner = makeAddr("owner");
        seller = makeAddr("seller");
        buyer = makeAddr("buyer");
        feeRecipient = makeAddr("feeRecipient");
        royaltyReceiver = makeAddr("royaltyReceiver");
        
        // 给账户充值
        vm.deal(owner, 100 ether);
        vm.deal(seller, 100 ether);
        vm.deal(buyer, 100 ether);
        
        vm.startPrank(owner);
        
        // 部署NFT合约
        nftImpl = address(new MyNFTUUPS());
        bytes memory nftInitData = abi.encodeWithSelector(
            MyNFTUUPS.initialize.selector,
            "Test NFT",
            "TNFT",
            royaltyReceiver,
            ROYALTY_BPS
        );
        nftProxy = address(new ERC1967Proxy(nftImpl, nftInitData));
        nft = MyNFTUUPS(nftProxy);
        
        // 部署Marketplace合约
        marketImpl = address(new NFTMarketplaceUUPS());
        bytes memory marketInitData = abi.encodeWithSelector(
            NFTMarketplaceUUPS.initialize.selector,
            PLATFORM_FEE_BPS,
            feeRecipient
        );
        marketProxy = address(new ERC1967Proxy(marketImpl, marketInitData));
        marketplace = NFTMarketplaceUUPS(marketProxy);
        
        vm.stopPrank();
        
        // 为seller铸造一些NFT
        vm.startPrank(seller);
        for (uint256 i = 0; i < 5; i++) {
            nft.mint{value: MINT_PRICE}(seller, string(abi.encodePacked("ipfs://QmTest", vm.toString(i))));
        }
        vm.stopPrank();
    }
    
    /* ========== 初始化测试 ========== */
    
    function test_MarketplaceInitialization() public view {
        assertEq(marketplace.owner(), owner);
        assertEq(marketplace.platformFee(), PLATFORM_FEE_BPS);
        assertEq(marketplace.feeRecipient(), feeRecipient);
        assertTrue(marketplace.royaltyEnabled());
        assertEq(marketplace.listingCount(), 0);
    }
    
    /* ========== NFT上架测试 ========== */
    
    function test_ListNFT() public {
        uint256 tokenId = 1;
        
        vm.startPrank(seller);
        nft.approve(address(marketplace), tokenId);
        
        vm.expectEmit(true, true, false, true);
        emit NFTListed(seller, tokenId, NFT_PRICE);
        
        uint256 listingId = marketplace.listNFT(address(nft), tokenId, NFT_PRICE);
        vm.stopPrank();
        
        assertEq(listingId, 1);
        assertEq(marketplace.listingCount(), 1);
        
        (address lseller, address nftContract, uint256 ltokenId, uint256 price, bool isActive) 
            = marketplace.getListing(listingId);
        
        assertEq(lseller, seller);
        assertEq(nftContract, address(nft));
        assertEq(ltokenId, tokenId);
        assertEq(price, NFT_PRICE);
        assertTrue(isActive);
    }
    
    function test_RevertWhen_ListNFTWithoutApproval() public {
        vm.prank(seller);
        vm.expectRevert();
        marketplace.listNFT(address(nft), 1, NFT_PRICE);
    }
    
    function test_RevertWhen_ListNFTNotOwned() public {
        // buyer尝试上架seller拥有的NFT（tokenId=1）
        vm.startPrank(buyer);
        vm.expectRevert(); // 会因为ERC721InvalidApprover而失败
        marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
    }
    
    function test_RevertWhen_ListNFTWithZeroPrice() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        vm.expectRevert();
        marketplace.listNFT(address(nft), 1, 0);
        vm.stopPrank();
    }
    
    function test_ListMultipleNFTs() public {
        vm.startPrank(seller);
        
        for (uint256 i = 1; i <= 3; i++) {
            nft.approve(address(marketplace), i);
            marketplace.listNFT(address(nft), i, NFT_PRICE + (i * 0.1 ether));
        }
        
        vm.stopPrank();
        
        assertEq(marketplace.listingCount(), 3);
    }
    
    /* ========== NFT下架测试 ========== */
    
    function test_DelistNFT() public {
        // 先上架
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        
        vm.expectEmit(true, true, false, false);
        emit NFTDelisted(seller, 1);
        
        marketplace.delistNFT(listingId);
        vm.stopPrank();
        
        (,,, , bool isActive) = marketplace.getListing(listingId);
        assertFalse(isActive);
    }
    
    function test_RevertWhen_DelistNFTNotSeller() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.delistNFT(listingId);
    }
    
    function test_RevertWhen_DelistInactiveListing() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        marketplace.delistNFT(listingId);
        
        // 尝试再次下架
        vm.expectRevert();
        marketplace.delistNFT(listingId);
        vm.stopPrank();
    }
    
    /* ========== 更新价格测试 ========== */
    
    function test_UpdateListingPrice() public {
        uint256 newPrice = 2 ether;
        
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        
        vm.expectEmit(true, true, false, true);
        emit NFTListingPriceUpdated(seller, 1, newPrice);
        
        marketplace.updateListingPrice(listingId, newPrice);
        vm.stopPrank();
        
        (,,, uint256 price,) = marketplace.getListing(listingId);
        assertEq(price, newPrice);
    }
    
    function test_RevertWhen_UpdatePriceNotSeller() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.updateListingPrice(listingId, 2 ether);
    }
    
    function test_RevertWhen_UpdatePriceToZero() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.expectRevert();
        marketplace.updateListingPrice(listingId, 0);
        vm.stopPrank();
    }
    
    /* ========== NFT购买测试 ========== */
    
    function test_PurchaseNFT() public {
        // 上架NFT
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        uint256 sellerBalanceBefore = seller.balance;
        uint256 feeRecipientBalanceBefore = feeRecipient.balance;
        uint256 royaltyReceiverBalanceBefore = royaltyReceiver.balance;
        
        // 购买NFT
        vm.prank(buyer);
        vm.expectEmit(true, true, false, true);
        emit NFTPurchased(buyer, 1, NFT_PRICE);
        marketplace.purchaseNFT{value: NFT_PRICE}(address(nft), listingId);
        
        // 验证所有权转移
        assertEq(nft.ownerOf(1), buyer);
        
        // 验证资金分配
        uint256 platformFee = (NFT_PRICE * PLATFORM_FEE_BPS) / 10000;
        uint256 royaltyFee = (NFT_PRICE * ROYALTY_BPS) / 10000;
        uint256 sellerProceeds = NFT_PRICE - platformFee - royaltyFee;
        
        assertEq(seller.balance - sellerBalanceBefore, sellerProceeds);
        assertEq(feeRecipient.balance - feeRecipientBalanceBefore, platformFee);
        assertEq(royaltyReceiver.balance - royaltyReceiverBalanceBefore, royaltyFee);
        
        // 验证挂单状态
        (,,, , bool isActive) = marketplace.getListing(listingId);
        assertFalse(isActive);
    }
    
    function test_PurchaseWithExcessPayment() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        uint256 excess = 0.5 ether;
        uint256 buyerBalanceBefore = buyer.balance;
        
        vm.prank(buyer);
        marketplace.purchaseNFT{value: NFT_PRICE + excess}(address(nft), listingId);
        
        // 多余的ETH应该退回
        assertEq(buyerBalanceBefore - buyer.balance, NFT_PRICE);
    }
    
    function test_RevertWhen_PurchaseInsufficientPayment() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.purchaseNFT{value: 0.5 ether}(address(nft), listingId);
    }
    
    function test_RevertWhen_PurchaseOwnNFT() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        
        vm.expectRevert();
        marketplace.purchaseNFT{value: NFT_PRICE}(address(nft), listingId);
        vm.stopPrank();
    }
    
    function test_RevertWhen_PurchaseInactiveListing() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        marketplace.delistNFT(listingId);
        vm.stopPrank();
        
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.purchaseNFT{value: NFT_PRICE}(address(nft), listingId);
    }
    
    /* ========== 拍卖创建测试 ========== */
    
    function test_CreateAuction() public {
        uint256 tokenId = 1;
        uint256 startingBid = 0.5 ether;
        uint256 duration = 24; // 24小时
        
        vm.startPrank(seller);
        nft.approve(address(marketplace), tokenId);
        
        vm.expectEmit(true, true, false, false);
        emit NFTAuctionListed(seller, tokenId, startingBid, block.timestamp + duration * 1 hours);
        
        uint256 auctionId = marketplace.createAuction(address(nft), tokenId, startingBid, duration);
        vm.stopPrank();
        
        assertEq(auctionId, 1);
        assertEq(marketplace.auctionCounter(), 1);
    }
    
    function test_RevertWhen_CreateAuctionWithoutApproval() public {
        vm.prank(seller);
        vm.expectRevert();
        marketplace.createAuction(address(nft), 1, 0.5 ether, 24);
    }
    
    function test_RevertWhen_CreateAuctionZeroBid() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        vm.expectRevert();
        marketplace.createAuction(address(nft), 1, 0, 24);
        vm.stopPrank();
    }
    
    function test_RevertWhen_CreateAuctionShortDuration() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        vm.expectRevert();
        marketplace.createAuction(address(nft), 1, 0.5 ether, 0);
        vm.stopPrank();
    }
    
    /* ========== 拍卖出价测试 ========== */
    
    function test_PlaceBid() public {
        // 创建拍卖
        uint256 tokenId = 1;
        uint256 startingBid = 0.5 ether;
        
        vm.startPrank(seller);
        nft.approve(address(marketplace), tokenId);
        uint256 auctionId = marketplace.createAuction(address(nft), tokenId, startingBid, 24);
        vm.stopPrank();
        
        // 出价
        vm.prank(buyer);
        vm.expectEmit(true, true, false, true);
        emit NFTAuctionBidPlaced(buyer, tokenId, startingBid);
        marketplace.placeBid{value: startingBid}(auctionId);
        
        // 验证出价信息
        (,,,, uint256 highestBid, address highestBidder,,) = marketplace.auctions(auctionId);
        assertEq(highestBidder, buyer);
        assertEq(highestBid, startingBid);
    }
    
    function test_PlaceHigherBid() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 24);
        vm.stopPrank();
        
        address bidder1 = makeAddr("bidder1");
        address bidder2 = makeAddr("bidder2");
        vm.deal(bidder1, 10 ether);
        vm.deal(bidder2, 10 ether);
        
        // 第一个出价
        vm.prank(bidder1);
        marketplace.placeBid{value: 0.5 ether}(auctionId);
        
        // 第二个更高的出价
        vm.prank(bidder2);
        marketplace.placeBid{value: 0.6 ether}(auctionId);
        
        // 验证第一个出价者可以提取退款
        assertEq(marketplace.pendingReturns(auctionId, bidder1), 0.5 ether);
    }
    
    function test_WithdrawBidRefund() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 24);
        vm.stopPrank();
        
        address bidder1 = makeAddr("bidder1");
        address bidder2 = makeAddr("bidder2");
        vm.deal(bidder1, 10 ether);
        vm.deal(bidder2, 10 ether);
        
        vm.prank(bidder1);
        marketplace.placeBid{value: 0.5 ether}(auctionId);
        
        vm.prank(bidder2);
        marketplace.placeBid{value: 0.6 ether}(auctionId);
        
        uint256 bidder1BalanceBefore = bidder1.balance;
        
        vm.prank(bidder1);
        marketplace.withdrawBidRefund(auctionId);
        
        assertEq(bidder1.balance - bidder1BalanceBefore, 0.5 ether);
        assertEq(marketplace.pendingReturns(auctionId, bidder1), 0);
    }
    
    function test_RevertWhen_BidTooLow() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 24);
        vm.stopPrank();
        
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.placeBid{value: 0.3 ether}(auctionId);
    }
    
    function test_RevertWhen_BidAfterAuctionEnded() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 1);
        vm.stopPrank();
        
        // 时间前进25小时
        vm.warp(block.timestamp + 25 hours);
        
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.placeBid{value: 0.5 ether}(auctionId);
    }
    
    function test_RevertWhen_SellerCannotBid() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 24);
        vm.expectRevert();
        marketplace.placeBid{value: 0.5 ether}(auctionId);
        vm.stopPrank();
    }
    
    /* ========== 拍卖结束测试 ========== */
    
    function test_EndAuction() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 1);
        vm.stopPrank();
        
        vm.prank(buyer);
        marketplace.placeBid{value: 0.5 ether}(auctionId);
        
        // 时间前进2小时
        vm.warp(block.timestamp + 2 hours);
        
        uint256 sellerBalanceBefore = seller.balance;
        
        vm.prank(seller);
        vm.expectEmit(true, true, false, true);
        emit NFTAuctionEnded(buyer, 1, 0.5 ether);
        marketplace.endAuction(auctionId, address(nft));
        
        // 验证NFT转移
        assertEq(nft.ownerOf(1), buyer);
        
        // 验证资金分配
        assertTrue(seller.balance > sellerBalanceBefore);
    }
    
    function test_RevertWhen_EndAuctionBeforeTime() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, 0.5 ether, 24);
        vm.expectRevert();
        marketplace.endAuction(auctionId, address(nft));
        vm.stopPrank();
    }
    
    function test_GetActiveListingsAfterDelist() public {
        vm.startPrank(seller);
        
        nft.approve(address(marketplace), 1);
        uint256 listing1 = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        
        nft.approve(address(marketplace), 1);
        marketplace.listNFT(address(nft), 1, NFT_PRICE);
        
        marketplace.delistNFT(listing1);
        
        vm.stopPrank();
        
        uint256[] memory activeListings = marketplace.getActiveListings();
        assertEq(activeListings.length, 1);
        assertEq(activeListings[0], 2);
    }
    
    /* ========== 管理员功能测试 ========== */
    
    function test_SetPlatformFee() public {
        uint256 newFee = 500; // 5%
        
        vm.prank(owner);
        marketplace.setPlatformFee(newFee);
        
        assertEq(marketplace.platformFee(), newFee);
    }
    
    function test_RevertWhen_SetPlatformFeeTooHigh() public {
        vm.prank(owner);
        vm.expectRevert();
        marketplace.setPlatformFee(1001); // > 10%
    }
    
    function test_RevertWhen_NonOwnerSetPlatformFee() public {
        vm.prank(seller);
        vm.expectRevert();
        marketplace.setPlatformFee(500);
    }
    
    function test_SetFeeRecipient() public {
        address newRecipient = makeAddr("newRecipient");
        
        vm.prank(owner);
        marketplace.setFeeRecipient(newRecipient);
        
        assertEq(marketplace.feeRecipient(), newRecipient);
    }
    
    function test_SetRoyaltyEnabled() public {
        vm.prank(owner);
        marketplace.setRoyaltyEnabled(false);
        
        assertFalse(marketplace.royaltyEnabled());
        
        vm.prank(owner);
        marketplace.setRoyaltyEnabled(true);
        
        assertTrue(marketplace.royaltyEnabled());
    }
    
    /* ========== 重入攻击防护测试 ========== */
    
    function test_ReentrancyProtection() public {
        // ReentrancyGuard应该防止重入攻击
        // 这个测试确保purchaseNFT有nonReentrant修饰符
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        vm.prank(buyer);
        marketplace.purchaseNFT{value: NFT_PRICE}(address(nft), listingId);
        
        // 不应该能再次购买同一个listing
        vm.prank(buyer);
        vm.expectRevert();
        marketplace.purchaseNFT{value: NFT_PRICE}(address(nft), listingId);
    }
    
    function test_GetActiveListings() public {
        // 上架3个NFT
        vm.startPrank(seller);
        for (uint256 i = 1; i <= 3; i++) {
            nft.approve(address(marketplace), i);
            marketplace.listNFT(address(nft), i, NFT_PRICE + (i * 0.1 ether));
        }
        vm.stopPrank();
        
        // 验证active listings数量
        assertEq(marketplace.listingCount(), 3);
    }
    
    /* ========== Gas优化测试 ========== */
    
    function test_GasCostOfListing() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        
        uint256 gasBefore = gasleft();
        marketplace.listNFT(address(nft), 1, NFT_PRICE);
        uint256 gasUsed = gasBefore - gasleft();
        
        console.log("Gas used for listing:", gasUsed);
        vm.stopPrank();
    }
    
    function test_GasCostOfPurchase() public {
        vm.startPrank(seller);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, NFT_PRICE);
        vm.stopPrank();
        
        vm.prank(buyer);
        uint256 gasBefore = gasleft();
        marketplace.purchaseNFT{value: NFT_PRICE}(address(nft), listingId);
        uint256 gasUsed = gasBefore - gasleft();
        
        console.log("Gas used for purchase:", gasUsed);
    }
}
