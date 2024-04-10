package db

import (
	"database/sql"

	"gorm.io/gorm"
)

func getModels() []any {
	return []any{
		&User{},
		&Node{},
		&Content{},
		&Access{},
	}
}

type User struct {
	gorm.Model
	Name string `gorm:"notNull"`
}

type Node struct {
	gorm.Model
	Name     string `gorm:"notNull"`
	OwnerID  uint   `gorm:"notNull"`
	Owner    User
	ParentID sql.NullInt64
	Parent   *Node
}

type Content struct {
	gorm.Model
	FileID  uint `gorm:"notNull"`
	File    Node
	Content string `gorm:"notNull"`
}

type AccessType uint8

const (
	Read = iota
	Write
)

type Access struct {
	UserID   uint `gorm:"notNull"`
	User     User
	Type     AccessType `gorm:"notNull"`
	TargetID uint       `gorm:"notNull"`
	Target   Node
}
