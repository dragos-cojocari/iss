# BORK Backend

Spring Boot REST API for the BORK (Book Organization & Rental Kiosk) library management system.

## Technology Stack

- **Java**: 25
- **Framework**: Spring Boot 3.5.0
- **ORM**: Hibernate (via Spring Data JPA)
- **Database**: PostgreSQL 16
- **Build Tool**: Maven

## Project Structure

```
backend/
├── src/
│   ├── main/
│   │   ├── java/com/bork/
│   │   │   ├── BorkApplication.java      # Main application class
│   │   │   ├── config/                   # Configuration classes
│   │   │   ├── controller/               # REST controllers
│   │   │   ├── service/                  # Business logic (future)
│   │   │   ├── repository/               # Data access (future)
│   │   │   └── model/                    # JPA entities (future)
│   │   └── resources/
│   │       ├── application.yml           # Main configuration
│   │       └── application-dev.yml       # Development profile
│   └── test/                             # Unit and integration tests
├── pom.xml                               # Maven dependencies
├── Dockerfile                            # Container image
└── README.md
```

## Prerequisites

- Java 25 (JDK)
- Maven 3.9+
- PostgreSQL (running via Docker or locally)

## Running Locally

### 1. Start Database

```bash
# From project root
make dev-db
```

### 2. Run Application

```bash
cd backend
mvn spring-boot:run
```

Or with specific profile:

```bash
mvn spring-boot:run -Dspring-boot.run.profiles=dev
```

### 3. Verify

```bash
curl http://localhost:8080/api/health
```

Expected response:

```json
{
  "status": "UP",
  "application": "BORK Backend",
  "version": "0.1.0-SNAPSHOT",
  "database": "Connected",
  "databaseProductName": "PostgreSQL",
  "databaseProductVersion": "16.x"
}
```

## Building

```bash
# Build JAR
mvn clean package

# Skip tests
mvn clean package -DskipTests

# Build Docker image
docker build -t bork-backend:latest .
```

## Testing

```bash
# Run all tests
mvn test

# Run specific test
mvn test -Dtest=HealthControllerTest
```

## Interactive API Documentation

The BORK API includes interactive Swagger UI documentation:

**URL:** http://localhost:8080/swagger-ui.html

### Features

- **Try It Out**: Test endpoints directly in the browser
- **Authentication**: Login and use session cookies for protected endpoints
- **Schema Validation**: View request/response models
- **OpenAPI Spec**: Export OpenAPI JSON/YAML for client generation

### OpenAPI Endpoints

- **Swagger UI**: http://localhost:8080/swagger-ui.html
- **OpenAPI JSON**: http://localhost:8080/v3/api-docs
- **OpenAPI YAML**: http://localhost:8080/v3/api-docs.yaml

---

## API Endpoints

### Health Check (Public)

- **GET** `/api/health` - Application and database health status

### Authentication (Public)

- **POST** `/api/auth/login` - Login with username and password
  - Request body: `{"username": "student1", "password": "Test123!"}`
  - Returns session cookie and user info
- **POST** `/api/auth/logout` - Logout and invalidate session
- **GET** `/api/auth/me` - Get current authenticated user info

### Books (Protected)

- **GET** `/api/books` - Get all books
- **GET** `/api/books/available` - Get available books only
- **GET** `/api/books/{id}` - Get book by ID (UUID)
- **GET** `/api/books/search?q={term}` - Search books by title or author

### Users (Protected - Testing)

- **GET** `/api/users` - Get all users
- **GET** `/api/users/{id}` - Get user by ID (UUID)
- **GET** `/api/users/username/{username}` - Get user by username

### Future Endpoints

- `/api/rentals` - Rental management
- `/api/cart` - Shopping cart operations
- `/api/categories` - Category management

## Configuration

Environment variables (see `.env.example`):

- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `POSTGRES_DB` - Database name (default: bork_db)
- `POSTGRES_USER` - Database user (default: bork_user)
- `POSTGRES_PASSWORD` - Database password (default: bork_password)
- `BACKEND_PORT` - Server port (default: 8080)
- `SPRING_PROFILES_ACTIVE` - Active profile (default: dev)

## Development

### Hot Reload

Spring Boot DevTools is included for automatic restart on code changes.

### Database Migrations

Currently using Hibernate's `validate` mode. Future: Flyway or Liquibase for migrations.

## Troubleshooting

### Database Connection Issues

```bash
# Check if database is running
docker ps | grep bork-db

# Check database logs
docker logs bork-db

# Test connection manually
psql -h localhost -p 5432 -U bork_user -d bork_db
```

### Port Already in Use

```bash
# Find process using port 8080
lsof -i :8080

# Kill process
kill -9 <PID>
```
