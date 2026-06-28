package safety

import (
	"path/filepath"
	"strings"
)

var secretPatterns = []string{
	".env",
	".env.*",
	"*.pem",
	"id_rsa",
	"id_rsa.pub",
	"id_ed25519",
	"id_ed25519.pub",
	".secret*",
	"*.key",
	".credentials*",
	".aws/credentials",
	".gcp/credentials*",
	".token*",
}

type SecretDetector struct {
	patterns []string
}

func NewSecretDetector() *SecretDetector {
	return &SecretDetector{patterns: secretPatterns}
}

func (s *SecretDetector) IsSecretFile(filePath string) bool {
	name := filepath.Base(filePath)
	for _, pattern := range s.patterns {
		if matched, _ := filepath.Match(pattern, name); matched {
			return true
		}
		if strings.Contains(pattern, "*") {
			continue
		}
		if name == pattern {
			return true
		}
	}
	return false
}
