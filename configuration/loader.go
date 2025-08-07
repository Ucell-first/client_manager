package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadFromJson(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", filepath, err)
	}
	defer file.Close()
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode JSON config, %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate JSON configuration: %w", err)
	}
	return &config, nil
}
