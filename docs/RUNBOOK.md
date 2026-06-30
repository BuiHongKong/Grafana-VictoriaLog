# Operational Runbook

## Prerequisites
- Docker Desktop installed and running.
- Basic knowledge of Docker commands.

## How to Start the Stack
1. Open a terminal and navigate to the root directory of this project.
2. Build and start the containers using Docker Compose:
   ```bash
   docker compose up -d --build
   ```
3. Wait a few moments for the services to initialize. Grafana might take a minute to download and install the required plugin.

## How to Check Container Logs
You can monitor the logs of the individual services to ensure they are running correctly:

- **Check all logs**:
  ```bash
  docker compose logs -f
  ```
- **Check Golang App logs** (useful to see the generated JSON payloads):
  ```bash
  docker compose logs -f golang-app
  ```
- **Check VictoriaLogs logs**:
  ```bash
  docker compose logs -f victorialogs
  ```
- **Check Grafana logs**:
  ```bash
  docker compose logs -f grafana
  ```

## How to Stop & Clean Up
To stop the services and remove the containers and default network:
```bash
docker compose down
```

To stop the services and also remove the persistent data volume (this deletes all stored logs):
```bash
docker compose down -v
```
check