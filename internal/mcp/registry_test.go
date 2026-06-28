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

package mcp

import (
	"testing"

	"github.com/kazyamaz200/agentos/internal/tools"
)

func TestToolAdapter_Name(t *testing.T) {
	t.Parallel()

	adapter := &ToolAdapter{
		def: ToolDefinition{Name: "my_tool"},
	}
	if adapter.Name() != "mcp_my_tool" {
		t.Errorf("Name() = %q, want %q", adapter.Name(), "mcp_my_tool")
	}
}

func TestToolAdapter_Name_PreservesPrefix(t *testing.T) {
	t.Parallel()

	adapter := &ToolAdapter{
		def: ToolDefinition{Name: "read_file"},
	}
	if adapter.Name() != "mcp_read_file" {
		t.Errorf("Name() = %q, want %q", adapter.Name(), "mcp_read_file")
	}
}

func TestRegistry_AddAndList(t *testing.T) {
	t.Parallel()

	reg := tools.NewRegistry()
	adapter := &ToolAdapter{
		def: ToolDefinition{
			Name:        "test_tool",
			Description: "A test tool",
			InputSchema: map[string]interface{}{"type": "object"},
		},
	}

	reg.Register(adapter)
	names := reg.List()

	if len(names) != 1 {
		t.Fatalf("got %d tools, want 1", len(names))
	}
	if names[0] != "mcp_test_tool" {
		t.Errorf("tool name = %q, want %q", names[0], "mcp_test_tool")
	}

	tool, ok := reg.Get("mcp_test_tool")
	if !ok {
		t.Fatal("tool not found in registry")
	}
	if tool.Name() != "mcp_test_tool" {
		t.Errorf("tool.Name() = %q, want %q", tool.Name(), "mcp_test_tool")
	}
}

func TestRegistry_AddMultiple(t *testing.T) {
	t.Parallel()

	reg := tools.NewRegistry()
	reg.Register(&ToolAdapter{def: ToolDefinition{Name: "tool_a"}})
	reg.Register(&ToolAdapter{def: ToolDefinition{Name: "tool_b"}})
	reg.Register(&ToolAdapter{def: ToolDefinition{Name: "tool_c"}})

	names := reg.List()
	if len(names) != 3 {
		t.Fatalf("got %d tools, want 3", len(names))
	}
}

func TestRegistry_AddDuplicateOverwrites(t *testing.T) {
	t.Parallel()

	reg := tools.NewRegistry()
	reg.Register(&ToolAdapter{def: ToolDefinition{Name: "same"}})
	reg.Register(&ToolAdapter{def: ToolDefinition{Name: "same"}})

	names := reg.List()
	if len(names) != 1 {
		t.Errorf("got %d tools, want 1 (duplicate should overwrite)", len(names))
	}
}
