package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSegments(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		segments []string
	}{
		{
			name:     "empty path",
			path:     "",
			segments: []string{},
		},
		{
			name:     "root path",
			path:     "/",
			segments: []string{},
		},
		{
			name:     "/foo/bar/baz",
			path:     "/foo/bar/baz",
			segments: []string{"foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			segments := getPathSegments(test.path)
			assert.Equal(t, test.segments, segments)
		})
	}
}
