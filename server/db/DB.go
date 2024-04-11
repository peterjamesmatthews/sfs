package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"ariga.io/atlas-go-sdk/atlasexec"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/meta"
)

func New(config config.DatabaseConfig) (*gorm.DB, error) {
	// connect to database
	db, err := connect(config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize: %w", err)
	}

	// migrate schema
	if err = migrate(config); err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}

func connect(config config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()))
	if err != nil {
		return nil, fmt.Errorf(
			"failed to open gorm connection to dsn %s: %w", config.GetDSN(), err)
	}

	return db, nil
}

var migrationsPath = filepath.Join(meta.Root, "db", "migrations")

func migrate(config config.DatabaseConfig) error {
	dir := os.DirFS(migrationsPath)

	wd, err := atlasexec.NewWorkingDir(atlasexec.WithMigrations(dir))
	if err != nil {
		return fmt.Errorf(
			"failed to open atlast working directory at %s: %w", migrationsPath, err,
		)
	}
	defer wd.Close()

	client, err := atlasexec.NewClient(wd.Path(), "atlas")
	if err != nil {
		return fmt.Errorf(
			"failed to initialize atlas client at path %s: %w", wd.Path(), err,
		)
	}

	_, err = client.MigrateApply(
		context.Background(),
		&atlasexec.MigrateApplyParams{URL: config.GetDSN()},
	)
	if err != nil {
		return fmt.Errorf("failed to apply atlas migrations: %w", err)
	}

	return nil
}
