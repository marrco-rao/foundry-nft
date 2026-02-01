// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {MyNFTUUPS} from "../src/nft/MyNFTUUPS.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @title DeployMyNFT
 * @notice 部署MyNFT UUPS可升级合约的脚本
 * @dev 部署实现合约和代理合约，并初始化
 */
contract DeployMyNFT is Script {
    // 可以通过环境变量覆盖这些默认值
    string public constant DEFAULT_NAME = "My NFT Collection";
    string public constant DEFAULT_SYMBOL = "MYNFT";
    uint96 public constant DEFAULT_ROYALTY_BPS = 250; // 2.5%

    function run() external returns (address proxy, address implementation) {
        // 从环境变量读取部署者私钥
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Deploying MyNFT with account:", deployer);
        console.log("Account balance:", deployer.balance);

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署实现合约
        implementation = address(new MyNFTUUPS());
        console.log("Implementation deployed at:", implementation);

        // 2. 准备初始化数据
        // 版税接收者默认为部署者
        address royaltyReceiver = deployer;
        bytes memory initData = abi.encodeWithSelector(
            MyNFTUUPS.initialize.selector,
            DEFAULT_NAME,
            DEFAULT_SYMBOL,
            royaltyReceiver,
            DEFAULT_ROYALTY_BPS
        );

        // 3. 部署代理合约
        proxy = address(new ERC1967Proxy(implementation, initData));
        console.log("Proxy deployed at:", proxy);
        console.log("NFT Collection Name:", DEFAULT_NAME);
        console.log("NFT Collection Symbol:", DEFAULT_SYMBOL);
        console.log("Royalty Receiver:", royaltyReceiver);
        console.log("Royalty BPS:", DEFAULT_ROYALTY_BPS);

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("NFT Proxy Address:", proxy);
        console.log("NFT Implementation Address:", implementation);
    }

    /**
     * @notice 使用自定义参数部署NFT
     * @dev 可以在其他脚本中调用此函数
     */
    function deployWithParams(
        string memory name,
        string memory symbol,
        address royaltyReceiver,
        uint96 royaltyBps
    ) public returns (address proxy, address implementation) {
        // 部署实现合约
        implementation = address(new MyNFTUUPS());
        console.log("Implementation deployed at:", implementation);

        // 准备初始化数据
        bytes memory initData = abi.encodeWithSelector(
            MyNFTUUPS.initialize.selector,
            name,
            symbol,
            royaltyReceiver,
            royaltyBps
        );

        // 部署代理合约
        proxy = address(new ERC1967Proxy(implementation, initData));
        console.log("Proxy deployed at:", proxy);
        console.log("NFT Collection:", name);
        console.log("Symbol:", symbol);
        console.log("Royalty BPS:", royaltyBps);
        console.log("Royalty Receiver:", royaltyReceiver);

        return (proxy, implementation);
    }
}
