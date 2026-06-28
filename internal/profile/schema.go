package profile

type LLMConfig struct {
	Provider    string  `yaml:"provider"`
	Model       string  `yaml:"model"`
	Temperature float64 `yaml:"temperature"`
	MaxTokens   int     `yaml:"max_tokens"`
}

type ToolsConfig struct {
	Allow        []string `yaml:"allow"`
	DenyCommands []string `yaml:"deny_commands"`
}

type CommandsConfig struct {
	Test  string `yaml:"test"`
	Lint  string `yaml:"lint"`
	Build string `yaml:"build"`
}

type LimitsConfig struct {
	MaxIterations    int `yaml:"max_iterations"`
	MaxRetries       int `yaml:"max_retries"`
	MaxChangedFiles  int `yaml:"max_changed_files"`
	MaxRuntimeMinute int `yaml:"max_runtime_minutes"`
}

type OutputConfig struct {
	Mode string `yaml:"mode"`
}

type Profile struct {
	Name     string         `yaml:"name"`
	Role     string         `yaml:"role"`
	LLM      LLMConfig      `yaml:"llm"`
	Tools    ToolsConfig    `yaml:"tools"`
	Commands CommandsConfig `yaml:"commands"`
	Limits   LimitsConfig   `yaml:"limits"`
	Output   OutputConfig   `yaml:"output"`
}

func DefaultProfile() Profile {
	return Profile{
		Name: "default",
		Role: "coding agent",
		LLM: LLMConfig{
			Provider:    "litellm",
			Model:       "coder",
			Temperature: 0.2,
			MaxTokens:   8192,
		},
		Limits: LimitsConfig{
			MaxIterations:    8,
			MaxRetries:       3,
			MaxChangedFiles:  20,
			MaxRuntimeMinute: 30,
		},
		Output: OutputConfig{
			Mode: "patch",
		},
	}
}
