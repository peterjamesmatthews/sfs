package mem

import (
	"strings"
)

type URI struct {
	uri string
}

const URIDelimiter = "/"
const RootName = ""

var RootURI = URI{uri: URIDelimiter}

func (u URI) IsAbsolute() bool {
	return strings.HasPrefix(u.uri, URIDelimiter)
}

// GetNames returns the names of the URI split by the delimiter.
func (u URI) GetNames() []string {
	if u == RootURI {
		return []string{RootName}
	}

	return strings.Split(u.uri, URIDelimiter)
}

// AddNames returns a new URI with the given names appended to the end.
func (u URI) AddNames(names ...string) URI {
	if len(names) == 0 {
		return u
	}
	names = append(u.GetNames(), names...)
	return URI{uri: strings.Join(names, URIDelimiter)}
}
