# 🐘 PostgreSQL Docker Setup – Quick Start Guide

## 🔧 Requirements
- Docker & Docker Compose installed
- Clone this repo and navigate to the project directory

## ⚙️ Environment Setup
Copy the `.env.example` file to `.env` and configure:

```bash
cp .env.example .env
```

### Required Variables

| Variable           | Description                                     | Example               |
|--------------------|-------------------------------------------------|-----------------------|
| `APP_NAME`         | Prefix for container names                      | `myapp`               |
| `POSTGRES_USER`    | PostgreSQL superuser username                   | `postgres`            |
| `POSTGRES_PASSWORD`| Password for `POSTGRES_USER`                   | `secret123`           |
| `POSTGRES_DB`      | Default database name (used for backup)         | `mydatabase`          |
| `NETWORK_NAME`     | Docker network name (used internally)           | `app_net`             |
| `BACKUP_INTERVAL`  | Interval in seconds between automatic backups   | `3600` (1 hour)       |
| `BACKUP_PATH`      | Host path to store backup files (`.sql.gz`)     | `./backups`           |

💡 For production environments, never commit the `.env` file with real values.

---

## 🚀 Start the Setup

⚠️ Before running any `.sh` scripts, make sure they have execution permissions:

```bash
chmod +x setup-postgres.sh restore-backup.sh delete.sh
```
Run this command to build and start the services and initialize the databases:

```bash
./setup-postgres.sh
```

This script:
- Starts the `postgres` and `backup` containers
- Waits until PostgreSQL is ready
- Runs `init-multiple-db.sh` to create:
  - `<db>`, `<db>_backup`, `<db>_test`
  - Users: `tester` (with full access to test DB) and `readonly` (read-only access to main DB)
- Prompts you interactively for usernames and passwords

---

## 💾 Automatic Backups
The `backup` container automatically:
- Dumps `${POSTGRES_DB}` to `/backups/backup_YYYY-MM-DD_HH-MM-SS.sql.gz`
- Keeps only the latest 10 backup files
- Restores the latest dump to `${POSTGRES_DB}_backup` as a mirror

---

## ♻️ Restore a Backup
To manually restore a backup to a specific database:

```bash
DB_NAME=mydatabase DB_USER=postgres DB_PASSWORD=secret \
./restore-backup.sh backups/backup_2025-05-10_19-00-00.sql.gz
```

This will:
- Terminate all active connections to the database
- Drop and recreate the database
- Restore from the `.sql.gz` file

You must have `psql`, `createdb`, and `dropdb` available in your shell.

---

## 🧹 Clean Everything (optional)
To fully delete the PostgreSQL environment:

```bash
./delete.sh            # Removes containers, volumes, network (preserves backups)
./delete.sh --delete-backups   # Also removes local backup files
```

---

## 🔐 Production Tips
- Change default passwords
- Never expose port 5432 publicly
- Do not commit real `.env` credentials
- Offload backups to a secure location

---

## ✅ Ready to Use
This setup is ideal for development and local testing. For production, harden security, rotate credentials, and manage storage externally.
