package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"HTTP_PORT" envDefault:"8080"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env wasn't found")
	}
}

func New() (*Config, error) {
	config := &Config{}

	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("failed to parse config from environment variables: %w", err)
	}

	return config, nil
}
