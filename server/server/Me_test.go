package server_test

import (
	"errors"
	"testing"
)

// TestGetTokensFromAuth0 tests server's handling of the Me query.
//
// test cases:
//  1. Invalid token
//  2. Invalid token issuer
//  3. Missing subject claim
//  4. Invalid subject claim
//  5. User not found
//  6. User found
func TestMe(t *testing.T) {
	// TODO implement me
	t.Fatal(errors.ErrUnsupported)
}
