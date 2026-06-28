package tools

import "context"

type ToolInput map[string]interface{}
type ToolOutput struct {
	Success  bool
	Data     interface{}
	Error    string
}

type Tool interface {
	Name() string
	Run(ctx context.Context, input ToolInput) ToolOutput
}

type Registry struct {
	tools map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{tools: make(map[string]Tool)}
}

func (r *Registry) Register(t Tool) {
	r.tools[t.Name()] = t
}

func (r *Registry) Get(name string) (Tool, bool) {
	t, ok := r.tools[name]
	return t, ok
}

func (r *Registry) List() []string {
	var names []string
	for n := range r.tools {
		names = append(names, n)
	}
	return names
}
