package deepseek

import (
	"bufio"
	"context"
	"fmt"

	utils "github.com/EthanHosier/deepseek-go-fixed-function-calling/utils"
)

type Client struct {
	authToken string
	baseURL   string
}

// NewClient creates a new client with an authentication token.
func NewClient(authToken string) *Client {
	return &Client{
		authToken: authToken,
		baseURL:   "https://api.deepseek.com/",
	}
}

// CreateChatCompletion sends a chat completion request and returns the generated response.
func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request *ChatCompletionRequest,
) (*utils.ChatCompletionResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	req, err := utils.NewRequestBuilder(c.authToken).
		SetBaseURL(c.baseURL).
		SetPath("chat/completions").
		SetBodyFromStruct(request).
		Build(ctx)

	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	resp, err := utils.SendRequest(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, HandleAPIError(resp)
	}

	updatedResp, err := utils.HandleResponse(resp)

	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return updatedResp, err
}

// CreateStreamChatCompletion send a chat completion request with stream = true and returns the delta
func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request *StreamChatCompletionRequest,
) (ChatCompletionStream, error) {

	request.Stream = true
	req, err := utils.NewRequestBuilder(c.authToken).
		SetBaseURL(c.baseURL).
		SetPath("chat/completions/").
		SetBodyFromStruct(request).
		Build(ctx)

	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	resp, err := utils.SendRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, HandleAPIError(resp)
	}

	ctx, cancel := context.WithCancel(ctx)
	stream := &chatCompletionStream{
		ctx:    ctx,
		cancel: cancel,
		resp:   resp,
		reader: bufio.NewReader(resp.Body),
	}
	return stream, nil
}
