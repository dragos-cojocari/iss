-- Create Categories Table
-- Stores book categories/genres for classification and filtering

CREATE TABLE categories (
    category_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_categories_name ON categories(name);

COMMENT ON TABLE categories IS 'Book categories for classification';
COMMENT ON COLUMN categories.category_id IS 'UUID primary key';
COMMENT ON COLUMN categories.name IS 'Unique category name';
