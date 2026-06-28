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

package sandbox

import (
	"testing"
)

func TestNewPolicy_ReturnsDefaultValues(t *testing.T) {
	t.Parallel()

	p := NewPolicy()
	if p == nil {
		t.Fatal("NewPolicy returned nil")
	}

	if p.MaxFileSize != 1024*1024 {
		t.Errorf("MaxFileSize = %d, want %d", p.MaxFileSize, 1024*1024)
	}

	expectedPatterns := []string{
		".env", ".env.*",
		"*.pem", "id_rsa", "id_ed25519",
	}
	if len(p.DenyWritePatterns) != len(expectedPatterns) {
		t.Fatalf("got %d deny patterns, want %d", len(p.DenyWritePatterns), len(expectedPatterns))
	}
	for i, pattern := range p.DenyWritePatterns {
		if pattern != expectedPatterns[i] {
			t.Errorf("DenyWritePatterns[%d] = %q, want %q", i, pattern, expectedPatterns[i])
		}
	}
}

func TestPolicy_StructDefaults(t *testing.T) {
	t.Parallel()

	p := Policy{}
	if p.MaxFileSize != 0 {
		t.Errorf("zero value MaxFileSize = %d, want 0", p.MaxFileSize)
	}
	if p.DenyWritePatterns != nil {
		t.Errorf("zero value DenyWritePatterns = %v, want nil", p.DenyWritePatterns)
	}
}

func TestPolicy_NewPolicyHasCorrectMaxFileSize(t *testing.T) {
	t.Parallel()

	p := NewPolicy()
	if p.MaxFileSize != 1024*1024 {
		t.Errorf("MaxFileSize = %d, want %d", p.MaxFileSize, 1024*1024)
	}
}

func TestPolicy_NewPolicyHasDenyPatterns(t *testing.T) {
	t.Parallel()

	p := NewPolicy()
	if len(p.DenyWritePatterns) == 0 {
		t.Error("DenyWritePatterns should not be empty")
	}

	found := false
	for _, pattern := range p.DenyWritePatterns {
		if pattern == ".env" {
			found = true
			break
		}
	}
	if !found {
		t.Error(`expected ".env" in DenyWritePatterns`)
	}
}
