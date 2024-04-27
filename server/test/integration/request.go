package integration

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"pjm.dev/sfs/graph"
)

func newGQLRequest(t *testing.T, requestor graph.User, path string) *http.Request {
	t.Helper()

	query := gqlFileToString(t, path)

	return &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Scheme: "HTTP",
			Path:   "/graphql", // should be shared with test handler's path
		},
		Header: http.Header{
			"Accept":          {"application/json"},
			"Accept-Language": {"en-US,en;q=0.5"},
			"Accept-Encoding": {"gzip, deflate, br"},
			"Authorization":   {requestor.Name},
			"Content-Type":    {"application/json"},
			"Content-Length":  {strconv.Itoa(len(query))},
		},
		Body: io.NopCloser(strings.NewReader(query)),
	}
}
