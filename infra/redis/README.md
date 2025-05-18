## Redis – Local Docker Setup (Practice Only)

This folder provides a simple Redis setup using Docker Compose for development and testing purposes.

### 🚀 Getting Started

1. **Install Docker & Docker Compose**.
2. **Copy** `.env.example` to `.env` and set a strong `REDIS_PASSWORD`.
3. **Create network** (if not exists):

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
* `.env.example`: Template for Redis password.

### ⚠️ Production Limitations

* Password is set, but no encryption (TLS).
* Port 6379 is exposed—do **not** do this in production.
* Persistence is basic (single volume).

🔐 Use strong passwords and internal networks in production. This setup is **not secure for public deployment**.

---

This setup is ideal for learning and testing. Harden it for production use with better security and persistence strategies.
