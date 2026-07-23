-- Test if events can be queried with all fields
SELECT 
    id, 
    slug,
    title, 
    venue, 
    address, 
    full_location,
    price,
    event_date
FROM events 
LIMIT 2;

-- Check if full_location is a generated column
SELECT column_name, data_type, is_generated, generation_expression
FROM information_schema.columns
WHERE table_name = 'events' 
AND column_name IN ('full_location', 'slug', 'search_vector');
