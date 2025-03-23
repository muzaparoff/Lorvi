package ai

import (
	"fmt"
)

// AIClient is an interface for interacting with AI backends.
type AIClient interface {
	// Ask sends a question to the AI backend and returns a response.
	Ask(question string) (string, error)
}

// NewClient returns an AIClient based on the backend string.
func NewClient(backend string) AIClient {
	switch backend {
	case "ollama":
		return &OllamaClient{}
	case "openai":
		return &OpenAIClient{}
	case "claude":
		return &ClaudeClient{}
	case "gemini":
		return &GeminiClient{}
	default:
		fmt.Printf("Unknown AI backend '%s', defaulting to Ollama.\n", backend)
		return &OllamaClient{}
	}
}

// OllamaClient is a stub implementation for the Ollama AI backend.
type OllamaClient struct{}

func (c *OllamaClient) Ask(question string) (string, error) {
	// TODO: Implement integration with Ollama
	return "Ollama response for: " + question, nil
}

// OpenAIClient is a stub implementation for the OpenAI backend.
type OpenAIClient struct{}

func (c *OpenAIClient) Ask(question string) (string, error) {
	// TODO: Implement integration with OpenAI API
	return "OpenAI response for: " + question, nil
}

// ClaudeClient is a stub implementation for the Claude backend.
type ClaudeClient struct{}

func (c *ClaudeClient) Ask(question string) (string, error) {
	// TODO: Implement integration with Anthropic Claude
	return "Claude response for: " + question, nil
}

// GeminiClient is a stub implementation for the Gemini backend.
type GeminiClient struct{}

func (c *GeminiClient) Ask(question string) (string, error) {
	// TODO: Implement integration with Google Gemini
	return "Gemini response for: " + question, nil
}
