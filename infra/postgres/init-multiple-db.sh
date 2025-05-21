#!/bin/bash
# ============================================================
# 📦 Script to initialize PostgreSQL databases
# 📌 Usage:
#    ./init-db.sh
#
# This script:
# - Prompts for the base name of the database and test user credentials
# - Creates three databases:
#     - <name>
#     - <name>_backup
#     - <name>_test
# - Creates two users:
#     - tester (with full access to <name>_test)
#     - readonly (with read-only access to <name>)
#
# Requires the POSTGRES_USER environment variable to be set,
# for example via Docker Compose.
# ============================================================

set -e

# Prompt for base database name
read -p "📛 Enter the base name for your database: " DB_NAME
if [ -z "$DB_NAME" ]; then
  echo "❌ You must provide a database name."
  exit 1
fi

# Prompt for tester username and password with defaults
read -p "👤 Enter the username for the test user [tester]: " TEST_USER
TEST_USER="${TEST_USER:-tester}"

read -s -p "🔑 Enter the password for the test user [testerpass]: " TEST_PASS
echo ""
TEST_PASS="${TEST_PASS:-testerpass}"

# Ensure POSTGRES_USER is set
if [ -z "$POSTGRES_USER" ]; then
  echo "❌ Environment variable POSTGRES_USER is not set."
  exit 1
fi

# Run SQL commands
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE ${DB_NAME};
    CREATE DATABASE ${DB_NAME}_backup;
    CREATE DATABASE ${DB_NAME}_test;

    CREATE USER ${TEST_USER} WITH PASSWORD '${TEST_PASS}';
    GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME}_test TO ${TEST_USER};

    CREATE USER readonly WITH PASSWORD 'readonlypass';
    GRANT CONNECT ON DATABASE ${DB_NAME} TO readonly;
EOSQL

echo "✅ Databases and users created successfully."
