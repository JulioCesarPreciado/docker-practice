#!/bin/bash
# ============================================================
# 🚀 Global setup script to start PostgreSQL, Redis, API and Socket
# ============================================================

set -e

cd "$(dirname "$0")"  # go to ./infra

# 1. Start PostgreSQL
echo "🐘 Starting PostgreSQL setup..."
./postgres/setup-postgres.sh

# 2. Start Redis
echo "🔁 Starting Redis..."
cd redis
[ -f .env ] && export $(grep -v '^#' .env | xargs)
docker network inspect "${NETWORK_NAME:-app_net}" >/dev/null 2>&1 || \
  docker network create "${NETWORK_NAME:-app_net}"
docker-compose up -d
cd ..

# 3. Start API
echo "🚀 Starting API..."
cd api
[ -f .env ] && export $(grep -v '^#' .env | xargs)
API_PORT="${PORT:-3000}"
docker-compose up -d
cd ..

# 4. Start socket
echo "🔌 Starting socket..."
cd socket
[ -f .env ] && export $(grep -v '^#' .env | xargs)
SOCKET_PORT="${PORT:-4000}"
docker-compose up -d
cd ..

# 5. Final message
echo "✅ All services are up!"
echo "📌 PostgreSQL running in: ${APP_NAME:-myapp}-postgres"
echo "📌 Redis running in: ${APP_NAME:-myapp}-redis (port ${REDIS_PORT:-6379})"
echo "📌 API running in: ${APP_NAME:-myapp}-api (port ${API_PORT})"
echo "📌 Socket running in: ${APP_NAME:-myapp}-socket (port ${SOCKET_PORT})"
