package tmuxinator

import (
	"os/exec"
	"fmt"
	"strings"
)

func fetchProjectList() []string {
	stdout, err := exec.Command("tmuxinator", "list").Output()

	if err != nil {
			fmt.Println(err.Error())
			return []string{}
	}

	output := strings.Split(string(stdout), "\n")
	projectList := strings.Split(output[1], " ")
	
	var cleanProjectList []string

	for _, str := range projectList {
		cleanString := strings.ReplaceAll(str, "\n", "")

		if str != "" {
			cleanProjectList = append(cleanProjectList, cleanString)
		}
	}

	return cleanProjectList
}

func startTmuxProject(projectName string) {
	_, err := exec.Command("tmuxinator", "start", projectName).Output()

	if err != nil {
			fmt.Println(err.Error())
	}
}
