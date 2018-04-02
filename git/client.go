package git

import (
	//"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Client struct {
	CommitCount int
	Repository  string
}

func NewClient() *Client {
	client := Client{
		CommitCount: 10,
		Repository:  "/Users/Chris/Documents/Lendesk/core-api",
	}

	return &client
}

/* -------------------- Unexported Functions -------------------- */

func (client *Client) CurrentBranch() string {
	arg := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	cmd := exec.Command("git", arg...)
	str := executeCommand(cmd)

	return str
}

func (client *Client) ChangedFiles() []string {
	arg := []string{"status", "--porcelain"}
	cmd := exec.Command("git", arg...)
	str := executeCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

func (client *Client) Commits() []string {
	arg := []string{"log", "--date=format:\"%b %d, %Y\"", "-n 10", "--pretty=format:\"[forestgreen]%h [white]%s [grey]%an on %cd[white]\""}
	cmd := exec.Command("git", arg...)
	str := executeCommand(cmd)

	data := strings.Split(str, "\n")

	return data
}

/* -------------------- Unexported Functions -------------------- */

func executeCommand(cmd *exec.Cmd) string {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "err"
	}

	if err := cmd.Start(); err != nil {
		return "err"
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	return str
}
