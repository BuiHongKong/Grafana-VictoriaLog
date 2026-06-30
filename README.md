# Centralized Logging Simulation Environment

## Overview
This project provides a centralized logging simulation environment using Docker Compose. It is designed to demonstrate log ingestion, storage, and visualization using VictoriaLogs and Grafana, with a custom Golang backend acting as the log generator.

## Tech Stack
- **Infrastructure**: Docker Compose
- **Log Storage**: VictoriaLogs (`victoriametrics/victoria-logs:v1.3.0`)
- **Visualization**: Grafana OSS (latest) with `victoriametrics-logs-datasource` plugin
- **Log Generator App**: Golang (1.21+)

## System Architecture & Data Flow
1. **Golang App (Log Generator)**: Runs continuously, generating mock application logs (INFO, WARN, ERROR) containing simulated transaction data.
2. **Log Ingestion**: The Golang App sends these logs directly to VictoriaLogs via HTTP POST to the `/insert/jsonline` endpoint.
3. **Storage**: VictoriaLogs receives the JSON line payloads, parses the `_stream` field for labels (e.g., job, env), and stores the data.
4. **Visualization**: Grafana connects to VictoriaLogs to query and visualize the logs using LogsQL.
