package integration

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

func gqlFileToString(t *testing.T, path string) string {
	t.Helper()

	if !filepath.IsAbs(path) {
		path = filepath.Join(meta.Root, "test", "data", path)
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open %s: %s", path, err.Error())
	}

	// read file contents to string
	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read %s: %s", path, err.Error())
	}

	query := Query{Query: string(bytes)}

	json, err := json.Marshal(query)
	if err != nil {
		t.Fatalf("failed to marshal %s's content to JSON: %s", path, err.Error())
	}

	return string(json)
}
