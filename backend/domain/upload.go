package domain

type UploadRequest struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
}

type UploadedFile struct {
	URL      string `json:"url"`
	PublicID string `json:"public_id"`
}

type UploadResponse struct {
	Files []UploadedFile `json:"files"`
}
