package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

func New() (*Service, error) {
	envConfig := envConfig{}
	if err := env.Parse(&envConfig); err != nil {
		return nil, fmt.Errorf("environment parsing failed. err: %w", err)
	}

	return &Service{
		envConfig: envConfig,
	}, nil
}

type Service struct {
	envConfig envConfig
}

type envConfig struct {
	SessionSecret string `env:"SESSION_SECRET,required"`
}
