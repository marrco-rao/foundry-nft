// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {DeployMyNFT} from "./DeployMyNFT.s.sol";
import {DeployMarketplace} from "./DeployMarketplace.s.sol";
import {MockWETH} from "../src/mocks/MockWETH.sol";
import {MockV3Aggregator} from "../src/mocks/MockV3Aggregator.sol";

/**
 * @title DeployAll
 * @notice 一次性部署NFT合约和Marketplace合约的脚本
 * @dev 按顺序部署两个合约，并输出完整的部署信息
 */
contract DeployAll is Script {
    // Chainlink ETH/USD 价格预言机地址
    address public constant MAINNET_ETH_USD_FEED = 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419;
    address public constant SEPOLIA_ETH_USD_FEED = 0x694AA1769357215DE4FAC081bf1f309aDC325306;

    // WETH 地址
    address public constant MAINNET_WETH = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
    address public constant SEPOLIA_WETH = 0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9;

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        uint256 chainId = block.chainid;
        
        console.log("\n========================================");
        console.log("Deploying Complete NFT Ecosystem");
        console.log("========================================");
        console.log("Deployer:", deployer);
        console.log("Chain ID:", chainId);
        console.log("Balance:", deployer.balance);
        console.log("========================================\n");

        vm.startBroadcast(deployerPrivateKey);

        // 0. 根据网络配置支付代币和价格预言机地址
        address weth;
        address priceFeed;
        if (chainId == 1) {
            weth = MAINNET_WETH;
            priceFeed = MAINNET_ETH_USD_FEED;
            console.log("Step 0: Using Mainnet addresses...");
            console.log("Mainnet WETH:", weth);
            console.log("Mainnet ETH/USD Feed:", priceFeed);
        } else if (chainId == 11155111) {
            weth = SEPOLIA_WETH;
            priceFeed = SEPOLIA_ETH_USD_FEED;
            console.log("Step 0: Using Sepolia addresses...");
            console.log("Sepolia WETH:", weth);
            console.log("Sepolia ETH/USD Feed:", priceFeed);
        } else {
            console.log("Step 0: Deploying Mock contracts for local/test network...");
            weth = address(new MockWETH());
            priceFeed = address(new MockV3Aggregator(8, 2000 * 10**8)); // ETH = $2000
            console.log("Mock WETH:", weth);
            console.log("Mock Price Feed:", priceFeed);
        }
        console.log("Step 0 completed!\n");

        // 1. 部署NFT合约
        console.log("Step 1: Deploying NFT Contract...");
        DeployMyNFT nftDeployer = new DeployMyNFT();
        (address nftProxy, address nftImpl) = nftDeployer.deployWithParams(
            "My NFT Collection",
            "MYNFT",
            deployer, // 版税接收者
            250       // 2.5% 版税
        );
        console.log("NFT deployed successfully!\n");

        // 2. 部署Marketplace合约
        console.log("Step 2: Deploying Marketplace Contract...");
        DeployMarketplace marketDeployer = new DeployMarketplace();
        (address marketProxy, address marketImpl) = marketDeployer.deployWithParams(
            250,      // 2.5% 平台手续费
            deployer, // 手续费接收者
            weth,     // WETH 地址
            priceFeed // 价格预言机地址
        );
        console.log("Marketplace deployed successfully!\n");

        vm.stopBroadcast();

        // 3. 输出完整的部署信息
        console.log("\n========================================");
        console.log("DEPLOYMENT COMPLETED");
        console.log("========================================");
        console.log("\nNFT Contract:");
        console.log("  Proxy:", nftProxy);
        console.log("  Implementation:", nftImpl);
        console.log("\nMarketplace Contract:");
        console.log("  Proxy:", marketProxy);
        console.log("  Implementation:", marketImpl);
        console.log("\nSupporting Contracts:");
        console.log("  WETH:", weth);
        console.log("  ETH/USD Price Feed:", priceFeed);
        console.log("\nConfiguration:");
        console.log("  Deployer:", deployer);
        console.log("  NFT Royalty: 2.5%");
        console.log("  Marketplace Fee: 2.5%");
        console.log("  Supported Payments: ETH, WETH");
        console.log("========================================\n");

        // 4. 输出后续操作提示
        console.log("Next Steps:");
        console.log("1. Verify contracts on block explorer");
        console.log("2. Test NFT minting");
        console.log("3. Approve marketplace to handle NFTs");
        console.log("4. Test marketplace listing and purchase (ETH/WETH)");
        console.log("5. Query USD prices using Chainlink oracle");
    }
}
