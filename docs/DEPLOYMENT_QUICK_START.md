# 快速开始指南

## ⚡ 最快开始方式（推荐）

**本地测试（零配置，一键完成）：**
```bash
make anvil-test
```

此命令自动完成：启动节点 → 部署合约 → 运行测试 → 清理环境

**优势：**
- ✅ 无需配置私钥和 RPC
- ✅ 无需手动启动 Anvil
- ✅ 自动化部署和测试
- ✅ 包含完整的集成测试
- ✅ 自动清理，不留残留进程

---

## 部署脚本总结

### 📦 三种部署方式

| 脚本 | 用途 | 使用场景 |
|------|------|---------|
| `DeployMyNFT.s.sol` | 仅部署NFT合约 | 单独部署/测试NFT功能 |
| `DeployMarketplace.s.sol` | 仅部署市场合约 | 单独部署/测试市场功能 |
| `DeployAll.s.sol` | 部署完整生态 | 一次性部署所有合约（按链自动选真实地址或 Mock） |

## ✅ 推荐方案：分开编写

**原因：**
1. **灵活性** - 独立部署、升级和测试
2. **复用性** - Market可支持多个NFT集合
3. **维护性** - 代码清晰，职责分离
4. **成本优化** - 按需部署，节省gas

## 🚀 快速部署

### 选项A：分步部署（推荐用于生产）

```bash
# 1. 先部署NFT
forge script script/DeployMyNFT.s.sol:DeployMyNFT \
  --rpc-url $SEPOLIA_RPC --broadcast --verify

# 2. 测试NFT功能
# cast send <NFT_ADDRESS> "mintNFT(address,string)" <YOUR_ADDRESS> "ipfs://..."

# 3. 再部署Marketplace
forge script script/DeployMarketplace.s.sol:DeployMarketplace \
  --rpc-url $SEPOLIA_RPC --broadcast --verify
```

### 选项B：一键部署（适合快速测试）

```bash
# 方式1：环境变量
forge script script/DeployAll.s.sol:DeployAll \
  --rpc-url $SEPOLIA_RPC --broadcast --verify

# 方式2：端点别名（foundry.toml）
forge script script/DeployAll.s.sol:DeployAll \
  --rpc-url sepolia --broadcast --verify
```

## 📋 部署前检查清单

- [ ] 安装Foundry: `curl -L https://foundry.paradigm.xyz | bash`
- [ ] 创建`.env`文件并设置`PRIVATE_KEY`
- [ ] 设置RPC端点（Alchemy/Infura）
- [ ] 账户有足够测试ETH
- [ ] 设置Etherscan API密钥（用于验证）

## 🔧 环境配置

```bash
# .env 文件
PRIVATE_KEY=0x...
SEPOLIA_RPC=https://eth-sepolia.g.alchemy.com/v2/YOUR_KEY
ETHERSCAN_API_KEY=YOUR_KEY

# 加载环境变量
source .env
```

## 💡 最佳实践

### 开发阶段
```bash
# 推荐：使用自动化测试（零配置）
make anvil-test

# 或手动方式：使用本地测试网
anvil  # 启动本地节点
forge script script/DeployAll.s.sol:DeployAll --rpc-url http://localhost:8545 --broadcast
```

### 测试网部署
```bash
# Sepolia测试网
forge script script/DeployAll.s.sol:DeployAll \
  --rpc-url $SEPOLIA_RPC \
  --broadcast \
  --verify \
  -vvvv
```

### 生产环境
```bash
# 1. 先部署到测试网充分测试
# 2. 审计合约代码
# 3. 分步部署到主网
forge script script/DeployMyNFT.s.sol:DeployMyNFT \
  --rpc-url $MAINNET_RPC \
  --broadcast \
  --verify
```

## 📊 部署后验证

```bash
# 检查NFT合约
cast call <NFT_PROXY> "name()(string)" --rpc-url $SEPOLIA_RPC
cast call <NFT_PROXY> "owner()(address)" --rpc-url $SEPOLIA_RPC

# 检查Marketplace合约  
cast call <MARKET_PROXY> "platformFee()(uint256)" --rpc-url $SEPOLIA_RPC
cast call <MARKET_PROXY> "feeRecipient()(address)" --rpc-url $SEPOLIA_RPC
```

## ❓ 常见问题

**Q: 为什么分开部署？**  
A: 灵活性更高，可以单独升级，Market可以支持多个NFT集合。

**Q: DeployAll和分开部署有什么区别？**  
A: DeployAll 现在会按 chainId 自动处理依赖地址：在 Mainnet/Sepolia 使用真实 WETH + Chainlink，其他网络自动部署 Mock。分开部署仍然更灵活，适合精细化发布。

**Q: 如何升级合约？**  
A: 使用UUPS模式，只需部署新的实现合约，然后调用`upgradeTo`函数。

**Q: 部署成本是多少？**  
A: 在Sepolia测试网约0.025 ETH，主网取决于gas价格。

## 🗂️ 部署后记录清单（建议）

- 网络：chainId、RPC 提供商
- NFT：Proxy、Implementation
- Marketplace：Proxy、Implementation
- 依赖：WETH、ETH/USD Feed
- 部署交易哈希：部署与验证相关 tx

日常调用请使用 Proxy 地址；Implementation 地址用于升级与审计追踪。

## 📚 相关文档

- 详细部署指南: [docs/DEPLOYMENT.md](../docs/DEPLOYMENT.md)
- NFT需求文档: [docs/requirements/nft-requirements-v1.md](../docs/requirements/nft-requirements-v1.md)
- Foundry文档: https://book.getfoundry.sh/
