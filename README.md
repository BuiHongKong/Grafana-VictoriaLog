# Centralized Logging Simulation Environment

## Overview
This project provides a centralized logging simulation environment using Docker Compose. It is designed to demonstrate a decoupled log ingestion, storage, and visualization architecture using Fluent Bit, VictoriaLogs, and Grafana.

For a visual explanation of this architecture, please open the: [Architecture Documentation](file:///d:/PERSONAL/Grafana-VictoriaLog/docs/grafana-victoriaLog-fluentbit.html).

## Tech Stack
- **Infrastructure**: Docker Compose
- **Log Aggregator (Gateway)**: Fluent Bit
- **Log Storage**: VictoriaLogs
- **Visualization**: Grafana OSS with `victoriametrics-logs-datasource` plugin
- **Mock External App**: Golang Log Generator

## System Architecture & Data Flow
This environment models a production-ready Endpoint (Gateway) approach suitable for internal servers:

1. **Golang App (Mock App)**: Simulates an external or internal service. It generates mock logs and sends them over the network via HTTP POST to the Fluent Bit endpoint.
2. **Fluent Bit (Endpoint)**: Acts as a centralized listener on port `8888`. It receives HTTP JSON logs, processes them, and forwards them to the storage layer.
3. **VictoriaLogs (Storage)**: Receives the processed logs from Fluent Bit and stores them efficiently on disk.
4. **Grafana (Visualization)**: Queries VictoriaLogs using LogsQL to display real-time logs on pre-provisioned Dashboards.

## Usage
Start the complete stack:
```bash
docker compose down
docker compose up -d --build
```
Access Grafana at `http://localhost:3000`. The mock logs will automatically start flowing into the "Golang Application Logs" dashboard.
