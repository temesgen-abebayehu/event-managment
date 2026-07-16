package hasura

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	endpoint    string
	adminSecret string
}

func NewClient(endpoint, adminSecret string) *Client {
	return &Client{
		endpoint:    endpoint,
		adminSecret: adminSecret,
	}
}

func (c *Client) Mutate(query string, variables map[string]interface{}, result interface{}) error {
	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", c.adminSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response struct {
		Data   json.RawMessage `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if len(response.Errors) > 0 {
		return fmt.Errorf(response.Errors[0].Message)
	}

	return json.Unmarshal(response.Data, result)
}
