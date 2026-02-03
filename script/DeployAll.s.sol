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
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("\n========================================");
        console.log("Deploying Complete NFT Ecosystem");
        console.log("========================================");
        console.log("Deployer:", deployer);
        console.log("Balance:", deployer.balance);
        console.log("========================================\n");

        vm.startBroadcast(deployerPrivateKey);

        // 0. 部署 Mock WETH 和价格预言机（本地测试）
        console.log("Step 0: Deploying Mock contracts...");
        address weth = address(new MockWETH());
        address priceFeed = address(new MockV3Aggregator(8, 2000 * 10**8)); // ETH = $2000
        console.log("Mock WETH:", weth);
        console.log("Mock Price Feed:", priceFeed);
        console.log("Mock contracts deployed!\n");

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
