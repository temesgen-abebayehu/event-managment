-- Create follows table (composite PK)
CREATE TABLE follows (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, event_id)
);

-- Create bookmarks table (composite PK)
CREATE TABLE bookmarks (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, event_id)
);

CREATE INDEX idx_follows_event_id ON follows(event_id);
CREATE INDEX idx_bookmarks_event_id ON bookmarks(event_id);
