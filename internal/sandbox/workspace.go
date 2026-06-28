package sandbox

import (
	"fmt"
	"os"
	"path/filepath"
)

type Workspace struct {
	RootDir   string
	RunsDir   string
	TaskID    string
	RunDir    string
}

func NewWorkspace(rootDir string) *Workspace {
	return &Workspace{RootDir: rootDir}
}

func (w *Workspace) PrepareRun(taskID string) error {
	w.TaskID = taskID
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	w.RunsDir = filepath.Join(homeDir, ".agentos", "runs")
	w.RunDir = filepath.Join(w.RunsDir, taskID)

	if err := os.MkdirAll(w.RunDir, 0755); err != nil {
		return fmt.Errorf("create run dir: %w", err)
	}
	return nil
}

func (w *Workspace) RunPath() string {
	return w.RunDir
}

func (w *Workspace) SaveFile(name string, data []byte) error {
	path := filepath.Join(w.RunDir, name)
	return os.WriteFile(path, data, 0644)
}

func (w *Workspace) AbsPath(relative string) string {
	return filepath.Join(w.RootDir, relative)
}

func (w *Workspace) RepoAbsPath(repoRelative string) string {
	if filepath.IsAbs(repoRelative) {
		return repoRelative
	}
	return filepath.Join(w.RootDir, repoRelative)
}
