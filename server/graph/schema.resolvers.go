package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"errors"
	"fmt"
)

// User is the resolver for the user field.
func (r *accessResolver) User(ctx context.Context, obj *Access) (*User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Target is the resolver for the target field.
func (r *accessResolver) Target(ctx context.Context, obj *Access) (Node, error) {
	panic(fmt.Errorf("not implemented: Target - target"))
}

// Owner is the resolver for the owner field.
func (r *fileResolver) Owner(ctx context.Context, obj *File) (*User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Parent is the resolver for the parent field.
func (r *fileResolver) Parent(ctx context.Context, obj *File) (*Folder, error) {
	panic(fmt.Errorf("not implemented: Parent - parent"))
}

// Owner is the resolver for the owner field.
func (r *folderResolver) Owner(ctx context.Context, obj *Folder) (*User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Parent is the resolver for the parent field.
func (r *folderResolver) Parent(ctx context.Context, obj *Folder) (*Folder, error) {
	panic(fmt.Errorf("not implemented: Parent - parent"))
}

// Children is the resolver for the children field.
func (r *folderResolver) Children(ctx context.Context, obj *Folder) ([]Node, error) {
	panic(fmt.Errorf("not implemented: Children - children"))
}

// RefreshTokens is the resolver for the refreshTokens field.
func (r *mutationResolver) RefreshTokens(ctx context.Context, refresh string) (*Tokens, error) {
	tokens, err := r.SharedFileSystem.RefreshTokens(refresh)
	if errors.Is(err, ErrForbidden) {
		return nil, fmt.Errorf("refresh token is expired")
	} else if err != nil {
		return nil, fmt.Errorf("failed to refresh tokens: %w", err)
	}
	return &tokens, nil
}

// RenameNode is the resolver for the renameNode field.
func (r *mutationResolver) RenameNode(ctx context.Context, id string, name string) (Node, error) {
	panic(fmt.Errorf("not implemented: RenameNode - renameNode"))
}

// MoveNode is the resolver for the moveNode field.
func (r *mutationResolver) MoveNode(ctx context.Context, id string, parentID *string) (Node, error) {
	panic(fmt.Errorf("not implemented: MoveNode - moveNode"))
}

// ShareNode is the resolver for the shareNode field.
func (r *mutationResolver) ShareNode(ctx context.Context, userID string, accessType AccessType, targetID string) (*Access, error) {
	panic(fmt.Errorf("not implemented: ShareNode - shareNode"))
}

// CreateFolder is the resolver for the createFolder field.
func (r *mutationResolver) CreateFolder(ctx context.Context, path string) (*Folder, error) {
	panic(fmt.Errorf("not implemented: CreateFolder - createFolder"))
}

// CreateFile is the resolver for the createFile field.
func (r *mutationResolver) CreateFile(ctx context.Context, parentID *string, name string, content *string) (*File, error) {
	panic(fmt.Errorf("not implemented: CreateFile - createFile"))
}

// WriteFile is the resolver for the writeFile field.
func (r *mutationResolver) WriteFile(ctx context.Context, id string, content string) (*File, error) {
	panic(fmt.Errorf("not implemented: WriteFile - writeFile"))
}

// GetTokensFromAuth0 is the resolver for the getTokensFromAuth0 field.
func (r *queryResolver) GetTokensFromAuth0(ctx context.Context, token string) (*Tokens, error) {
	tokens, err := r.SharedFileSystem.GetTokensFromAuth0Token(token)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}
	return &tokens, nil
}

// Me is the resolver for the Me field.
func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}
	return &user, nil
}

// GetNodeFromPath is the resolver for the getNodeFromPath field.
func (r *queryResolver) GetNodeFromPath(ctx context.Context, path string) (Node, error) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	node, err := r.SharedFileSystem.GetNodeFromPath(user, path)
	if errors.Is(err, ErrNotFound) {
		return nil, fmt.Errorf("node not found")
	} else if errors.Is(err, ErrUnauthorized) {
		return nil, fmt.Errorf("unauthorized")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node: %w", err)
	}

	return node, nil
}

// Access returns AccessResolver implementation.
func (r *Resolver) Access() AccessResolver { return &accessResolver{r} }

// File returns FileResolver implementation.
func (r *Resolver) File() FileResolver { return &fileResolver{r} }

// Folder returns FolderResolver implementation.
func (r *Resolver) Folder() FolderResolver { return &folderResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type accessResolver struct{ *Resolver }
type fileResolver struct{ *Resolver }
type folderResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
