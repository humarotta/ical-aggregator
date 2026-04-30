package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Calendars []Calendar `yaml:"calendars"`
}

type Calendar struct {
	Name  string `yaml:"name"`
	Feeds []Feed `yaml:"feeds"`
}

type Feed struct {
	URL string `yaml:"url"`
}

func NewConfig() *Config {
	return &Config{}
}

func Load(filepath string) (*Config, error) {
	cfg := NewConfig()

	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(file, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}
