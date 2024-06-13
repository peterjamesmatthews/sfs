package meta

import (
	"path/filepath"
	"runtime"
)

// adapted from https://stackoverflow.com/a/58294680/15068300

var (
	_, file, _, _ = runtime.Caller(0)
	// Absolute path to to the root of this go module.
	Root = filepath.Join(filepath.Dir(file), "..")
)
