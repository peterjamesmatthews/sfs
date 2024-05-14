package app

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type App struct {
	db    *gorm.DB
	uuids []uuid.UUID
}

func New(db *gorm.DB) App {
	return App{db: db}
}
