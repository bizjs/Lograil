# Lograil Architecture Design

## Overview
Lograil is a lightweight log management platform built on VictoriaLogs, designed for resource-constrained environments and small to mid-sized teams. The system consists of three main components: Control Plane Backend, Ingestion Backend, and Web UI.

## System Components

### 1. Control Plane Backend
**Purpose**: Manages system configuration, user authentication, authorization, and metadata operations.

**Responsibilities**:
- User management and authentication
- Project/tenant management
- Configuration management (retention policies, indexing rules)
- API key management
- Monitoring and health checks
- Metadata queries (log sources, schemas)

**Technology Stack**:
- Language: Go (for performance and resource efficiency)
- Framework: Gin or Echo for HTTP API
- Database: PostgreSQL for metadata storage
- Authentication: JWT with refresh tokens
- Authorization: RBAC (Role-Based Access Control)

### 2. Ingestion Backend
**Purpose**: Receives, processes, and stores log data from various sources.

**Responsibilities**:
- Log ingestion via multiple protocols (HTTP, TCP, UDP, gRPC)
- Log parsing and transformation
- Data validation and filtering
- Batch processing and buffering
- Integration with VictoriaLogs for storage
- Rate limiting and backpressure handling

**Technology Stack**:
- Language: Go (high throughput for ingestion)
- Framework: Custom HTTP server with goroutines
- Protocols: HTTP/JSON, Syslog, Fluentd, OpenTelemetry
- Queue: In-memory buffer with optional Redis for high volume
- Storage: VictoriaLogs client integration

### 3. Web UI
**Purpose**: Provides a user-friendly interface for log exploration, management, and monitoring.

**Responsibilities**:
- Log search and filtering interface
- Real-time log streaming
- Dashboard and visualization
- User management interface
- Configuration management UI
- Alert and notification management

**Technology Stack**:
- Frontend: React with TypeScript
- UI Framework: Ant Design or Material-UI
- State Management: Redux Toolkit or Zustand
- API Client: Axios or React Query
- Charts: Recharts or D3.js
- Build Tool: Vite

### 4. Storage Layer - VictoriaLogs
**Purpose**: High-performance log storage and querying.

**Features**:
- Columnar storage optimized for logs
- Efficient compression
- Fast full-text search
- Time-series indexing
- Distributed architecture support

## Data Flow

```
Log Sources → Ingestion Backend → VictoriaLogs
                    ↓
Web UI ← Control Plane Backend ← VictoriaLogs
```

### Detailed Flow:
1. **Ingestion**:
   - Applications send logs to Ingestion Backend via HTTP/gRPC
   - Ingestion Backend validates, parses, and enriches logs
   - Logs are batched and written to VictoriaLogs

2. **Query**:
   - Web UI sends query requests to Control Plane Backend
   - Control Plane validates permissions and forwards to VictoriaLogs
   - Results are streamed back through Control Plane to Web UI

3. **Management**:
   - Administrative operations go through Control Plane Backend
   - Metadata stored in PostgreSQL, logs in VictoriaLogs

## API Design

### Control Plane API
```
POST   /api/v1/auth/login
GET    /api/v1/projects
POST   /api/v1/projects
GET    /api/v1/projects/{id}/logs?query={query}&start={time}&end={time}
PUT    /api/v1/config/retention
GET    /api/v1/users
POST   /api/v1/users
```

### Ingestion API
```
POST   /ingest/logs
POST   /ingest/batch
GET    /health
```

## Security Considerations

- **Authentication**: JWT tokens with short expiration
- **Authorization**: Project-based access control
- **Encryption**: TLS for all external communications
- **API Keys**: For service-to-service authentication
- **Audit Logs**: All administrative actions logged

## Scalability & Performance

- **Horizontal Scaling**: All components designed for stateless operation
- **Load Balancing**: Nginx or Kubernetes ingress
- **Caching**: Redis for session and query result caching
- **Rate Limiting**: Token bucket algorithm at ingress points
- **Monitoring**: Prometheus metrics, Grafana dashboards

## Deployment Architecture

### Component Diagram
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Web Browser   │    │  Log Sources    │    │   Monitoring    │
│                 │    │  (Apps, Sys)    │    │   (Prometheus)  │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                     │                      │
          ▼                     ▼                      ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│     Web UI      │    │ Ingestion       │    │ Control Plane   │
│   (React)       │    │ Backend (Go)    │    │ Backend (Go)     │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                     │                      │
          └─────────────────────┼──────────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │  VictoriaLogs   │
                       │   (Storage)     │
                       └─────────┬───────┘
                                 │
                       ┌─────────┴─────────┐
                       │                   │
              ┌────────▼────────┐ ┌────────▼────────┐
              │ PostgreSQL      │ │     Redis       │
              │ (Metadata)      │ │   (Cache)       │
              └─────────────────┘ └─────────────────┘
```

### Development Environment
- Docker Compose with local VictoriaLogs instance
- Hot reload for development
- Local PostgreSQL and Redis

### Production Environment
- Docker Compose for container orchestration
- VictoriaLogs cluster for high availability
- External PostgreSQL and Redis
- Nginx reverse proxy for load balancing
- Docker volumes for data persistence

## Monitoring & Observability

- **Metrics**: Request latency, throughput, error rates
- **Logs**: Structured logging with correlation IDs
- **Tracing**: OpenTelemetry integration
- **Alerts**: Prometheus Alertmanager for critical issues

## Future Extensions

- **Alerting Engine**: Rule-based log alerting
- **Log Analytics**: Advanced analytics and ML insights
- **Multi-tenancy**: Enhanced isolation and resource management
- **Integration APIs**: Webhooks, Slack notifications, etc.
