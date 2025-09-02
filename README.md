# ITQ HR Landing Page Backend API

Go backend API built with Fiber framework for ITQ HR Landing Page application.

## Tech Stack
- **Language**: Go 1.23.2
- **Framework**: Fiber v2
- **Database**: MySQL with GORM
- **Cache**: Redis
- **Documentation**: Swagger/OpenAPI
- **Authentication**: JWT

## Prerequisites
- Go 1.23.2+
- MySQL 8.0+
- Redis 6.0+
- Docker (optional)

## Quick Start

### 1. Install Dependencies
```bash
go mod download
```

### 2. Install Swagger CLI
```bash
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

### 3. Configure Application
Edit `config.yaml`:
```yaml
App:
  port: 3000

database:
  master_host: localhost
  master_port: 3306
  master_username: your_username
  master_password: your_password
  master_dbname: your_database

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  access_token: your_secret
  refresh_token: your_secret
```

### 4. Run Application
```bash
# Generate Swagger docs
swag init -g ./bootstrap/swagger.go

# Run
go run main.go
```

## Docker

### Build & Run
```bash
docker build -t ithq-kkl-v2 .
docker run -p 3000:3000 ithq-kkl-v2
```

### Docker Compose
```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "3000:3000"
    depends_on:
      - mysql
      - redis

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: ithq_hr
    ports:
      - "3306:3306"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
```

Run: `docker-compose up -d`

## Project Structure
```
ithq-kkl-v2/
├── api/              # Controllers, middleware, routes
├── bootstrap/        # App initialization
├── core/            # Utilities and logging
├── data/            # Repositories and services
├── domain/          # Entities and models
├── config.yaml      # Configuration
├── dockerfile       # Docker config
└── main.go         # Entry point
```

## API Documentation
- **Swagger UI**: `http://localhost:3000/swagger/index.html`
- **Base Path**: `/api/v1`
- **Auth**: Bearer JWT Token

## Development

### Git Workflow
```bash
chmod +x git_workflow.sh
./git_workflow.sh "commit message"
```

### Testing
```bash
go test ./...
go test -cover ./...
```

## Common Issues

**Swagger not found:**
```bash
go install github.com/swaggo/swag/cmd/swag@latest
source ~/.zshrc
```

**Port in use:**
```bash
lsof -ti:3000 | xargs kill -9
```

## License
Apache 2.0
