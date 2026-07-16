-- Trigger: Update updated_at for users
CREATE TRIGGER update_users_timestamp
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_updated_at();

-- Trigger: Update updated_at for events
CREATE TRIGGER update_events_timestamp
BEFORE UPDATE ON events
FOR EACH ROW EXECUTE FUNCTION update_updated_at();

-- Trigger: Update quantity_sold when order status changes
CREATE TRIGGER update_ticket_quantity
AFTER INSERT OR UPDATE ON orders
FOR EACH ROW EXECUTE FUNCTION update_quantity_sold();
