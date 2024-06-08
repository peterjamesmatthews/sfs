package app

import (
	"crypto/rand"
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
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
