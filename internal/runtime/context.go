package runtime

import (
	"context"

	"github.com/kazyamaz200/agentos/internal/llm"
	"github.com/kazyamaz200/agentos/internal/profile"
	"github.com/kazyamaz200/agentos/internal/sandbox"
	"github.com/kazyamaz200/agentos/internal/task"
	"github.com/kazyamaz200/agentos/internal/tools"
	"github.com/kazyamaz200/agentos/internal/state"
)

type RunContext struct {
	Context   context.Context
	Task      *task.Task
	Profile   *profile.Profile
	LLM       llm.LLMClient
	Workspace *sandbox.Workspace
	Registry  *tools.Registry
	Logger    *state.Logger
	Store     *state.RunStore
	Config    *Config
	Iteration int
	MaxRetries int
}

type Config struct {
	DryRun  bool
	Verbose bool
}

func NewRunContext(ctx context.Context, tk *task.Task, r *Runtime) *RunContext {
	return &RunContext{
		Context:    ctx,
		Task:       tk,
		Profile:    r.Profile,
		LLM:        r.LLM,
		Workspace:  r.Workspace,
		Registry:   r.Registry,
		Logger:     r.Logger,
		Store:      r.Store,
		Config:     r.Config,
		MaxRetries: r.Profile.Limits.MaxRetries,
	}
}
