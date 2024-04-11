package app

import (
	"strings"
)

type uri string

const rootName = ""
const uriDelimiter = "/"
const rootURI uri = "/"

// IsAbsolute reports whether a URI starts from the root URI.
func (u uri) IsAbsolute() bool {
	return strings.HasPrefix(string(u), string(rootURI))
}

// GetNames returns the names of the URI split by the delimiter.
func (u uri) GetNames() []string {
	if u == rootURI {
		return []string{rootName}
	}

	names := strings.Split(string(u), uriDelimiter)
	if len(names) > 1 && names[len(names)-1] == "" {
		return names[:len(names)-1]
	}

	return names
}

// AddNames returns a new URI with the given names appended to the end.
func (u uri) AddNames(names ...string) uri {
	if len(names) == 0 {
		return u
	}
	names = append(u.GetNames(), names...)
	return uri(strings.Join(names, uriDelimiter))
}
