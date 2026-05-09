# BORK - Book Organization & Rental Kiosk

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Java](https://img.shields.io/badge/Java-25-orange.svg)](https://openjdk.org/projects/jdk/25/)
[![Go](https://img.shields.io/badge/Go-1.26.3-00ADD8.svg)](https://go.dev/)
[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.5-6DB33F.svg)](https://spring.io/projects/spring-boot)
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

- **Docker** & **Docker Compose**
- **Java 25** (local dev)
- **Go 1.26.3** (local dev)
- **Maven 3.9+** (backend builds)

### Local Development (Recommended)

```bash
# 1. Start database
make dev-db

# 2. Start backend (new terminal)
make dev-backend

# 3. Start frontend TUI (new terminal)
make dev-frontend
```

**Login:** `student1` / `Test123!`

### Docker (Full Stack)

```bash
make docker-build
make docker-up
docker attach bork-frontend
```

### 📖 API Documentation

**Swagger UI:** http://localhost:8080/swagger-ui.html

Test endpoints, view schemas, and authenticate directly in the browser.

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

See [Development Guide](docs/development.md) for detailed instructions on:

- Make targets and commands
- Backend/Frontend development
- Testing and Docker usage
- Configuration and environment variables
- API testing with curl

## � Documentation

See [Documentation Index](docs/index.md) for complete documentation.

**Quick Links:**

- [Quick Start Guide](docs/quickstart.md) - Get running in under 5 minutes
- [Development Guide](docs/development.md) - Setup, testing, Docker, configuration
- [API Guide](docs/api_guide.md) - REST API reference with curl examples
- [Design & Architecture](docs/design/) - Requirements, UML diagrams, UI prototypes

## 📝 Lab Assignments

See [Lab Assignments](docs/design/lab_assignments.md) for a complete timeline of deliverables.

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
