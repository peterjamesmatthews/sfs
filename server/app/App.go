package app

import (
	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/db/models"
)

type App struct {
	db *pgx.Conn
	q  *models.Queries
}

func New(conn *pgx.Conn) App {
	return App{
		db: conn,
		q:  models.New(conn),
	}
}
