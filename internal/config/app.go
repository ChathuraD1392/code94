package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// App represents the root application configuration structure.
type App struct {
	Server Server `yaml:"server"` // Server configuration section.
}

// Server represents the server-specific configuration settings.
type Server struct {
	Port int `yaml:"port"` // Port on which the server will run.
}

// LoadConfig loads the application configuration from a YAML file specified by an environment variable.
func LoadConfig(envVar string, defaultConfig App) (App, error) {
	filePath := os.Getenv(envVar)
	if filePath == "" {
		return defaultConfig, errors.New("environment variable not set")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return defaultConfig, err
	}
	defer file.Close()
	var config App
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return defaultConfig, err
	}

	return config, nil
}
