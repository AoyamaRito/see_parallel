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

func GetContext() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	contextFile := filepath.Join(wd, ".see_parallel", "context")
	data, err := os.ReadFile(contextFile)
	if err != nil {
		return ""
	}

	return string(data)
}

func SetContext(context string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	configDir := filepath.Join(wd, ".see_parallel")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}

	contextFile := filepath.Join(configDir, "context")
	return os.WriteFile(contextFile, []byte(context), 0600)
}

func ClearContext() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	contextFile := filepath.Join(wd, ".see_parallel", "context")
	return os.Remove(contextFile)
}