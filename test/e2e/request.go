package e2e

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"pjm.dev/sfs/graph/model"
)

func newRequest(requestor model.User, path string, query string) *http.Request {
	return &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Scheme: "HTTP",
			Path:   path,
		},
		Header: map[string][]string{
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
