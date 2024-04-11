package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pjm.dev/sfs/config"
)

func New(config config.DatabaseConfig) (*gorm.DB, error) {
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

	return db, nil
}

func connect(config config.DatabaseConfig) (*gorm.DB, error) {
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
	models := getModels()

	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to migrate\n%w", err)
	}

	return nil
}

func seed(*gorm.DB) error {
	return nil
}
