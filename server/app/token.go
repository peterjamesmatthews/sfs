package app

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"pjm.dev/sfs/db/models"
)

func (a *App) getTokenFromAuthorization(auth string) string {
	return strings.TrimPrefix(auth, "Bearer ")
}

// GetTokensForUser generates access and refresh tokens for a user.
func (a *App) GetTokensForUser(user models.User) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    a.Config.JWT_Issuer,
		Subject:   user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	access, err := accessToken.SignedString(a.Config.JWT_Secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    a.Config.JWT_Issuer,
		Subject:   user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(14 * 24 * time.Hour)), // two weeks
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	refresh, err := refreshToken.SignedString(a.Config.JWT_Secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	// return access and refresh tokens
	return access, refresh, nil
}

func (a *App) getUserFromToken(tokenString string) (models.User, error) {
	// parse token
	token, err := jwt.ParseWithClaims(
		tokenString,
		new(jwt.RegisteredClaims),
		func(t *jwt.Token) (interface{}, error) {
			return a.Config.JWT_Secret, nil
		},
	)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to parse access token: %w", err)
	} else if !token.Valid {
		return models.User{}, errors.New("invalid token")
	}

	// verify token issuer
	iss, err := token.Claims.GetIssuer()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get issuer from access token: %w", err)
	} else if iss != a.Config.JWT_Issuer {
		return models.User{}, errors.New("invalid token issuer")
	}

	// get subject from token
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get subject from access token: %w", err)
	}

	// get user's id from subject subject
	var id uuid.UUID
	err = id.Scan(sub)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to scan user ID from access token: %w", err)
	}

	// get user by id
	user, err := a.Queries.GetUserByID(context.Background(), id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user by ID: %w", err)
	}

	// return user
	return user, nil
}

var errExpired = errors.New("expired")

func (a *App) refreshTokens(refresh string) (string, string, error) {
	// parse refresh token
	token, err := jwt.ParseWithClaims(refresh, new(jwt.RegisteredClaims), func(t *jwt.Token) (interface{}, error) {
		return a.Config.JWT_Secret, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to parse refresh token: %w", err)
	}

	// verify token issuer
	iss, err := token.Claims.GetIssuer()
	if err != nil {
		return "", "", fmt.Errorf("failed to get issuer from refresh token: %w", err)
	} else if iss != a.Config.JWT_Issuer {
		return "", "", errors.New("invalid token issuer")
	}

	// verify token is not expired
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return "", "", fmt.Errorf("failed to get expiration from refresh token: %w", err)
	} else if time.Now().After(exp.Time) {
		return "", "", errExpired
	}

	// get subject from token
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return "", "", fmt.Errorf("failed to get subject from refresh token: %w", err)
	}

	// get id from user
	var id uuid.UUID
	err = id.Scan(sub)
	if err != nil {
		return "", "", fmt.Errorf("failed to scan user ID from refresh token: %w", err)
	}

	// get user by id
	user, err := a.Queries.GetUserByID(context.Background(), id)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user by ID: %w", err)
	}

	// generate new access and refresh tokens
	access, newRefresh, err := a.GetTokensForUser(user)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	return access, newRefresh, nil
}
