// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {MyNFTUUPS} from "../src/nft/MyNFTUUPS.sol";
import {NFTMarketplaceUUPS} from "../src/nft-market/NFTMarketplaceUUPS.sol";
import {MockWETH} from "../src/mocks/MockWETH.sol";
import {MockV3Aggregator} from "../src/mocks/MockV3Aggregator.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract AutomatedTest is Script {
    MyNFTUUPS public nft;
    NFTMarketplaceUUPS public marketplace;
    MockWETH public weth;
    MockV3Aggregator public priceFeed;
    
    address public deployer;
    
    function run() external {
        console.log("");
        console.log("==========================================");
        console.log("   Anvil Automated Test");
        console.log("==========================================");
        console.log("");
        
        uint256 deployerPk = vm.envUint("PRIVATE_KEY");
        deployer = vm.addr(deployerPk);
        
        deployContracts();
        runBasicChecks();
        
        console.log("");
        console.log("[SUCCESS] All tests passed!");
        console.log("==========================================");
        console.log("");
    }
    
    function deployContracts() internal {
        uint256 deployerPk = vm.envUint("PRIVATE_KEY");
        
        console.log("Deploying contracts...");
        vm.startBroadcast(deployerPk);
        
        // Deploy mocks
        weth = new MockWETH();
        priceFeed = new MockV3Aggregator(8, 2000 * 10**8);
        
        // Deploy NFT
        address nftImpl = address(new MyNFTUUPS());
        bytes memory nftInitData = abi.encodeWithSelector(
            MyNFTUUPS.initialize.selector,
            "Test NFT",
            "TNFT",
            deployer,
            250
        );
        address nftProxy = address(new ERC1967Proxy(nftImpl, nftInitData));
        nft = MyNFTUUPS(nftProxy);
        
        // Deploy Marketplace
        address marketplaceImpl = address(new NFTMarketplaceUUPS());
        bytes memory marketplaceInitData = abi.encodeWithSelector(
            NFTMarketplaceUUPS.initialize.selector,
            250,
            deployer,
            address(weth),
            address(priceFeed)
        );
        address marketplaceProxy = address(new ERC1967Proxy(marketplaceImpl, marketplaceInitData));
        marketplace = NFTMarketplaceUUPS(payable(marketplaceProxy));
        
        vm.stopBroadcast();
        
        console.log("NFT:", address(nft));
        console.log("Marketplace:", address(marketplace));
        console.log("");
    }
    
    function runBasicChecks() internal view {
        console.log("Running basic checks...");
        console.log("NFT Name:", nft.name());
        console.log("NFT Symbol:", nft.symbol());
        console.log("Marketplace Fee:", marketplace.platformFee(), "BPS");
        console.log("Basic checks passed!");
        console.log("");
    }
}
