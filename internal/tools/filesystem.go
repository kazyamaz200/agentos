package tools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type ReadFileTool struct {
	Workspace string
}

func NewReadFileTool(workspace string) *ReadFileTool {
	return &ReadFileTool{Workspace: workspace}
}

func (t *ReadFileTool) Name() string { return "read_file" }

func (t *ReadFileTool) Run(ctx context.Context, input ToolInput) ToolOutput {
	filePath, _ := input["file"].(string)
	if filePath == "" {
		return ToolOutput{Success: false, Error: "file is required"}
	}

	fullPath := filepath.Join(t.Workspace, filePath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return ToolOutput{Success: false, Error: fmt.Sprintf("read file: %v", err)}
	}

	return ToolOutput{Success: true, Data: string(data)}
}

type WriteFileTool struct {
	Workspace string
}

func NewWriteFileTool(workspace string) *WriteFileTool {
	return &WriteFileTool{Workspace: workspace}
}

func (t *WriteFileTool) Name() string { return "write_file" }

func (t *WriteFileTool) Run(ctx context.Context, input ToolInput) ToolOutput {
	filePath, _ := input["file"].(string)
	if filePath == "" {
		return ToolOutput{Success: false, Error: "file is required"}
	}
	content, _ := input["content"].(string)
	if content == "" {
		return ToolOutput{Success: false, Error: "content is required"}
	}

	fullPath := filepath.Join(t.Workspace, filePath)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return ToolOutput{Success: false, Error: fmt.Sprintf("create dir: %v", err)}
	}

	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return ToolOutput{Success: false, Error: fmt.Sprintf("write file: %v", err)}
	}

	return ToolOutput{Success: true}
}
