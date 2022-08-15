// Copied verbatim from:
//
//   https://github.com/otiai10/copy/blob/master/copy.go

package cfg

import (
	"io"
	"os"
	"path/filepath"
)

// Copy copies src to dest, doesn't matter if src is a directory or a file
func Copy(src, dest string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	return locationCopy(src, dest, info)
}

// "info" must be given here, NOT nil.
func locationCopy(src, dest string, info os.FileInfo) error {
	if info.IsDir() {
		return directoryCopy(src, dest, info)
	}
	return fileCopy(src, dest, info)
}

func fileCopy(src, dest string, info os.FileInfo) error {

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	if err = os.Chmod(f.Name(), info.Mode()); err != nil {
		return err
	}

	s, err := os.Open(filepath.Clean(src))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = io.Copy(f, s)
	return err
}

func directoryCopy(src, dest string, info os.FileInfo) error {

	if err := os.MkdirAll(dest, info.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		if err := locationCopy(
			filepath.Join(src, info.Name()),
			filepath.Join(dest, info.Name()),
			info,
		); err != nil {
			return err
		}
	}

	return nil
}
