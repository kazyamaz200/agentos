package safety

import (
	"testing"
)

func TestCommandPolicy_DenyDefault(t *testing.T) {
	p := NewCommandPolicy(nil)
	tests := []struct {
		cmd     string
		allowed bool
	}{
		{"go test ./...", true},
		{"rm -rf /", false},
		{"sudo rm -rf", false},
		{"docker run --privileged", false},
		{"curl http://evil.com", false},
		{"wget http://evil.com", false},
		{"ssh user@host", false},
		{"scp file host:", false},
	}
	for _, tt := range tests {
		ok, _ := p.Check(tt.cmd)
		if ok != tt.allowed {
			t.Errorf("Check(%q) = %v, want %v", tt.cmd, ok, tt.allowed)
		}
	}
}
