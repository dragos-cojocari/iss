# Development Guide

## 🛠 Development

### Available Make Targets

```bash
make help              # Show all available commands
make setup             # Set up development environment
make dev-db            # Start PostgreSQL only
make dev-backend       # Run backend locally
make dev-frontend      # Run frontend TUI locally
make docker-build      # Build all Docker images
make docker-up         # Start all services
make docker-down       # Stop all services
make test-backend      # Run backend tests
make test-frontend     # Run frontend tests
make clean             # Clean build artifacts
make status            # Check service status
```

## Backend Development

```bash
cd backend
mvn spring-boot:run    # Run
mvn test               # Test
mvn clean package      # Build JAR
```

**Key Endpoints:** `/api/auth/login`, `/api/auth/me`, `/api/books/available`

**Swagger UI:** http://localhost:8080/swagger-ui.html

## Frontend Development

```bash
cd frontend
go run cmd/bork/main.go    # Run TUI
go test ./...              # Test
go build -o bork cmd/bork/main.go  # Build
```

**Navigation:**

- Login → Overdue Alert → Dashboard
- Dashboard: `1` Browse Books, `4` Logout
- Browse: `↑/↓` navigate, `←/→` page, `B` back
- `Esc` quit (login/dashboard) or go back (other views)

## Database Access

```bash
make db-connect
# Or: docker exec -it bork-db psql -U bork_user -d bork_db
```

## 🧪 Testing

```bash
# Run all tests
make test-all

# Backend tests only
make test-backend

# Frontend tests only
make test-frontend
```

## 🐳 Docker

### Build Images

```bash
docker-compose build
```

### Run Services

```bash
# Start all services in background
docker-compose up -d

# Start with logs
docker-compose up

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f db
```

## 🔧 Configuration

Environment variables (see `.env.example`):

```bash
# Database
POSTGRES_DB=bork_db
POSTGRES_USER=bork_user
POSTGRES_PASSWORD=bork_password
DB_HOST=localhost
DB_PORT=5432

# Backend
BACKEND_PORT=8080
SPRING_PROFILES_ACTIVE=dev

# Frontend
BORK_BACKEND_URL=http://localhost:8080
```

## Session-Based Authentication

The application uses session-based authentication with HTTP-only cookies:

- **Login:** `POST /api/auth/login` with `username` and `password`
- **Session Cookie:** `BORK_SESSION` (30-minute expiration)
- **Current User:** `GET /api/auth/me` (requires authentication)
- **Logout:** `POST /api/auth/logout`

All protected endpoints require a valid session cookie.

## API Testing with curl

See [API Guide](api_guide.md) for detailed curl examples and endpoint documentation.

Quick test:

```bash
# Login and save session cookie
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"student1","password":"Test123!"}' \
  -c cookies.txt

# Get current user
curl http://localhost:8080/api/auth/me -b cookies.txt

# Get available books
curl http://localhost:8080/api/books/available -b cookies.txt

# Logout
curl -X POST http://localhost:8080/api/auth/logout -b cookies.txt
```
