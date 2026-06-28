package agent

type RetryConfig struct {
	MaxRetries int
}

type RetryHandler struct {
	config RetryConfig
}

func NewRetryHandler(config RetryConfig) *RetryHandler {
	return &RetryHandler{config: config}
}

func (h *RetryHandler) ShouldRetry(attempt int, testFailed bool, lintFailed bool) bool {
	if attempt >= h.config.MaxRetries {
		return false
	}
	return testFailed || lintFailed
}
