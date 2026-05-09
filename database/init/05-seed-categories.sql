-- Seed Categories
-- Initial book categories for the library system

INSERT INTO categories (name, description) VALUES
    ('Fiction', 'Literary works of imaginative narration'),
    ('Non-Fiction', 'Factual and informative works'),
    ('Science', 'Scientific and technical literature'),
    ('History', 'Historical accounts and biographies'),
    ('Biography', 'Life stories and memoirs'),
    ('Technology', 'Computing, engineering, and technical books'),
    ('Arts', 'Visual arts, music, and performing arts'),
    ('Philosophy', 'Philosophical works and critical thinking'),
    ('Business', 'Business, economics, and management'),
    ('Self-Help', 'Personal development and improvement');

SELECT 'Seeded ' || COUNT(*) || ' categories' AS status FROM categories;
