# BORK - Book Organization & Rental Kiosk

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Java](https://img.shields.io/badge/Java-21-orange.svg)](https://openjdk.org/projects/jdk/21/)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8.svg)](https://go.dev/)
[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.2-6DB33F.svg)](https://spring.io/projects/spring-boot)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791.svg)](https://www.postgresql.org/)

A library management system designed for university use, allowing students to browse, filter, and rent books through a terminal-based interface.

## 📚 Project Overview

BORK is a simplified library management system with:

- **Backend**: Spring Boot REST API with Hibernate ORM
- **Frontend**: Go-based Terminal User Interface (TUI)
- **Database**: PostgreSQL
- **Architecture**: Three-tier (Presentation, Business, Data layers)

### Key Features

- 📖 Browse and filter books by title, author, or category
- 🛒 Shopping cart-style rental system
- 📋 Rental management (max 3 books, 30-day period)
- ⚠️ Overdue notifications
- 🔐 User authentication
- 📊 Import books and users via CSV/JSON

## 🏗 Architecture

```
┌─────────────────┐
│   Go TUI        │  ← Terminal User Interface
│   (Frontend)    │
└────────┬────────┘
         │ HTTP/REST
         ▼
┌─────────────────┐
│  Spring Boot    │  ← Business Logic & REST API
│   (Backend)     │
└────────┬────────┘
         │ JDBC/Hibernate
         ▼
┌─────────────────┐
│   PostgreSQL    │  ← Data Persistence
│   (Database)    │
└─────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- **Docker** and **Docker Compose**
- **Java 21** (for local backend development)
- **Go 1.23** (for local frontend development)
- **Maven 3.9+** (for backend builds)

### Option 1: Run Everything with Docker

```bash
# Clone the repository
git clone <repository-url>
cd iss

# Set up environment
make setup

# Build and start all services
make docker-build
make docker-up

# Access the frontend TUI
docker attach bork-frontend

# Check backend health
curl http://localhost:8080/api/health
```

### Option 2: Local Development Mode

```bash
# 1. Start database only
make dev-db

# 2. In a new terminal, start backend
make dev-backend

# 3. In another terminal, install frontend dependencies and start TUI
make dev-frontend-deps
make dev-frontend
```

## 📁 Project Structure

```
iss/
├── backend/              # Spring Boot REST API
│   ├── src/
│   │   ├── main/java/com/bork/
│   │   │   ├── BorkApplication.java
│   │   │   ├── controller/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   └── model/
│   │   └── resources/
│   │       └── application.yml
│   ├── pom.xml
│   └── Dockerfile
├── frontend/             # Go TUI Application
│   ├── cmd/bork-tui/
│   ├── internal/ui/
│   ├── go.mod
│   └── Dockerfile
├── database/             # Database initialization
│   └── init/
├── docs/                 # Documentation
│   ├── specification.md
│   ├── class_diagram.md
│   ├── usecase_diagram.md
│   └── ...
├── docker-compose.yml    # Full stack
├── docker-compose.dev.yml # DB only
├── Makefile              # Build & run commands
└── README.md
```

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

### Backend Development

```bash
cd backend

# Run with Maven
mvn spring-boot:run

# Run tests
mvn test

# Build JAR
mvn clean package
```

**API Endpoints:**

- `GET /api/health` - Health check

### Frontend Development

```bash
cd frontend

# Install dependencies
go mod download

# Run TUI
go run cmd/bork-tui/main.go

# Run tests
go test ./...

# Build binary
go build -o bork-tui cmd/bork-tui/main.go
```

### Database Access

```bash
# Connect to database
make db-connect

# Or manually
docker exec -it bork-db psql -U bork_user -d bork_db
```

## 📖 Documentation

All functional specifications and technical documentation can be found in [docs/](./docs/index.md):

- [Specification](./docs/specification.md) - Functional requirements
- [Non-Functional Requirements](./docs/nfr.md) - Quality attributes
- [Use Case Diagram](./docs/usecase_diagram.md) - UML use cases
- [Class Diagram](./docs/class_diagram.md) - Conceptual model
- [Sequence Diagrams](./docs/sd_login.md) - Login and rental flows
- [TUI Prototype](./docs/tui_prototype.md) - UI mockups

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
API_BASE_URL=http://localhost:8080/api
```

## 📝 Lab Assignments

| Date       | Description                                                                             | Files                                                                                             |
| ---------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| 2026-05-09 | Lab 3 & 4: add class diagram and sequence diagrams for 2 usecases. Add the UI prototype | `docs/class_diagram.md`<br>`docs/sd_login.md`<br>`docs/sd_rentbook.md`<br>`docs/tui_prototype.md` |
| 2026-03-21 | Lab 2: Use Cases and Use Case Diagram                                                   | `docs/[ISS 2025-2026] Use case template.docx`<br>`docs/nfr.md`<br>`docs/usecase_diagram.md`       |
| 2026-03-07 | Lab 1: Flow diagram                                                                     | `docs/flow_diagram.md`                                                                            |
| 2026-03-07 | Lab 1: Project goals                                                                    | `Makefile`<br>`docs/specification.md`                                                             |

## 🤝 Contribution Guidelines

- **Pull Requests**: All contributions must be made via pull requests
- **Commit Format**: Follow [Conventional Commits](https://www.conventionalcommits.org/) specification
- **Signed Commits**: All commits must be signed with GPG/SSH keys
- **Pre-commit Hooks**: Run pre-commit hooks before submitting (`pre-commit run --all-files`)
- **Squash Commits**: Squash all commits into a single commit before merging
- **Stay Updated**: Keep your branch up to date with the main branch

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- University project for Software Systems Engineering course
- Built with Spring Boot, Go, PostgreSQL, and Docker
