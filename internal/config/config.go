package config

// Config stores configuration required by the CLI.
type Config struct {
	Version string
}

// NewConfig creates a new Config type.
func NewConfig(version string) Config {
	return Config{
		Version: version,
	}
}
