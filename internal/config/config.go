package config

import (
	"os"
	"path/filepath"
)

func GetAPIKey() string {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey != "" {
		return apiKey
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	configFile := filepath.Join(homeDir, ".see_parallel", "config")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return ""
	}

	return string(data)
}

func SetAPIKey(apiKey string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".see_parallel")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "config")
	return os.WriteFile(configFile, []byte(apiKey), 0600)
}