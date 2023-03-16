package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config stores configuration required by the CLI.
type Config struct {
	Path    string `validate:"required,dir"`
	Version string
}

// CLIName is the name of the project.
const CLIName = "atlas"

// NewConfig creates a new Config type.
func NewConfig(version string) (*Config, error) {
	viper.SetConfigName(configName())
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to get user home directory path: %w", err)
	}

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok { //nolint:errorlint
		return nil, fmt.Errorf("config does not exist (%s): %w", homeDir, err)
	}
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshalling viper config to struct: %w", err)
	}

	config.Version = version

	validate := validator.New()

	err = validate.Struct(config)
	if err != nil {
		return nil, fmt.Errorf("config is invalid: %w", err)
	}

	return &config, nil
}

func configName() string {
	return fmt.Sprintf(".%s-notes.yml", CLIName)
}
