package snbgo

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Account string
	Key     string
	Server  string
	Port    string
	Schema  string
	Timeout int // milliseconds
}

func LoadFromEnv() (*Config, error) {
	account := os.Getenv("SNB_ACCOUNT")
	key := os.Getenv("SNB_KEY")
	if account == "" || key == "" {
		return nil, fmt.Errorf("snbgo: SNB_ACCOUNT and SNB_KEY are required")
	}

	c := &Config{
		Account: account,
		Key:     key,
		Server:  getEnvDefault("SNB_SERVER", "sandbox.snbsecurities.com"),
		Port:    getEnvDefault("SNB_PORT", "443"),
		Schema:  getEnvDefault("SNB_SCHEMA", "https"),
		Timeout: 10000,
	}

	if t := os.Getenv("SNB_TIMEOUT"); t != "" {
		v, err := strconv.Atoi(t)
		if err != nil {
			return nil, fmt.Errorf("snbgo: invalid SNB_TIMEOUT: %v", err)
		}
		c.Timeout = v
	}

	return c, nil
}

func (c *Config) BaseURL() string {
	return fmt.Sprintf("%s://%s:%s", c.Schema, c.Server, c.Port)
}

func getEnvDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
