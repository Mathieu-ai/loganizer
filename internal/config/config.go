package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LogConfig represents a single log file configuration entry
type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

/**
 * LoadConfig deserializes JSON configuration file into LogConfig slice.
 * Returns error if file access or JSON parsing fails.
 *
 * @param configPath Path to the JSON configuration file
 * @return Slice of LogConfig and error if any
 */
func LoadConfig(configPath string) ([]LogConfig, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var configs []LogConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return configs, nil
}
