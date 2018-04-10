package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

/* -------------------- Exported Functions -------------------- */

func (client *Client) Branch() string {
	arg := []string{client.gitDir(), client.workTree(), "rev-parse", "--abbrev-ref", "HEAD"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}

func (client *Client) ChangedFiles() []string {
	arg := []string{client.gitDir(), client.workTree(), "status", "--porcelain"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (client *Client) Commits() []string {
	numStr := fmt.Sprintf("-n %d", Config.UInt("wtf.mods.git.commitCount", 10))

	arg := []string{client.gitDir(), client.workTree(), "log", "--date=format:\"%b %d, %Y\"", numStr, "--pretty=format:\"[forestgreen]%h [white]%s [grey]%an on %cd[white]\""}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (client *Client) Repository() string {
	arg := []string{client.gitDir(), client.workTree(), "rev-parse", "--show-toplevel"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}

/* -------------------- Exported Functions -------------------- */

func (client *Client) gitDir() string {
	return fmt.Sprintf("--git-dir=%s/.git", Config.UString("wtf.mods.git.repository"))
}

func (client *Client) workTree() string {
	return fmt.Sprintf("--work-tree=%s", Config.UString("wtf.mods.git.repository"))
}
