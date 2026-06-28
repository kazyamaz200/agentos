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

package orchestrator

import (
	"testing"
)

func TestStrategy_Constants(t *testing.T) {
	t.Parallel()

	if StrategySequential != Strategy("sequential") {
		t.Errorf("StrategySequential = %q, want %q", StrategySequential, "sequential")
	}
	if StrategyParallel != Strategy("parallel") {
		t.Errorf("StrategyParallel = %q, want %q", StrategyParallel, "parallel")
	}
}

func TestStrategy_Sequential(t *testing.T) {
	t.Parallel()

	var s Strategy = "sequential"
	if s != StrategySequential {
		t.Errorf("s = %q, want %q", s, StrategySequential)
	}
}

func TestStrategy_Parallel(t *testing.T) {
	t.Parallel()

	var s Strategy = "parallel"
	if s != StrategyParallel {
		t.Errorf("s = %q, want %q", s, StrategyParallel)
	}
}

func TestSubtask_Defaults(t *testing.T) {
	t.Parallel()

	st := Subtask{}
	if st.ID != "" {
		t.Errorf("ID = %q, want empty", st.ID)
	}
	if st.Description != "" {
		t.Errorf("Description = %q, want empty", st.Description)
	}
	if st.AgentName != "" {
		t.Errorf("AgentName = %q, want empty", st.AgentName)
	}
	if st.Deps != nil {
		t.Errorf("Deps = %v, want nil", st.Deps)
	}
}

func TestSubtaskResult_Defaults(t *testing.T) {
	t.Parallel()

	sr := SubtaskResult{}
	if sr.SubtaskID != "" {
		t.Errorf("SubtaskID = %q, want empty", sr.SubtaskID)
	}
	if sr.Output != "" {
		t.Errorf("Output = %q, want empty", sr.Output)
	}
	if sr.Error != "" {
		t.Errorf("Error = %q, want empty", sr.Error)
	}
	if sr.Success {
		t.Error("Success should be false")
	}
}

func TestTaskPlan_Defaults(t *testing.T) {
	t.Parallel()

	tp := TaskPlan{}
	if tp.Description != "" {
		t.Errorf("Description = %q, want empty", tp.Description)
	}
	if tp.Subtasks != nil {
		t.Errorf("Subtasks = %v, want nil", tp.Subtasks)
	}
}

func TestMergeResults(t *testing.T) {
	t.Parallel()

	o := &Orchestrator{}
	results := []SubtaskResult{
		{SubtaskID: "step-1", Output: "done", Success: true},
	}
	merged := o.MergeResults(results)
	if merged == "" {
		t.Error("MergeResults returned empty string")
	}
}
