package app

import (
	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/db/models"
)

type App struct {
	config  Config
	queries *models.Queries
	auth0   Auth0
}

func New(config Config, conn *pgx.Conn) App {
	app := App{config: config, queries: models.New(conn)}
	app.auth0 = &app
	return app
}
