package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

func New(ctx context.Context) (Config, error) {
	var config Config
	err := envconfig.Process(ctx, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

type Config struct {
	Server   ServerConfig   `env:", prefix=SERVER_"`
	Database DatabaseConfig `env:", prefix=DATABASE_"`
}

type ServerConfig struct {
	Hostname      string `env:"HOSTNAME"`
	Port          string `env:"PORT"`
	GraphEndpoint string `env:"GRAPH_ENDPOINT"`
}

type DatabaseConfig struct {
	Hostname string `env:"HOSTNAME"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
}
