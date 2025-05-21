#!/bin/bash
# ============================================================
# 🚀 Script to deploy PostgreSQL with Docker and run the
#    multi-database initialization script.
# ============================================================

# Exit immediately on error
set -e

# 1. Check if docker-compose is installed
if ! command -v docker-compose >/dev/null 2>&1; then
  echo "❌ docker-compose is not installed. Please install it before continuing."
  exit 1
fi

# 2. Check that the required directories and files exist
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

if [ ! -f "$SCRIPT_DIR/init-multiple-db.sh" ]; then
  echo "❌ Initialization script 'init-multiple-db.sh' not found in the folder."
  exit 1
fi

# Navigate to the Docker Compose directory
cd "$SCRIPT_DIR"

# 3. Start the PostgreSQL container in the background
echo "🚀 Starting PostgreSQL container (docker-compose up -d)..."
docker-compose up -d
if [ $? -ne 0 ]; then
  echo "❌ Failed to run 'docker-compose up -d'. Aborting."
  exit 1
fi

# 4. Wait for PostgreSQL to be ready to accept connections
echo "⏳ Waiting for PostgreSQL to be ready to accept connections..."
# Read POSTGRES_USER from .env if available, default to 'postgres'
DB_USER="postgres"
if [ -f .env ]; then
  ENV_POSTGRES_USER=$(grep -E '^POSTGRES_USER=' .env | cut -d '=' -f2 | tr -d '\"')
  if [ -n "$ENV_POSTGRES_USER" ]; then
    DB_USER="$ENV_POSTGRES_USER"
  fi
fi

# Poll pg_isready until PostgreSQL is available
MAX_RETRIES=30  # ~60 seconds max wait
RETRY_COUNT=0
until docker-compose exec -T postgres pg_isready -U "$DB_USER" -h localhost >/dev/null 2>&1; do
  RETRY_COUNT=$((RETRY_COUNT+1))
  if [ $RETRY_COUNT -ge $MAX_RETRIES ]; then
    echo "❌ PostgreSQL container was not ready after $((MAX_RETRIES*2)) seconds. Aborting."
    exit 1
  fi
  sleep 2
done
echo "✅ PostgreSQL container is ready to accept connections."

# 5. Execute the initialization script inside the container (interactive)
echo "🏗️ Running the database initialization script inside the container..."
docker-compose exec postgres /docker-entrypoint-initdb.d/init-multiple-db.sh

echo "✅ PostgreSQL setup process completed successfully."
