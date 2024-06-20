package app

import (
	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/db/models"
)

type App struct {
	config  Config
	queries *models.Queries
}

func New(config Config, conn *pgx.Conn) App {
	return App{config: config, queries: models.New(conn)}
}
