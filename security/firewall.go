package security

import (
	"os/exec"
	"runtime"
	"strings"
	"bytes"
	"os/user"

	"github.com/senorprogrammer/wtf/wtf"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

/* -------------------- Exported Functions -------------------- */

func FirewallState() string {
	switch runtime.GOOS {
	case "linux":
		return firewallStateLinux()
	case "darwin":
		return firewallStateMacOS()
	default:
		return ""
	}
}

func FirewallStealthState() string {
	switch runtime.GOOS {
	case "linux":
		return firewallStealthStateLinux()
	case "darwin":
		return firewallStealthStateMacOS()
	default:
		return ""
	}
}

/* -------------------- Unexported Functions -------------------- */

func firewallStateLinux() string { // might be very Ubuntu specific
	user, _ := user.Current()

	if strings.Contains(user.Username, "root") {
		cmd := exec.Command("ufw", "status")

		var o bytes.Buffer
		cmd.Stdout = &o
		if err := cmd.Run(); err != nil {
			return "[red]NA[white]"
		}

		if strings.Contains(o.String(), "active") {
			return "[green]Enabled[white]"
		} else {
			return "[red]Disabled[white]"
		}
	} else {
		return "[red]NA[white]"
	}
}

func firewallStateMacOS() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")
	str := wtf.ExecuteCommand(cmd)

	return statusLabel(str)
}

func firewallStealthStateLinux() string {
	return "[red]NA[white]"
}

func firewallStealthStateMacOS() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")
	str := wtf.ExecuteCommand(cmd)

	return statusLabel(str)
}

func statusLabel(str string) string {
	label := "off"

	if strings.Contains(str, "enabled") {
		label = "on"
	}

	return label
}
