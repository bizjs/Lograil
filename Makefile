.PHONY: build build-control-plane build-ingestion build-web-ui clean test docker-up docker-down

# Build all components
build: build-control-plane build-ingestion build-web-ui

# Build Control Plane
build-control-plane:
	@echo "Building Control Plane..."
	@cd control-plane && go mod tidy && go build -o ../bin/control-plane ./cmd/server

# Build Ingestion Backend
build-ingestion:
	@echo "Building Ingestion Backend..."
	@cd ingestion && go mod tidy && go build -o ../bin/ingestion ./cmd/server

# Build Web UI
build-web-ui:
	@echo "Building Web UI..."
	@cd web-ui && pnpm install && pnpm build

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -rf web-ui/dist/
	@rm -rf web-ui/node_modules/

# Run tests
test:
	@echo "Running tests..."
	@cd control-plane && go test ./...
	@cd ingestion && go test ./...
	@cd web-ui && pnpm test

# Start Docker services
docker-up:
	@echo "Starting Docker services..."
	@cd docker && docker-compose up -d

# Stop Docker services
docker-down:
	@echo "Stopping Docker services..."
	@cd docker && docker-compose down

# View logs
logs:
	@cd docker && docker-compose logs -f

# Development setup
dev-setup: docker-up
	@echo "Development environment is ready!"
	@echo "Control Plane: http://localhost:8080"
	@echo "Ingestion: http://localhost:8081"
	@echo "Web UI: http://localhost:3000"
	@echo "VictoriaLogs: http://localhost:9428"

# Quick start for development
dev: build docker-up
	@echo "Lograil development environment started!"
	@echo "Run 'make logs' to view service logs"
