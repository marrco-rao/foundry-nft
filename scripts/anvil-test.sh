#!/bin/bash

##############################################################################
# Anvil 自动化测试脚本
# 功能：启动 Anvil -> 部署合约 -> 运行测试 -> 清理
##############################################################################

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Anvil 配置
ANVIL_PORT=8545
ANVIL_RPC="http://127.0.0.1:${ANVIL_PORT}"
ANVIL_PID_FILE=".anvil.pid"

# 测试账户（Anvil 默认第一个账户）
export PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
export ANVIL_RPC=$ANVIL_RPC

# 合约地址文件
DEPLOYMENT_FILE=".deployment-addresses.json"

##############################################################################
# 函数：清理环境
##############################################################################
cleanup() {
    log_info "清理环境..."
    
    # 停止 Anvil
    if [ -f "$ANVIL_PID_FILE" ]; then
        ANVIL_PID=$(cat "$ANVIL_PID_FILE")
        if ps -p $ANVIL_PID > /dev/null 2>&1; then
            log_info "停止 Anvil (PID: $ANVIL_PID)..."
            kill $ANVIL_PID 2>/dev/null || true
            sleep 1
        fi
        rm -f "$ANVIL_PID_FILE"
    fi
    
    # 清理临时文件
    rm -f "$DEPLOYMENT_FILE"
    
    log_success "清理完成"
}

# 设置退出时自动清理
trap cleanup EXIT INT TERM

##############################################################################
# 函数：启动 Anvil
##############################################################################
start_anvil() {
    log_info "启动 Anvil 本地节点..."
    
    # 检查端口是否被占用
    if lsof -Pi :$ANVIL_PORT -sTCP:LISTEN -t >/dev/null 2>&1; then
        log_warning "端口 $ANVIL_PORT 已被占用"
        log_info "尝试停止占用该端口的进程..."
        lsof -ti:$ANVIL_PORT | xargs kill -9 2>/dev/null || true
        sleep 2
        
        # 再次检查端口
        if lsof -Pi :$ANVIL_PORT -sTCP:LISTEN -t >/dev/null 2>&1; then
            log_error "无法释放端口 $ANVIL_PORT，请手动停止占用该端口的进程"
            exit 1
        fi
    fi
    
    # 启动 Anvil（后台运行）
    anvil --port $ANVIL_PORT > anvil.log 2>&1 &
    ANVIL_PID=$!
    echo $ANVIL_PID > "$ANVIL_PID_FILE"
    
    log_info "Anvil PID: $ANVIL_PID"
    log_info "等待 Anvil 启动..."
    
    # 等待 Anvil 完全启动
    max_attempts=30
    attempt=0
    while ! curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        $ANVIL_RPC > /dev/null 2>&1; do
        
        attempt=$((attempt + 1))
        if [ $attempt -ge $max_attempts ]; then
            log_error "Anvil 启动超时"
            cat anvil.log
            exit 1
        fi
        
        sleep 0.5
        echo -n "."
    done
    
    echo ""
    log_success "Anvil 启动成功！RPC: $ANVIL_RPC"
}

##############################################################################
# 函数：编译合约
##############################################################################
compile_contracts() {
    log_info "编译合约..."
    forge build
    
    if [ $? -eq 0 ]; then
        log_success "合约编译成功"
    else
        log_error "合约编译失败"
        exit 1
    fi
}

##############################################################################
# 函数：部署合约并运行集成测试
##############################################################################
deploy_and_test() {
    log_info "运行集成测试（含部署）..."
    
    # 运行自动化测试脚本
    forge script script/AutomatedTest.s.sol:AutomatedTest \
        --rpc-url $ANVIL_RPC \
        --broadcast \
        -v
    
    if [ $? -ne 0 ]; then
        log_error "集成测试失败"
        exit 1
    fi
    
    log_success "集成测试通过"
}

##############################################################################
# 函数：运行 Foundry 测试
##############################################################################
run_foundry_tests() {
    log_info "运行 Foundry 测试套件..."
    
    forge test --fork-url $ANVIL_RPC -vv
    
    if [ $? -eq 0 ]; then
        log_success "所有测试通过 ✓"
    else
        log_error "测试失败"
        exit 1
    fi
}

##############################################################################
# 函数：运行集成测试
##############################################################################
run_integration_tests() {
    log_info "运行集成测试脚本..."
    
    # 运行自定义的交互测试脚本
    if [ -f "script/AutomatedTest.s.sol" ]; then
        log_info "找到集成测试脚本，开始执行..."
        forge script script/AutomatedTest.s.sol:AutomatedTest \
            --rpc-url $ANVIL_RPC \
            --broadcast \
            -vv
        
        if [ $? -eq 0 ]; then
            log_success "集成测试通过 ✓"
        else
            log_warning "集成测试失败（可能需要修复）"
            # 不退出，继续执行
        fi
    else
        log_warning "未找到集成测试脚本，跳过"
    fi
}

##############################################################################
# 函数：生成测试报告
##############################################################################
generate_report() {
    log_info "生成测试报告..."
    
    echo ""
    echo "=========================================="
    echo "           测试报告"
    echo "=========================================="
    forge test --fork-url $ANVIL_RPC --gas-report | tail -n 20
    echo "=========================================="
}

##############################################################################
# 主流程
##############################################################################
main() {
    echo ""
    echo "=========================================="
    echo "   Anvil 自动化测试"
    echo "=========================================="
    echo ""
    
    # 1. 启动 Anvil
    start_anvil
    
    # 2. 编译合约
    compile_contracts
    
    # 3. 部署并运行集成测试
    deploy_and_test
    
    echo ""
    log_info "等待 2 秒让部署稳定..."
    sleep 2
    
    # 4. 生成报告
    generate_report
    
    echo ""
    log_success "所有测试完成！"
    echo ""
}

# 执行主流程
main
