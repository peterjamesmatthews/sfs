package app

import (
	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/auth0"
	"pjm.dev/sfs/db/models"
)

type App struct {
	queries *models.Queries
	auth0   Auth0
}

func New(conn *pgx.Conn) App {
	app := App{
		queries: models.New(conn),
		auth0:   &auth0.App{},
	}
	return app
}
