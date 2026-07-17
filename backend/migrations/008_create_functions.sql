-- Function: Update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function: Update quantity_sold when order is completed
CREATE OR REPLACE FUNCTION update_quantity_sold()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' AND NEW.status = 'completed' THEN
        UPDATE tickets 
        SET quantity_sold = quantity_sold + NEW.quantity
        WHERE id = NEW.ticket_id;
    ELSIF TG_OP = 'UPDATE' AND OLD.status != 'completed' AND NEW.status = 'completed' THEN
        UPDATE tickets 
        SET quantity_sold = quantity_sold + NEW.quantity
        WHERE id = NEW.ticket_id;
    ELSIF TG_OP = 'UPDATE' AND OLD.status = 'completed' AND NEW.status = 'refunded' THEN
        UPDATE tickets 
        SET quantity_sold = quantity_sold - OLD.quantity
        WHERE id = NEW.ticket_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function: Search events with tags support
CREATE OR REPLACE FUNCTION search_events(search_term TEXT)
RETURNS SETOF events AS $$
BEGIN
    RETURN QUERY
    SELECT DISTINCT e.*
    FROM events e
    LEFT JOIN event_tags et ON e.id = et.event_id
    LEFT JOIN tags t ON et.tag_id = t.id
    WHERE 
        e.search_vector @@ plainto_tsquery('english', search_term)
        OR t.name ILIKE '%' || search_term || '%'
    ORDER BY ts_rank(e.search_vector, plainto_tsquery('english', search_term)) DESC;
END;
$$ LANGUAGE plpgsql STABLE;

-- Function: Get nearby events using Haversine formula
CREATE OR REPLACE FUNCTION get_nearby_events(
    lat DECIMAL,
    lng DECIMAL,
    radius_km DECIMAL DEFAULT 10
)
RETURNS TABLE (
    event_id UUID,
    distance_km DECIMAL
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        e.id,
        ROUND(
            CAST(
                (6371 * acos(
                    cos(radians(lat)) * 
                    cos(radians(e.latitude)) * 
                    cos(radians(e.longitude) - radians(lng)) + 
                    sin(radians(lat)) * 
                    sin(radians(e.latitude))
                )) AS NUMERIC
            ), 
            2
        ) AS distance_km
    FROM events e
    WHERE (6371 * acos(
        cos(radians(lat)) * 
        cos(radians(e.latitude)) * 
        cos(radians(e.longitude) - radians(lng)) + 
        sin(radians(lat)) * 
        sin(radians(e.latitude))
    )) <= radius_km
    ORDER BY distance_km ASC;
END;
$$ LANGUAGE plpgsql STABLE;
