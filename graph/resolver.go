package graph

//go:generate go run github.com/99designs/gqlgen generate

import "pjm.dev/sfs/auth"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SFS   SharedFileSystemer
	AuthN auth.Authenticator
}
