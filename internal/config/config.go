package config

import (
	"fmt"
	"os"
	"path/filepath"

	yamlV3 "gopkg.in/yaml.v3"
)

// StatsConfig ...
type statsConfig struct {
	IsCPU     bool `yaml:"is_cpu"`
	IsDisk    bool `yaml:"is_disk"`
	IsLoadAvg bool `yaml:"is_load_avg"`
}

// Config ...
type Config struct {
	Stats *statsConfig `yaml:"stats"`
}

// GetConfig ...
func GetConfig(path string) (*Config, error) {
	configContent, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	config := &Config{}
	err = yamlV3.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file %s: %w", path, err)
	}

	return config, nil
}

// IsCPU ...
func (c *Config) IsCPU() bool {
	return c.Stats.IsCPU
}

// IsDisk ...
func (c *Config) IsDisk() bool {
	return c.Stats.IsDisk
}

// IsLoadAvg ...
func (c *Config) IsLoadAvg() bool {
	return c.Stats.IsLoadAvg
}
