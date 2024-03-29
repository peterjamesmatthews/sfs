package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
)

// GetRoot is the resolver for the getRoot field.
func (r *queryResolver) GetRoot(ctx context.Context) (*Folder, error) {
	user, err := handleGettingUserFromContext(ctx, r.AuthN)
	if err != nil {
		return nil, err
	}

	root, err := r.SFS.GetRoot(user)
	if err != nil {
		return nil, fmt.Errorf("failed to get root: %w", err)
	}

	return &root, nil
}

// GetNodeByID is the resolver for the getNodeById field.
func (r *queryResolver) GetNodeByID(ctx context.Context, id string) (Node, error) {
	user, err := handleGettingUserFromContext(ctx, r.AuthN)
	if err != nil {
		return nil, err
	}

	node, err := r.SFS.GetNodeByID(user, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get node by id: %w", err)
	}

	return node, nil
}

// GetNodeByURI is the resolver for the getNodeByURI field.
func (r *queryResolver) GetNodeByURI(ctx context.Context, uri string) (Node, error) {
	user, err := handleGettingUserFromContext(ctx, r.AuthN)
	if err != nil {
		return nil, err
	}

	node, err := r.SFS.GetNodeByURI(user, uri)
	if err != nil {
		return nil, fmt.Errorf("failed to get node at uri: %w", err)
	}

	return node, nil
}

// GetFileByID is the resolver for the getFileById field.
func (r *queryResolver) GetFileByID(ctx context.Context, id string) (*File, error) {
	user, err := handleGettingUserFromContext(ctx, r.AuthN)
	if err != nil {
		return nil, err
	}

	file, err := r.SFS.GetFileByID(user, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get file by id: %w", err)
	}

	return &file, nil
}

// GetFolderByID is the resolver for the getFolderById field.
func (r *queryResolver) GetFolderByID(ctx context.Context, id string) (*Folder, error) {
	user, err := handleGettingUserFromContext(ctx, r.AuthN)
	if err != nil {
		return nil, err
	}

	folder, err := r.SFS.GetFolderByID(user, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get folder by id: %w", err)
	}

	return &folder, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
