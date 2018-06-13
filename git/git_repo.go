package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type GitRepo struct {
	Branch       string
	ChangedFiles []string
	Commits      []string
	Repository   string
	Path         string
}

func NewGitRepo(repoPath string) *GitRepo {
	repo := GitRepo{Path: repoPath}

	repo.Branch = repo.branch()
	repo.ChangedFiles = repo.changedFiles()
	repo.Commits = repo.commits()
	repo.Repository = strings.TrimSpace(repo.repository())

	return &repo
}

/* -------------------- Unexported Functions -------------------- */

func (repo *GitRepo) branch() string {
	arg := []string{repo.gitDir(), repo.workTree(), "rev-parse", "--abbrev-ref", "HEAD"}

	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}

func (repo *GitRepo) changedFiles() []string {
	arg := []string{repo.gitDir(), repo.workTree(), "status", "--porcelain"}

	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *GitRepo) commits() []string {
	numStr := fmt.Sprintf("-n %d", Config.UInt("wtf.mods.git.commitCount", 10))

	arg := []string{repo.gitDir(), repo.workTree(), "log", "--date=format:\"%b %d, %Y\"", numStr, "--pretty=format:\"[forestgreen]%h [white]%s [grey]%an on %cd[white]\""}

	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *GitRepo) repository() string {
	arg := []string{repo.gitDir(), repo.workTree(), "rev-parse", "--show-toplevel"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}
func (repo *GitRepo) pull() string {
	arg := []string{repo.gitDir(), repo.workTree(), "pull"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)
	return str
}

func (repo *GitRepo) checkout(branch string) string {
	arg := []string{repo.gitDir(), repo.workTree(), "checkout", branch}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)
	return str
}

func (repo *GitRepo) gitDir() string {
	return fmt.Sprintf("--git-dir=%s/.git", repo.Path)
}

func (repo *GitRepo) workTree() string {
	return fmt.Sprintf("--work-tree=%s", repo.Path)
}
