package config

import (
	"fmt"
	"os"
	"path/filepath"

	yamlV3 "gopkg.in/yaml.v3"
)

// StatsConfig ...
type StatsConfig struct {
	CPU     bool
	Disk    bool
	LoadAvg bool
}

// Config ...
type Config struct {
	Stats StatsConfig
}

// GetConfig ...
func GetConfig(path string) (*Config, error) {
	configContent, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	config := Config{}
	err = yamlV3.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file %s: %w", path, err)
	}

	return &config, nil
}
