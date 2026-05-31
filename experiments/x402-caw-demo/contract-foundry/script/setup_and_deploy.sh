#!/bin/bash
echo "🚀 Starting Anvil and Deploying Contracts..."

# 清空代理（大小写全部清掉，curl 会读大写的版本）
unset http_proxy https_proxy HTTP_PROXY HTTPS_PROXY ALL_PROXY no_proxy NO_PROXY
git config --global --unset http.proxy 2>/dev/null || true
git config --global --unset https.proxy 2>/dev/null || true
export FOUNDRY_DISABLE_NIGHTLY_WARNING=1

# 杀掉可能残留的 anvil 和 provider 进程
pkill -f "anvil" || true
pkill -f "provider/main.go" || true
sleep 1

# 1. 后台启动 Anvil
anvil --host 0.0.0.0 &
sleep 2

# 2. 设置环境变量 (Anvil 默认账户)
export RPC_URL="http://127.0.0.1:8545"
export OWNER_PK="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
export PROVIDER_ADDR="0x70997970C51812dc3A010C7d01b50e0d17dc79C8" 
export EVIL_ADDR="0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

# 3. 编译合约
echo "🔨 Compiling contracts..."

# 确保依赖存在
forge install OpenZeppelin/openzeppelin-contracts --no-commit || true

# 生成 remappings.txt
forge remappings > remappings.txt
echo "📝 Generated remappings:"
cat remappings.txt

forge build

# 4. 部署 MockUSDC（使用更健壮的地址提取）
echo "📄 Deploying MockUSDC..."
USDC_DEPLOY_OUTPUT=$(forge create src/MockUSDC.sol:MockUSDC \
  --private-key "$OWNER_PK" \
  --rpc-url "$RPC_URL" \
  --broadcast 2>&1)

# 提取以太坊地址（只取 Deployed at 那一行）
USDC_ADDR=$(echo "$USDC_DEPLOY_OUTPUT" | grep -oE 'Deployed to: 0x[0-9a-fA-F]{40}' | grep -oE '0x[0-9a-fA-F]{40}')

if [ -z "$USDC_ADDR" ]; then
  echo "❌ MockUSDC 部署失败"
  echo "Output: $USDC_DEPLOY_OUTPUT"
  exit 1
fi
echo "✅ USDC deployed at: $USDC_ADDR"

# 5. 部署 SimpleCAW
echo "📄 Deploying SimpleCAW..."
# --constructor-args 直接传地址，forge 会自己 ABI 编码
CAW_DEPLOY_OUTPUT=$(forge create src/SimpleCAW.sol:SimpleCAW \
  --private-key "$OWNER_PK" \
  --rpc-url "$RPC_URL" \
  --broadcast \
  --constructor-args "$USDC_ADDR" 2>&1)

CAW_ADDR=$(echo "$CAW_DEPLOY_OUTPUT" | grep -oE 'Deployed to: 0x[0-9a-fA-F]{40}' | grep -oE '0x[0-9a-fA-F]{40}')

if [ -z "$CAW_ADDR" ]; then
  echo "❌ SimpleCAW 部署失败"
  echo "Output: $CAW_DEPLOY_OUTPUT"
  exit 1
fi
echo "✅ CAW deployed at: $CAW_ADDR"

# 6. 环境初始化：给 CAW 充值 & 设置策略
echo "⚙️ Setting up policies..."

# 给 CAW 充值 1000 USDC（USDC decimals=6）
AMOUNT_TO_CAW=$((1000 * 10**6))
echo "Transferring $AMOUNT_TO_CAW wei (1000 USDC) to CAW contract..."

cast send "$USDC_ADDR" \
  "transfer(address,uint256)" \
  "$CAW_ADDR" \
  "$AMOUNT_TO_CAW" \
  --private-key "$OWNER_PK" \
  --rpc-url "$RPC_URL" > /dev/null

# 设置策略：允许向 PROVIDER_ADDR 支付最多 100 USDC/天
DAILY_LIMIT=$((100 * 10**6))
echo "Setting daily limit $DAILY_LIMIT wei (100 USDC) for $PROVIDER_ADDR..."

cast send "$CAW_ADDR" \
  "setLimit(address,uint256)" \
  "$PROVIDER_ADDR" \
  "$DAILY_LIMIT" \
  --private-key "$OWNER_PK" \
  --rpc-url "$RPC_URL" > /dev/null

# 7. 将地址写入 .env 供 Go 程序读取
SBACKEND="$(cd "$(dirname "$0")/../../backend-go" && pwd)"
echo "CAW_ADDR=$CAW_ADDR" > "$SBACKEND/.env"
echo "USDC_ADDR=$USDC_ADDR" >> "$SBACKEND/.env"
echo "OWNER_PK=$OWNER_PK" >> "$SBACKEND/.env"
echo "PROVIDER_ADDR=$PROVIDER_ADDR" >> "$SBACKEND/.env"
echo "EVIL_ADDR=$EVIL_ADDR" >> "$SBACKEND/.env"

echo "✅ Setup complete! Environment variables saved to .env"