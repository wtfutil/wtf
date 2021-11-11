package mercurial

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

type MercurialRepo struct {
	Branch       string
	Bookmark     string
	ChangedFiles []string
	Commits      []string
	Repository   string
	Path         string
}

func NewMercurialRepo(repoPath string, commitCount int, commitFormat string) *MercurialRepo {
	repo := MercurialRepo{Path: repoPath}

	repo.Branch = strings.TrimSpace(repo.branch())
	repo.Bookmark = strings.TrimSpace(repo.bookmark())
	repo.ChangedFiles = repo.changedFiles()
	repo.Commits = repo.commits(commitCount, commitFormat)
	repo.Repository = strings.TrimSpace(repo.Path)

	return &repo
}

/* -------------------- Unexported Functions -------------------- */

func (repo *MercurialRepo) branch() string {
	arg := []string{"branch", repo.repoPath()}

	cmd := exec.Command("hg", arg...)
	str := utils.ExecuteCommand(cmd)

	return str
}

func (repo *MercurialRepo) bookmark() string {
	bookmark, err := os.ReadFile(path.Join(repo.Path, ".hg", "bookmarks.current"))
	if err != nil {
		return ""
	}
	return string(bookmark)
}

func (repo *MercurialRepo) changedFiles() []string {
	arg := []string{"status", repo.repoPath()}

	cmd := exec.Command("hg", arg...)
	str := utils.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *MercurialRepo) commits(commitCount int, commitFormat string) []string {
	numStr := fmt.Sprintf("-l %d", commitCount)
	commitStr := fmt.Sprintf("--template=\"%s\n\"", commitFormat)

	arg := []string{"log", repo.repoPath(), numStr, commitStr}

	cmd := exec.Command("hg", arg...)
	str := utils.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *MercurialRepo) pull() string {
	arg := []string{"pull", repo.repoPath()}
	cmd := exec.Command("hg", arg...)
	str := utils.ExecuteCommand(cmd)
	return str
}

func (repo *MercurialRepo) checkout(branch string) string {
	arg := []string{"checkout", repo.repoPath(), branch}
	cmd := exec.Command("hg", arg...)
	str := utils.ExecuteCommand(cmd)
	return str
}

func (repo *MercurialRepo) repoPath() string {
	return fmt.Sprintf("--repository=%s", repo.Path)
}
