-- Create event_images table
CREATE TABLE event_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    public_id VARCHAR(255) NOT NULL,
    is_featured BOOLEAN DEFAULT false
);

CREATE INDEX idx_event_images_event_id ON event_images(event_id);
