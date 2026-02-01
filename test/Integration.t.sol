// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Test, console} from "forge-std/Test.sol";
import {NFTMarketplaceUUPS} from "../src/nft-market/NFTMarketplaceUUPS.sol";
import {MyNFTUUPS} from "../src/nft/MyNFTUUPS.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @title Integration Test Suite
 * @notice 测试NFT和Marketplace的完整集成流程
 */
contract IntegrationTest is Test {
    NFTMarketplaceUUPS public marketplace;
    MyNFTUUPS public nft;
    
    address public owner;
    address public artist;
    address public collector1;
    address public collector2;
    address public collector3;
    address public platformOwner;
    address public royaltyReceiver;
    
    uint256 public constant MINT_PRICE = 0.01 ether;
    uint256 public constant PLATFORM_FEE = 250; // 2.5%
    uint256 public constant ROYALTY_FEE = 250; // 2.5%
    
    function setUp() public {
        // 设置角色
        owner = makeAddr("owner");
        artist = makeAddr("artist");
        collector1 = makeAddr("collector1");
        collector2 = makeAddr("collector2");
        collector3 = makeAddr("collector3");
        platformOwner = makeAddr("platformOwner");
        royaltyReceiver = artist; // 艺术家接收版税
        
        // 充值
        vm.deal(owner, 100 ether);
        vm.deal(artist, 100 ether);
        vm.deal(collector1, 100 ether);
        vm.deal(collector2, 100 ether);
        vm.deal(collector3, 100 ether);
        
        // 部署NFT合约
        vm.startPrank(owner);
        address nftImpl = address(new MyNFTUUPS());
        bytes memory nftInitData = abi.encodeWithSelector(
            MyNFTUUPS.initialize.selector,
            "Artist Collection",
            "ART",
            royaltyReceiver,
            ROYALTY_FEE
        );
        address nftProxy = address(new ERC1967Proxy(nftImpl, nftInitData));
        nft = MyNFTUUPS(nftProxy);
        vm.stopPrank();
        
        // 部署Marketplace合约
        vm.startPrank(platformOwner);
        address marketImpl = address(new NFTMarketplaceUUPS());
        bytes memory marketInitData = abi.encodeWithSelector(
            NFTMarketplaceUUPS.initialize.selector,
            PLATFORM_FEE,
            platformOwner
        );
        address marketProxy = address(new ERC1967Proxy(marketImpl, marketInitData));
        marketplace = NFTMarketplaceUUPS(marketProxy);
        vm.stopPrank();
    }
    
    /* ========== 完整用户旅程测试 ========== */
    
    function test_CompleteNFTLifecycle() public {
        console.log("\n=== Complete NFT Lifecycle Test ===\n");
        
        // 1. 艺术家铸造NFT
        console.log("Step 1: Artist mints NFT");
        vm.prank(artist);
        nft.mint{value: MINT_PRICE}(artist, "ipfs://QmArtwork1");
        assertEq(nft.ownerOf(1), artist);
        console.log("  NFT minted to artist");
        
        // 2. 艺术家在市场上架NFT
        console.log("\nStep 2: Artist lists NFT on marketplace");
        uint256 listPrice = 1 ether;
        vm.startPrank(artist);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, listPrice);
        vm.stopPrank();
        console.log("  NFT listed for", listPrice);
        
        // 3. 收藏家1购买NFT
        console.log("\nStep 3: Collector1 purchases NFT");
        uint256 artist_balance_before = artist.balance;
        uint256 platform_balance_before = platformOwner.balance;
        uint256 royalty_balance_before = royaltyReceiver.balance;
        
        vm.prank(collector1);
        marketplace.purchaseNFT{value: listPrice}(address(nft), listingId);
        
        assertEq(nft.ownerOf(1), collector1);
        console.log("  NFT transferred to collector1");
        
        // 验证第一次销售的资金分配
        // artist既是卖家又是版税接收者，所以收到全部（除了平台费）
        uint256 platformFee = (listPrice * PLATFORM_FEE) / 10000;
        uint256 royaltyFee = (listPrice * ROYALTY_FEE) / 10000;
        uint256 sellerAmount = listPrice - platformFee - royaltyFee;
        uint256 artistTotalReceives = sellerAmount + royaltyFee; // artist收到全部除了平台费
        
        assertApproxEqAbs(artist.balance - artist_balance_before, artistTotalReceives, 1);
        assertApproxEqAbs(platformOwner.balance - platform_balance_before, platformFee, 1);
        console.log("  Funds distributed correctly");
        
        // 4. 收藏家1转售NFT
        console.log("\nStep 4: Collector1 resells NFT");
        uint256 resalePrice = 2 ether;
        vm.startPrank(collector1);
        nft.approve(address(marketplace), 1);
        uint256 resaleListingId = marketplace.listNFT(address(nft), 1, resalePrice);
        vm.stopPrank();
        console.log("  NFT relisted for", resalePrice);
        
        // 5. 收藏家2购买NFT（验证版税支付）
        console.log("\nStep 5: Collector2 purchases from Collector1");
        uint256 collector1_balance_before = collector1.balance;
        artist_balance_before = royaltyReceiver.balance;
        platform_balance_before = platformOwner.balance;
        
        vm.prank(collector2);
        marketplace.purchaseNFT{value: resalePrice}(address(nft), resaleListingId);
        
        assertEq(nft.ownerOf(1), collector2);
        console.log("  NFT transferred to collector2");
        
        // 验证二次销售的资金分配（包含版税）
        platformFee = (resalePrice * PLATFORM_FEE) / 10000;
        royaltyFee = (resalePrice * ROYALTY_FEE) / 10000;
        uint256 sellerProceeds = resalePrice - platformFee - royaltyFee;
        
        assertApproxEqAbs(collector1.balance - collector1_balance_before, sellerProceeds, 1);
        assertApproxEqAbs(platformOwner.balance - platform_balance_before, platformFee, 1);
        assertApproxEqAbs(royaltyReceiver.balance - artist_balance_before, royaltyFee, 1);
        console.log("  Royalty paid to artist:", royaltyFee);
        console.log("  Platform fee:", platformFee);
        console.log("  Seller proceeds:", sellerProceeds);
    }
    
    function test_MultipleNFTsMarketplace() public {
        console.log("\n=== Multiple NFTs Marketplace Test ===\n");
        
        // 艺术家铸造多个NFT
        console.log("Step 1: Artist mints multiple NFTs");
        vm.startPrank(artist);
        for (uint256 i = 1; i <= 5; i++) {
            nft.mint{value: MINT_PRICE}(
                artist,
                string(abi.encodePacked("ipfs://QmArtwork", vm.toString(i)))
            );
        }
        vm.stopPrank();
        console.log("  5 NFTs minted");
        
        // 上架所有NFT
        console.log("\nStep 2: List all NFTs with different prices");
        vm.startPrank(artist);
        uint256[] memory listingIds = new uint256[](5);
        for (uint256 i = 1; i <= 5; i++) {
            nft.approve(address(marketplace), i);
            uint256 price = i * 0.5 ether;
            listingIds[i-1] = marketplace.listNFT(address(nft), i, price);
            console.log("  NFT", i, "listed for", price);
        }
        vm.stopPrank();
        
        // 多个收藏家购买不同NFT
        console.log("\nStep 3: Multiple collectors purchase different NFTs");
        
        vm.prank(collector1);
        marketplace.purchaseNFT{value: 0.5 ether}(address(nft), listingIds[0]);
        console.log("  Collector1 bought NFT #1");
        
        vm.prank(collector2);
        marketplace.purchaseNFT{value: 1 ether}(address(nft), listingIds[1]);
        console.log("  Collector2 bought NFT #2");
        
        vm.prank(collector3);
        marketplace.purchaseNFT{value: 1.5 ether}(address(nft), listingIds[2]);
        console.log("  Collector3 bought NFT #3");
        
        // 验证所有权
        assertEq(nft.ownerOf(1), collector1);
        assertEq(nft.ownerOf(2), collector2);
        assertEq(nft.ownerOf(3), collector3);
        assertEq(nft.ownerOf(4), artist); // 未售出
        assertEq(nft.ownerOf(5), artist); // 未售出
        
        console.log("\n  All NFTs transferred correctly");
    }
    
    function test_AuctionCompleteFlow() public {
        console.log("\n=== Complete Auction Flow Test ===\n");
        
        // 1. 艺术家铸造NFT
        console.log("Step 1: Artist mints rare NFT");
        vm.prank(artist);
        nft.mint{value: MINT_PRICE}(artist, "ipfs://QmRareArtwork");
        
        // 2. 创建拍卖
        console.log("\nStep 2: Artist creates auction");
        uint256 startingBid = 1 ether;
        uint256 duration = 48; // 48小时
        
        vm.startPrank(artist);
        nft.approve(address(marketplace), 1);
        uint256 auctionId = marketplace.createAuction(address(nft), 1, startingBid, duration);
        vm.stopPrank();
        console.log("  Auction created with starting bid:", startingBid);
        
        // 3. 多个竞拍者出价
        console.log("\nStep 3: Multiple bidders place bids");
        
        vm.prank(collector1);
        marketplace.placeBid{value: 1 ether}(auctionId);
        console.log("  Collector1 bids: 1 ETH");
        
        vm.prank(collector2);
        marketplace.placeBid{value: 1.5 ether}(auctionId);
        console.log("  Collector2 bids: 1.5 ETH");
        
        vm.prank(collector3);
        marketplace.placeBid{value: 2 ether}(auctionId);
        console.log("  Collector3 bids: 2 ETH");
        
        // 验证待退款
        assertEq(marketplace.pendingReturns(auctionId, collector1), 1 ether);
        assertEq(marketplace.pendingReturns(auctionId, collector2), 1.5 ether);
        
        // 4. 失败的竞拍者提取退款
        console.log("\nStep 4: Losing bidders withdraw refunds");
        
        uint256 collector1_balance_before = collector1.balance;
        vm.prank(collector1);
        marketplace.withdrawBidRefund(auctionId);
        assertEq(collector1.balance - collector1_balance_before, 1 ether);
        console.log("  Collector1 withdrew: 1 ETH");
        
        uint256 collector2_balance_before = collector2.balance;
        vm.prank(collector2);
        marketplace.withdrawBidRefund(auctionId);
        assertEq(collector2.balance - collector2_balance_before, 1.5 ether);
        console.log("  Collector2 withdrew: 1.5 ETH");
        
        // 5. 拍卖结束
        console.log("\nStep 5: Auction ends");
        vm.warp(block.timestamp + 49 hours);
        
        uint256 artist_balance_before = artist.balance;
        uint256 platform_balance_before = platformOwner.balance;
        
        vm.prank(artist);
        marketplace.endAuction(auctionId, address(nft));
        
        assertEq(nft.ownerOf(1), collector3);
        console.log("  Winner: Collector3");
        console.log("  Final bid: 2 ETH");
        
        // 验证资金分配
        // artist既是卖家又是版税接收者，所以收到: sellerAmount + royaltyAmount
        uint256 platformFee = (2 ether * PLATFORM_FEE) / 10000;
        uint256 royaltyFee = (2 ether * ROYALTY_FEE) / 10000;
        uint256 sellerAmount = 2 ether - platformFee - royaltyFee;
        uint256 artistTotalReceives = sellerAmount + royaltyFee; // artist收到卖家款+版税
        
        assertApproxEqAbs(artist.balance - artist_balance_before, artistTotalReceives, 1);
        assertApproxEqAbs(platformOwner.balance - platform_balance_before, platformFee, 1);
        console.log("  Artist received (seller + royalty):", artistTotalReceives);
        console.log("  Platform fee:", platformFee);
    }
    
    function test_PriceUpdateAndCancel() public {
        console.log("\n=== Price Update and Cancel Test ===\n");
        
        // 铸造并上架
        vm.prank(artist);
        nft.mint{value: MINT_PRICE}(artist, "ipfs://QmArtwork");
        
        vm.startPrank(artist);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, 1 ether);
        console.log("Listed NFT for: 1 ETH");
        
        // 更新价格
        marketplace.updateListingPrice(listingId, 2 ether);
        console.log("Updated price to: 2 ETH");
        
        (,,, uint256 price,) = marketplace.getListing(listingId);
        assertEq(price, 2 ether);
        
        // 再次更新
        marketplace.updateListingPrice(listingId, 0.5 ether);
        console.log("Updated price to: 0.5 ETH");
        
        (,,, price,) = marketplace.getListing(listingId);
        assertEq(price, 0.5 ether);
        
        // 取消挂单
        marketplace.delistNFT(listingId);
        console.log("Delisted NFT");
        vm.stopPrank();
        
        (,,,, bool isActive) = marketplace.getListing(listingId);
        assertFalse(isActive);
        
        // NFT仍然属于艺术家
        assertEq(nft.ownerOf(1), artist);
        console.log("NFT returned to artist");
    }
    
    function test_CrossCollectionMarketplace() public {
        console.log("\n=== Cross-Collection Marketplace Test ===\n");
        
        // 部署第二个NFT合约
        vm.prank(owner);
        address nft2Impl = address(new MyNFTUUPS());
        bytes memory nft2InitData = abi.encodeWithSelector(
            MyNFTUUPS.initialize.selector,
            "Second Collection",
            "SEC",
            royaltyReceiver,
            ROYALTY_FEE
        );
        address nft2Proxy = address(new ERC1967Proxy(nft2Impl, nft2InitData));
        MyNFTUUPS nft2 = MyNFTUUPS(nft2Proxy);
        
        console.log("Deployed second NFT collection");
        
        // 在两个集合中各铸造NFT
        vm.startPrank(artist);
        nft.mint{value: MINT_PRICE}(artist, "ipfs://Collection1/1");
        nft2.mint{value: MINT_PRICE}(artist, "ipfs://Collection2/1");
        
        // 在同一个市场上架两个集合的NFT
        nft.approve(address(marketplace), 1);
        nft2.approve(address(marketplace), 1);
        
        uint256 listing1 = marketplace.listNFT(address(nft), 1, 1 ether);
        uint256 listing2 = marketplace.listNFT(address(nft2), 1, 1.5 ether);
        vm.stopPrank();
        
        console.log("Listed NFTs from both collections");
        
        // 不同收藏家购买不同集合的NFT
        vm.prank(collector1);
        marketplace.purchaseNFT{value: 1 ether}(address(nft), listing1);
        
        vm.prank(collector2);
        marketplace.purchaseNFT{value: 1.5 ether}(address(nft2), listing2);
        
        assertEq(nft.ownerOf(1), collector1);
        assertEq(nft2.ownerOf(1), collector2);
        
        console.log("Cross-collection trading successful");
    }
    
    function test_MarketplaceUpgrade() public {
        console.log("\n=== Marketplace Upgrade Test ===\n");
        
        // 在市场上有活跃挂单
        vm.prank(artist);
        nft.mint{value: MINT_PRICE}(artist, "ipfs://QmArtwork");
        
        vm.startPrank(artist);
        nft.approve(address(marketplace), 1);
        uint256 listingId = marketplace.listNFT(address(nft), 1, 1 ether);
        vm.stopPrank();
        
        console.log("Created listing before upgrade");
        
        // 升级市场合约（使用marketplace的owner即platformOwner）
        vm.startPrank(platformOwner);
        NFTMarketplaceUUPS newImpl = new NFTMarketplaceUUPS();
        marketplace.upgradeToAndCall(address(newImpl), "");
        vm.stopPrank();
        
        console.log("Marketplace upgraded");
        
        // 验证数据保持不变
        assertEq(marketplace.listingCount(), 1);
        assertEq(marketplace.platformFee(), PLATFORM_FEE);
        assertEq(marketplace.feeRecipient(), platformOwner);
        
        (address seller,,, uint256 price, bool isActive) = marketplace.getListing(listingId);
        assertEq(seller, artist);
        assertEq(price, 1 ether);
        assertTrue(isActive);
        
        console.log("All data preserved after upgrade");
        
        // 升级后仍可正常购买
        vm.prank(collector1);
        marketplace.purchaseNFT{value: 1 ether}(address(nft), listingId);
        
        assertEq(nft.ownerOf(1), collector1);
        console.log("Trading works after upgrade");
    }
    
    function test_StressTest() public {
        console.log("\n=== Stress Test ===\n");
        
        uint256 nftCount = 10;
        
        // 批量铸造
        console.log("Minting", nftCount, "NFTs...");
        vm.startPrank(artist);
        for (uint256 i = 1; i <= nftCount; i++) {
            nft.mint{value: MINT_PRICE}(
                artist,
                string(abi.encodePacked("ipfs://Batch", vm.toString(i)))
            );
        }
        vm.stopPrank();
        
        // 批量上架
        console.log("Listing all NFTs...");
        vm.startPrank(artist);
        for (uint256 i = 1; i <= nftCount; i++) {
            nft.approve(address(marketplace), i);
            marketplace.listNFT(address(nft), i, i * 0.1 ether);
        }
        vm.stopPrank();
        
        assertEq(marketplace.listingCount(), nftCount);
        
        console.log("All", nftCount, "NFTs listed successfully");
        console.log("Marketplace stress test passed");
    }
}
