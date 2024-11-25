package config

import (
	"fmt"
	"os"
)

type Config struct {
	serverPort  string
	databaseDSN string
}

func LoadConfig() (*Config, error) {

	dbURL := getEnv("DATABASE_URL", "")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	return &Config{
		databaseDSN: dbURL,
	}, nil
}

func getEnv(key string, fallback ...string) string {
	if len(fallback) == 0 {
		fmt.Println("fallback required")
	}
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback[0]
}

func (c *Config) ServerPort() string {
	return c.serverPort
}

func (c *Config) DatabaseDSN() string {
	return c.databaseDSN
}
