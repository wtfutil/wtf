package security

import (
	"os/exec"
	"runtime"
	"strings"

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

func firewallStateLinux() string {
	return "[red]NA[white]"
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
