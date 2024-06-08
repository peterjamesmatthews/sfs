package app

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"pjm.dev/sfs/db/models"
)

// userNameRegex is a regular expression that matches valid user names.
//
// A valid user name must:
//   - Be between 3 and 64 characters long.
//   - Contain only letters, numbers, and spaces.
var userNameRegex = regexp.MustCompile(`^[a-zA-z\d\ ]{3,64}$`)

// isValidUserName reports whether the provided name is a valid user name.
func (a *App) isValidUserName(name string) bool {
	return userNameRegex.MatchString(name)
}

// generateSalt generates a random 32-byte slice.
func (a *App) generateSalt() ([]byte, error) {
	N := 32
	b := make([]byte, N)
	n, err := rand.Read(b)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read random bytes: %w", err)
	} else if n < N {
		return []byte{}, errors.New("failed to read enough random bytes")
	}
	return b, nil
}

// hashPasswordWithSalt hashes a password with a salt using bcrypt.
func (a *App) hashPasswordWithSalt(password string, salt []byte) ([]byte, error) {
	// hash password with salt using bcrypt
	hash, err := bcrypt.GenerateFromPassword(append([]byte(password), salt...), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return []byte{}, errors.New("password is too long")
	} else if err != nil {
		return []byte{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// return hash
	return hash, nil
}

// comparePasswordWithSalt compares a password with a salt and hash using bcrypt.
func (a *App) comparePasswordWithSalt(password string, salt []byte, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, append([]byte(password), salt...))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to compare hash and password: %w", err)
	}
	return true, nil
}

const JWT_ISSUER = "sfs" // TODO move to config

var JWT_SECRET = []byte("secret") // TODO move to config

// generateTokensForUser generates access and refresh tokens for a user.
func (a *App) generateTokensForUser(user models.User) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    JWT_ISSUER,
		Subject:   uuid.UUID(user.ID.Bytes).String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	access, err := accessToken.SignedString(JWT_SECRET)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    JWT_ISSUER,
		Subject:   uuid.UUID(user.ID.Bytes).String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	refresh, err := refreshToken.SignedString(JWT_SECRET)
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
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		},
	)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to parse access token: %w", err)
	} else if !token.Valid {
		return models.User{}, errors.New("token is invalid")
	}

	// verify token issuer
	iss, err := token.Claims.GetIssuer()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get issuer from access token: %w", err)
	} else if iss != JWT_ISSUER {
		return models.User{}, errors.New("invalid token issuer")
	}

	// get subject from token
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get subject from access token: %w", err)
	}

	id := &pgtype.UUID{}
	err = id.Scan(sub)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to scan user ID from access token: %w", err)
	}

	user, err := a.q.GetUserByID(context.Background(), *id)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}
