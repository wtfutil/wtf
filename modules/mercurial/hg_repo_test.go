package mercurial

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repoPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{name: "simple path", path: "/home/user/repo", want: "--repository=/home/user/repo"},
		{name: "empty path", path: "", want: "--repository="},
		{name: "path with spaces", path: "/home/user/my repo", want: "--repository=/home/user/my repo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MercurialRepo{Path: tt.path}
			assert.Equal(t, tt.want, repo.repoPath())
		})
	}
}

func Test_bookmark(t *testing.T) {
	t.Run("bookmarks file exists", func(t *testing.T) {
		dir := t.TempDir()
		hgDir := filepath.Join(dir, ".hg")
		assert.NoError(t, os.MkdirAll(hgDir, 0o755))
		assert.NoError(t, os.WriteFile(filepath.Join(hgDir, "bookmarks.current"), []byte("my-feature"), 0o644))

		repo := &MercurialRepo{Path: dir}
		assert.Equal(t, "my-feature", repo.bookmark())
	})

	t.Run("bookmarks file missing", func(t *testing.T) {
		dir := t.TempDir()

		repo := &MercurialRepo{Path: dir}
		assert.Equal(t, "", repo.bookmark())
	})
}

// fakeHg installs a fake "hg" executable on PATH for the duration of the
// test, so exec-based methods can be exercised without a real Mercurial
// installation or repository.
func fakeHg(t *testing.T) {
	t.Helper()

	dir := t.TempDir()

	var scriptPath, contents string
	if runtime.GOOS == "windows" {
		scriptPath = filepath.Join(dir, "hg.bat")
		contents = "@echo off\r\n" +
			"if \"%~1\"==\"branch\" (\r\n echo feature-branch\r\n exit /b 0\r\n)\r\n" +
			"if \"%~1\"==\"status\" (\r\n echo M file1.txt\r\n echo A file2.txt\r\n exit /b 0\r\n)\r\n" +
			"if \"%~1\"==\"log\" (\r\n echo rev1: first commit\r\n echo rev2: second commit\r\n exit /b 0\r\n)\r\n" +
			"if \"%~1\"==\"pull\" (\r\n echo pulling changes\r\n exit /b 0\r\n)\r\n" +
			"if \"%~1\"==\"checkout\" (\r\n echo checked out %~4\r\n exit /b 0\r\n)\r\n" +
			"echo unknown command\r\nexit /b 1\r\n"
	} else {
		scriptPath = filepath.Join(dir, "hg")
		contents = "#!/bin/sh\n" +
			"case \"$1\" in\n" +
			"  branch) echo 'feature-branch' ;;\n" +
			"  status) echo 'M file1.txt'; echo 'A file2.txt' ;;\n" +
			"  log) echo 'rev1: first commit'; echo 'rev2: second commit' ;;\n" +
			"  pull) echo 'pulling changes' ;;\n" +
			"  checkout) echo \"checked out $3\" ;;\n" +
			"  *) echo 'unknown command'; exit 1 ;;\n" +
			"esac\n"
	}

	assert.NoError(t, os.WriteFile(scriptPath, []byte(contents), 0o755))

	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

// trimLines removes trailing carriage returns from each line, which appear
// on Windows when the fake hg script echoes output.
func trimLines(lines []string) []string {
	trimmed := make([]string, len(lines))
	for i, line := range lines {
		trimmed[i] = strings.TrimRight(line, "\r")
	}
	return trimmed
}

func Test_NewMercurialRepo_WithFakeHg(t *testing.T) {
	fakeHg(t)

	repoPath := t.TempDir()
	repo := NewMercurialRepo(repoPath, 2, "{rev}:{desc}")

	assert.Equal(t, "feature-branch", repo.Branch)
	assert.Equal(t, "", repo.Bookmark)
	assert.Equal(t, strings.TrimSpace(repoPath), repo.Repository)
	assert.Equal(t, []string{"M file1.txt", "A file2.txt", ""}, trimLines(repo.ChangedFiles))
	assert.Equal(t, []string{"rev1: first commit", "rev2: second commit", ""}, trimLines(repo.Commits))
}

func Test_pull_WithFakeHg(t *testing.T) {
	fakeHg(t)

	repo := &MercurialRepo{Path: t.TempDir()}
	result := repo.pull()

	assert.Contains(t, result, "pulling changes")
}

func Test_checkout_WithFakeHg(t *testing.T) {
	fakeHg(t)

	repo := &MercurialRepo{Path: t.TempDir()}
	result := repo.checkout("my-branch")

	assert.Contains(t, result, "checked out my-branch")
}
