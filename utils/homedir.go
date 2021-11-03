// Package homedir helps with detecting and expanding the user's home directory

// Copied (mostly) verbatim from https://github.com/Atrox/homedir

package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// ExpandHomeDir expands the path to include the home directory if the path
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
// returned as-is.
func ExpandHomeDir(path string) (string, error) {
	if path == "" {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}
