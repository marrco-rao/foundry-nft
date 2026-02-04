# NFT Marketplace with UUPS Upgradeable Contracts

[![Tests](https://img.shields.io/badge/tests-72%20passed-brightgreen)](docs/test_report/)
[![Solidity](https://img.shields.io/badge/solidity-0.8.29-blue)](https://soliditylang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Foundry](https://img.shields.io/badge/Built%20with-Foundry-FFDB1C.svg)](https://getfoundry.sh/)

一个功能完整的 NFT 市场智能合约系统，支持 ERC721 NFT 铸造、交易、拍卖和版税分配，集成 Chainlink 价格预言机。

## ✨ 核心特性

- 🎨 **ERC721 NFT 合约** - 支持铸造、转移和元数据管理
- 🛒 **NFT 市场** - 完整的上架、购买和拍卖功能
- 💎 **版税系统** - 实现 ERC2981 标准，自动版税分配
- 🔄 **UUPS 可升级** - 所有合约支持安全升级
- 💰 **双支付方式** - 支持 ETH 和 WETH 支付
- 📊 **Chainlink 集成** - ETH/USD 价格查询
- 🔐 **安全防护** - 重入攻击保护和权限控制

## 🚀 快速开始

### 前置要求

```bash
# 安装 Foundry
curl -L https://foundry.paradigm.xyz | bash
foundryup
```

### 一键测试

```bash
# 自动化测试（启动节点 + 部署 + 测试）
make anvil-test

# 输出: 72/72 测试通过 ✅
```

### 编译和测试

```bash
# 安装依赖
make install

# 编译合约
make build

# 运行测试
make test

# 带 Gas 报告
make test-gas
```

## 📁 项目结构

```
├── src/
│   ├── nft/
│   │   ├── MyNFT.sol                    # 基础 NFT 合约
│   │   └── MyNFTUUPS.sol                # 可升级 NFT 合约
│   └── nft-market/
│       └── NFTMarketplaceUUPS.sol       # 可升级市场合约
├── test/
│   ├── MyNFTUUPS.t.sol                  # NFT 单元测试
│   ├── NFTMarketplaceUUPS.t.sol         # 市场单元测试
│   └── Integration.t.sol                # 集成测试
├── script/
│   ├── DeployMyNFT.s.sol               # NFT 部署脚本
│   ├── DeployMarketplace.s.sol         # 市场部署脚本
│   ├── DeployAll.s.sol                 # 完整部署脚本
│   └── AutomatedTest.s.sol             # 自动化测试脚本
└── docs/                                # 完整文档
```

## 🔧 合约功能

### NFT 合约 (MyNFTUUPS)

- `mint(address to, string uri)` - 铸造 NFT（需支付费用）
- `setMintPrice(uint256 price)` - 设置铸造价格
- `withdraw()` - 提取合约资金
- 支持 ERC2981 版税标准

### 市场合约 (NFTMarketplaceUUPS)

**交易功能：**
- `listNFT()` - 上架 NFT
- `delistNFT()` - 下架 NFT
- `purchaseNFT()` - 购买 NFT
- `updateListingPrice()` - 更新价格

**拍卖功能：**
- `createAuction()` - 创建拍卖
- `placeBid()` - 竞价
- `endAuction()` - 结束拍卖

**价格查询（Chainlink）：**
- `getETHPrice()` - 获取 ETH/USD 价格
- `getListingPriceInUSD()` - 挂单美元价格
- `getAuctionPriceInUSD()` - 拍卖美元价格

## 🧪 测试覆盖

| 测试类型 | 测试数量 | 状态 |
|---------|---------|------|
| NFT 合约测试 | ~25 | ✅ |
| 市场合约测试 | ~40 | ✅ |
| 集成测试 | ~7 | ✅ |
| **总计** | **72** | **✅ 100%** |

查看详细测试报告: [docs/test_report/](docs/test_report/)

## 📚 部署

### 本地测试（Anvil）

```bash
# 方式1：自动化部署和测试
make anvil-test

# 方式2：手动部署
make deploy-local
```

### 测试网部署（Sepolia）

```bash
# 配置环境变量
cp .env.example .env
# 编辑 .env 设置 PRIVATE_KEY 和 SEPOLIA_RPC

# 部署
make deploy-sepolia
```

详细部署指南: [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md)

## 📖 文档

| 文档 | 说明 |
|------|------|
| [TESTING.md](docs/TESTING.md) | 测试指南和使用说明 |
| [DEPLOYMENT.md](docs/DEPLOYMENT.md) | 部署流程和配置 |
| [DEPLOYMENT_QUICK_START.md](docs/DEPLOYMENT_QUICK_START.md) | 快速开始指南 |
| [nft-requirements-v1.md](docs/requirements/nft-requirements-v1.md) | 需求文档 |

## 🛠️ 开发命令

```bash
make help              # 显示所有可用命令
make install           # 安装依赖
make build             # 编译合约
make test              # 运行测试
make test-gas          # Gas 报告
make clean             # 清理构建文件
make fmt               # 格式化代码
make anvil-test        # 完整自动化测试
```

## 📊 Gas 报告

关键操作 Gas 消耗：

| 操作 | Gas 消耗 |
|------|----------|
| NFT 铸造 | ~80,000 |
| NFT 上架 | ~50,000 |
| NFT 购买 | ~100,000 |
| 创建拍卖 | ~70,000 |
| 竞价 | ~60,000 |

## 🔐 安全特性

- ✅ OpenZeppelin 合约库
- ✅ UUPS 可升级模式
- ✅ ReentrancyGuard 防重入
- ✅ Ownable 权限控制
- ✅ SafeERC20 安全转账
- ✅ 完整的单元测试和集成测试

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 License

MIT License - 详见 [LICENSE](LICENSE)

## 🔗 相关链接

- [Foundry Book](https://book.getfoundry.sh/)
- [OpenZeppelin Docs](https://docs.openzeppelin.com/)
- [ERC721 Standard](https://eips.ethereum.org/EIPS/eip-721)
- [ERC2981 Royalty Standard](https://eips.ethereum.org/EIPS/eip-2981)
- [Chainlink Price Feeds](https://docs.chain.link/data-feeds)

---

**开发团队** | 使用 [Foundry](https://getfoundry.sh/) 构建
