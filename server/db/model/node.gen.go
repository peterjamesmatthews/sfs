// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameNode = "node"

// Node mapped from table <node>
type Node struct {
	ID     string `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name"`
	Owner  string `gorm:"column:owner;not null" json:"owner"`
	Parent string `gorm:"column:parent" json:"parent"`
}

// TableName Node's table name
func (*Node) TableName() string {
	return TableNameNode
}