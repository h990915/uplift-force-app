# Uplift Force Backend

Go + Gin + GORM backend service

## Quick Start

```bash
# 1. Install dependencies
go mod download

# 2. Configure environment
cp .env.example .env

# 3. Start service
go run main.go
```

Access: http://localhost:8080

## Environment Configuration

Create `.env` file:
```env
# Database
DIRECT_URL="postgresql://postgres.xxx:password@host:5432/postgres"
DATABASE_URL="postgresql://postgres.xxx:password@host:6543/postgres?pgbouncer=true"

# Service
PORT=8080
APP_ENV=development
JWT_SECRET=your-jwt-secret

# Blockchain
SEPOLIA_RPC_URL="https://sepolia.infura.io/v3/your-project-id"
AVALANCHE_FUJI_RPC_URL="https://avalanche-fuji.infura.io/v3/your-project-id"
MAIN_CONTRACT_ADDRESS="0x9BEa73dE49536643f8749871D39bE9dFCCeC35D5"

# External APIs
RIOT_API_KEY=RGAPI-your-key
PRIVATE_KEY_1=your-private-key
```

## Development

### Hot Reload
```bash
# Install air
go install github.com/cosmtrek/air@latest

# Start with hot reload
air
```

### Testing
```bash
go test ./...
```

## Deployment

### Local Build
```bash
go build -o bin/uplift-backend main.go
./bin/uplift-backend
```

### Docker
```bash
docker build -t uplift-backend .
docker run -p 8080:8080 --env-file .env uplift-backend
```

### Production (Railway)
```bash
railway login
railway init
railway up
```

## Troubleshooting

**Port in use**
```bash
lsof -i :8080
```

**Database connection**
```bash
# Test connection
psql "your-database-url"
```

**Dependencies**
```bash
go mod tidy
```

---
**Health check**: `curl http://localhost:8080/health`
