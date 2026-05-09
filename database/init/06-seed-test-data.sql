-- Seed Test Data
-- Sample users and books for development and testing
-- NOTE: Passwords are BCrypt hashed "Test123!" with cost factor 12

-- Test Users
-- Password for all test users: Test123!
-- BCrypt hash generated with cost factor 12
INSERT INTO users (username, email, first_name, last_name, password_hash) VALUES
    ('student1', 'student1@university.edu', 'Alice', 'Johnson', '$2a$12$npQ4eUifBcveY3V3nm3p7O7HoTPwE/b8zJF29X7nnstARkmxPGL8q'),
    ('student2', 'student2@university.edu', 'Bob', 'Smith', '$2a$12$npQ4eUifBcveY3V3nm3p7O7HoTPwE/b8zJF29X7nnstARkmxPGL8q'),
    ('student3', 'student3@university.edu', 'Carol', 'Williams', '$2a$12$npQ4eUifBcveY3V3nm3p7O7HoTPwE/b8zJF29X7nnstARkmxPGL8q'),
    ('jdoe', 'john.doe@university.edu', 'John', 'Doe', '$2a$12$npQ4eUifBcveY3V3nm3p7O7HoTPwE/b8zJF29X7nnstARkmxPGL8q'),
    ('testuser', 'test@university.edu', 'Test', 'User', '$2a$12$npQ4eUifBcveY3V3nm3p7O7HoTPwE/b8zJF29X7nnstARkmxPGL8q');

-- Test Books
-- Get category IDs for foreign keys
DO $$
DECLARE
    cat_fiction UUID;
    cat_science UUID;
    cat_history UUID;
    cat_tech UUID;
    cat_philosophy UUID;
    cat_business UUID;
BEGIN
    SELECT category_id INTO cat_fiction FROM categories WHERE name = 'Fiction';
    SELECT category_id INTO cat_science FROM categories WHERE name = 'Science';
    SELECT category_id INTO cat_history FROM categories WHERE name = 'History';
    SELECT category_id INTO cat_tech FROM categories WHERE name = 'Technology';
    SELECT category_id INTO cat_philosophy FROM categories WHERE name = 'Philosophy';
    SELECT category_id INTO cat_business FROM categories WHERE name = 'Business';

    INSERT INTO books (title, author, isbn, category_id, is_available) VALUES
        ('The Great Gatsby', 'F. Scott Fitzgerald', '9780743273565', cat_fiction, TRUE),
        ('To Kill a Mockingbird', 'Harper Lee', '9780061120084', cat_fiction, TRUE),
        ('1984', 'George Orwell', '9780451524935', cat_fiction, TRUE),
        ('Pride and Prejudice', 'Jane Austen', '9780141439518', cat_fiction, FALSE),
        ('A Brief History of Time', 'Stephen Hawking', '9780553380163', cat_science, TRUE),
        ('The Selfish Gene', 'Richard Dawkins', '9780198788607', cat_science, TRUE),
        ('Sapiens', 'Yuval Noah Harari', '9780062316097', cat_history, TRUE),
        ('Guns, Germs, and Steel', 'Jared Diamond', '9780393317558', cat_history, FALSE),
        ('Clean Code', 'Robert C. Martin', '9780132350884', cat_tech, TRUE),
        ('The Pragmatic Programmer', 'Andrew Hunt', '9780135957059', cat_tech, TRUE),
        ('Design Patterns', 'Gang of Four', '9780201633610', cat_tech, TRUE),
        ('Thinking, Fast and Slow', 'Daniel Kahneman', '9780374533557', cat_philosophy, TRUE),
        ('The Lean Startup', 'Eric Ries', '9780307887894', cat_business, TRUE),
        ('Good to Great', 'Jim Collins', '9780066620992', cat_business, FALSE);
END $$;

-- Display summary
SELECT 'Seeded ' || COUNT(*) || ' test users' AS status FROM users;
SELECT 'Seeded ' || COUNT(*) || ' test books' AS status FROM books;
SELECT 'Available books: ' || COUNT(*) AS status FROM books WHERE is_available = TRUE;
