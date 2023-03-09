package config

// Config stores configuration required by the CLI.
type Config struct {
	Version string
}

// CLIName is the name of the project.
const CLIName = "atlas"

// NewConfig creates a new Config type.
func NewConfig(version string) Config {
	return Config{
		Version: version,
	}
}
