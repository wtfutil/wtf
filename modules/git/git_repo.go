package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

type GitRepo struct {
	Branch       string
	ChangedFiles []string
	Commits      []string
	Repository   string
	Path         string
}

func NewGitRepo(repoPath string, commitCount int, commitFormat, dateFormat string) *GitRepo {
	repo := GitRepo{Path: repoPath}

	repo.Branch = repo.branch()
	repo.ChangedFiles = repo.changedFiles()
	repo.Commits = repo.commits(commitCount, commitFormat, dateFormat)
	repo.Repository = strings.TrimSpace(repo.repository())

	return &repo
}

/* -------------------- Unexported Functions -------------------- */

func (repo *GitRepo) branch() string {
	arg := []string{repo.gitDir(), repo.workTree(), "rev-parse", "--abbrev-ref", "HEAD"}

	cmd := exec.Command(__go_cmd, arg...)
	str := utils.ExecuteCommand(cmd)

	return str
}

func (repo *GitRepo) changedFiles() []string {
	arg := []string{repo.gitDir(), repo.workTree(), "status", "--porcelain"}

	cmd := exec.Command(__go_cmd, arg...)
	str := utils.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *GitRepo) commits(commitCount int, commitFormat, dateFormat string) []string {
	dateStr := fmt.Sprintf("--date=format:\"%s\"", dateFormat)
	numStr := fmt.Sprintf("-n %d", commitCount)
	commitStr := fmt.Sprintf("--pretty=format:\"%s\"", commitFormat)

	arg := []string{repo.gitDir(), repo.workTree(), "log", dateStr, numStr, commitStr}

	cmd := exec.Command(__go_cmd, arg...)
	str := utils.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *GitRepo) repository() string {
	arg := []string{repo.gitDir(), repo.workTree(), "rev-parse", "--show-toplevel"}
	cmd := exec.Command(__go_cmd, arg...)
	str := utils.ExecuteCommand(cmd)

	return str
}
func (repo *GitRepo) pull() string {
	arg := []string{repo.gitDir(), repo.workTree(), "pull"}
	cmd := exec.Command(__go_cmd, arg...)
	str := utils.ExecuteCommand(cmd)
	return str
}

func (repo *GitRepo) checkout(branch string) string {
	arg := []string{repo.gitDir(), repo.workTree(), "checkout", branch}
	cmd := exec.Command(__go_cmd, arg...)
	str := utils.ExecuteCommand(cmd)
	return str
}

func (repo *GitRepo) gitDir() string {
	return fmt.Sprintf("--git-dir=%s/.git", repo.Path)
}

func (repo *GitRepo) workTree() string {
	return fmt.Sprintf("--work-tree=%s", repo.Path)
}
