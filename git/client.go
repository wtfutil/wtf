package git

import (
	//"fmt"
	"io/ioutil"
	"os/exec"
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
	files := []string{}

	return files
}

func (client *Client) Commits() []string {
	files := []string{}

	return files
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
