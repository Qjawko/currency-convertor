package config

import (
	"errors"
	"os"
)

type Config struct {
	APIKey string
}

// LoadConfig загружает конфигурацию из переменных окружения.
func LoadConfig() (*Config, error) {
	apiKey, exists := os.LookupEnv("CMC_API_KEY")
	if !exists {
		return nil, ErrMissingAPIKey
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}

var ErrMissingAPIKey = errors.New("missing CMC_API_KEY environment variable")
