package agent

import (
	"github.com/kazyamaz200/agentos/internal/runtime"
)

type Agent interface {
	Name() string
	Plan(ctx *runtime.RunContext) (*runtime.Plan, error)
	Execute(ctx *runtime.RunContext, plan *runtime.Plan) (*runtime.ExecutionResult, error)
	Review(ctx *runtime.RunContext, result *runtime.ExecutionResult) (*runtime.ReviewResult, error)
}

type BaseAgent struct {
	name string
}

func NewBaseAgent(name string) *BaseAgent {
	return &BaseAgent{name: name}
}

func (a *BaseAgent) Name() string { return a.name }
