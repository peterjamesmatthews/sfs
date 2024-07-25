package app

import (
	"database/sql"

	"pjm.dev/sfs/db/models"
)

type App struct {
	Config  Config
	Queries *models.Queries
	Auth0   Auth0er
}

func New(config Config, db *sql.DB, auth0 Auth0er) App {
	return App{Config: config, Queries: models.New(db), Auth0: auth0}
}

func (a *App) SetDatabase(db *sql.DB) {
	a.Queries = models.New(db)
}
