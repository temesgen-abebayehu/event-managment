-- Migration 012: Fix search_vector to work with tags while keeping GENERATED constraint

-- =====================================================
-- PROBLEM: 
-- search_vector is GENERATED ALWAYS (only includes title + description)
-- BUT we need it to also include tags dynamically
-- Triggers in migration 010 tried to SET search_vector, causing error
-- =====================================================

-- SOLUTION:
-- 1. Keep search_vector as GENERATED (requirement)
-- 2. Remove triggers that try to manually set search_vector
-- 3. Create a SEPARATE computed field for tags in search
-- 4. Combine them at query time OR use a view
-- =====================================================

-- Step 1: Drop problematic triggers from migration 010
DROP TRIGGER IF EXISTS trigger_update_search_vector ON events;
DROP TRIGGER IF EXISTS trigger_update_search_on_tags ON event_tags;

-- Step 2: Drop problematic functions from migration 010
DROP FUNCTION IF EXISTS update_event_search_vector();
DROP FUNCTION IF EXISTS update_event_search_on_tags();

-- Step 3: Create a function to get tags as tsvector (for search)
CREATE OR REPLACE FUNCTION get_event_tags_tsvector(event_id_param UUID)
RETURNS tsvector AS $$
DECLARE
    tags_text TEXT;
BEGIN
    SELECT string_agg(t.name, ' ') INTO tags_text
    FROM event_tags et
    JOIN tags t ON et.tag_id = t.id
    WHERE et.event_id = event_id_param;
    
    IF tags_text IS NULL THEN
        RETURN to_tsvector('english', '');
    END IF;
    
    RETURN setweight(to_tsvector('english', tags_text), 'B');
END;
$$ LANGUAGE plpgsql STABLE;

-- Step 4: Create enhanced search function that includes tags
CREATE OR REPLACE FUNCTION search_events_with_tags(search_term TEXT)
RETURNS SETOF events AS $$
BEGIN
    RETURN QUERY
    SELECT e.*
    FROM events e
    WHERE (e.search_vector || get_event_tags_tsvector(e.id)) 
          @@ plainto_tsquery('english', search_term)
    ORDER BY ts_rank(
        (e.search_vector || get_event_tags_tsvector(e.id)), 
        plainto_tsquery('english', search_term)
    ) DESC;
END;
$$ LANGUAGE plpgsql STABLE;

-- Step 5: Create a materialized view for better performance (optional, for heavy usage)
-- This pre-computes search_vector with tags
CREATE MATERIALIZED VIEW IF NOT EXISTS events_search_index AS
SELECT 
    e.id,
    e.search_vector || get_event_tags_tsvector(e.id) as full_search_vector
FROM events e;

-- Create index on materialized view
CREATE INDEX IF NOT EXISTS idx_events_search_full ON events_search_index USING GIN(full_search_vector);

-- Step 6: Function to refresh the materialized view
CREATE OR REPLACE FUNCTION refresh_events_search_index()
RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY events_search_index;
END;
$$ LANGUAGE plpgsql;

-- Step 7: Trigger to refresh materialized view when events change
CREATE OR REPLACE FUNCTION trigger_refresh_search_index()
RETURNS TRIGGER AS $$
BEGIN
    PERFORM refresh_events_search_index();
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Note: These triggers refresh the materialized view
-- For high-traffic apps, consider scheduled refresh instead
CREATE TRIGGER trigger_refresh_on_event_change
AFTER INSERT OR UPDATE OR DELETE ON events
FOR EACH STATEMENT
EXECUTE FUNCTION trigger_refresh_search_index();

CREATE TRIGGER trigger_refresh_on_tags_change
AFTER INSERT OR UPDATE OR DELETE ON event_tags
FOR EACH STATEMENT
EXECUTE FUNCTION trigger_refresh_search_index();

-- Step 8: Initial refresh of materialized view
REFRESH MATERIALIZED VIEW events_search_index;

-- =====================================================
-- HASURA COMPUTED FIELD FOR TAGS
-- =====================================================
-- Add this in Hasura Console -> Data -> events table -> Computed Fields:
-- 
-- Field Name: search_with_tags
-- Function: get_event_tags_tsvector
-- 
-- This allows GraphQL queries like:
-- query SearchEvents($term: String!) {
--   search_events_with_tags(args: {search_term: $term}) {
--     id
--     title
--     search_with_tags
--   }
-- }

-- =====================================================
-- VERIFICATION QUERIES
-- =====================================================

-- Check that search_vector is still GENERATED:
-- SELECT column_name, is_generated, generation_expression
-- FROM information_schema.columns 
-- WHERE table_name = 'events' AND column_name = 'search_vector';
-- (Should show is_generated = 'ALWAYS')

-- Test search with tags:
-- SELECT * FROM search_events_with_tags('music');

-- Test tag insertion (should work without errors):
-- INSERT INTO event_tags (event_id, tag_id) 
-- VALUES ('your-event-id', 'your-tag-id')
-- ON CONFLICT (event_id, tag_id) DO NOTHING;
