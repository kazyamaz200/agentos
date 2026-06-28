package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Config struct {
	BaseURL    string
	APIKey     string
	ModelCoder string
	Timeout    time.Duration
}

type LiteLLMClient struct {
	config Config
	http   *http.Client
}

func NewLiteLLMClient(config Config) *LiteLLMClient {
	if config.Timeout == 0 {
		config.Timeout = 5 * time.Minute
	}
	return &LiteLLMClient{
		config: config,
		http: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (c *LiteLLMClient) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.config.BaseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.config.APIKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("LLM API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return nil, fmt.Errorf("no choices in LLM response")
	}

	return &chatResp, nil
}

func (c *LiteLLMClient) ModelName() string {
	return c.config.ModelCoder
}

func (c *LiteLLMClient) Config() Config {
	return c.config
}

type MockLLMClient struct {
	Responses []ChatResponse
	Index     int
}

func NewMockLLMClient(responses []ChatResponse) *MockLLMClient {
	return &MockLLMClient{Responses: responses}
}

func (m *MockLLMClient) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	if m.Index >= len(m.Responses) {
		return nil, fmt.Errorf("no more mock responses")
	}
	resp := m.Responses[m.Index]
	m.Index++
	return &resp, nil
}

func (m *MockLLMClient) ModelName() string {
	return "mock-model"
}
