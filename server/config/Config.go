package config

import (
	"context"
	"encoding/json"
	"fmt"

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

func (c Config) String() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("failed to marshal config: %s", err.Error())
	}
	return string(bytes)
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

func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		d.User,
		d.Password,
		d.Hostname,
		d.Port,
		d.Name,
	)
}
