package app

import (
	"slices"
	"strings"
)

func getPathSegments(path string) []string {
	// split path by /
	segments := strings.Split(path, "/")

	// remove empty strings
	segments = slices.DeleteFunc(segments, func(s string) bool { return s == "" })

	// return segments
	return segments
}
