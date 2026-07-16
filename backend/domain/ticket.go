package domain

type Ticket struct {
	ID            string  `json:"id"`
	EventID       string  `json:"event_id"`
	Price         float64 `json:"price"`
	QuantityTotal int     `json:"quantity_total"`
	QuantitySold  int     `json:"quantity_sold"`
}

type Order struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	TicketID    string  `json:"ticket_id"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"` // pending, completed, failed, refunded
	ChapaTxRef  string  `json:"chapa_tx_ref"`
	CreatedAt   string  `json:"created_at"`
}
