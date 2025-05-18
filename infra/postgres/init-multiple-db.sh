#!/bin/bash
# ============================================================
# 📦 Script to initialize PostgreSQL databases
# 📌 Usage:
#    ./init-db.sh <database_name>
#    Example:
#    ./init-db.sh example
#
# This script:
# - Creates three databases: <name>, <name>_backup, and <name>_test
# - Creates two users:
#     - tester (with full access to <name>_test)
#     - readonly (with read-only access to <name>)
# Requires the environment variable POSTGRES_USER (and optionally POSTGRES_PASSWORD) to be set.
# ============================================================

set -e

DB_NAME="$1"

if [ -z "$DB_NAME" ]; then
  echo "❌ You must provide a database name."
  echo "🔧 Usage: ./init-db.sh <database_name>"
  exit 1
fi

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE ${DB_NAME};
    CREATE DATABASE ${DB_NAME}_backup;
    CREATE DATABASE ${DB_NAME}_test;

    CREATE USER tester WITH PASSWORD 'testerpass';
    GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME}_test TO tester;

    CREATE USER readonly WITH PASSWORD 'readonlypass';
    GRANT CONNECT ON DATABASE ${DB_NAME} TO readonly;
EOSQL
