# Lograil
A lightweight, project-oriented log management platform built on VictoriaLogs, designed for resource-constrained environments, modern backend services, and small to mid-sized teams.

## Architecture

Lograil consists of three main components:

- **Control Plane Backend**: Manages authentication, authorization, user management, and system configuration
- **Ingestion Backend**: High-performance log ingestion service supporting multiple protocols
- **Web UI**: Modern React-based interface for log exploration and management

For detailed architecture information, see:
- [Architecture Overview](docs/architecture.md)
- [Project Structure](docs/project-structure.md)
- [Deployment Guide](docs/deployment.md)

## Quick Start

### Prerequisites

- Go 1.21+
- Node.js 24+
- pnpm
- Docker and Docker Compose
- SQLite3 (automatically handled in containers)

### Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/bizjs/Lograil.git
   cd lograil
   ```

2. **Build all components**
   ```bash
   make build
   ```

3. **Start development environment**
   ```bash
   make dev-setup
   ```

4. **Access the services**
   - Web UI: http://localhost:3000
   - Control Plane API: http://localhost:9012
   - Ingestion API: http://localhost:9011
   - VictoriaLogs: http://localhost:9428

### Manual Development

If you prefer to run services individually:

```bash
# Terminal 1: Start databases (Redis and VictoriaLogs)
cd docker && docker-compose up redis victorialogs -d

# Terminal 2: Start Control Plane (SQLite database will be created automatically)
cd control-plane && go run cmd/server/main.go

# Terminal 3: Start Ingestion Backend
cd ingestion && go run cmd/server/main.go

# Terminal 4: Start Web UI
cd web-ui && pnpm dev
```

## API Usage

### Ingest Logs

```bash
# Single log entry
curl -X POST http://localhost:9011/ingest/logs \
  -H "Content-Type: application/json" \
  -d '{
    "level": "info",
    "message": "User logged in",
    "source": "auth-service",
    "fields": {"user_id": 123}
  }'

# Batch log entries
curl -X POST http://localhost:9011/ingest/batch \
  -H "Content-Type: application/json" \
  -d '{
    "logs": [
      {
        "level": "info",
        "message": "Request processed",
        "source": "api-gateway"
      },
      {
        "level": "error",
        "message": "Database connection failed",
        "source": "user-service"
      }
    ]
  }'
```

### Query Logs

```bash
# Get project logs
curl "http://localhost:9012/api/v1/projects/1/logs?query=error&start=2024-01-01T00:00:00Z"
```

## Development

### Project Structure

```
lograil/
├── control-plane/     # Go backend for control plane
├── ingestion/         # Go backend for log ingestion
├── web-ui/           # React frontend
├── docker/           # Docker configurations
├── docs/             # Documentation
└── scripts/          # Build and utility scripts
```

### Available Commands

```bash
# Build all components
make build

# Build individual components
make build-control-plane
make build-ingestion
make build-web-ui

# Start/stop Docker services
make docker-up
make docker-down

# View logs
make logs

# Clean build artifacts
make clean
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
