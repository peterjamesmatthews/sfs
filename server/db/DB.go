package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/config"
)

func New(config config.DatabaseConfig) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), config.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
