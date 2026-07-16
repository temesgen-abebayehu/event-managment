package repository

import (
	"event-management-backend/domain"
	"event-management-backend/infrastructure/hasura"
)

type OrderRepository struct {
	hasura *hasura.Client
}

func NewOrderRepository(hasura *hasura.Client) *OrderRepository {
	return &OrderRepository{hasura: hasura}
}

// Create - Used by payment action to create pending order
func (r *OrderRepository) Create(userID, ticketID string, quantity int, totalPrice float64, chapaTxRef string) (*domain.Order, error) {
	query := `
		mutation($user_id: uuid!, $ticket_id: uuid!, $quantity: Int!, $total_price: numeric!, $chapa_tx_ref: String!) {
			insert_orders_one(object: {
				user_id: $user_id
				ticket_id: $ticket_id
				quantity: $quantity
				total_price: $total_price
				status: "pending"
				chapa_tx_ref: $chapa_tx_ref
			}) {
				id user_id ticket_id quantity total_price status chapa_tx_ref created_at
			}
		}
	`

	variables := map[string]interface{}{
		"user_id":      userID,
		"ticket_id":    ticketID,
		"quantity":     quantity,
		"total_price":  totalPrice,
		"chapa_tx_ref": chapaTxRef,
	}

	var result struct {
		InsertOrdersOne domain.Order `json:"insert_orders_one"`
	}

	err := r.hasura.Mutate(query, variables, &result)
	return &result.InsertOrdersOne, err
}

// UpdateStatus - Used by payment callback to update order status
func (r *OrderRepository) UpdateStatus(chapaTxRef, status string) error {
	query := `
		mutation($chapa_tx_ref: String!, $status: String!) {
			update_orders(
				where: {chapa_tx_ref: {_eq: $chapa_tx_ref}}
				_set: {status: $status}
			) {
				affected_rows
			}
		}
	`

	variables := map[string]interface{}{
		"chapa_tx_ref": chapaTxRef,
		"status":       status,
	}

	var result interface{}
	return r.hasura.Mutate(query, variables, &result)
}

// FindByTxRef - Used by payment callback to get order
func (r *OrderRepository) FindByTxRef(chapaTxRef string) (*domain.Order, error) {
	query := `
		query($chapa_tx_ref: String!) {
			orders(where: {chapa_tx_ref: {_eq: $chapa_tx_ref}}, limit: 1) {
				id user_id ticket_id quantity total_price status chapa_tx_ref created_at
			}
		}
	`

	variables := map[string]interface{}{"chapa_tx_ref": chapaTxRef}

	var result struct {
		Orders []domain.Order `json:"orders"`
	}

	err := r.hasura.Mutate(query, variables, &result)
	if err != nil || len(result.Orders) == 0 {
		return nil, err
	}

	return &result.Orders[0], nil
}
