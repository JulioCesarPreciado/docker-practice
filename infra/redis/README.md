## Redis – Local Docker Setup (Practice Only)

This folder provides a simple Redis setup using Docker Compose for development and testing purposes.

### 🚀 Getting Started

1. **Install Docker & Docker Compose**.
2. **Copy** `.env.example` to `.env` and configure the following:

   ```env
   # Application configuration
   APP_NAME=autocotiza

   # Redis configuration
   REDIS_PASSWORD=password

   # Network configuration
   NETWORK_NAME=app_net
   ```

3. **Create network** (if it doesn't exist):

   ```bash
   docker network create app_net
   ```

4. **Start Redis**:

   ```bash
   docker-compose up -d
   ```

   Redis will be available on `localhost:6379`.

### 📦 What's Included

* `docker-compose.yml`: Runs Redis v8, sets password via `REDIS_PASSWORD`, maps port 6379, and uses volume `redis_data` for persistence.
* `.env.example`: Template for environment configuration.
* `flush-redis.sh`: Script to flush all Redis data securely, useful for debugging or resetting the database.

   To use:
   ```bash
   chmod +x flush-redis.sh
   ./flush-redis.sh
   ```

### ⚠️ Production Limitations

* Password is set, but there is **no encryption (no TLS)**.
* Port 6379 is exposed—do **not** expose Redis publicly.
* Persistence is basic (a single local volume).

🔐 Use strong passwords, internal networks, and encrypted connections in production environments. This setup is **only intended for local development and learning**.

---

This setup is ideal for practicing Docker, Redis, and basic infrastructure automation. For production, extend it with secure networking, backup strategies, monitoring, and failover.
