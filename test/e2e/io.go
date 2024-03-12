package e2e

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"pjm.dev/sfs/meta"
)

type Query struct {
	Query string `json:"query"`
}

func GQLFileToString(t *testing.T, path string) string {
	t.Helper()

	if !filepath.IsAbs(path) {
		path = filepath.Join(meta.Root, "test", "data", path)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	// read file contents to string
	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	query := Query{Query: string(bytes)}

	json, err := json.Marshal(query)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}

	s := string(json)

	return s
}
