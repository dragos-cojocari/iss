-- Create Users Table
-- Stores user account information with BCrypt password hashing

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    failed_login_attempts INTEGER NOT NULL DEFAULT 0,

    CONSTRAINT chk_username_length CHECK (LENGTH(username) >= 3),
    CONSTRAINT chk_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_is_locked ON users(is_locked);

COMMENT ON TABLE users IS 'User accounts for library system';
COMMENT ON COLUMN users.user_id IS 'UUID primary key';
COMMENT ON COLUMN users.password_hash IS 'BCrypt hashed password (cost factor 12)';
COMMENT ON COLUMN users.is_locked IS 'Account locked after 3 failed login attempts';
COMMENT ON COLUMN users.failed_login_attempts IS 'Counter for failed login attempts';
