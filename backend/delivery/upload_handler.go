package delivery

import (
	"encoding/json"
	"event-management-backend/usecase"
	"event-management-backend/utils"
	"net/http"
	"strings"
)

type UploadHandler struct {
	uploadUsecase *usecase.UploadUsecase
	jwtSecret     string
}

func NewUploadHandler(uploadUsecase *usecase.UploadUsecase, jwtSecret string) *UploadHandler {
	return &UploadHandler{uploadUsecase: uploadUsecase, jwtSecret: jwtSecret}
}

func (h *UploadHandler) UploadFiles(w http.ResponseWriter, r *http.Request) {
	// Verify JWT token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondError(w, 401, "Authorization header required")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondError(w, 401, "Bearer token required")
		return
	}

	claims, err := utils.ValidateJWT(tokenString, h.jwtSecret)
	if err != nil {
		respondError(w, 401, "Invalid or expired token")
		return
	}

	// Extract authenticated user ID from JWT claims
	authenticatedUserID, err := utils.ExtractUserIDFromJWT(claims)
	if err != nil {
		respondError(w, 401, "Invalid token claims")
		return
	}

	// Parse multipart form (32MB max)
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		respondError(w, 400, "Failed to parse form")
		return
	}

	// Get event_id from form — user_id comes from the verified JWT
	eventID := r.FormValue("event_id")

	if eventID == "" {
		respondError(w, 400, "event_id is required")
		return
	}

	// Get files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		respondError(w, 400, "No files provided")
		return
	}

	// Upload files using the authenticated user's ID (not from form body)
	result, err := h.uploadUsecase.UploadFiles(files, authenticatedUserID, eventID)
	if err != nil {
		respondError(w, 400, err.Error())
		return
	}

	respondJSON(w, 200, result)
}

func (h *UploadHandler) DeleteFiles(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Input struct {
			PublicIDs []string `json:"public_ids"`
		} `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, 400, "Invalid request")
		return
	}

	if len(payload.Input.PublicIDs) == 0 {
		respondError(w, 400, "No public_ids provided")
		return
	}

	err := h.uploadUsecase.DeleteFiles(payload.Input.PublicIDs)
	if err != nil {
		respondError(w, 500, "Failed to delete files")
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Files deleted successfully"})
}
