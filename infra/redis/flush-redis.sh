#!/bin/bash
# ============================================================
# 🧹 Secure script to flush all Redis data from the container
# ============================================================

set -e

# Load .env if it exists
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

CONTAINER_NAME="${APP_NAME:-myapp}-redis"
REDIS_PASSWORD="${REDIS_PASSWORD:-changeme}"

echo "⚠️ This will delete ALL data from Redis in container '$CONTAINER_NAME'."
read -p "Do you want to continue? (y/N): " CONFIRM
[[ "$CONFIRM" != "y" ]] && echo "❌ Operation canceled." && exit 1

echo "🧹 Flushing all Redis databases..."
docker exec -e REDISCLI_AUTH="$REDIS_PASSWORD" -i "$CONTAINER_NAME" redis-cli FLUSHALL

echo "✅ Redis has been flushed securely."
