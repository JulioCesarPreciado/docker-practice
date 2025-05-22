#!/bin/bash
# ============================================================
# 📦 Script to initialize PostgreSQL databases
# 📌 Usage:
#    ./init-db.sh
#
# This script:
# - Prompts for the base name of the database
# - Prompts for credentials of two users:
#     - tester (full access to <name>_test)
#     - readonly (read-only access to <name>)
# - Creates three databases:
#     - <name>
#     - <name>_backup
#     - <name>_test
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

# Prompt for tester user credentials
read -p "👤 Enter the username for the test user [tester]: " TEST_USER
TEST_USER="${TEST_USER:-tester}"

read -s -p "🔑 Enter the password for the test user [testerpass]: " TEST_PASS
echo ""
TEST_PASS="${TEST_PASS:-testerpass}"

# Prompt for readonly user credentials
read -p "👤 Enter the username for the readonly user [readonly]: " READONLY_USER
READONLY_USER="${READONLY_USER:-readonly}"

read -s -p "🔑 Enter the password for the readonly user [readonlypass]: " READONLY_PASS
echo ""
READONLY_PASS="${READONLY_PASS:-readonlypass}"

# Ensure POSTGRES_USER is set
if [ -z "$POSTGRES_USER" ]; then
  echo "❌ Environment variable POSTGRES_USER is not set."
  exit 1
fi

# Execute SQL commands
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" --host "localhost" <<-EOSQL
    CREATE DATABASE ${DB_NAME};
    CREATE DATABASE ${DB_NAME}_backup;
    CREATE DATABASE ${DB_NAME}_test;

    CREATE USER ${TEST_USER} WITH PASSWORD '${TEST_PASS}';
    GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME}_test TO ${TEST_USER};

    CREATE USER ${READONLY_USER} WITH PASSWORD '${READONLY_PASS}';
    GRANT CONNECT ON DATABASE ${DB_NAME} TO ${READONLY_USER};
EOSQL

echo "✅ Databases and users created successfully."
