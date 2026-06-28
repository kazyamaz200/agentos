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

package task

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewLoader(t *testing.T) {
	t.Parallel()

	l := NewLoader()
	if l == nil {
		t.Fatal("NewLoader() returned nil")
	}
}

func TestLoader_Load(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "task.yaml")
	content := []byte(`id: loader-test
repo: org/repo
title: Loader test
`)
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}

	l := NewLoader()
	task, err := l.Load(path)
	if err != nil {
		t.Fatalf("Loader.Load() error = %v", err)
	}
	if task.ID != "loader-test" {
		t.Errorf("ID = %q, want %q", task.ID, "loader-test")
	}
	if task.Repo != "org/repo" {
		t.Errorf("Repo = %q, want %q", task.Repo, "org/repo")
	}
	if task.Title != "Loader test" {
		t.Errorf("Title = %q, want %q", task.Title, "Loader test")
	}
}

func TestLoader_Load_MissingFile(t *testing.T) {
	t.Parallel()

	l := NewLoader()
	_, err := l.Load("/nonexistent/path.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
