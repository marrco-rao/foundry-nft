#!/usr/bin/env bash
set -euo pipefail

# ------------------------------------------------------------
# DeployMarketplace-repair.sh
# Purpose:
#   Repair deployment for marketplace when DeployAll/DeployMarketplace
#   script deployment fails due to contract-size issue of script contract.
#
# Usage:
#   1) Fill PRIVATE_KEY and SEPOLIA_RPC below (or export env vars).
#   2) chmod +x script/DeployMarketplace-repair.sh
#   3) ./script/DeployMarketplace-repair.sh
#
# Optional env vars:
#   PLATFORM_FEE_BPS (default: 250)
#   FEE_RECIPIENT    (default: derived from PRIVATE_KEY)
#   WETH             (default: Sepolia WETH)
#   PRICE_FEED       (default: Sepolia ETH/USD feed)
#   VERIFY_IMPL      (default: false)
#   ETHERSCAN_API_KEY (required only when VERIFY_IMPL=true)
# ------------------------------------------------------------

PRIVATE_KEY="${PRIVATE_KEY:-0xYOUR_PRIVATE_KEY_HERE}"
SEPOLIA_RPC="${SEPOLIA_RPC:-https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY_HERE}"
PLATFORM_FEE_BPS="${PLATFORM_FEE_BPS:-250}"
FEE_RECIPIENT="${FEE_RECIPIENT:-}"
WETH="${WETH:-0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9}"
PRICE_FEED="${PRICE_FEED:-0x694AA1769357215DE4FAC081bf1f309aDC325306}"
VERIFY_IMPL="${VERIFY_IMPL:-false}"
ETHERSCAN_API_KEY="${ETHERSCAN_API_KEY:-}"

extract_address() {
  # Only accept explicit deployed-address markers from forge output.
  local raw="$1"
  local addr
  local clean

  clean=$(echo "$raw" | sed -E 's/\x1b\[[0-9;]*m//g')

  # forge create commonly prints: "Deployed to: 0x..."
  addr=$(echo "$clean" | awk '/Deployed to:/{print $3}' | tail -n1)
  if [[ "$addr" =~ ^0x[a-fA-F0-9]{40}$ ]]; then
    echo "$addr"
    return 0
  fi

  # Some outputs may include: "Contract Address: 0x..."
  addr=$(echo "$clean" | awk '/Contract Address:/{print $3}' | tail -n1)
  if [[ "$addr" =~ ^0x[a-fA-F0-9]{40}$ ]]; then
    echo "$addr"
    return 0
  fi

  echo ""
}

if [[ "$PRIVATE_KEY" == "0xYOUR_PRIVATE_KEY_HERE" ]]; then
  echo "ERROR: Please set PRIVATE_KEY first."
  exit 1
fi

if [[ "$SEPOLIA_RPC" == *"YOUR_API_KEY_HERE"* ]]; then
  echo "ERROR: Please set SEPOLIA_RPC first."
  exit 1
fi

if ! command -v forge >/dev/null 2>&1; then
  echo "ERROR: forge not found in PATH"
  exit 1
fi

if ! command -v cast >/dev/null 2>&1; then
  echo "ERROR: cast not found in PATH"
  exit 1
fi

echo "==> Checking connected chain"
CHAIN_ID=$(cast chain-id --rpc-url "$SEPOLIA_RPC")
if [[ "$CHAIN_ID" != "11155111" ]]; then
  echo "ERROR: RPC is not Sepolia (chainId=$CHAIN_ID)."
  exit 1
fi

if [[ -z "$FEE_RECIPIENT" ]]; then
  FEE_RECIPIENT=$(cast wallet address --private-key "$PRIVATE_KEY")
fi

echo "==> Deploy config"
echo "Chain ID          : $CHAIN_ID"
echo "Fee Recipient     : $FEE_RECIPIENT"
echo "Platform Fee (bps): $PLATFORM_FEE_BPS"
echo "WETH              : $WETH"
echo "Price Feed        : $PRICE_FEED"

echo "==> Deploying NFTMarketplaceUUPS implementation"
IMPL_OUT=$(forge create src/nft-market/NFTMarketplaceUUPS.sol:NFTMarketplaceUUPS \
  --rpc-url "$SEPOLIA_RPC" \
  --private-key "$PRIVATE_KEY" \
  --broadcast 2>&1)

echo "$IMPL_OUT"
MARKET_IMPL=$(extract_address "$IMPL_OUT")
if [[ -z "$MARKET_IMPL" ]]; then
  echo "ERROR: Failed to parse MARKET_IMPL from forge output"
  echo "Hint: ensure forge create actually broadcasted and did not fail before deployment."
  echo "Raw forge output:"
  echo "$IMPL_OUT"
  exit 1
fi

echo "==> Building initialize calldata"
INIT_DATA=$(cast calldata "initialize(uint256,address,address,address)" \
  "$PLATFORM_FEE_BPS" "$FEE_RECIPIENT" "$WETH" "$PRICE_FEED")

echo "==> Deploying ERC1967Proxy"
PROXY_OUT=$(forge create lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol:ERC1967Proxy \
  --rpc-url "$SEPOLIA_RPC" \
  --private-key "$PRIVATE_KEY" \
  --broadcast \
  --constructor-args "$MARKET_IMPL" "$INIT_DATA" 2>&1)

echo "$PROXY_OUT"
MARKET_PROXY=$(extract_address "$PROXY_OUT")
if [[ -z "$MARKET_PROXY" ]]; then
  echo "ERROR: Failed to parse MARKET_PROXY from forge output"
  echo "Hint: ensure --broadcast is present and constructor args are valid."
  echo "Raw forge output:"
  echo "$PROXY_OUT"
  exit 1
fi

echo "==> Basic on-chain checks"
cast code "$MARKET_PROXY" --rpc-url "$SEPOLIA_RPC" >/dev/null
PF=$(cast call "$MARKET_PROXY" "platformFee()(uint256)" --rpc-url "$SEPOLIA_RPC")
FR=$(cast call "$MARKET_PROXY" "feeRecipient()(address)" --rpc-url "$SEPOLIA_RPC")
OW=$(cast call "$MARKET_PROXY" "owner()(address)" --rpc-url "$SEPOLIA_RPC")

echo
echo "================ Deployment Result ================"
echo "MARKET_IMPL  : $MARKET_IMPL"
echo "MARKET_PROXY : $MARKET_PROXY"
echo "platformFee  : $PF"
echo "feeRecipient : $FR"
echo "owner        : $OW"
echo "=================================================="
echo "Note: Addresses above are on-chain because --broadcast is enabled."

echo
echo "Export helpers:"
echo "export MARKET_IMPL=$MARKET_IMPL"
echo "export MARKET_PROXY=$MARKET_PROXY"

if [[ "$VERIFY_IMPL" == "true" ]]; then
  if [[ -z "$ETHERSCAN_API_KEY" ]]; then
    echo "ERROR: VERIFY_IMPL=true but ETHERSCAN_API_KEY is empty"
    exit 1
  fi

  echo "==> Verifying implementation on Etherscan"
  forge verify-contract "$MARKET_IMPL" src/nft-market/NFTMarketplaceUUPS.sol:NFTMarketplaceUUPS \
    --chain sepolia \
    --etherscan-api-key "$ETHERSCAN_API_KEY" \
    --watch
fi
