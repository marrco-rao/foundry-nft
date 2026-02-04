# Makefile for Foundry NFT Project
# 简化常用命令

.PHONY: help install build test clean anvil-test deploy-local interact-local

# 默认显示帮助
help:
	@echo "可用命令:"
	@echo "  make install        - 安装依赖"
	@echo "  make build          - 编译合约"
	@echo "  make test           - 运行测试"
	@echo "  make test-gas       - 运行测试并显示 gas 报告"
	@echo "  make clean          - 清理构建文件"
	@echo ""
	@echo "Anvil 本地测试:"
	@echo "  make anvil-test     - 完整自动化测试（启动 Anvil + 部署 + 测试）"
	@echo "  make anvil-quick    - 快速测试（需要先启动 Anvil）"
	@echo "  make deploy-local   - 部署到本地 Anvil"
	@echo "  make interact-local - 运行交互测试"
	@echo ""
	@echo "其他:"
	@echo "  make fmt            - 格式化代码"
	@echo "  make lint           - 代码检查"

# 安装依赖
install:
	forge install

# 编译合约
build:
	forge build

# 运行测试
test:
	forge test -vv

# 测试 + Gas 报告
test-gas:
	forge test --gas-report

# 清理
clean:
	forge clean
	rm -rf cache out broadcast

# 代码格式化
fmt:
	forge fmt

# 代码检查
lint:
	forge fmt --check

##############################################################################
# Anvil 相关命令
##############################################################################

# 完整自动化测试（启动 Anvil + 部署 + 测试）
anvil-test:
	@chmod +x scripts/anvil-test.sh
	@./scripts/anvil-test.sh

# 快速测试（假设 Anvil 已在运行）
anvil-quick:
	@chmod +x scripts/anvil-quick-test.sh
	@./scripts/anvil-quick-test.sh

# 部署到本地 Anvil（需要先启动 Anvil）
deploy-local:
	@echo "部署到本地 Anvil..."
	@export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 && \
	forge script script/DeployAll.s.sol:DeployAll \
		--rpc-url http://127.0.0.1:8545 \
		--broadcast \
		-vv

# 运行交互测试脚本
interact-local:
	@echo "运行交互测试..."
	@export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 && \
	forge script script/AutomatedTest.s.sol:AutomatedTest \
		--rpc-url http://127.0.0.1:8545 \
		--broadcast \
		-vv

##############################################################################
# 测试网部署（示例）
##############################################################################

# 部署到 Sepolia（需要配置 .env）
deploy-sepolia:
	@echo "部署到 Sepolia..."
	@forge script script/DeployAll.s.sol:DeployAll \
		--rpc-url $(SEPOLIA_RPC) \
		--broadcast \
		--verify \
		-vvvv

# 查看部署的合约
verify-sepolia:
	@forge verify-contract $(CONTRACT_ADDRESS) \
		--chain-id 11155111 \
		--watch
