package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type EnvConfig struct {
	Port                  string `env:"PORT, default=8080"`
	Ip                    string `env:"IP, default=0.0.0.0"`
	PostgresHost          string `env:"POSTGRES_HOST, default=localhost"`
	PostgresPort          string `env:"POSTGRES_PORT, default=5432"`
	PostgresUser          string `env:"POSTGRES_USER, default=postgres"`
	PostgresPassword      string `env:"POSTGRES_PASSWORD"`
	PostgresDatabase      string `env:"POSTGRES_DATABASE, default=postgres"`
	NatsUrl               string `env:"NATS_URL, default=nats://localhost:4222"`
}

var Config *EnvConfig

func Init(ctx context.Context) error {
	Config = &EnvConfig{}
	if err := envconfig.Process(ctx, Config); err != nil {
		return err
	}
	return nil
}
