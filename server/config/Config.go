package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-envconfig"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/auth0"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/server"
)

type Config struct {
	App      app.Config    `env:", prefix="`
	Auth0    auth0.Config  `env:", prefix=AUTH0_"`
	Database db.Config     `env:", prefix=DATABASE_"`
	Server   server.Config `env:", prefix=SERVER_"`
}

func New(ctx context.Context) (Config, error) {
	var config Config
	err := envconfig.Process(ctx, &config)
	return config, err
}

func (c Config) String() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("failed to marshal config: %s", err.Error())
	}
	return string(bytes)
}
