package usecase

import (
	"errors"
	"event-management-backend/domain"
	"event-management-backend/infrastructure/chapa"
	"event-management-backend/repository"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PaymentUsecase struct {
	ticketRepo  *repository.TicketRepository
	orderRepo   *repository.OrderRepository
	userRepo    *repository.UserRepository
	chapaClient *chapa.Client
	callbackURL string
	returnURL   string
}

func NewPaymentUsecase(
	ticketRepo *repository.TicketRepository,
	orderRepo *repository.OrderRepository,
	userRepo *repository.UserRepository,
	chapaClient *chapa.Client,
	callbackURL, returnURL string,
) *PaymentUsecase {
	return &PaymentUsecase{
		ticketRepo:  ticketRepo,
		orderRepo:   orderRepo,
		userRepo:    userRepo,
		chapaClient: chapaClient,
		callbackURL: callbackURL,
		returnURL:   returnURL,
	}
}

func (u *PaymentUsecase) InitiatePayment(userID string, req domain.InitiatePaymentRequest) (*domain.InitiatePaymentResponse, error) {
	// Validate
	if req.Quantity < 1 {
		return nil, errors.New("quantity must be at least 1")
	}

	// Get the authenticated user's real info for Chapa
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Get ticket info
	ticket, err := u.ticketRepo.GetByEventID(req.EventID)
	if err != nil {
		return nil, errors.New("ticket not found")
	}

	// Check availability
	available := ticket.QuantityTotal - ticket.QuantitySold
	if req.Quantity > available {
		return nil, fmt.Errorf("only %d tickets available", available)
	}

	// Calculate total price
	totalPrice := ticket.Price * float64(req.Quantity)

	// Generate unique transaction reference
	txRef := fmt.Sprintf("EVENT-%d-%s", time.Now().Unix(), uuid.New().String()[:8])

	// Create pending order
	order, err := u.orderRepo.Create(userID, ticket.ID, req.Quantity, totalPrice, txRef)
	if err != nil {
		return nil, errors.New("failed to create order")
	}

	// Split full_name into first/last for Chapa
	firstName, lastName := splitName(user.FullName)

	// Initialize Chapa payment with real user data
	chapaReq := chapa.InitializePaymentRequest{
		Amount:      fmt.Sprintf("%.2f", totalPrice),
		Currency:    "ETB",
		Email:       user.Email,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: getPhoneOrDefault(user.Phone),
		TxRef:       txRef,
		CallbackURL: u.callbackURL,
		ReturnURL:   u.returnURL,
	}

	chapaResp, err := u.chapaClient.InitializePayment(chapaReq)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize payment: %v", err)
	}

	return &domain.InitiatePaymentResponse{
		CheckoutURL: chapaResp.Data.CheckoutURL,
		TxRef:       txRef,
		TotalPrice:  totalPrice,
		OrderID:     order.ID,
	}, nil
}

func (u *PaymentUsecase) VerifyPayment(txRef string) (*domain.VerifyPaymentResponse, error) {
	// Verify with Chapa
	chapaResp, err := u.chapaClient.VerifyPayment(txRef)
	if err != nil {
		return nil, errors.New("failed to verify payment")
	}

	// Update order status based on Chapa response
	var status string
	if chapaResp.Data.Status == "success" {
		status = "completed"
	} else {
		status = "failed"
	}

	err = u.orderRepo.UpdateStatus(txRef, status)
	if err != nil {
		return nil, errors.New("failed to update order")
	}

	// Get updated order
	order, err := u.orderRepo.FindByTxRef(txRef)
	if err != nil {
		return nil, errors.New("order not found")
	}

	return &domain.VerifyPaymentResponse{
		Status: status,
		Order:  order,
	}, nil
}

func (u *PaymentUsecase) HandleCallback(txRef, status string) error {
	// Update order status from Chapa callback
	var orderStatus string
	if status == "success" {
		orderStatus = "completed"
	} else {
		orderStatus = "failed"
	}

	return u.orderRepo.UpdateStatus(txRef, orderStatus)
}

// splitName splits "John Doe" into ("John", "Doe"). If only one name, uses "." as last name.
func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "User", "."
	}
	if len(parts) == 1 {
		return parts[0], "."
	}
	return parts[0], strings.Join(parts[1:], " ")
}

// getPhoneOrDefault returns user's phone or default Ethiopian number format
func getPhoneOrDefault(phone *string) string {
	if phone != nil && *phone != "" {
		return *phone
	}
	return "0900000000" // Default placeholder for testing
}
