package delivery

import (
	"encoding/json"
	"event-management-backend/domain"
	"event-management-backend/usecase"
	"net/http"
)

type PaymentHandler struct {
	paymentUsecase *usecase.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase *usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUsecase: paymentUsecase}
}

func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Input           domain.InitiatePaymentRequest `json:"input"`
		SessionVariables struct {
			UserID string `json:"x-hasura-user-id"`
		} `json:"session_variables"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, 400, "Invalid request")
		return
	}

	userID := payload.SessionVariables.UserID
	if userID == "" {
		respondError(w, 401, "User not authenticated")
		return
	}

	result, err := h.paymentUsecase.InitiatePayment(userID, payload.Input)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}

	respondJSON(w, 200, result)
}

func (h *PaymentHandler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Input domain.VerifyPaymentRequest `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, 400, "Invalid request")
		return
	}

	result, err := h.paymentUsecase.VerifyPayment(payload.Input.TxRef)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}

	respondJSON(w, 200, result)
}

func (h *PaymentHandler) PaymentCallback(w http.ResponseWriter, r *http.Request) {
	var callback struct {
		TxRef  string `json:"tx_ref"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		respondError(w, 400, "Invalid callback")
		return
	}

	err := h.paymentUsecase.HandleCallback(callback.TxRef, callback.Status)
	if err != nil {
		respondError(w, 500, "Failed to process callback")
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Callback processed"})
}
