package agent

import (
	"context"
	"fmt"

	"github.com/kazyamaz200/agentos/internal/llm"
	"github.com/kazyamaz200/agentos/internal/runtime"
)

type Planner struct {
	llm llm.LLMClient
}

func NewPlanner(llmClient llm.LLMClient) *Planner {
	return &Planner{llm: llmClient}
}

func (p *Planner) Plan(ctx *runtime.RunContext) (*runtime.Plan, error) {
	systemMsg := llm.Message{Role: llm.RoleSystem, Content: llm.SystemPromptPlanner}
	userMsg := llm.Message{
		Role: llm.RoleUser,
		Content: fmt.Sprintf(`Task: %s
Description:
%s

Repository: %s
Base branch: %s

Create a plan to implement this task.`, ctx.Task.Title, ctx.Task.Description, ctx.Task.Repo, ctx.Task.BaseBranch),
	}

	resp, err := p.llm.Chat(context.Background(), llm.ChatRequest{
		Model:       p.llm.ModelName(),
		Messages:    []llm.Message{systemMsg, userMsg},
		Temperature: ctx.Profile.LLM.Temperature,
		MaxTokens:   ctx.Profile.LLM.MaxTokens,
	})
	if err != nil {
		return nil, fmt.Errorf("LLM plan request: %w", err)
	}

	return runtime.ParsePlan(resp)
}
