package delivery

import (
	"encoding/json"
	"net/http"
	"strings"

	"event-management-backend/repository"
	"event-management-backend/utils"
)

type TagHandler struct {
	hasuraClient repository.HasuraClient
	jwtSecret    string
}

func NewTagHandler(hasuraClient repository.HasuraClient, jwtSecret string) *TagHandler {
	return &TagHandler{
		hasuraClient: hasuraClient,
		jwtSecret:    jwtSecret,
	}
}

func (h *TagHandler) LinkTags(w http.ResponseWriter, r *http.Request) {
	// Verify JWT
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondError(w, 401, "Authorization required")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	_, err := utils.ValidateJWT(tokenString, h.jwtSecret)
	if err != nil {
		respondError(w, 401, "Invalid token")
		return
	}

	var payload struct {
		Input struct {
			EventID  string   `json:"event_id"`
			TagNames []string `json:"tag_names"`
		} `json:"input"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, 400, "Invalid request")
		return
	}

	// Insert tags and get their IDs
	tagIDs := []string{}
	for _, tagName := range payload.Input.TagNames {
		query := `
			mutation InsertTag($name: String!) {
				insert_tags_one(
					object: {name: $name}
					on_conflict: {constraint: tags_name_key, update_columns: [name]}
				) {
					id
				}
			}
		`
		variables := map[string]interface{}{"name": tagName}
		
		var result struct {
			Data struct {
				InsertTagsOne struct {
					ID string `json:"id"`
				} `json:"insert_tags_one"`
			} `json:"data"`
		}
		
		if err := h.hasuraClient.Query(query, variables, &result); err == nil {
			if result.Data.InsertTagsOne.ID != "" {
				tagIDs = append(tagIDs, result.Data.InsertTagsOne.ID)
			}
		}
	}

	// Link tags to event
	for _, tagID := range tagIDs {
		query := `
			mutation LinkTag($event_id: uuid!, $tag_id: uuid!) {
				insert_event_tags_one(
					object: {event_id: $event_id, tag_id: $tag_id}
					on_conflict: {constraint: event_tags_pkey, update_columns: []}
				) {
					event_id
				}
			}
		`
		variables := map[string]interface{}{
			"event_id": payload.Input.EventID,
			"tag_id":   tagID,
		}
		
		h.hasuraClient.Query(query, variables, nil)
	}

	respondJSON(w, 200, map[string]interface{}{
		"linked_tags": len(tagIDs),
	})
}
