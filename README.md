# 🐳 Docker Practice Project

This repository contains a practice environment using **Docker Compose** to orchestrate four key services:

1. **PostgreSQL** – Relational database
2. **Redis** – In-memory key-value store
3. **API (Go)** – RESTful backend service
4. **Socket Service (Go)** – WebSocket or real-time communication layer

Each service is containerized and located inside the `infra/` directory.

## 🔧 Service Structure

```
infra/
├── postgres   # PostgreSQL with backups and DB init script
├── redis      # Redis with password protection
├── api        # REST API built in Go
└── socket     # Socket service built in Go
```

## 🚀 Startup Instructions

To run the full environment, each service must be started in the following order:

```bash
cd infra/postgres
docker-compose up -d

cd ../redis
docker-compose up -d

cd ../api
docker-compose up -d

cd ../socket
docker-compose up -d
```

This ensures that the database and Redis services are ready before the API and Socket containers start.

## 🌐 Networking

All services connect via a shared Docker network called `app_net`. Make sure it's created before starting:

```bash
docker network create app_net  # only needed the first time
```

## 📂 Environment Configuration

Each service uses its own `.env` file. Make sure to configure:

* PostgreSQL credentials
* Redis password
* API and Socket port numbers
* JWT secrets, DB/Redis hostnames, etc.

## ⚠️ Notes

* This project is intended for **learning and development purposes**.
* Security hardening and production optimizations are not applied by default.
* Logs and build context are kept simple to ease debugging and testing.

---

Feel free to expand or modify this template as your stack grows. 🚀
