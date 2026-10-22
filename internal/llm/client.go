package llm

import "net/http"

// Client sends requests to an OpenAI-compatible API.
type Client struct {
	baseURL string
	apiKey  string
	model   string
	http    *http.Client
}

// NewClient creates a client for the API at baseURL.
func NewClient(baseURL, apiKey, model string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		model:   model,
		http:    &http.Client{}, // TODO: add a timeout
	}
}
