// Package homedir helps with detecting and expanding the user's home directory

// Copied (mostly) verbatim from https://github.com/Atrox/homedir

package wtf

import (
	"errors"
	"os/user"
	"path/filepath"
)

// Dir returns the home directory for the executing user.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	if currentUser.HomeDir == "" {
		return "", errors.New("cannot find user-specific home dir")
	}

	return currentUser.HomeDir, nil
}

// Expand expands the path to include the home directory if the path
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
// returned as-is.
func ExpandHomeDir(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := Home()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}
