// +build !windows

package security

// http://applehelpwriter.com/2017/05/21/how-to-reveal-hidden-users/

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

/* -------------------- Exported Functions -------------------- */

func LoggedInUsers() []string {
	switch runtime.GOOS {
	case "linux":
		return loggedInUsersLinux()
	case "darwin":
		return loggedInUsersMacOs()
	default:
		return []string{}
	}
}

/* -------------------- Unexported Functions -------------------- */

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

func loggedInUsersLinux() []string {
	cmd := exec.Command("who", "-us")
	users := wtf.ExecuteCommand(cmd)

	cleaned := []string{}

	for _, user := range strings.Split(users, "\n") {
		clean := true
		col := strings.Split(user, " ")

		if len(col) > 0 {
			for _, cleanedU := range cleaned {
				u := strings.TrimSpace(col[0])
				if len(u) == 0 || strings.Compare(cleanedU, col[0]) == 0 {
					clean = false
				}
			}

			if clean {
				cleaned = append(cleaned, col[0])
			}
		}

	}

	return cleaned
}

func loggedInUsersMacOs() []string {
	cmd := exec.Command("dscl", []string{".", "-list", "/Users"}...)
	users := wtf.ExecuteCommand(cmd)

	return cleanUsers(strings.Split(users, "\n"))
}
