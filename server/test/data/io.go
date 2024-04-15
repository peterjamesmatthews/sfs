package data

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"pjm.dev/sfs/meta"
)

func Read(t *testing.T, path string) []byte {
	t.Helper()

	// if path is relative, make it absolute to the test data directory
	if !filepath.IsAbs(path) {
		path = filepath.Join(meta.Root, "test", "data", path)
	}

	// open file
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}

	// read file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	return bytes
}
