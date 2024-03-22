// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
	GetID() string
	GetName() string
	GetOwner() *User
	GetParent() *Folder
}

type Access struct {
	User   *User      `json:"user"`
	Type   AccessType `json:"type"`
	Target Node       `json:"target"`
}

type File struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Owner   *User   `json:"owner"`
	Parent  *Folder `json:"parent,omitempty"`
	Content string  `json:"content"`
}

func (File) IsNode()                 {}
func (this File) GetID() string      { return this.ID }
func (this File) GetName() string    { return this.Name }
func (this File) GetOwner() *User    { return this.Owner }
func (this File) GetParent() *Folder { return this.Parent }

type Folder struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Owner    *User   `json:"owner"`
	Parent   *Folder `json:"parent,omitempty"`
	Children []Node  `json:"children"`
}

func (Folder) IsNode()                 {}
func (this Folder) GetID() string      { return this.ID }
func (this Folder) GetName() string    { return this.Name }
func (this Folder) GetOwner() *User    { return this.Owner }
func (this Folder) GetParent() *Folder { return this.Parent }

type Mutation struct {
}

type Query struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccessType string

const (
	AccessTypeRead  AccessType = "READ"
	AccessTypeWrite AccessType = "WRITE"
)

var AllAccessType = []AccessType{
	AccessTypeRead,
	AccessTypeWrite,
}

func (e AccessType) IsValid() bool {
	switch e {
	case AccessTypeRead, AccessTypeWrite:
		return true
	}
	return false
}

func (e AccessType) String() string {
	return string(e)
}

func (e *AccessType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccessType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccessType", str)
	}
	return nil
}

func (e AccessType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}