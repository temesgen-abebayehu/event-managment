package repository

import (
	"event-management-backend/domain"
	"event-management-backend/infrastructure/hasura"
)

type TicketRepository struct {
	hasura *hasura.Client
}

func NewTicketRepository(hasura *hasura.Client) *TicketRepository {
	return &TicketRepository{hasura: hasura}
}

// GetByEventID - Used by payment action to get ticket info
func (r *TicketRepository) GetByEventID(eventID string) (*domain.Ticket, error) {
	query := `
		query($event_id: uuid!) {
			tickets(where: {event_id: {_eq: $event_id}}, limit: 1) {
				id event_id price quantity_total quantity_sold
			}
		}
	`

	variables := map[string]interface{}{"event_id": eventID}

	var result struct {
		Tickets []domain.Ticket `json:"tickets"`
	}

	err := r.hasura.Mutate(query, variables, &result)
	if err != nil || len(result.Tickets) == 0 {
		return nil, err
	}

	return &result.Tickets[0], nil
}
