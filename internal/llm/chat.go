package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Chat sends one user message and returns the raw JSON response.
// TODO: take []Message, build a chatRequest instead of the map, decode into chatResponse and return a Message.
func (c *Client) Chat(content string) ([]byte, error) {
	// Request
	requestData := map[string]any{
		"model": c.model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": content,
			},
		},
	}

	jsonBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error while encoding request to JSON: %w", err)
	}

	// TODO: use http.NewRequestWithContext
	req, err := http.NewRequest("POST", c.baseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error while creating the request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending the request: %w", err)
	}
	defer resp.Body.Close()

	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(bodyResponse))
	}

	return bodyResponse, nil
}
