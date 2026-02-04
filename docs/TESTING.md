# 测试指南

## 🚀 快速开始：Anvil 自动化测试

本项目提供了完全自动化的测试流程，**一键完成所有测试**：

```bash
make anvil-test
```

此命令会自动完成：
1. ✅ 启动 Anvil 本地节点
2. ✅ 编译所有合约
3. ✅ 部署完整生态系统（NFT + Marketplace + Mock合约）
4. ✅ 运行集成测试
5. ✅ 生成 Gas 报告
6. ✅ 自动清理环境

**其他可用命令：**
```bash
make anvil-quick    # 快速测试（需预先启动 Anvil）
make deploy-local   # 仅部署到本地
make interact-local # 运行交互测试
make test          # 运行单元测试
make test-gas      # 显示 gas 报告
```

---

## 测试文件结构

本项目包含三个测试文件，全面覆盖NFT和Marketplace的所有功能：

```
test/
├── MyNFTUUPS.t.sol          # NFT合约单元测试
├── NFTMarketplaceUUPS.t.sol # Marketplace合约单元测试
└── Integration.t.sol         # 集成测试（完整用户旅程）
```

## 测试覆盖范围

### 1. MyNFTUUPS.t.sol - NFT合约测试

✅ **初始化测试**
- 合约初始化验证
- 版税信息配置
- 防止重复初始化

✅ **NFT铸造测试**
- 基本铸造功能
- 批量铸造
- 铸造到其他地址
- 支付验证
- 最大供应量限制

✅ **所有权管理测试**
- 设置铸造价格
- 提取资金
- 权限控制

✅ **ERC721标准功能**
- NFT转移
- 授权和批准
- 批准所有操作符

✅ **ERC2981版税测试**
- 版税信息查询
- 版税接收者配置

✅ **UUPS升级测试**
- 合约升级
- 升级权限控制
- 状态保持验证

✅ **接口支持测试**
- ERC721接口
- ERC2981接口
- ERC165接口

✅ **Gas优化测试**
- 铸造gas消耗

✅ **Fuzz测试**
- 随机价格测试
- 随机URI测试

✅ **边界条件测试**
- 精确支付
- 超额支付退款

### 2. NFTMarketplaceUUPS.t.sol - Marketplace测试

✅ **初始化测试**
- Marketplace初始化
- 配置验证

✅ **NFT上架测试**
- 单个NFT上架
- 批量上架
- 授权验证
- 价格验证

✅ **NFT下架测试**
- 下架功能
- 权限控制
- 状态验证

✅ **价格更新测试**
- 更新挂单价格
- 权限验证

✅ **NFT购买测试**
- 购买流程
- 资金分配（平台费、版税、卖家收益）
- 超额支付退款
- 权限验证

✅ **拍卖创建测试**
- 创建拍卖
- 参数验证
- 授权验证

✅ **拍卖出价测试**
- 出价功能
- 价格递增验证
- 退款机制
- 提取退款

✅ **拍卖结束测试**
- 结束拍卖
- NFT转移
- 资金分配
- 时间验证

✅ **查询功能测试**
- 获取活跃挂单
- 挂单信息查询

✅ **管理员功能测试**
- 设置平台费率
- 设置手续费接收地址
- 启用/禁用版税

✅ **安全测试**
- 重入攻击防护

✅ **Gas优化测试**
- 上架gas消耗
- 购买gas消耗

### 3. Integration.t.sol - 集成测试

✅ **完整NFT生命周期**
- 铸造 → 上架 → 购买 → 转售
- 版税支付验证
- 多次交易流程

✅ **多NFT市场**
- 批量铸造和上架
- 多买家购买
- 活跃挂单查询

✅ **完整拍卖流程**
- 创建拍卖
- 多人竞拍
- 出价退款
- 拍卖结束和资金分配

✅ **价格更新和取消**
- 价格调整
- 取消挂单
- 状态验证

✅ **跨集合市场**
- 多个NFT合约
- 统一市场交易

✅ **合约升级**
- Marketplace升级
- 数据保持
- 升级后功能验证

✅ **压力测试**
- 批量操作
- 性能验证

## 运行测试

### 🎯 推荐：使用 Anvil 自动化测试（一键测试）

这是最简单、最完整的测试方式：

```bash
make anvil-test
```

**输出示例：**
```
==========================================
   Anvil 自动化测试
==========================================

[INFO] 启动 Anvil 本地节点...
[SUCCESS] Anvil 启动成功！
[INFO] 编译合约...
[SUCCESS] 合约编译成功
[INFO] 运行集成测试（含部署）...
[SUCCESS] 集成测试通过
[INFO] 生成测试报告...
[SUCCESS] 所有测试完成！
```

### 运行单元测试

```bash
# 使用 Makefile（推荐）
make test

# 或直接使用 forge
forge test
```

### 运行特定测试文件

```bash
# NFT合约测试
forge test --match-path test/MyNFTUUPS.t.sol

# Marketplace测试
forge test --match-path test/NFTMarketplaceUUPS.t.sol

# 集成测试
forge test --match-path test/Integration.t.sol
```

### 运行特定测试函数

```bash
# 运行特定测试
forge test --match-test test_MintNFT

# 运行包含特定关键词的测试
forge test --match-test "Purchase"
```

### 详细输出

```bash
# -v: 基本输出
# -vv: 显示所有测试
# -vvv: 显示失败的stack trace
# -vvvv: 显示所有trace和setup
# -vvvvv: 显示所有opcode

forge test -vv                    # 推荐
forge test -vvv                   # 调试时使用
forge test --match-path test/Integration.t.sol -vv  # 查看集成测试详情
```

### Gas报告

```bash
forge test --gas-report
```

### 覆盖率报告

```bash
forge coverage
```

### 生成详细覆盖率报告

```bash
forge coverage --report lcov
genhtml lcov.info -o coverage
open coverage/index.html
```

## 测试最佳实践

### 1. 测试命名规范

```solidity
// ✅ 正确
function test_MintNFT() public { }
function test_RevertWhen_InsufficientPayment() public { }
function testFuzz_MintPrice(uint256 price) public { }

// ❌ 错误
function testFail_SomethingBad() public { }  // 已弃用
```

### 2. 使用 vm.expectRevert

```solidity
// 新写法
function test_RevertWhen_InsufficientPayment() public {
    vm.prank(user1);
    vm.expectRevert("Insufficient payment for minting");
    nft.mint{value: 0.001 ether}(user1, "ipfs://QmTest");
}
```

### 3. 测试组织

```solidity
contract MyContractTest is Test {
    // 1. State variables
    MyContract public myContract;
    
    // 2. setUp function
    function setUp() public { }
    
    // 3. Tests grouped by feature
    /* ========== Feature A Tests ========== */
    function test_FeatureA1() public { }
    function test_FeatureA2() public { }
    
    /* ========== Feature B Tests ========== */
    function test_FeatureB1() public { }
}
```

## 已知问题和注意事项

### 1. testFail* 已弃用

Foundry已弃用`testFail*`前缀。如果看到以下警告：

```
[FAIL: `testFail*` has been removed. Consider changing to test_Revert[If|When]_Condition and expecting a revert]
```

**解决方案**：将这些测试改写为使用`vm.expectRevert()`

### 2. NFT合约函数名

- ✅ 实际函数名：`mint(address to, string calldata uri)`
- ❌ 不是：`mintNFT()`

### 3. 版税功能

当前NFT合约中版税在初始化时设置，没有`setDefaultRoyalty`函数。

## 测试统计

### 当前测试数量

- **总测试数**: 72 个测试 ✅
- **MyNFTUUPS.t.sol**: ~25个测试
- **NFTMarketplaceUUPS.t.sol**: ~40个测试
- **Integration.t.sol**: ~7个集成测试
- **测试通过率**: 100%

### 总覆盖率目标

- 函数覆盖率: >90%
- 分支覆盖率: >85%
- 行覆盖率: >90%

## 示例：运行和查看测试

```bash
# 1. 最推荐：完整自动化测试（包括部署）
make anvil-test

# 2. 快速运行单元测试
make test

# 3. 带 Gas 报告的测试
make test-gas

# 4. 查看详细的集成测试输出
forge test --match-path test/Integration.t.sol -vv

# 5. 运行特定功能的测试
forge test --match-test "Auction" -vv

# 6. 生成覆盖率报告
forge coverage

# 7. 调试失败的测试
forge test --match-test test_PurchaseNFT -vvvv
```

## 持续集成

建议在CI/CD中运行：

```yaml
# .github/workflows/test.yml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          submodules: recursive
      
      - name: Install Foundry
        uses: foundry-rs/foundry-toolchain@v1
      
      - name: Run tests
        run: forge test -vv
      
      - name: Generate coverage
        run: forge coverage
```

## 下一步

1. 根据需要添加更多边界条件测试
2. 增加更多Fuzz测试覆盖
3. 添加不变量测试（Invariant Tests）
4. 定期检查和更新测试以匹配合约变更
5. 保持测试覆盖率在90%以上

## 相关文档

- [Foundry Book - Testing](https://book.getfoundry.sh/forge/tests)
- [Foundry Book - Cheatcodes](https://book.getfoundry.sh/cheatcodes/)
- [Foundry Book - Coverage](https://book.getfoundry.sh/reference/forge/forge-coverage)
