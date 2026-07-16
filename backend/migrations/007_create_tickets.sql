-- Create tickets table (event ticket template)
CREATE TABLE tickets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL DEFAULT 0,
    quantity_total INTEGER NOT NULL,
    quantity_sold INTEGER NOT NULL DEFAULT 0,
    CONSTRAINT check_quantity_sold CHECK (quantity_sold <= quantity_total)
);

-- Create orders table (actual purchases)
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    ticket_id UUID NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL DEFAULT 1,
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    chapa_tx_ref VARCHAR(255) UNIQUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT check_status CHECK (status IN ('pending', 'completed', 'failed', 'refunded'))
);

CREATE INDEX idx_tickets_event_id ON tickets(event_id);
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_ticket_id ON orders(ticket_id);
CREATE INDEX idx_orders_tx_ref ON orders(chapa_tx_ref);
