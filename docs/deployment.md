# Lograil Deployment Guide

## Overview
Lograil can be deployed using Docker Compose for both development and production environments. The system consists of multiple services that work together to provide a complete log management solution.

## Prerequisites

- Docker and Docker Compose
- At least 4GB RAM available
- 20GB free disk space
- Linux/macOS/Windows with Docker support

## Quick Start

### Development Environment

1. **Clone the repository**
   ```bash
   git clone https://github.com/bizjs/Lograil.git
   cd lograil
   ```

2. **Start development environment**
   ```bash
   cd docker
   docker-compose up -d
   ```

3. **Access the services**
   - Web UI: http://localhost:3000
   - Control Plane API: http://localhost:8080
   - Ingestion API: http://localhost:8081
   - VictoriaLogs: http://localhost:9428

### Production Environment

1. **Prepare production configuration**
   ```bash
   cd docker
   cp docker-compose.prod.yml docker-compose.override.yml
   # Edit docker-compose.override.yml with your production settings
   ```

2. **Configure environment variables**
   Create a `.env` file in the docker directory:
   ```env
   POSTGRES_PASSWORD=your_secure_password
   REDIS_PASSWORD=your_redis_password
   JWT_SECRET=your_jwt_secret_key
   VICTORIA_LOGS_DATA_PATH=/data/victoria-logs
   ```

3. **Start production environment**
   ```bash
   docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
   ```

## Service Configuration

### Control Plane Backend
- **Port**: 8080
- **Environment Variables**:
  - `DATABASE_URL`: PostgreSQL connection string
  - `REDIS_URL`: Redis connection string
  - `JWT_SECRET`: Secret key for JWT tokens
  - `SERVER_PORT`: Port to listen on (default: 8080)

### Ingestion Backend
- **Port**: 8081
- **Environment Variables**:
  - `VICTORIA_LOGS_URL`: VictoriaLogs endpoint
  - `REDIS_URL`: Redis connection for buffering
  - `SERVER_PORT`: Port to listen on (default: 8081)
  - `BATCH_SIZE`: Batch size for log writes (default: 100)
  - `BUFFER_SIZE`: In-memory buffer size (default: 1000)

### Web UI
- **Port**: 3000
- **Environment Variables**:
  - `VITE_API_BASE_URL`: Control Plane API base URL
  - `VITE_INGESTION_URL`: Ingestion API base URL

### VictoriaLogs
- **Port**: 9428
- **Data Path**: `/data/victoria-logs`
- **Configuration**: Custom VictoriaLogs configuration for log storage

### PostgreSQL
- **Port**: 5432
- **Database**: lograil
- **User**: lograil
- **Data Path**: `/data/postgres`

### Redis
- **Port**: 6379
- **Data Path**: `/data/redis`

## Data Persistence

### Docker Volumes
The production setup uses named Docker volumes for data persistence:
- `lograil_postgres_data`: PostgreSQL database files
- `lograil_redis_data`: Redis data files
- `lograil_victoria_logs_data`: VictoriaLogs data files

### Backup Strategy
1. **Database Backup**:
   ```bash
   docker exec lograil_postgres pg_dump -U lograil lograil > backup.sql
   ```

2. **VictoriaLogs Backup**:
   ```bash
   docker run --rm -v lograil_victoria_logs_data:/data alpine tar czf - /data > victoria-logs-backup.tar.gz
   ```

## Monitoring

### Health Checks
All services include health check endpoints:
- Control Plane: `GET /health`
- Ingestion: `GET /health`
- VictoriaLogs: `GET /health`

### Logs
View service logs:
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f control-plane
```

### Metrics
Services expose Prometheus metrics:
- Control Plane: `/metrics`
- Ingestion: `/metrics`

## Scaling

### Horizontal Scaling
For high-traffic deployments:

1. **Run multiple ingestion instances**:
   ```yaml
   services:
     ingestion:
       deploy:
         replicas: 3
   ```

2. **Use load balancer**:
   Add Nginx or HAProxy in front of multiple instances.

### Vertical Scaling
Increase resource limits in docker-compose.prod.yml:
```yaml
services:
  ingestion:
    deploy:
      resources:
        limits:
          memory: 2G
          cpus: '2.0'
```

## Security

### Network Security
- Services communicate over internal Docker network
- Only expose necessary ports (Web UI, APIs)
- Use TLS termination at load balancer level

### Data Security
- Encrypt sensitive data at rest
- Use strong passwords for databases
- Regular security updates of base images

### Access Control
- JWT-based authentication
- Role-based access control (RBAC)
- API key authentication for service-to-service communication

## Troubleshooting

### Common Issues

1. **Port conflicts**:
   - Check if ports 3000, 8080, 8081, 9428, 5432, 6379 are available
   - Use `docker-compose ps` to see port mappings

2. **Database connection issues**:
   - Ensure PostgreSQL is healthy: `docker-compose logs postgres`
   - Check connection string in environment variables

3. **VictoriaLogs not starting**:
   - Check disk space and permissions
   - Review VictoriaLogs configuration

4. **Out of memory**:
   - Increase Docker memory limit
   - Reduce batch sizes in ingestion service

### Logs and Debugging
```bash
# View all logs
docker-compose logs

# Follow logs in real-time
docker-compose logs -f

# View specific service logs
docker-compose logs control-plane

# View last 100 lines
docker-compose logs --tail=100 ingestion
```

## Maintenance

### Updates
1. Pull latest images:
   ```bash
   docker-compose pull
   ```

2. Restart services:
   ```bash
   docker-compose up -d
   ```

### Cleanup
Remove unused containers and volumes:
```bash
docker system prune -a --volumes
```

### Backup and Restore
Regular backups are essential for production deployments. Schedule automated backups using cron jobs or backup tools.
