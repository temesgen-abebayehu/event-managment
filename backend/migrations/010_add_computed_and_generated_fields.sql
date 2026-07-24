-- Migration: Add computed fields and generated properties

-- =====================================================
-- GENERATED PROPERTIES
-- =====================================================

-- Add slug column to events (if not exists)
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'events' AND column_name = 'slug'
    ) THEN
        ALTER TABLE events ADD COLUMN slug TEXT;
    END IF;
END $$;

-- Add unique constraint on slug (if not exists)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'events_slug_key'
    ) THEN
        ALTER TABLE events ADD CONSTRAINT events_slug_key UNIQUE (slug);
    END IF;
END $$;

-- Function to generate slug from title
CREATE OR REPLACE FUNCTION generate_slug(title TEXT, event_id UUID)
RETURNS TEXT AS $$
DECLARE
    base_slug TEXT;
    final_slug TEXT;
    counter INTEGER := 0;
BEGIN
    -- Convert title to lowercase, replace spaces with hyphens, remove special chars
    base_slug := lower(regexp_replace(title, '[^a-zA-Z0-9\s-]', '', 'g'));
    base_slug := regexp_replace(base_slug, '\s+', '-', 'g');
    base_slug := regexp_replace(base_slug, '-+', '-', 'g');
    base_slug := trim(both '-' from base_slug);
    
    -- Add short UUID suffix for uniqueness
    final_slug := base_slug || '-' || substring(event_id::TEXT from 1 for 8);
    
    RETURN final_slug;
END;
$$ LANGUAGE plpgsql IMMUTABLE;

-- Generate slugs for existing events
UPDATE events 
SET slug = generate_slug(title, id)
WHERE slug IS NULL;

-- Make slug NOT NULL after populating existing data
ALTER TABLE events 
ALTER COLUMN slug SET NOT NULL;

-- Create index on slug for fast lookups (if not exists)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes WHERE indexname = 'idx_events_slug'
    ) THEN
        CREATE INDEX idx_events_slug ON events(slug);
    END IF;
END $$;

-- =====================================================
-- TRIGGERS FOR AUTO-GENERATING SLUG
-- =====================================================

-- Function to auto-generate slug on insert/update
CREATE OR REPLACE FUNCTION auto_generate_slug()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.slug IS NULL OR (TG_OP = 'UPDATE' AND OLD.title != NEW.title) THEN
        NEW.slug := generate_slug(NEW.title, NEW.id);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to auto-generate slug
CREATE TRIGGER trigger_auto_generate_slug
BEFORE INSERT OR UPDATE ON events
FOR EACH ROW
EXECUTE FUNCTION auto_generate_slug();

-- =====================================================
-- COMPUTED PROPERTIES (via Hasura)
-- =====================================================
-- Note: These are configured in Hasura Console, but we document them here:
-- 
-- 1. tickets_remaining: 
--    SQL: SELECT (quantity_total - quantity_sold) FROM tickets WHERE event_id = $1 LIMIT 1
--
-- 2. is_sold_out:
--    SQL: SELECT (quantity_sold >= quantity_total) FROM tickets WHERE event_id = $1 LIMIT 1
--
-- 3. event_status:
--    SQL: SELECT CASE 
--           WHEN event_date < NOW() THEN 'past'
--           WHEN event_date > NOW() + INTERVAL '1 hour' THEN 'upcoming'
--           ELSE 'ongoing'
--         END FROM events WHERE id = $1

-- =====================================================
-- ENHANCED SEARCH VECTOR (Already exists from migration 003)
-- =====================================================

-- search_vector already exists as a generated column from migration 003
-- We just need to update the trigger to include tags

-- Function to generate search vector from title, description, and tags
CREATE OR REPLACE FUNCTION update_event_search_vector()
RETURNS TRIGGER AS $$
DECLARE
    tags_text TEXT := '';
BEGIN
    -- Get all tags for this event
    SELECT string_agg(t.name, ' ') INTO tags_text
    FROM event_tags et
    JOIN tags t ON et.tag_id = t.id
    WHERE et.event_id = NEW.id;
    
    -- Generate tsvector from title, description, venue, address, and tags
    NEW.search_vector := 
        setweight(to_tsvector('english', coalesce(NEW.title, '')), 'A') ||
        setweight(to_tsvector('english', coalesce(NEW.description, '')), 'B') ||
        setweight(to_tsvector('english', coalesce(NEW.venue, '')), 'C') ||
        setweight(to_tsvector('english', coalesce(NEW.address, '')), 'C') ||
        setweight(to_tsvector('english', coalesce(tags_text, '')), 'B');
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to auto-update search vector
CREATE TRIGGER trigger_update_search_vector
BEFORE INSERT OR UPDATE ON events
FOR EACH ROW
EXECUTE FUNCTION update_event_search_vector();

-- Also update search vector when tags change
CREATE OR REPLACE FUNCTION update_event_search_on_tags()
RETURNS TRIGGER AS $$
BEGIN
    -- Just touch the updated_at to trigger the search_vector recalculation
    UPDATE events 
    SET updated_at = NOW()
    WHERE id = COALESCE(NEW.event_id, OLD.event_id);
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_search_on_tags
AFTER INSERT OR UPDATE OR DELETE ON event_tags
FOR EACH ROW
EXECUTE FUNCTION update_event_search_on_tags();

-- Create GIN index for fast full-text search (if not exists)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_indexes WHERE indexname = 'idx_events_search_vector'
    ) THEN
        CREATE INDEX idx_events_search_vector ON events USING GIN(search_vector);
    END IF;
END $$;

-- =====================================================
-- UPDATE SEARCH FUNCTION TO USE SEARCH VECTOR
-- =====================================================

-- Drop old search function
DROP FUNCTION IF EXISTS search_events(TEXT);

-- New optimized search function using search_vector
CREATE OR REPLACE FUNCTION search_events(search_term TEXT)
RETURNS SETOF events AS $$
BEGIN
    RETURN QUERY
    SELECT e.*
    FROM events e
    WHERE e.search_vector @@ plainto_tsquery('english', search_term)
    ORDER BY ts_rank(e.search_vector, plainto_tsquery('english', search_term)) DESC;
END;
$$ LANGUAGE plpgsql STABLE;

-- =====================================================
-- HELPER VIEWS (Optional - for convenience)
-- =====================================================

-- View for events with computed data
CREATE OR REPLACE VIEW events_with_status AS
SELECT 
    e.*,
    CASE 
        WHEN e.event_date < NOW() THEN 'past'
        WHEN e.event_date > NOW() + INTERVAL '1 hour' THEN 'upcoming'
        ELSE 'ongoing'
    END as status,
    COALESCE(t.quantity_total - t.quantity_sold, 0) as tickets_remaining,
    COALESCE(t.quantity_sold >= t.quantity_total, false) as is_sold_out
FROM events e
LEFT JOIN tickets t ON e.id = t.event_id;

-- Grant permissions on view (create role if needed)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'hasura_user') THEN
        CREATE ROLE hasura_user WITH LOGIN PASSWORD 'hasura_user_password';
    END IF;
END $$;

GRANT SELECT ON events_with_status TO hasura_user;
