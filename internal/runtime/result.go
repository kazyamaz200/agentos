package runtime

import (
	"time"
)

type Plan struct {
	Summary             string  `json:"plan_summary"`
	Steps               []Step  `json:"steps"`
	EstimatedFilesChanged int   `json:"estimated_files_changed"`
}

type Step struct {
	StepNumber  int      `json:"step_number"`
	Action      string   `json:"action"`
	Description string   `json:"description"`
	TargetFiles []string `json:"target_files"`
	Reasoning   string   `json:"reasoning"`
}

type ExecutionResult struct {
	StepResults []StepResult `json:"step_results"`
	Diff        string       `json:"diff"`
	Success     bool         `json:"success"`
	Error       string       `json:"error,omitempty"`
}

type StepResult struct {
	StepNumber int    `json:"step_number"`
	Action     string `json:"action"`
	Success    bool   `json:"success"`
	Output     string `json:"output"`
	Error      string `json:"error,omitempty"`
	Duration   time.Duration `json:"duration"`
}

type ReviewResult struct {
	Approved bool          `json:"approved"`
	Issues   []ReviewIssue `json:"issues"`
	Summary  string        `json:"summary"`
}

type ReviewIssue struct {
	Severity string `json:"severity"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
}
