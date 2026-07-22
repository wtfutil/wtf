//go:build darwin

package security

import (
	"os/exec"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

func FirewallState() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")
	str := utils.ExecuteCommand(cmd)

	return statusLabel(str)
}

func FirewallStealthState() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")
	str := utils.ExecuteCommand(cmd)

	return statusLabel(str)
}

func statusLabel(str string) string {
	if strings.Contains(str, "enabled") {
		return "on"
	}
	return "off"
}
