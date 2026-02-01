// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {NFTMarketplaceUUPS} from "../src/nft-market/NFTMarketplaceUUPS.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @title DeployMarketplace
 * @notice 部署NFT Marketplace UUPS可升级合约的脚本
 * @dev 部署实现合约和代理合约，并初始化
 */
contract DeployMarketplace is Script {
    // 默认配置
    uint256 public constant DEFAULT_PLATFORM_FEE_BPS = 250; // 2.5%

    function run() external returns (address proxy, address implementation) {
        // 从环境变量读取部署者私钥
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Deploying NFT Marketplace with account:", deployer);
        console.log("Account balance:", deployer.balance);

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署实现合约
        implementation = address(new NFTMarketplaceUUPS());
        console.log("Marketplace Implementation deployed at:", implementation);

        // 2. 准备初始化数据
        // 手续费接收者默认为部署者
        address feeRecipient = deployer;
        bytes memory initData = abi.encodeWithSelector(
            NFTMarketplaceUUPS.initialize.selector,
            DEFAULT_PLATFORM_FEE_BPS,
            feeRecipient
        );

        // 3. 部署代理合约
        proxy = address(new ERC1967Proxy(implementation, initData));
        console.log("Marketplace Proxy deployed at:", proxy);
        console.log("Platform Fee:", DEFAULT_PLATFORM_FEE_BPS, "BPS");
        console.log("Fee Recipient:", feeRecipient);

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Marketplace Proxy Address:", proxy);
        console.log("Marketplace Implementation Address:", implementation);
    }

    /**
     * @notice 使用自定义参数部署Marketplace
     * @dev 可以在其他脚本中调用此函数
     */
    function deployWithParams(
        uint256 platformFeeBps,
        address feeRecipient
    ) public returns (address proxy, address implementation) {
        require(platformFeeBps <= 1000, "Fee too high"); // 最大10%
        require(feeRecipient != address(0), "Invalid fee recipient");

        // 部署实现合约
        implementation = address(new NFTMarketplaceUUPS());
        console.log("Implementation deployed at:", implementation);

        // 准备初始化数据
        bytes memory initData = abi.encodeWithSelector(
            NFTMarketplaceUUPS.initialize.selector,
            platformFeeBps,
            feeRecipient
        );

        // 部署代理合约
        proxy = address(new ERC1967Proxy(implementation, initData));
        console.log("Proxy deployed at:", proxy);
        console.log("Platform Fee:", platformFeeBps, "BPS");
        console.log("Fee Recipient:", feeRecipient);

        return (proxy, implementation);
    }
}
