# 测试报告目录

本目录包含项目的自动化测试报告和日志文件。

---

## 📋 文件说明

### 测试报告

| 文件 | 类型 | 说明 |
|------|------|------|
| `TEST_REPORT_*.md` | Markdown | 格式化的测试报告，包含统计数据和分析 |
| `anvil-test-*.log` | 日志文件 | 完整的测试输出日志（包含所有细节） |

### 当前报告

- **最新测试报告:** [TEST_REPORT_20260204.md](./TEST_REPORT_20260204.md)
- **最新测试日志:** [anvil-test-20260204-092737.log](./anvil-test-20260204-092737.log)

---

## 🚀 如何生成新的测试报告

### 方法1：自动保存日志（推荐）

```bash
# 运行测试并保存日志
make anvil-test 2>&1 | tee docs/test_report/anvil-test-$(date +%Y%m%d-%H%M%S).log
```

### 方法2：仅运行测试

```bash
# 运行测试（不保存）
make anvil-test
```

### 方法3：保存到自定义文件

```bash
# 保存到指定文件名
make anvil-test > docs/test_report/my-test.log 2>&1
```

---

## 📊 测试报告内容

每个测试报告包含：

### 1. 测试结果概览
- 总测试数
- 通过/失败/跳过数量
- 通过率
- 执行时间

### 2. 测试流程
- Anvil 节点启动
- 合约编译
- 合约部署
- 集成测试执行

### 3. Gas 使用报告
- 关键函数 Gas 消耗
- 最小/平均/最大值统计
- 优化建议

### 4. 测试套件详情
- 测试文件统计
- 测试用例列表
- 覆盖范围

### 5. 编译器警告
- Linting 提示
- 警告分析
- 优化建议

---

## 📈 测试历史

| 日期 | 测试数 | 通过率 | 执行时间 | 报告 |
|------|--------|--------|----------|------|
| 2026-02-04 | 72 | 100% | 122.42ms | [查看](./TEST_REPORT_20260204.md) |

---

## 🔍 查看测试详情

### 查看格式化报告

```bash
# 在终端查看
cat docs/test_report/TEST_REPORT_20260204.md

# 使用 Markdown 预览器
open docs/test_report/TEST_REPORT_20260204.md
```

### 查看原始日志

```bash
# 查看完整日志
cat docs/test_report/anvil-test-20260204-092737.log

# 分页查看
less docs/test_report/anvil-test-20260204-092737.log

# 搜索关键词
grep "test_" docs/test_report/anvil-test-20260204-092737.log
```

### 分析 Gas 使用

```bash
# 提取 Gas 报告
grep -A 20 "测试报告" docs/test_report/anvil-test-20260204-092737.log

# 查看部署信息
grep "NFT:" docs/test_report/anvil-test-20260204-092737.log
```

---

## 🎯 最佳实践

### 定期测试

建议在以下情况运行测试并保存报告：

1. **代码变更后** - 确保功能正常
2. **部署前** - 验证所有功能
3. **每日构建** - 持续集成
4. **发布前** - 完整回归测试

### 报告命名规范

```bash
# 推荐格式
anvil-test-YYYYMMDD-HHMMSS.log
TEST_REPORT_YYYYMMDD.md

# 示例
anvil-test-20260204-092737.log
TEST_REPORT_20260204.md
```

### 保留历史

- 保留至少最近 10 次测试报告
- 重要里程碑的报告永久保留
- 定期清理旧报告（>30天）

---

## 🛠️ 故障排查

### 测试失败

如果测试失败，检查日志文件中的：

```bash
# 查找错误
grep "FAIL\|Error\|failed" docs/test_report/*.log

# 查看失败的测试
grep "test_.*FAIL" docs/test_report/*.log
```

### 部署失败

```bash
# 查看部署错误
grep "deployment.*fail\|Failed.*deploy" docs/test_report/*.log
```

### Gas 问题

```bash
# 查看 Gas 消耗
grep -A 10 "Gas used:" docs/test_report/*.log
```

---

## 📚 相关文档

- [测试指南](../TESTING.md)
- [部署指南](../DEPLOYMENT.md)
- [快速开始](../DEPLOYMENT_QUICK_START.md)
- [需求文档](../requirements/nft-requirements-v1.md)

---

## 🔗 持续集成

项目可以集成到 CI/CD 流程：

```yaml
# .github/workflows/test.yml 示例
- name: Run Tests
  run: |
    make anvil-test | tee test-report-${{ github.run_number }}.log
    
- name: Upload Report
  uses: actions/upload-artifact@v3
  with:
    name: test-reports
    path: test-report-*.log
```

---

**维护者:** NFT 项目团队  
**最后更新:** 2026年2月4日
