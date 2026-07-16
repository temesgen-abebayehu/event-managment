-- Create categories table
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL
);

-- Insert default categories
INSERT INTO categories (name, slug) VALUES
    ('Music', 'music'),
    ('Sports', 'sports'),
    ('Business', 'business'),
    ('Arts', 'arts'),
    ('Food', 'food'),
    ('Technology', 'technology'),
    ('Education', 'education'),
    ('Community', 'community');
