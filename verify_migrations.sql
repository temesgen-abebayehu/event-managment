-- Verify Migration Status

-- 1. Check if slug column exists and has data
SELECT 
    'Slug column check:' as check_name,
    COUNT(*) as total_events,
    COUNT(slug) as events_with_slug,
    COUNT(CASE WHEN slug IS NULL THEN 1 END) as events_without_slug
FROM events;

-- 2. Show sample slugs
SELECT id, title, slug 
FROM events 
LIMIT 5;

-- 3. Check if search_vector exists
SELECT 
    'Search vector check:' as check_name,
    COUNT(*) as total_events,
    COUNT(search_vector) as events_with_search_vector
FROM events;

-- 4. Check if computed field functions exist
SELECT 
    'Computed functions check:' as check_name,
    COUNT(*) as function_count
FROM pg_proc 
WHERE proname IN (
    'event_tickets_remaining',
    'event_is_sold_out', 
    'event_status',
    'event_attendee_count'
);

-- 5. Check if triggers exist
SELECT 
    'Triggers check:' as check_name,
    COUNT(*) as trigger_count
FROM pg_trigger 
WHERE tgname IN (
    'trigger_auto_generate_slug',
    'trigger_update_search_vector',
    'trigger_update_search_on_tags'
);

-- 6. Check if indexes exist
SELECT 
    'Indexes check:' as check_name,
    COUNT(*) as index_count
FROM pg_indexes 
WHERE indexname IN (
    'idx_events_slug',
    'idx_events_search_vector'
);
