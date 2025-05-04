package config

// Config holds the application configuration
type Config struct {
	WindowWidth  int
	WindowHeight int
	Title        string
}

// New creates a new configuration with default values
func New() *Config {
	return &Config{
		WindowWidth:  800,
		WindowHeight: 600,
		Title:        "Catch My File",
	}
}
