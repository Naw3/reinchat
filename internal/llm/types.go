package llm

// Message represents a single message in a conversation with the LLM.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type chatResponse struct {
	Choices []choice `json:"choices"`
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}
