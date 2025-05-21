#!/bin/bash
# ============================================================
# 🧹 Script to completely clean up PostgreSQL environment
# Removes containers, volumes, networks, and backups
# ============================================================

set -e

# Go to the postgres infra directory
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# 1. Stop and remove containers, volumes, networks, and images
echo "🧨 Stopping and removing containers, networks, images, and volumes..."
docker-compose down --volumes --remove-orphans --rmi all

# 2. Remove named volume if defined
if docker volume ls | grep -q pgdata; then
  echo "🗑️ Removing volume 'pgdata'..."
  docker volume rm pgdata || true
fi

# 3. Remove backup files (optional, if using local folder)
BACKUP_PATH=$(grep -E '^BACKUP_PATH=' .env | cut -d '=' -f2 | tr -d '"')
BACKUP_PATH=${BACKUP_PATH:-./backups}

if [ -d "$BACKUP_PATH" ]; then
  echo "🗑️ Removing backup folder: $BACKUP_PATH..."
  rm -rf "$BACKUP_PATH"
fi

echo "✅ PostgreSQL environment has been fully cleaned up."