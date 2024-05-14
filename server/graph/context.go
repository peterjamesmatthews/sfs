package graph

import (
	"context"
	"errors"
	"fmt"
)

func handleGettingUserFromContext(ctx context.Context, authN Authenticator) (User, error) {
	user, err := authN.FromContext(ctx)
	if errors.Is(err, ErrUnauthorized) {
		return User{}, err
	} else if err != nil {
		return User{}, fmt.Errorf("failed to get authenticated user: %w", err)
	}
	return user, nil
}
