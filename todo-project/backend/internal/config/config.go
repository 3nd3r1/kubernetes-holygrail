package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port string `env:"PORT, default=8080"`
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) ParseEnv(ctx context.Context) error {
	if err := envconfig.Process(ctx, config); err != nil {
		return err
	}

	return nil
}
