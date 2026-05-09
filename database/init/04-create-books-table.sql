-- Create Books Table
-- Stores book inventory with availability tracking

CREATE TABLE books (
    book_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(13) UNIQUE,
    category_id UUID NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,
    added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_books_category FOREIGN KEY (category_id)
        REFERENCES categories(category_id) ON DELETE RESTRICT,
    CONSTRAINT chk_isbn_format CHECK (isbn IS NULL OR isbn ~ '^\d{13}$')
);

CREATE INDEX idx_books_title ON books(title);
CREATE INDEX idx_books_author ON books(author);
CREATE INDEX idx_books_isbn ON books(isbn);
CREATE INDEX idx_books_category_id ON books(category_id);
CREATE INDEX idx_books_is_available ON books(is_available);

COMMENT ON TABLE books IS 'Book inventory for library system';
COMMENT ON COLUMN books.book_id IS 'UUID primary key';
COMMENT ON COLUMN books.isbn IS 'ISBN-13 format (13 digits), optional';
COMMENT ON COLUMN books.is_available IS 'TRUE if book can be rented';
COMMENT ON COLUMN books.category_id IS 'Foreign key to categories table';
