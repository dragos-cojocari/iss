# BORK - Quick Start Guide

This guide will help you get the BORK application running in under 5 minutes.

## Prerequisites Check

Before starting, ensure you have:

```bash
# Check Docker
docker --version
docker-compose --version

# Check Java (for local development)
java --version  # Should be 25

# Check Go (for local development)
go version  # Should be 1.26.3

# Check Maven (for local development)
mvn --version  # Should be 3.9+
```

## Option 1: Full Docker Stack (Recommended for First Run)

This runs everything in Docker containers.

```bash
# 1. Clone and navigate to project
cd /Users/dragos.cojocari/ubb/iss

# 2. Set up environment
make setup

# 3. Build all Docker images (this may take a few minutes)
make docker-build

# 4. Start all services
make docker-up

# 5. Wait for services to be ready (~30 seconds)
# Check status
make status

# 6. Verify backend is running
curl http://localhost:8080/api/health

# Expected output:
# {
#   "status": "UP",
#   "application": "BORK Backend",
#   "version": "0.1.0-SNAPSHOT",
#   "database": "Connected",
#   ...
# }

# 7. Access the frontend TUI
docker attach bork-frontend

# You should see the BORK login screen!
# Use Esc or Ctrl+C to exit
```

### Stopping Services

```bash
# Stop all services
make docker-down

# Stop and remove all data
make docker-down-volumes
```

## Option 2: Local Development Mode (Recommended for Development)

This runs the database in Docker, but backend and frontend locally for faster development.

### Terminal 1: Database

```bash
# Start PostgreSQL
make dev-db

# Verify it's running
docker ps | grep bork-db
```

### Terminal 2: Backend

```bash
# Run Spring Boot backend
make dev-backend

# Wait for "Started BorkApplication" message
# Backend will be available at http://localhost:8080
```

### Terminal 3: Frontend

```bash
# Install Go dependencies (first time only)
make dev-frontend-deps

# Run the TUI
make dev-frontend

# You should see the BORK login screen!
```

### Stopping Services

```bash
# Terminal 2 & 3: Ctrl+C to stop backend and frontend

# Terminal 1: Stop database
make dev-db-stop
```

## Verification Steps

### 1. Check Database

```bash
# Connect to database
make db-connect

# Inside psql, run:
\l              # List databases (should see bork_db)
\q              # Quit
```

### 2. Check Backend

```bash
# Health check
curl http://localhost:8080/api/health | jq

# Or use make target
make backend-health
```

### 3. Check Frontend

The frontend TUI should display:

- BORK ASCII logo
- Username and password fields
- Tab navigation working
- Esc to exit

## Common Issues & Solutions

### Issue: Port 5432 already in use

```bash
# Check what's using the port
lsof -i :5432

# Stop existing PostgreSQL
brew services stop postgresql  # macOS
sudo systemctl stop postgresql  # Linux

# Or change port in .env file
DB_PORT=5433
```

### Issue: Port 8080 already in use

```bash
# Check what's using the port
lsof -i :8080

# Kill the process or change port in .env
BACKEND_PORT=8081
```

### Issue: Docker build fails

```bash
# Clean Docker cache
docker system prune -a

# Rebuild
make docker-build
```

### Issue: Go dependencies not found

```bash
# Re-download dependencies
cd frontend
go mod download
go mod tidy
```

### Issue: Maven build fails

```bash
# Clean and rebuild
cd backend
mvn clean install -U
```

## Next Steps

Now that you have BORK running:

1. **Explore the Code**

   - Backend: `backend/src/main/java/com/bork/`
   - Frontend: `frontend/internal/ui/`
   - Database: `database/init/`

2. **Read the Documentation**

   - [Full README](./README.md)
   - [Project Specification](./docs/specification.md)
   - [Use Cases](./docs/usecase_diagram.md)
   - [Class Diagram](./docs/class_diagram.md)

3. **Start Development**

   - Add new API endpoints in `backend/src/main/java/com/bork/controller/`
   - Add new TUI screens in `frontend/internal/ui/`
   - Create database entities in `backend/src/main/java/com/bork/model/`

4. **Run Tests**
   ```bash
   make test-backend
   make test-frontend
   ```

## Development Workflow

```bash
# 1. Start database
make dev-db

# 2. Make changes to backend code
# Backend auto-reloads with Spring DevTools

# 3. Make changes to frontend code
# Restart frontend: Ctrl+C and make dev-frontend

# 4. Test changes
make test-backend
make test-frontend

# 5. Commit changes
git add .
git commit -m "feat: add new feature"

# 6. Stop services when done
make dev-db-stop
```

## Useful Commands

```bash
# View all available commands
make help

# Check service status
make status

# View logs
make docker-logs
make docker-logs-backend
make docker-logs-db

# Clean everything
make clean
make clean-docker

# Database operations
make db-connect
```

## Getting Help

- Check [README.md](./README.md) for detailed documentation
- Check [docs/](./docs/) for specifications and diagrams
- Check individual component READMEs:
  - [Backend README](./backend/README.md)
  - [Frontend README](./frontend/README.md)
  - [Database README](./database/README.md)

## Success Checklist

- [ ] Database running on port 5432
- [ ] Backend responding at http://localhost:8080/api/health
- [ ] Frontend TUI displays login screen
- [ ] Can navigate with Tab key
- [ ] Can exit with Esc
- [ ] No error messages in logs

If all checkboxes are checked, you're ready to develop! 🎉
