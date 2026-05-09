# BORK Database

This directory contains database initialization scripts and migrations for the BORK application.

## Structure

- `init/` - Initial database setup scripts (run once on container creation)
  - `01-init.sql` - Database initialization

## PostgreSQL Configuration

- **Database**: `bork_db`
- **User**: `bork_user`
- **Password**: `bork_password` (development only)
- **Port**: 5432

## Running Locally

The database runs in a Docker container. Use the following commands:

```bash
# Start database only (for local development)
make dev-db

# Or using docker-compose directly
docker-compose -f docker-compose.dev.yml up -d
```

## Connecting to Database

```bash
# Using psql
psql -h localhost -p 5432 -U bork_user -d bork_db

# Using Docker exec
docker exec -it bork-db psql -U bork_user -d bork_db
```

## Future Migrations

Database migrations will be managed by Hibernate (via Spring Boot) or Flyway/Liquibase.
