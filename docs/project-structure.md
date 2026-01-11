# Lograil Project Structure

## Overall Project Layout
```
lograil/
├── docs/                          # Documentation
│   ├── architecture.md           # System architecture design
│   ├── api-spec.md               # API specifications
│   └── deployment.md             # Deployment guides
├── control-plane/                # Control Plane Backend
│   ├── cmd/
│   │   └── server/               # Main application entry
│   ├── internal/
│   │   ├── api/                  # HTTP API handlers
│   │   ├── auth/                 # Authentication & authorization
│   │   ├── config/               # Configuration management
│   │   ├── database/             # Database operations
│   │   ├── models/               # Data models
│   │   └── service/              # Business logic
│   ├── pkg/                      # Shared packages
│   ├── docker/                   # Docker files
│   ├── migrations/               # Database migrations
│   └── go.mod
├── ingestion/                    # Ingestion Backend
│   ├── cmd/
│   │   └── server/
│   ├── internal/
│   │   ├── api/                  # Ingestion endpoints
│   │   ├── parser/               # Log parsing logic
│   │   ├── processor/            # Log processing pipeline
│   │   ├── storage/              # VictoriaLogs integration
│   │   └── buffer/               # Buffering and queuing
│   ├── pkg/
│   ├── docker/
│   └── go.mod
├── web-ui/                       # Web UI Frontend
│   ├── public/                   # Static assets
│   ├── src/
│   │   ├── components/           # React components
│   │   ├── pages/                # Page components
│   │   ├── hooks/                # Custom React hooks
│   │   ├── services/             # API client services
│   │   ├── store/                # State management
│   │   ├── utils/                # Utility functions
│   │   └── types/                # TypeScript type definitions
│   ├── docker/
│   ├── package.json
│   └── vite.config.ts
├── docker/                       # Docker Compose and configs
│   ├── docker-compose.yml        # Development environment
│   ├── docker-compose.prod.yml   # Production environment
│   └── victoria-logs/            # VictoriaLogs config
├── scripts/                      # Build and deployment scripts
│   ├── build.sh
│   ├── deploy.sh
│   └── setup-dev.sh
├── .github/                      # GitHub Actions
│   └── workflows/
├── LICENSE
├── README.md
└── Makefile                      # Build automation
```

## Directory Explanations

### Control Plane Backend (`control-plane/`)
- **cmd/server/**: Application entry point with main function
- **internal/api/**: REST API handlers for user management, projects, configuration
- **internal/auth/**: JWT authentication, RBAC authorization logic
- **internal/config/**: Configuration loading and validation
- **internal/database/**: PostgreSQL connection and query operations
- **internal/models/**: Go structs for users, projects, configurations
- **internal/service/**: Business logic layer coordinating between API and database
- **pkg/**: Reusable packages that could be shared across services
- **migrations/**: Database schema migration files

### Ingestion Backend (`ingestion/`)
- **cmd/server/**: High-performance server entry point
- **internal/api/**: Log ingestion endpoints (HTTP, gRPC)
- **internal/parser/**: Log format parsers (JSON, syslog, etc.)
- **internal/processor/**: Log enrichment, filtering, validation
- **internal/storage/**: VictoriaLogs client and write operations
- **internal/buffer/**: In-memory buffering and batching logic
- **pkg/**: Shared ingestion utilities

### Web UI (`web-ui/`)
- **src/components/**: Reusable UI components (LogViewer, SearchBar, etc.)
- **src/pages/**: Route-based page components
- **src/hooks/**: Custom hooks for data fetching, real-time updates
- **src/services/**: API client functions for Control Plane communication
- **src/store/**: Redux/Zustand state management
- **src/utils/**: Helper functions for formatting, validation
- **src/types/**: TypeScript interfaces and types

### Infrastructure (`docker/`, `scripts/`)
- **docker/**: Container definitions and orchestration with Docker Compose
- **scripts/**: Automation scripts for CI/CD pipelines

## Key Design Principles

### Separation of Concerns
- Each backend service has clear boundaries and responsibilities
- Frontend is completely decoupled from backend services
- Infrastructure concerns separated from application code

### Go Project Structure
- Follows standard Go project layout conventions
- Uses `internal/` for private application code
- Uses `pkg/` for potentially shareable code
- Clear separation between command, business logic, and data access

### Frontend Organization
- Component-based architecture with clear separation
- Custom hooks for reusable logic
- Centralized state management
- Type-safe API interactions

### Infrastructure as Code
- Docker for containerization
- Docker Compose for orchestration
- Scripts for automation

This structure provides:
- **Scalability**: Services can be developed and deployed independently
- **Maintainability**: Clear organization makes code easy to find and modify
- **Testability**: Each layer can be unit tested in isolation
- **DevOps**: Easy containerization and deployment automation
