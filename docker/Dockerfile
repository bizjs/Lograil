# Multi-stage build for Go applications
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata sqlite-dev gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files from root
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build stage - accepts build target as argument
ARG SERVICE_PATH
RUN if [ "$SERVICE_PATH" = "control-plane" ]; then \
        CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app ./control-plane/cmd/server; \
    elif [ "$SERVICE_PATH" = "ingestion" ]; then \
        CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app ./ingestion/cmd/server; \
    else \
        echo "Unknown service: $SERVICE_PATH" && exit 1; \
    fi

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata sqlite

# Create app directory and data directory for SQLite
WORKDIR /app
RUN mkdir -p /app/data

# Copy binary from builder stage
COPY --from=builder /app/app .

# Copy any migration files if they exist
COPY --from=builder /app/control-plane/migrations ./migrations 2>/dev/null || true
COPY --from=builder /app/ingestion/migrations ./migrations 2>/dev/null || true

# Expose port (will be overridden in docker-compose)
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./app"]
