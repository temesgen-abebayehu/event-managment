-- Migration: Add Hasura Computed Field Functions
-- These functions are used by Hasura to compute fields dynamically

-- =====================================================
-- COMPUTED FIELD: tickets_remaining
-- =====================================================
-- Returns the number of remaining tickets for an event
CREATE OR REPLACE FUNCTION event_tickets_remaining(event_row events)
RETURNS INTEGER AS $$
  SELECT COALESCE(quantity_total - quantity_sold, 0)::INTEGER
  FROM tickets
  WHERE event_id = event_row.id
  LIMIT 1;
$$ LANGUAGE SQL STABLE;

-- Grant execute permission (create role if needed)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'hasura_user') THEN
        CREATE ROLE hasura_user WITH LOGIN PASSWORD 'hasura_user_password';
    END IF;
END $$;

GRANT EXECUTE ON FUNCTION event_tickets_remaining(events) TO hasura_user;

-- =====================================================
-- COMPUTED FIELD: is_sold_out
-- =====================================================
-- Returns true if the event is sold out
CREATE OR REPLACE FUNCTION event_is_sold_out(event_row events)
RETURNS BOOLEAN AS $$
  SELECT COALESCE(quantity_sold >= quantity_total, FALSE)
  FROM tickets
  WHERE event_id = event_row.id
  LIMIT 1;
$$ LANGUAGE SQL STABLE;

-- Grant execute permission
GRANT EXECUTE ON FUNCTION event_is_sold_out(events) TO hasura_user;

-- =====================================================
-- COMPUTED FIELD: event_status
-- =====================================================
-- Returns 'past', 'upcoming', or 'ongoing' based on event date
CREATE OR REPLACE FUNCTION event_status(event_row events)
RETURNS TEXT AS $$
  SELECT CASE
    WHEN event_row.event_date < NOW() THEN 'past'
    WHEN event_row.event_date > NOW() + INTERVAL '1 hour' THEN 'upcoming'
    ELSE 'ongoing'
  END;
$$ LANGUAGE SQL STABLE;

-- Grant execute permission
GRANT EXECUTE ON FUNCTION event_status(events) TO hasura_user;

-- =====================================================
-- COMPUTED FIELD: attendee_count (bonus)
-- =====================================================
-- Returns the number of people who bought tickets
CREATE OR REPLACE FUNCTION event_attendee_count(event_row events)
RETURNS INTEGER AS $$
  SELECT COALESCE(quantity_sold, 0)::INTEGER
  FROM tickets
  WHERE event_id = event_row.id
  LIMIT 1;
$$ LANGUAGE SQL STABLE;

-- Grant execute permission
GRANT EXECUTE ON FUNCTION event_attendee_count(events) TO hasura_user;

-- =====================================================
-- TEST QUERIES (for verification)
-- =====================================================

-- Test tickets_remaining
-- SELECT id, title, event_tickets_remaining(events.*) as tickets_remaining FROM events LIMIT 5;

-- Test is_sold_out
-- SELECT id, title, event_is_sold_out(events.*) as is_sold_out FROM events LIMIT 5;

-- Test event_status
-- SELECT id, title, event_date, event_status(events.*) as status FROM events LIMIT 5;

-- Test attendee_count
-- SELECT id, title, event_attendee_count(events.*) as attendees FROM events LIMIT 5;
