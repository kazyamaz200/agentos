package sandbox

type Policy struct {
	DenyWritePatterns []string
	MaxFileSize       int64
}

func NewPolicy() *Policy {
	return &Policy{
		DenyWritePatterns: []string{
			".env", ".env.*",
			"*.pem", "id_rsa", "id_ed25519",
		},
		MaxFileSize: 1024 * 1024,
	}
}
