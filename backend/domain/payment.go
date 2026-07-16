package domain

type InitiatePaymentRequest struct {
	EventID  string `json:"event_id"`
	Quantity int    `json:"quantity"`
}

type InitiatePaymentResponse struct {
	CheckoutURL string  `json:"checkout_url"`
	TxRef       string  `json:"tx_ref"`
	TotalPrice  float64 `json:"total_price"`
	OrderID     string  `json:"order_id"`
}

type VerifyPaymentRequest struct {
	TxRef string `json:"tx_ref"`
}

type VerifyPaymentResponse struct {
	Status string       `json:"status"`
	Order  *Order       `json:"order,omitempty"`
}
