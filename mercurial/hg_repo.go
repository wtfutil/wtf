package mercurial

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type MercurialRepo struct {
	Branch       string
	Bookmark     string
	ChangedFiles []string
	Commits      []string
	Repository   string
	Path         string
}

func NewMercurialRepo(repoPath string) *MercurialRepo {
	repo := MercurialRepo{Path: repoPath}

	repo.Branch = strings.TrimSpace(repo.branch())
	repo.Bookmark = strings.TrimSpace(repo.bookmark())
	repo.ChangedFiles = repo.changedFiles()
	repo.Commits = repo.commits()
	repo.Repository = strings.TrimSpace(repo.Path)

	return &repo
}

/* -------------------- Unexported Functions -------------------- */

func (repo *MercurialRepo) branch() string {
	arg := []string{"branch", repo.repoPath()}

	cmd := exec.Command("hg", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}

func (repo *MercurialRepo) bookmark() string {
	bookmark, err := ioutil.ReadFile(path.Join(repo.Path, ".hg", "bookmarks.current"))
	if err != nil {
		return ""
	}
	return string(bookmark)
}

func (repo *MercurialRepo) changedFiles() []string {
	arg := []string{"status", repo.repoPath()}

	cmd := exec.Command("hg", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *MercurialRepo) commits() []string {
	numStr := fmt.Sprintf("-l %d", wtf.Config.UInt("wtf.mods.mercurial.commitCount", 10))

	commitFormat := wtf.Config.UString("wtf.mods.mercurial.commitFormat", "[forestgreen]{rev}:{phase} [white]{desc|firstline|strip} [grey]{author|person} {date|age}[white]")
	commitStr := fmt.Sprintf("--template=\"%s\n\"", commitFormat)

	arg := []string{"log", repo.repoPath(), numStr, commitStr}

	cmd := exec.Command("hg", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (repo *MercurialRepo) pull() string {
	arg := []string{"pull", repo.repoPath()}
	cmd := exec.Command("hg", arg...)
	str := wtf.ExecuteCommand(cmd)
	return str
}

func (repo *MercurialRepo) checkout(branch string) string {
	arg := []string{"checkout", repo.repoPath(), branch}
	cmd := exec.Command("hg", arg...)
	str := wtf.ExecuteCommand(cmd)
	return str
}

func (repo *MercurialRepo) repoPath() string {
	return fmt.Sprintf("--repository=%s", repo.Path)
}
