package client

import (
	"os/exec"
	"fmt"
	"strings"
	"github.com/wtfutil/wtf/utils"
)

func ProjectList() []string {
	cmdString := `tmuxinator list | grep -v "tmuxinator projects:" | tr -s ' ' | tr '\n' ' '`

	cmd := exec.Command("sh", "-c", cmdString)

	return strings.Split(utils.ExecuteCommand(cmd), " ")
}

func StartProject(projectName string) {
	_, err := exec.Command("tmuxinator", "start", projectName).Output()

	if err != nil {
			fmt.Println(err.Error())
	}
}

func EditProject(projectName string) {
	subcommand := fmt.Sprintf("tmuxinator edit %s", projectName)
	_, err := exec.Command("tmux", "new-window", subcommand).Output()

	if err != nil {
			fmt.Println(err.Error())
	}
}

func DeleteProject(projectName string) {
	_, err := exec.Command("tmuxinator", "delete", projectName).Output()

	if err != nil {
			fmt.Println(err.Error())
	}
}

func CopyProject(leftProj, rightProj string) {
	subcommand := fmt.Sprintf("tmuxinator copy %s %s", leftProj, rightProj)
	_, err := exec.Command("tmux", "new-window", subcommand).Output()

	if err != nil {
			fmt.Println(err.Error())
	}
}
