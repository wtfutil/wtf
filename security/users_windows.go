// +build windows

package security

import (
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

func LoggedInUsers() []string {
	cmd := exec.Command("powershell.exe", "(query user) -replace '\\s{2,}', ','")
	users := wtf.ExecuteCommand(cmd)
	return cleanUsers(strings.Split(users, "\n")[1:])
}

func cleanUsers(users []string) []string {
	cleaned := make([]string, 0)
	for _, user := range users {
		usr := strings.Split(user, ",")
		cleaned = append(cleaned, usr[0])
	}
	return cleaned
}
