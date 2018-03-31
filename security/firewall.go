package security

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

func FirewallState() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return firewallIcon("err")
	}

	if err := cmd.Start(); err != nil {
		return firewallIcon("err")
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	return firewallIcon(str)
}

func FirewallStealthState() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return firewallIcon("err")
	}

	if err := cmd.Start(); err != nil {
		return firewallIcon("err")
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	return firewallIcon(str)
}

func firewallIcon(str string) string {
	icon := "[red]off[white]"

	if strings.Contains(str, "enabled") {
		icon = "[green]on[white]"
	}

	return icon
}
