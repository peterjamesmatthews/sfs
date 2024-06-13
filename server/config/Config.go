package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-envconfig"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/server"
)

type Config struct {
	Server   server.Config `env:", prefix=SERVER_"`
	App      app.Config    `env:", prefix="`
	Database db.Config     `env:", prefix=DATABASE_"`
}

func New(ctx context.Context) (Config, error) {
	var config Config
	err := envconfig.Process(ctx, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c Config) String() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("failed to marshal config: %s", err.Error())
	}
	return string(bytes)
}
