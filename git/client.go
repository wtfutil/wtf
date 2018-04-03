package git

import (
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type Client struct {
	CommitCount int
}

func NewClient() *Client {
	client := Client{
		CommitCount: 10,
	}

	return &client
}

/* -------------------- Unexported Functions -------------------- */

func (client *Client) Branch() string {
	arg := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}

func (client *Client) ChangedFiles() []string {
	arg := []string{"status", "--porcelain"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (client *Client) Commits() []string {
	arg := []string{"log", "--date=format:\"%b %d, %Y\"", "-n 10", "--pretty=format:\"[forestgreen]%h [white]%s [grey]%an on %cd[white]\""}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (client *Client) Repository() string {
	arg := []string{"rev-parse", "--show-toplevel"}
	cmd := exec.Command("git", arg...)
	str := wtf.ExecuteCommand(cmd)

	return str
}
