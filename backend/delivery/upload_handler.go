package delivery

import (
	"encoding/json"
	"event-management-backend/usecase"
	"net/http"
)

type UploadHandler struct {
	uploadUsecase *usecase.UploadUsecase
}

func NewUploadHandler(uploadUsecase *usecase.UploadUsecase) *UploadHandler {
	return &UploadHandler{uploadUsecase: uploadUsecase}
}

func (h *UploadHandler) UploadFiles(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form (32MB max)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		respondError(w, 400, "Failed to parse form")
		return
	}

	// Get user_id and event_id from form
	userID := r.FormValue("user_id")
	eventID := r.FormValue("event_id")

	if userID == "" || eventID == "" {
		respondError(w, 400, "user_id and event_id are required")
		return
	}

	// Get files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		respondError(w, 400, "No files provided")
		return
	}

	// Upload files
	result, err := h.uploadUsecase.UploadFiles(files, userID, eventID)
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
