package app

import (
	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/db/models"
)

type App struct {
	config  Config
	queries *models.Queries
	auth0   Auth0er
}

func New(config Config, conn *pgx.Conn, auth0 Auth0er) App {
	return App{config: config, queries: models.New(conn), auth0: auth0}
}

func (a *App) SetDatabase(conn *pgx.Conn) {
	a.queries = models.New(conn)
}
