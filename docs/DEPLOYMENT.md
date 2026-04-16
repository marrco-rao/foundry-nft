# NFT 项目部署指南

## 部署脚本说明

本项目提供了三个部署脚本，满足不同的部署需求：

### 1. DeployMyNFT.s.sol
**用途**：单独部署NFT合约

**适用场景**：
- 只想部署NFT合约
- 需要部署多个不同的NFT集合
- 测试NFT合约功能

**运行命令**：
```bash
forge script script/DeployMyNFT.s.sol:DeployMyNFT --rpc-url <RPC_URL> --broadcast --verify
```

### 2. DeployMarketplace.s.sol
**用途**：单独部署NFT Marketplace合约

**适用场景**：
- 只想部署Marketplace合约
- NFT合约已存在，只需要市场功能
- 升级或重新部署市场合约

**运行命令**：
```bash
forge script script/DeployMarketplace.s.sol:DeployMarketplace --rpc-url <RPC_URL> --broadcast --verify
```

### 3. DeployAll.s.sol
**用途**：一次性部署完整的NFT生态系统

**适用场景**：
- 首次部署完整项目
- 部署到新的测试网或主网
- 需要完整的NFT + Marketplace生态

**网络行为说明**：
- Mainnet/Sepolia：自动使用链上真实 WETH 和 Chainlink ETH/USD Feed 地址
- 本地或其他测试链：自动部署 MockWETH 和 MockV3Aggregator

**运行命令**：
```bash
forge script script/DeployAll.s.sol:DeployAll --rpc-url <RPC_URL> --broadcast --verify
```

## 部署前准备

### 1. 设置环境变量

创建 `.env` 文件：
```bash
# 部署者私钥
PRIVATE_KEY=你的私钥

# RPC端点
SEPOLIA_RPC=https://eth-sepolia.g.alchemy.com/v2/your-api-key
MAINNET_RPC=https://eth-mainnet.g.alchemy.com/v2/your-api-key

# Etherscan API密钥（用于验证合约）
ETHERSCAN_API_KEY=你的etherscan-api-key
```

> 注意：`foundry.toml` 中 Sepolia 端点变量名为 `SEPOLIA_RPC`。

加载环境变量：
```bash
source .env
```

### 2. 确保账户有足够的ETH
检查余额：
```bash
cast balance <你的地址> --rpc-url $SEPOLIA_RPC
```

## 部署示例

### 部署到Sepolia测试网

#### 方式1：部署所有合约
```bash
# 写法A：直接使用环境变量
forge script script/DeployAll.s.sol:DeployAll \
  --rpc-url $SEPOLIA_RPC \
  --broadcast \
  --verify \
  -vvvv

# 写法B：使用 foundry.toml 里的端点别名
forge script script/DeployAll.s.sol:DeployAll \
  --rpc-url sepolia \
  --broadcast \
  --verify \
  -vvvv
```

#### 方式2：分别部署

**先部署NFT：**
```bash
forge script script/DeployMyNFT.s.sol:DeployMyNFT \
  --rpc-url $SEPOLIA_RPC \
  --broadcast \
  --verify \
  -vvvv
```

**再部署Marketplace：**
```bash
forge script script/DeployMarketplace.s.sol:DeployMarketplace \
  --rpc-url $SEPOLIA_RPC \
  --broadcast \
  --verify \
  -vvvv
```

### 本地测试部署

#### 方式1：自动化测试部署（推荐）

**一键完成启动、部署和测试：**
```bash
make anvil-test
```

#### 方式2：手动部署

使用Anvil本地测试网：
```bash
# 终端1：启动本地节点
anvil

# 终端2：部署合约
make deploy-local

# 或使用完整命令
forge script script/DeployAll.s.sol:DeployAll \
  --rpc-url http://localhost:8545 \
  --broadcast
```

#### 方式3：快速测试（Anvil 已运行）

如果 Anvil 已在另一个终端运行：
```bash
make anvil-quick
```

## DeployAll 网络地址策略

部署 `DeployAll.s.sol` 时，会按 `chainId` 自动选择依赖地址：

- `1`（Mainnet）：使用主网 WETH 与主网 ETH/USD Feed
- `11155111`（Sepolia）：使用 Sepolia WETH 与 Sepolia ETH/USD Feed
- 其他 chainId：部署 MockWETH 与 MockV3Aggregator

这保证了在 Sepolia 不会误用 Mock 地址。

## 部署后验证

### 1. 验证NFT合约
```bash
# 获取NFT名称
cast call <NFT_PROXY_ADDRESS> "name()(string)" --rpc-url $SEPOLIA_RPC

# 获取NFT符号
cast call <NFT_PROXY_ADDRESS> "symbol()(string)" --rpc-url $SEPOLIA_RPC

# 获取所有者
cast call <NFT_PROXY_ADDRESS> "owner()(address)" --rpc-url $SEPOLIA_RPC
```

### 2. 验证Marketplace合约
```bash
# 获取平台手续费
cast call <MARKETPLACE_PROXY_ADDRESS> "platformFee()(uint256)" --rpc-url $SEPOLIA_RPC

# 获取手续费接收地址
cast call <MARKETPLACE_PROXY_ADDRESS> "feeRecipient()(address)" --rpc-url $SEPOLIA_RPC
```

## 部署后必须记录的信息

建议把以下信息写入你的部署记录（如 Notion/表格/JSON）：

- 网络信息：`chainId`、RPC 提供商
- 部署账户：EOA 地址
- NFT：Proxy 地址、Implementation 地址
- Marketplace：Proxy 地址、Implementation 地址
- 依赖地址：WETH、ETH/USD Price Feed
- 部署交易哈希：每个创建交易的 tx hash

说明：日常交互与业务调用应使用 Proxy 地址，Implementation 地址主要用于升级和审计追踪。

## 部署架构说明

### UUPS代理模式
两个合约都使用UUPS（Universal Upgradeable Proxy Standard）代理模式：

```
用户 → 代理合约 (Proxy) → 实现合约 (Implementation)
         [存储状态]            [业务逻辑]
```

**优点**：
- 可升级：可以更新业务逻辑而不改变合约地址
- 节省gas：代理合约轻量
- 保持状态：升级时保留所有数据

### 合约关系图
```
MyNFTUUPS (NFT合约)
    ↓ (NFT持有者可以在市场上交易)
NFTMarketplaceUUPS (市场合约)
    - 支持任意ERC721合约
    - 自动处理ERC2981版税
```

## 最佳实践建议

### 1. 部署顺序
对于生产环境，推荐：
1. 先部署并测试NFT合约
2. 铸造一些测试NFT
3. 再部署Marketplace合约
4. 测试市场功能
5. 最后公开上线

### 2. 分开部署的优势
- ✅ **灵活性**：可以单独升级每个合约
- ✅ **复用性**：Marketplace可以支持多个NFT集合
- ✅ **风险控制**：分阶段部署降低风险
- ✅ **测试便利**：可以单独测试每个合约

### 3. 统一部署的优势
- ✅ **便捷性**：一次性部署完成
- ✅ **一致性**：确保所有合约参数一致
- ✅ **演示友好**：适合快速演示或测试

## 故障排查

### 编译失败
```bash
forge clean
forge build
```

### Gas不足
增加gas limit：
```bash
--gas-limit 5000000
```

### 验证失败
手动验证：
```bash
forge verify-contract <CONTRACT_ADDRESS> <CONTRACT_NAME> \
  --chain sepolia \
  --etherscan-api-key $ETHERSCAN_API_KEY
```

## 成本估算

基于Sepolia测试网的估算（实际成本取决于gas价格）：
- NFT合约部署：~0.01 ETH
- Marketplace部署：~0.015 ETH
- 总计：~0.025 ETH

## 相关文档

- [Foundry Book](https://book.getfoundry.sh/)
- [OpenZeppelin UUPS](https://docs.openzeppelin.com/contracts/4.x/api/proxy#UUPSUpgradeable)
- [ERC721 Standard](https://eips.ethereum.org/EIPS/eip-721)
- [ERC2981 Royalty Standard](https://eips.ethereum.org/EIPS/eip-2981)
