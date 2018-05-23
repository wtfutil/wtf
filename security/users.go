package security

import (
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

// http://applehelpwriter.com/2017/05/21/how-to-reveal-hidden-users/

func LoggedInUsers() []string {
	cmd := exec.Command("dscl", []string{".", "-list", "/Users"}...)
	users := wtf.ExecuteCommand(cmd)

	return cleanUsers(strings.Split(users, "\n"))
}

func cleanUsers(users []string) []string {
	rejects := []string{"_", "root", "nobody", "daemon", "Guest"}
	cleaned := []string{}

	for _, user := range users {
		clean := true

		for _, reject := range rejects {
			if strings.HasPrefix(user, reject) {
				clean = false
				continue
			}
		}

		if clean && user != "" {
			cleaned = append(cleaned, user)
		}
	}

	return cleaned
}
