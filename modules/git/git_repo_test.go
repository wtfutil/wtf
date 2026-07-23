package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_gitDir(t *testing.T) {
	repo := &GitRepo{Path: "/home/user/project"}
	assert.Equal(t, "--git-dir=/home/user/project/.git", repo.gitDir())
}

func Test_workTree(t *testing.T) {
	repo := &GitRepo{Path: "/home/user/project"}
	assert.Equal(t, "--work-tree=/home/user/project", repo.workTree())
}

func Test_commitLogArgs(t *testing.T) {
	tests := []struct {
		name         string
		commitCount  int
		commitFormat string
		dateFormat   string
		want         []string
	}{
		{
			name:         "basic formatting",
			commitCount:  5,
			commitFormat: "%h %s",
			dateFormat:   "%b %d, %Y",
			want:         []string{"log", `--date=format:"%b %d, %Y"`, "-n 5", `--pretty=format:"%h %s"`},
		},
		{
			name:         "zero commit count",
			commitCount:  0,
			commitFormat: "%h",
			dateFormat:   "%Y",
			want:         []string{"log", `--date=format:"%Y"`, "-n 0", `--pretty=format:"%h"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, commitLogArgs(tt.commitCount, tt.commitFormat, tt.dateFormat))
		})
	}
}

// initTestRepo creates a temporary git repository with a single commit and returns its path.
func initTestRepo(t *testing.T) string {
	t.Helper()

	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git binary not available")
	}

	dir := t.TempDir()

	run := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(),
			"GIT_AUTHOR_NAME=Test", "GIT_AUTHOR_EMAIL=test@example.com",
			"GIT_COMMITTER_NAME=Test", "GIT_COMMITTER_EMAIL=test@example.com",
		)
		out, err := cmd.CombinedOutput()
		require.NoErrorf(t, err, "git %v failed: %s", args, out)
	}

	run("init", "-b", "main")
	run("config", "user.email", "test@example.com")
	run("config", "user.name", "Test")

	readme := filepath.Join(dir, "README.md")
	require.NoError(t, os.WriteFile(readme, []byte("hello\n"), 0o644))

	run("add", "README.md")
	run("commit", "-m", "Initial commit")

	return dir
}

func Test_NewGitRepo(t *testing.T) {
	dir := initTestRepo(t)

	repo := NewGitRepo(dir, 5, "%s", "%Y-%m-%d")

	assert.Equal(t, "main\n", repo.Branch)
	assert.Contains(t, filepath.ToSlash(repo.Repository), filepath.ToSlash(dir))
	require.Len(t, repo.Commits, 1)
	assert.Contains(t, repo.Commits[0], "Initial commit")
	// A clean repo reports only the trailing empty line from `git status --porcelain`.
	assert.Equal(t, []string{""}, repo.ChangedFiles)
}

func Test_GitRepo_changedFiles(t *testing.T) {
	dir := initTestRepo(t)

	untracked := filepath.Join(dir, "untracked.txt")
	require.NoError(t, os.WriteFile(untracked, []byte("new\n"), 0o644))

	repo := &GitRepo{Path: dir}
	files := repo.changedFiles()

	require.Len(t, files, 2)
	assert.Contains(t, files[0], "untracked.txt")
	assert.True(t, strings.HasPrefix(strings.TrimSpace(files[0]), "??"))
}

func Test_GitRepo_checkout(t *testing.T) {
	dir := initTestRepo(t)

	cmd := exec.Command("git", "branch", "feature")
	cmd.Dir = dir
	require.NoError(t, cmd.Run())

	repo := &GitRepo{Path: dir}
	repo.checkout("feature")

	assert.Equal(t, "feature\n", repo.branch())
}
