package security

import (
	"os/exec"

	"github.com/senorprogrammer/wtf/wtf"
)

const dnsCmd = "networksetup"

func DnsServers() string {
	cmd := exec.Command(dnsCmd, "-getdnsservers", "Wi-Fi")
	return wtf.ExecuteCommand(cmd)
}
