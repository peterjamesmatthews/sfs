package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

func New() (Config, error) {
	var config Config
	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

type Config struct {
	Server   Server   `env:", prefix=SERVER_"`
	Database Database `env:", prefix=DATABASE_"j`
}

func (c Config) String() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("failed to marshal config: %s", err.Error())
	}
	return string(bytes)
}

type Server struct {
	Hostname      string `env:"HOSTNAME"`
	Port          string `env:"PORT"`
	GraphEndpoint string `env:"GRAPH_ENDPOINT"`
}

type Database struct{}
