package app

import (
	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/db/models"
)

type App struct {
	Config  Config
	Queries *models.Queries
	Auth0   Auth0er
}

func New(config Config, conn *pgx.Conn, auth0 Auth0er) App {
	return App{Config: config, Queries: models.New(conn), Auth0: auth0}
}

func (a *App) SetDatabase(conn *pgx.Conn) {
	a.Queries = models.New(conn)
}
