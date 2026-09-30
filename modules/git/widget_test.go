package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wtfutil/wtf/cfg"
)

func Test_currentData(t *testing.T) {
	repoA := &GitRepo{Repository: "a"}
	repoB := &GitRepo{Repository: "b"}

	tests := []struct {
		name  string
		repos []*GitRepo
		idx   int
		want  *GitRepo
	}{
		{"no repos returns nil", nil, 0, nil},
		{"negative index returns nil", []*GitRepo{repoA}, -1, nil},
		{"index beyond bounds returns nil", []*GitRepo{repoA}, 1, nil},
		{"valid index returns repo", []*GitRepo{repoA, repoB}, 1, repoB},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{GitRepos: tt.repos}
			widget.Idx = tt.idx

			assert.Equal(t, tt.want, widget.currentData())
		})
	}
}

func Test_gitRepos(t *testing.T) {
	dir := initTestRepo(t)

	widget := &Widget{
		settings: &Settings{
			Common:       &cfg.Common{},
			commitCount:  1,
			commitFormat: "%s",
			dateFormat:   "%Y",
		},
	}

	repos := widget.gitRepos([]string{dir})

	require.Len(t, repos, 1)
	assert.Equal(t, "main\n", repos[0].Branch)
}

func Test_findGitRepositories(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git binary not available")
	}

	root := t.TempDir()

	repoOne := filepath.Join(root, "repo-one")
	repoTwo := filepath.Join(root, "nested", "repo-two")
	ignored := filepath.Join(root, "vendor", "repo-three")

	for _, path := range []string{repoOne, repoTwo, ignored} {
		require.NoError(t, os.MkdirAll(path, 0o755))
		cmd := exec.Command("git", "init", "-b", "main")
		cmd.Dir = path
		require.NoError(t, cmd.Run())
	}

	widget := &Widget{
		settings: &Settings{
			Common:       &cfg.Common{},
			commitCount:  1,
			commitFormat: "%s",
			dateFormat:   "%Y",
		},
	}

	repos := widget.findGitRepositories(make([]*GitRepo, 0), root)

	require.Len(t, repos, 2)

	var found []string
	for _, r := range repos {
		found = append(found, filepath.Base(r.Path))
	}
	assert.ElementsMatch(t, []string{"repo-one", "repo-two"}, found)
}
