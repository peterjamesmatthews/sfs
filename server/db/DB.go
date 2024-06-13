package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func New(config Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), config.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
