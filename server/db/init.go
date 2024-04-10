package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pjm.dev/sfs/env"
)

func Initialize(config env.DatabaseConfig) (*gorm.DB, error) {
	// connect to database
	db, err := connect(config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize\n%v", err)
	}

	// migrate schema
	if err = migrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate\n%v", err)
	}

	// seed data
	if err = seed(db); err != nil {
		return nil, fmt.Errorf("failed to seed\n%v", err)
	}

	return nil, errors.New("db.Initialize is not implemented")
}

func connect(config env.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		"postgres",
		config.User,
		config.Password,
		config.Hostname,
		config.Port,
		config.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect\n%w", err)
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	return errors.New("db.migrate is not implemented")
}

func seed(db *gorm.DB) error {
	return errors.New("db.seed is not implemented")
}
