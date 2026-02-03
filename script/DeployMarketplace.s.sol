// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {NFTMarketplaceUUPS} from "../src/nft-market/NFTMarketplaceUUPS.sol";
import {MockWETH} from "../src/mocks/MockWETH.sol";
import {MockV3Aggregator} from "../src/mocks/MockV3Aggregator.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @title DeployMarketplace
 * @notice 部署NFT Marketplace UUPS可升级合约的脚本
 * @dev 部署实现合约、WETH、价格预言机和代理合约
 */
contract DeployMarketplace is Script {
    // 默认配置
    uint256 public constant DEFAULT_PLATFORM_FEE_BPS = 250; // 2.5%
    
    // Chainlink ETH/USD 价格预言机地址 (主网和测试网)
    address public constant MAINNET_ETH_USD_FEED = 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419;
    address public constant SEPOLIA_ETH_USD_FEED = 0x694AA1769357215DE4FAC081bf1f309aDC325306;
    
    // WETH 地址 (主网和测试网)
    address public constant MAINNET_WETH = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
    address public constant SEPOLIA_WETH = 0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9;

    function run() external returns (
        address marketplaceProxy,
        address marketplaceImpl,
        address weth,
        address priceFeed
    ) {
        // 从环境变量读取部署者私钥
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        console.log("Deploying NFT Marketplace with account:", deployer);
        console.log("Account balance:", deployer.balance);
        
        // 检测当前链
        uint256 chainId = block.chainid;
        console.log("Chain ID:", chainId);

        vm.startBroadcast(deployerPrivateKey);

        // 根据链 ID 选择地址或部署 Mock
        if (chainId == 1) { // Mainnet
            weth = MAINNET_WETH;
            priceFeed = MAINNET_ETH_USD_FEED;
            console.log("Using Mainnet WETH:", weth);
            console.log("Using Mainnet ETH/USD Feed:", priceFeed);
        } else if (chainId == 11155111) { // Sepolia
            weth = SEPOLIA_WETH;
            priceFeed = SEPOLIA_ETH_USD_FEED;
            console.log("Using Sepolia WETH:", weth);
            console.log("Using Sepolia ETH/USD Feed:", priceFeed);
        } else { // Local/Test Network - 部署 Mock
            console.log("Deploying Mock contracts for local/test network...");
            
            // 部署 Mock WETH
            weth = address(new MockWETH());
            console.log("Mock WETH deployed at:", weth);
            
            // 部署 Mock 价格预言机 (ETH = $2000, 8 decimals)
            priceFeed = address(new MockV3Aggregator(8, 2000 * 10**8));
            console.log("Mock ETH/USD Price Feed deployed at:", priceFeed);
        }

        // 1. 部署市场实现合约
        marketplaceImpl = address(new NFTMarketplaceUUPS());
        console.log("Marketplace Implementation deployed at:", marketplaceImpl);

        // 2. 准备初始化数据
        address feeRecipient = deployer; // 默认手续费接收者为部署者
        bytes memory initData = abi.encodeWithSelector(
            NFTMarketplaceUUPS.initialize.selector,
            DEFAULT_PLATFORM_FEE_BPS,
            feeRecipient,
            weth,
            priceFeed
        );

        // 3. 部署代理合约
        marketplaceProxy = address(new ERC1967Proxy(marketplaceImpl, initData));
        console.log("Marketplace Proxy deployed at:", marketplaceProxy);
        console.log("Platform Fee (bps):", DEFAULT_PLATFORM_FEE_BPS);
        console.log("Fee Recipient:", feeRecipient);

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Marketplace Proxy Address:", marketplaceProxy);
        console.log("Marketplace Implementation Address:", marketplaceImpl);
        console.log("WETH Address:", weth);
        console.log("ETH/USD Price Feed:", priceFeed);
        
        return (marketplaceProxy, marketplaceImpl, weth, priceFeed);
    }

    /**
     * @notice 使用自定义参数部署Marketplace
     * @dev 可以在其他脚本中调用此函数
     */
    function deployWithParams(
        uint256 platformFeeBps,
        address feeRecipient,
        address _weth,
        address _priceFeed
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
            feeRecipient,
            _weth,
            _priceFeed
        );

        // 部署代理合约
        proxy = address(new ERC1967Proxy(implementation, initData));
        console.log("Proxy deployed at:", proxy);
        console.log("Platform Fee:", platformFeeBps, "BPS");
        console.log("Fee Recipient:", feeRecipient);
        console.log("WETH:", _weth);
        console.log("Price Feed:", _priceFeed);

        return (proxy, implementation);
    }
}
