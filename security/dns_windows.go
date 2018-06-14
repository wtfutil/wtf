// +build windows

package security

import (
	"os/exec"

	"github.com/senorprogrammer/wtf/wtf"
)

func DnsServers() []string {
	cmd := exec.Command("powershell.exe", "Get-DnsClientServerAddress | Select-Object –ExpandProperty ServerAddresses")
	return []string{wtf.ExecuteCommand(cmd)}
}
