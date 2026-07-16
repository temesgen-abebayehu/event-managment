package delivery

import (
	"encoding/json"
	"event-management-backend/domain"
	"event-management-backend/usecase"
	"net/http"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Input domain.SignupRequest `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, 400, "Invalid request")
		return
	}

	result, err := h.authUsecase.Signup(payload.Input)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}

	respondJSON(w, 200, result)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Input domain.LoginRequest `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, 400, "Invalid request")
		return
	}

	result, err := h.authUsecase.Login(payload.Input)
	if err != nil {
		respondError(w, 401, err.Error())
		return
	}

	respondJSON(w, 200, result)
}

func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func respondJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
