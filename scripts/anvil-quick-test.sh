#!/bin/bash

##############################################################################
# Anvil 快速测试脚本 - 简化版
# 用法: ./scripts/anvil-quick-test.sh
##############################################################################

set -e

# 检查 Anvil 是否在运行
check_anvil() {
    if curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        http://127.0.0.1:8545 > /dev/null 2>&1; then
        echo "✓ Anvil is running"
        return 0
    else
        echo "✗ Anvil is not running"
        echo ""
        echo "请先在另一个终端启动 Anvil:"
        echo "  anvil"
        echo ""
        exit 1
    fi
}

echo ""
echo "=========================================="
echo "   Anvil 快速测试"
echo "=========================================="
echo ""

# 检查 Anvil
check_anvil

# 设置环境变量
export PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
export ANVIL_RPC="http://127.0.0.1:8545"

echo "1. 编译合约..."
forge build --silent

echo "2. 运行测试..."
forge test --fork-url $ANVIL_RPC -vv

echo ""
echo "✓ 测试完成！"
echo ""
