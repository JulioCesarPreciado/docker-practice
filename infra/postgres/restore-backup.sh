#!/bin/bash
# ============================================================
# 🔄 Script to restore a PostgreSQL database from a .sql.gz file
# 📌 Usage:
#    ./restore-backup.sh <backup_file.sql.gz>
#    Example:
#    ./restore-backup.sh backup_2025-05-10_19-00-00.sql.gz
#
# This script:
# - Terminates all active connections to the target database
# - Drops and recreates the specified database
# - Restores the database from a compressed SQL backup file (.sql.gz)
#
# Optional environment variables:
# - DB_NAME: name of the database to restore (default: "mydatabase")
# - DB_HOST: database host (default: "postgres")
# - DB_USER: database user (default: "postgres")
# - DB_PASSWORD: database password (default: "postgres")
# ============================================================

set -e

# Configurable parameters with defaults
BACKUP_FILE="$1"
DB_NAME="${DB_NAME:-mydatabase}"
DB_HOST="${DB_HOST:-postgres}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"

if [ -z "$BACKUP_FILE" ]; then
    echo "❌ You must provide the name of the backup file as an argument."
    echo "🔧 Usage: ./restore-backup.sh backup_2025-05-10_19-00-00.sql.gz"
    exit 1
fi

echo "⚠️ This will overwrite the database '$DB_NAME'."
read -p "Do you want to continue? (y/N): " CONFIRM
[[ "$CONFIRM" != "y" ]] && echo "❌ Operation canceled." && exit 1

echo "🔌 Terminating active connections..."
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d postgres -c "
  SELECT pg_terminate_backend(pid)
  FROM pg_stat_activity
  WHERE datname = '$DB_NAME' AND pid <> pg_backend_pid();
"

echo "🧹 Dropping database..."
PGPASSWORD=$DB_PASSWORD dropdb -h "$DB_HOST" -U "$DB_USER" "$DB_NAME"

echo "🛠️ Creating new empty database..."
PGPASSWORD=$DB_PASSWORD createdb -h "$DB_HOST" -U "$DB_USER" "$DB_NAME"

echo "📦 Restoring from $BACKUP_FILE..."
gunzip -c "$BACKUP_FILE" | PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME"

echo "✅ Restore completed successfully."
