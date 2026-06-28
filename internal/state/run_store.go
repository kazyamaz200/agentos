package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type RunStatus string

const (
	RunStatusPending    RunStatus = "pending"
	RunStatusPlanning   RunStatus = "planning"
	RunStatusExecuting  RunStatus = "executing"
	RunStatusTesting    RunStatus = "testing"
	RunStatusReviewing  RunStatus = "reviewing"
	RunStatusCompleted  RunStatus = "completed"
	RunStatusFailed     RunStatus = "failed"
	RunStatusCancelled  RunStatus = "cancelled"
)

type RunRecord struct {
	TaskID      string    `json:"task_id"`
	Status      RunStatus `json:"status"`
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at,omitempty"`
	ProfileName string    `json:"profile_name"`
	Branch      string    `json:"branch"`
	Error       string    `json:"error,omitempty"`
	Iteration   int       `json:"iteration"`
}

type RunStore struct {
	runDir string
}

func NewRunStore(runDir string) *RunStore {
	return &RunStore{runDir: runDir}
}

func (s *RunStore) Save(record *RunRecord) error {
	path := filepath.Join(s.runDir, "run_state.json")
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal run record: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

func (s *RunStore) Load() (*RunRecord, error) {
	path := filepath.Join(s.runDir, "run_state.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read run record: %w", err)
	}
	var record RunRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, fmt.Errorf("parse run record: %w", err)
	}
	return &record, nil
}
