.PHONY: help setup dev-db dev-backend dev-frontend dev-all docker-build docker-up docker-down docker-logs test-backend test-frontend lint-backend lint-frontend lint-all clean

# Default target
help:
	@echo "BORK - Book Organization & Rental Kiosk"
	@echo ""
	@echo "Available targets:"
	@echo "  setup            - Set up development environment (pre-commit hooks)"
	@echo "  dev-db           - Start PostgreSQL database only"
	@echo "  dev-backend      - Run backend locally (requires dev-db)"
	@echo "  dev-frontend     - Run frontend locally"
	@echo "  dev-all          - Run all services locally"
	@echo "  docker-build     - Build all Docker images"
	@echo "  docker-up        - Start all services with Docker Compose"
	@echo "  docker-down      - Stop all Docker services"
	@echo "  docker-logs      - View Docker logs"
	@echo "  test-backend     - Run backend tests"
	@echo "  test-frontend    - Run frontend tests"
	@echo "  lint-backend     - Lint backend code (Checkstyle)"
	@echo "  lint-frontend    - Lint frontend code (golangci-lint)"
	@echo "  lint-all         - Lint all code"
	@echo "  clean            - Clean build artifacts"

setup:
	@echo "Setting up development environment..."
	brew install maven
	@echo "Installing pre-commit hooks..."
	pre-commit install
	@echo "Creating .env file from example..."
	@if [ ! -f .env ]; then cp .env.example .env; echo ".env created"; else echo ".env already exists"; fi
	@echo "Setup complete!"

# Development - Database only
dev-db:
	@echo "Starting PostgreSQL database..."
	@echo "Pulling PostgreSQL image (if needed)..."
	@mkdir -p .docker-tmp && echo '{}' > .docker-tmp/config.json
	@export DOCKER_CONFIG=.docker-tmp && docker pull postgres:16-alpine 2>/dev/null || true
	docker compose -f docker-compose.dev.yml up -d
	@echo "Waiting for database to be ready..."
	@sleep 5
	@echo "Database is ready at localhost:5432"

dev-db-stop:
	@echo "Stopping PostgreSQL database..."
	docker compose -f docker-compose.dev.yml down

# Development - Backend
dev-backend:
	@echo "Starting backend (make sure database is running with 'make dev-db')..."
	cd backend && mvn spring-boot:run

# Development - Frontend
dev-frontend: dev-frontend-deps
	@echo "Starting frontend TUI..."
	cd frontend && go run cmd/bork-tui/main.go

# Development - Frontend (with dependencies)
dev-frontend-deps:
	@echo "Installing frontend dependencies..."
	cd frontend && go mod download
	@echo "Dependencies installed!"

# Development - All services locally
dev-all:
	@echo "Starting all services locally..."
	@echo "1. Starting database..."
	@make dev-db
	@echo "2. Backend and frontend must be started in separate terminals:"
	@echo "   Terminal 2: make dev-backend"
	@echo "   Terminal 3: make dev-frontend"

# Docker - Fetch images to workaround some config issues
docker-fetch-images:
	@echo "Fetching Docker images..."
	@mkdir -p .docker-tmp && echo '{}' > .docker-tmp/config.json
	@export DOCKER_CONFIG=.docker-tmp && docker pull postgres:16-alpine 2>/dev/null || true
	@export DOCKER_CONFIG=.docker-tmp && docker pull maven:3.9-eclipse-temurin-25-alpine 2>/dev/null || true
	@export DOCKER_CONFIG=.docker-tmp && docker pull eclipse-temurin:25-jre-alpine 2>/dev/null || true
	@export DOCKER_CONFIG=.docker-tmp && docker pull golang:1.26.3-alpine 2>/dev/null || true
	@export DOCKER_CONFIG=.docker-tmp && docker pull alpine:latest 2>/dev/null || true
	@export DOCKER_CONFIG=.docker-tmp && docker pull alpine:3.19 2>/dev/null || true
	docker compose build

# Docker - Build images. Uncomment the line below if you get a creds helper error
docker-build: docker-fetch-images
	@echo "Building Docker images..."
	docker compose build

# Docker - Start all services
docker-up:
	@echo "Starting all services with Docker Compose..."
	docker compose up -d
	@echo "Services started!"
	@echo "  - Database: localhost:5432"
	@echo "  - Backend: http://localhost:8080"
	@echo "  - Frontend: docker attach bork-frontend"

# Docker - Start with logs
docker-up-logs:
	@echo "Starting all services with logs..."
	docker compose up

# Docker - Stop all services
docker-down:
	@echo "Stopping all services..."
	docker compose down

# Docker - Stop and remove volumes
docker-down-volumes:
	@echo "Stopping all services and removing volumes..."
	docker compose down -v

# Docker - View logs
docker-logs:
	docker compose logs -f

# Docker - View specific service logs
docker-logs-db:
	docker compose logs -f db

docker-logs-backend:
	docker compose logs -f backend

docker-logs-frontend:
	docker compose logs -f frontend

# Docker - Restart services
docker-restart:
	@echo "Restarting all services..."
	docker compose restart

# Testing - Backend
test-backend:
	@echo "Running backend tests..."
	cd backend && mvn test

# Testing - Frontend
test-frontend:
	@echo "Running frontend tests..."
	cd frontend && go test ./...

# Testing - All
test-all: test-backend test-frontend

# Linting - Backend (Java with Checkstyle)
lint-backend:
	@echo "Linting backend code with Checkstyle..."
	cd backend && mvn checkstyle:check

# Linting - Frontend (Go with golangci-lint)
lint-frontend:
	@echo "Linting frontend code with golangci-lint..."
	cd frontend && golangci-lint run ./...

# Linting - All
lint-all: lint-backend lint-frontend

# Clean - Build artifacts
clean:
	@echo "Cleaning build artifacts..."
	cd backend && mvn clean
	cd frontend && go clean
	@echo "Clean complete!"

# Clean - Docker
clean-docker:
	@echo "Cleaning Docker resources..."
	docker compose down -v --rmi local
	@echo "Docker clean complete!"

# Database - Connect with psql
db-connect:
	@echo "Connecting to database..."
	docker exec -it bork-db psql -U bork_user -d bork_db

# Backend - Health check
backend-health:
	@echo "Checking backend health..."
	@curl -s http://localhost:8080/api/health | jq . || echo "Backend not responding or jq not installed"

# Status - Check all services
status:
	@echo "Checking service status..."
	@echo "\n=== Docker Containers ==="
	@docker ps -a | grep bork || echo "No BORK containers running"
	@echo "\n=== Backend Health ==="
	@curl -s http://localhost:8080/api/health 2>/dev/null | jq -r '.status // "Not responding"' || echo "Not responding"
	@echo "\n=== Database ==="
	@docker exec bork-db pg_isready -U bork_user -d bork_db 2>/dev/null || echo "Database not running"
