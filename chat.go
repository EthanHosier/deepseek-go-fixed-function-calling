package deepseek

import (
	"errors"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
)

var (
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method")
	ErrUnexpectedResponseFormat         = errors.New("unexpected response format")
)

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type StreamOptions struct {
	IncludeUsage bool
}

type Parameters struct {
	Type       string                 `json:"type"` // Type of the parameters, e.g., "object" (required)
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
}

// Function defines the structure of a function tool
type Function struct {
	Name        string      `json:"name"`                 // The name of the function (required)
	Description string      `json:"description"`          // Description of the function (required)
	Parameters  *Parameters `json:"parameters,omitempty"` // Parameters schema (optional)
}

// Tool defines the structure for a tool
type Tools struct {
	Type     string   `json:"type"`     // Type of the tool, e.g., "function" (required)
	Function Function `json:"function"` // The function details (required)
}

type ResponseFormat struct {
	Type string `json:"type"` //either text or json_object. If json_object, please mention "json" anywhere in your prompt.
}

// make a different struct for streaming with streaming options parameter
type ChatCompletionRequest struct {
	Model            string                  `json:"model"`                       // Required: Model ID, e.g., "deepseek-chat"
	Messages         []ChatCompletionMessage `json:"messages"`                    // Required: List of messages
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"` // Optional: Frequency penalty, >= -2 and <= 2
	MaxTokens        int                     `json:"max_tokens,omitempty"`        // Optional: Maximum tokens, > 1
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`  // Optional: Presence penalty, >= -2 and <= 2
	Temperature      float32                 `json:"temperature,omitempty"`       // Optional: Sampling temperature, <= 2
	TopP             float32                 `json:"top_p,omitempty"`             // Optional: Nucleus sampling parameter, <= 1
	ResponseFormat   *ResponseFormat         `json:"response_format,omitempty"`   // Optional: Custom response format
	Stop             []string                `json:"stop,omitempty"`              // Optional: Stop signals
	Tools            []Tools                 `json:"tools,omitempty"`             // Optional: List of tools
	LogProbs         bool                    `json:"logprobs,omitempty"`          // Optional: Enable log probabilities
	TopLogProbs      int                     `json:"top_logprobs,omitempty"`      // Optional: Number of top tokens with log probabilities, <= 20
}
