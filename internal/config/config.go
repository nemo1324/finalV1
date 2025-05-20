package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GRPC     *GRPC     `envconfig:"GRPC"`
	HTTP     *HTTP     `envconfig:"HTTP"`
	Postgres *Postgres `envconfig:"POSTGRES"`
	Logger   *Logger   `envconfig:"LOGGER"`
	JWT      *JWT      `envconfig:"JWT"`
}

func Load() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("process load config: %w", err)
	}

	return cfg, nil
}
