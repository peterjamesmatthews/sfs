package app

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pjm.dev/sfs/graph"
)

type Apper interface {
	graph.SharedFileSystemer
	graph.Authenticator
	graph.UUIDGenerator
}

type app struct {
	db    *gorm.DB
	uuids []uuid.UUID
}

func New(db *gorm.DB) Apper {
	return &app{db: db}
}
