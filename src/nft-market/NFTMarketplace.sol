// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

contract NFTMarketplace {
    // 这里可以添加NFT市场的功能，例如列出NFT、购买NFT等
    // NFT挂单结构体
    struct Listing {
        address seller; // 卖家地址
        uint256 tokenId; // NFT的Token ID   
        uint256 price;  // 售价（wei）
        bool isActive;  // 挂单是否激活 
    }

    // 拍卖单结构体
    struct Auction {
        address seller;      // 卖家地址
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
    // 事件：当NFT拍卖被d取消时触发
    event NFTAuctionCancelled(address indexed seller, uint256 indexed tokenId);

}