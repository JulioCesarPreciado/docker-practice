## PostgreSQL Docker Setup – Usage Guide

### Requirements

* Docker & Docker Compose installed.
* Clone this repository and navigate to the `postgresql` folder.

### Environment Variables

Copy `.env.example` to `.env` and configure:

* `APP_NAME`: Name prefix for containers.
* `POSTGRES_USER`, `POSTGRES_PASSWORD`: Database superuser credentials.
* `BACKUP_INTERVAL`: Backup interval in seconds (default: 3600).
* `BACKUP_PATH`: Path for storing `.sql.gz` backup files.

### Starting the Containers

Run:

```bash
docker-compose up -d
```

This will start two containers:

* **postgres**: PostgreSQL server with persistent volume.
* **backup**: Periodic backup service that keeps the last 10 dumps.

Use `docker-compose ps` to verify they are running. Data is persisted in the `pgdata` Docker volume.

### Initializing Databases

Run the following script to create multiple databases and roles:

```bash
docker-compose exec postgres /docker-entrypoint-initdb.d/init-multiple-db.sh mydatabase
```

This creates:

* `mydatabase`, `mydatabase_backup`, `mydatabase_test`
* `tester` (full access to test DB)
* `readonly` (read-only access to main DB)

Passwords are hardcoded (`testerpass`, `readonlypass`) and should be changed for production.

### Restoring from Backup

Use the `restore-backup.sh` script:

```bash
DB_NAME=mydatabase DB_USER=postgres DB_PASSWORD=secret ./restore-backup.sh backups/backup_2025-05-10_19-00-00.sql.gz
```

The script:

* Terminates active connections.
* Drops and recreates the DB.
* Restores from the given `.sql.gz` backup.

Ensure you have `psql`, `createdb`, `dropdb` CLI tools installed.

### Production Notes

* Never commit `.env` with real credentials.
* Replace default user passwords.
* Adjust network exposure: do **not** expose port 5432 publicly.
* Mount volumes to safe locations.
* Backup files should be moved off-server for long-term storage.

This setup is ideal for local development and testing. For production, harden credentials, secure access, and adapt persistence and backup policies accordingly.
