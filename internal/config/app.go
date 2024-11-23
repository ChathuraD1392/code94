package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type App struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port int `yaml:"port"`
}

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
