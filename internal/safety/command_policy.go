package safety

import (
	"strings"
)

type CommandPolicy struct {
	DenyCommands []string
	AllowList    []string
}

func NewCommandPolicy(denyCommands []string) *CommandPolicy {
	defaultDeny := []string{
		"rm -rf", "rm -rf /", "rm -rf /*",
		"sudo", "sudo ",
		"docker run --privileged",
		"curl", "wget",
		"scp", "ssh",
	}
	if len(denyCommands) > 0 {
		defaultDeny = append(defaultDeny, denyCommands...)
	}
	return &CommandPolicy{
		DenyCommands: defaultDeny,
	}
}

func (p *CommandPolicy) Check(command string) (bool, string) {
	cmdLower := strings.TrimSpace(strings.ToLower(command))
	for _, denied := range p.DenyCommands {
		if strings.Contains(cmdLower, strings.ToLower(denied)) {
			return false, denied
		}
	}
	return true, ""
}
