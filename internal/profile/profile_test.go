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

package profile

import (
	"testing"
)

func TestDefaultProfile(t *testing.T) {
	t.Parallel()

	p := DefaultProfile()
	if p.Name != "default" {
		t.Errorf("Name = %q, want %q", p.Name, "default")
	}
	if p.Role != "coding agent" {
		t.Errorf("Role = %q, want %q", p.Role, "coding agent")
	}
	if p.LLM.Provider != "litellm" {
		t.Errorf("LLM.Provider = %q, want %q", p.LLM.Provider, "litellm")
	}
	if p.LLM.Model != "coder" {
		t.Errorf("LLM.Model = %q, want %q", p.LLM.Model, "coder")
	}
	if p.LLM.Temperature != 0.2 {
		t.Errorf("LLM.Temperature = %f, want %f", p.LLM.Temperature, 0.2)
	}
	if p.LLM.MaxTokens != 8192 {
		t.Errorf("LLM.MaxTokens = %d, want %d", p.LLM.MaxTokens, 8192)
	}
	if p.Limits.MaxIterations != 8 {
		t.Errorf("Limits.MaxIterations = %d, want %d", p.Limits.MaxIterations, 8)
	}
	if p.Limits.MaxRetries != 3 {
		t.Errorf("Limits.MaxRetries = %d, want %d", p.Limits.MaxRetries, 3)
	}
	if p.Limits.MaxChangedFiles != 20 {
		t.Errorf("Limits.MaxChangedFiles = %d, want %d", p.Limits.MaxChangedFiles, 20)
	}
	if p.Limits.MaxRuntimeMinute != 30 {
		t.Errorf("Limits.MaxRuntimeMinute = %d, want %d", p.Limits.MaxRuntimeMinute, 30)
	}
	if p.Output.Mode != "patch" {
		t.Errorf("Output.Mode = %q, want %q", p.Output.Mode, "patch")
	}
}

func TestNewProfileService(t *testing.T) {
	t.Parallel()

	p := DefaultProfile()
	s := NewProfileService(&p)
	if s.Profile() != &p {
		t.Errorf("Profile() = %p, want %p", s.Profile(), &p)
	}
}

func TestIsToolAllowed_EmptyAllowAll(t *testing.T) {
	t.Parallel()

	p := DefaultProfile()
	s := NewProfileService(&p)

	if !s.IsToolAllowed("anything") {
		t.Error("expected all tools allowed when Allow is empty")
	}
	if !s.IsToolAllowed("") {
		t.Error("expected empty string allowed when Allow is empty")
	}
}

func TestIsToolAllowed_WithAllowList(t *testing.T) {
	t.Parallel()

	p := DefaultProfile()
	p.Tools.Allow = []string{"read", "write", "search"}
	s := NewProfileService(&p)

	if !s.IsToolAllowed("read") {
		t.Error("expected 'read' to be allowed")
	}
	if !s.IsToolAllowed("write") {
		t.Error("expected 'write' to be allowed")
	}
	if !s.IsToolAllowed("search") {
		t.Error("expected 'search' to be allowed")
	}
	if s.IsToolAllowed("shell") {
		t.Error("expected 'shell' to be denied")
	}
	if s.IsToolAllowed("git") {
		t.Error("expected 'git' to be denied")
	}
}

func TestDenyCommands(t *testing.T) {
	t.Parallel()

	p := DefaultProfile()
	p.Tools.DenyCommands = []string{"rm -rf", "sudo"}
	s := NewProfileService(&p)

	denied := s.DenyCommands()
	if len(denied) != 2 {
		t.Fatalf("len(DenyCommands()) = %d, want 2", len(denied))
	}
	if denied[0] != "rm -rf" {
		t.Errorf("DenyCommands()[0] = %q, want %q", denied[0], "rm -rf")
	}
	if denied[1] != "sudo" {
		t.Errorf("DenyCommands()[1] = %q, want %q", denied[1], "sudo")
	}
}

func TestDenyCommands_Empty(t *testing.T) {
	t.Parallel()

	p := DefaultProfile()
	s := NewProfileService(&p)

	denied := s.DenyCommands()
	if len(denied) != 0 {
		t.Errorf("len(DenyCommands()) = %d, want 0", len(denied))
	}
}
