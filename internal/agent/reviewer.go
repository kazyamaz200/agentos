package agent

import (
	"context"
	"fmt"

	"github.com/kazyamaz200/agentos/internal/llm"
	"github.com/kazyamaz200/agentos/internal/runtime"
)

type Reviewer struct {
	llm llm.LLMClient
}

func NewReviewer(llmClient llm.LLMClient) *Reviewer {
	return &Reviewer{llm: llmClient}
}

func (r *Reviewer) Review(ctx *runtime.RunContext, result *runtime.ExecutionResult) (*runtime.ReviewResult, error) {
	if result.Diff == "" {
		return &runtime.ReviewResult{Approved: true, Summary: "No changes to review"}, nil
	}

	systemMsg := llm.Message{Role: llm.RoleSystem, Content: llm.SystemPromptReviewer}
	userMsg := llm.Message{
		Role: llm.RoleUser,
		Content: fmt.Sprintf(`Review the following diff for task: %s

Description: %s

Diff:
%s`, ctx.Task.Title, ctx.Task.Description, result.Diff),
	}

	resp, err := r.llm.Chat(context.Background(), llm.ChatRequest{
		Model:       r.llm.ModelName(),
		Messages:    []llm.Message{systemMsg, userMsg},
		Temperature: 0.1,
		MaxTokens:   ctx.Profile.LLM.MaxTokens,
	})
	if err != nil {
		return nil, fmt.Errorf("LLM review request: %w", err)
	}

	return runtime.ParseReview(resp)
}
