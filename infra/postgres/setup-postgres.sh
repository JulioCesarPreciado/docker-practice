#!/bin/bash
# ============================================================
# 🚀 Script to set up PostgreSQL using Docker Compose
#    - Reads the desired Docker network from .env (default: app_net)
#    - Ensures the Docker network exists
#    - Appends the external network definition to docker-compose.yml
#    - Starts the PostgreSQL container
#    - Waits for PostgreSQL to be ready
#    - Executes the init-multiple-db.sh script inside the container
# ============================================================

# Exit immediately on error
set -e

# Step 1: Check prerequisites
if ! command -v docker-compose >/dev/null 2>&1; then
  echo "❌ docker-compose is not installed. Please install it before continuing."
  exit 1
fi

# Step 2: Load configuration
NETWORK_NAME="app_net"
DB_USER="postgres"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

if [ -f .env ]; then
  ENV_NETWORK_NAME=$(grep -E '^NETWORK_NAME=' .env | cut -d '=' -f2 | tr -d '"')
  [ -n "$ENV_NETWORK_NAME" ] && NETWORK_NAME="$ENV_NETWORK_NAME"

  ENV_POSTGRES_USER=$(grep -E '^POSTGRES_USER=' .env | cut -d '=' -f2 | tr -d '\"')
  [ -n "$ENV_POSTGRES_USER" ] && DB_USER="$ENV_POSTGRES_USER"
fi

# Step 3: Append external network to docker-compose.yml if missing
if ! grep -q "networks:" "$SCRIPT_DIR/docker-compose.yml"; then
  echo -e "\nnetworks:\n  $NETWORK_NAME:\n    external: true" >> "$SCRIPT_DIR/docker-compose.yml"
else
  if ! grep -q "$NETWORK_NAME:" "$SCRIPT_DIR/docker-compose.yml"; then
    echo -e "  $NETWORK_NAME:\n    external: true" >> "$SCRIPT_DIR/docker-compose.yml"
  fi
fi

# Step 4: Ensure Docker network exists
if ! docker network ls | grep -q "$NETWORK_NAME"; then
  echo "🔧 Creating Docker network '$NETWORK_NAME'..."
  docker network create "$NETWORK_NAME"
fi

# Step 5: Start PostgreSQL container
echo "🚀 Starting PostgreSQL container (docker-compose up -d)..."
docker-compose up -d
if [ $? -ne 0 ]; then
  echo "❌ Failed to run 'docker-compose up -d'. Aborting."
  exit 1
fi

# Step 6: Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready to accept connections..."
MAX_RETRIES=30
RETRY_COUNT=0
until docker-compose exec -T postgres pg_isready -U "$DB_USER" -h localhost >/dev/null 2>&1; do
  RETRY_COUNT=$((RETRY_COUNT+1))
  if [ $RETRY_COUNT -ge $MAX_RETRIES ]; then
    echo "❌ PostgreSQL container was not ready after $((MAX_RETRIES*2)) seconds. Aborting."
    exit 1
  fi
  sleep 2
done
echo "✅ PostgreSQL container is ready."

# Step 7: Execute DB initialization script
if [ ! -f "$SCRIPT_DIR/init-multiple-db.sh" ]; then
  echo "❌ Initialization script 'init-multiple-db.sh' not found in the folder."
  exit 1
fi

echo "🏗️ Running the database initialization script inside the container..."
docker-compose exec postgres /docker-entrypoint-initdb.d/init-multiple-db.sh

echo "✅ PostgreSQL setup process completed successfully."
