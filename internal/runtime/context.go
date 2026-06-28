package runtime

import (
	"github.com/kazyamaz200/agentos/internal/llm"
	"github.com/kazyamaz200/agentos/internal/profile"
	"github.com/kazyamaz200/agentos/internal/sandbox"
	"github.com/kazyamaz200/agentos/internal/task"
	"github.com/kazyamaz200/agentos/internal/tools"
	"github.com/kazyamaz200/agentos/internal/state"
)

type RunContext struct {
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
	DryRun bool
	Verbose bool
}
