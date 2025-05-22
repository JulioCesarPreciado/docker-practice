#!/bin/bash
# ============================================================
# 🧹 Global cleanup script to remove PostgreSQL, Redis, API and Socket
# ============================================================

set -e

cd "$(dirname "$0")"  # go to ./infra

# 1. Remove PostgreSQL
echo "🐘 Cleaning up PostgreSQL..."
./postgres/delete.sh "$@"

# 2. Remove Redis
echo "🔁 Cleaning up Redis..."
cd redis

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

docker-compose down --volumes --remove-orphans

if docker volume ls | grep -q redisdata; then
  echo "🗑️ Removing volume 'redisdata'..."
  docker volume rm redisdata || true
fi

cd ..

# 3. Remove API
echo "🧨 Cleaning up API..."
cd api

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi
API_PORT="${PORT:-3000}"
docker-compose down --volumes --remove-orphans

cd ..

# 4. Remove Socket
echo "🔌 Cleaning up Socket..."
cd socket

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi
SOCKET_PORT="${PORT:-4000}"
docker-compose down --volumes --remove-orphans

cd ..

# 5. Remove dangling Docker networks
echo "🔎 Checking for dangling Docker networks..."
ORPHANED=$(docker network ls -q -f dangling=true)
if [ -n "$ORPHANED" ]; then
  echo "🗑️ Removing dangling networks..."
  docker network rm $ORPHANED
fi

# 6. Summary
echo "✅ All services have been stopped and cleaned:"
echo "🧹 Redis flushed (port ${REDIS_PORT:-6379})"
echo "🧹 API removed (port ${API_PORT})"
echo "🧹 Socket removed (port ${SOCKET_PORT})"
