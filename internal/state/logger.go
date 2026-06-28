package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Component string    `json:"component"`
	Message   string    `json:"message"`
	Data      any       `json:"data,omitempty"`
}

type ToolLogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Tool      string    `json:"tool"`
	Input     any       `json:"input"`
	Output    any       `json:"output"`
	Duration  string    `json:"duration"`
}

type LLMLogEntry struct {
	Timestamp  time.Time `json:"timestamp"`
	Request    any       `json:"request"`
	Response   any       `json:"response"`
	Model      string    `json:"model"`
	Duration   string    `json:"duration"`
	PromptTokens int     `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
}

type Logger struct {
	runDir string
}

func NewLogger(runDir string) *Logger {
	return &Logger{runDir: runDir}
}

func (l *Logger) Log(level, component, message string, data any) error {
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Component: component,
		Message:   message,
		Data:      data,
	}
	return l.appendJSON("run.log", entry)
}

func (l *Logger) LogTool(tool string, input, output any, duration time.Duration) error {
	entry := ToolLogEntry{
		Timestamp: time.Now(),
		Tool:      tool,
		Input:     input,
		Output:    output,
		Duration:  duration.String(),
	}
	return l.appendJSON("tool_log.jsonl", entry)
}

func (l *Logger) LogLLM(req, resp any, model string, duration time.Duration, promptTokens, completionTokens int) error {
	entry := LLMLogEntry{
		Timestamp:        time.Now(),
		Request:          req,
		Response:         resp,
		Model:            model,
		Duration:         duration.String(),
		PromptTokens:     promptTokens,
		CompletionTokens: completionTokens,
	}
	return l.appendJSON("llm_log.jsonl", entry)
}

func (l *Logger) appendJSON(filename string, entry any) error {
	path := filepath.Join(l.runDir, filename)
	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("marshal log entry: %w", err)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open log file: %w", err)
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("write log: %w", err)
	}
	if _, err := f.Write([]byte("\n")); err != nil {
		return fmt.Errorf("write newline: %w", err)
	}
	return nil
}
