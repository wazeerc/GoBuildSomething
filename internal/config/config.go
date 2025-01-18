package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultPath string `json:"defaultPath"`
}

const configFileName = "config.json"

func GetConfigPath() string {
	execPath, _ := os.Executable()
	return filepath.Join(filepath.Dir(execPath), configFileName)
}

func LoadConfig() (*Config, error) {
	configPath := GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(GetConfigPath(), data, 0644)
}
