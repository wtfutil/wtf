package security

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

/* -------------------- Exported Functions -------------------- */

func FirewallState() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")
	str := executeCommand(cmd)

	return str
}

func FirewallStealthState() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")
	str := executeCommand(cmd)

	return str
}

/* -------------------- Unexported Functions -------------------- */

func executeCommand(cmd *exec.Cmd) string {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return firewallStr("err")
	}

	if err := cmd.Start(); err != nil {
		return firewallStr("err")
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	return firewallStr(str)
}

func firewallStr(str string) string {
	icon := "[red]off[white]"

	if strings.Contains(str, "enabled") {
		icon = "[green]on[white]"
	}

	return icon
}
