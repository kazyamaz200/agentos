// Copyright 2026 AgentOS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tools provides the tool abstraction for agent actions including file
// I/O, shell commands, git operations, code search, and test execution.
package tools

import "context"

// ToolInput is a flexible input map passed to a Tool when it is Run.
type ToolInput map[string]interface{}

// ToolOutput is the result returned by a Tool after execution.
type ToolOutput struct {
	Success  bool
	Data     interface{}
	Error    string
}

// Tool is the interface that all agent tools must implement.
type Tool interface {
	Name() string
	Run(ctx context.Context, input ToolInput) ToolOutput
}

// Registry holds a set of named Tool instances and provides lookup and
// registration operations.
type Registry struct {
	tools map[string]Tool
}

// NewRegistry creates a new empty Registry.
func NewRegistry() *Registry {
	return &Registry{tools: make(map[string]Tool)}
}

// Register adds a Tool to the registry under its Name.
func (r *Registry) Register(t Tool) {
	r.tools[t.Name()] = t
}

// Get returns the Tool associated with name, and a boolean indicating whether
// it was found.
func (r *Registry) Get(name string) (Tool, bool) {
	t, ok := r.tools[name]
	return t, ok
}

// List returns the names of all registered tools.
func (r *Registry) List() []string {
	var names []string
	for n := range r.tools {
		names = append(names, n)
	}
	return names
}
