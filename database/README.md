# BORK Database

This directory contains database initialization scripts and migrations for the BORK application.

## Structure

- `init/` - Initial database setup scripts (run once on container creation)
  - `01-init.sql` - Database initialization
  - `02-create-users-table.sql` - Users table with BCrypt password hashing
  - `03-create-categories-table.sql` - Book categories table
  - `04-create-books-table.sql` - Books inventory table
  - `05-seed-categories.sql` - Initial category data
  - `06-seed-test-data.sql` - Test users and books

## PostgreSQL Configuration

- **Database**: `bork_db`
- **User**: `bork_user`
- **Password**: `bork_password` (development only)
- **Port**: 5432

## Database Schema

### Tables

#### users

- **Primary Key**: `user_id` (UUID)
- **Unique Constraints**: `username`, `email`
- **Password**: BCrypt hashed with cost factor 12
- **Business Rules**:
  - Account locks after 3 failed login attempts
  - Maximum 3 active rentals per user

#### categories

- **Primary Key**: `category_id` (UUID)
- **Unique Constraints**: `name`
- **Purpose**: Book classification and filtering

#### books

- **Primary Key**: `book_id` (UUID)
- **Foreign Keys**: `category_id` → categories
- **Unique Constraints**: `isbn` (optional)
- **Business Rules**:
  - Only one active rental per book
  - Availability tracked via `is_available` flag

## Test Data

### Test Users

All test users have password: `Test123!`

- `student1` / `student1@university.edu` - Alice Johnson
- `student2` / `student2@university.edu` - Bob Smith
- `student3` / `student3@university.edu` - Carol Williams
- `jdoe` / `john.doe@university.edu` - John Doe
- `testuser` / `test@university.edu` - Test User

### Test Books

14 books across various categories (Fiction, Science, History, Technology, etc.)
Mix of available and unavailable books for testing rental scenarios

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
