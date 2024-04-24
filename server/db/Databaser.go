package db

import (
	"errors"

	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
)

func New(config config.Database) (app.Databaser, error) {
	return nil, errors.ErrUnsupported
}
