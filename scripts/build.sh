#!/bin/bash

set -e

echo "Building Lograil..."

# Build Control Plane
echo "Building Control Plane..."
cd control-plane
go mod tidy
go build -o ../bin/control-plane ./cmd/server
cd ..

# Build Ingestion Backend
echo "Building Ingestion Backend..."
cd ingestion
go mod tidy
go build -o ../bin/ingestion ./cmd/server
cd ..

# Build Web UI
echo "Building Web UI..."
cd web-ui
pnpm install
pnpm build
cd ..

echo "Build completed successfully!"
echo "Binaries are available in the bin/ directory"
echo "Web UI build is available in web-ui/dist/"
