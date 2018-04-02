package security

import (
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

/* -------------------- Exported Functions -------------------- */

func FirewallState() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")
	str := wtf.ExecuteCommand(cmd)

	return status(str)
}

func FirewallStealthState() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")
	str := wtf.ExecuteCommand(cmd)

	return status(str)
}

/* -------------------- Unexported Functions -------------------- */

func status(str string) string {
	icon := "[red]off[white]"

	if strings.Contains(str, "enabled") {
		icon = "[green]on[white]"
	}

	return icon
}
