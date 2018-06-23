// +build !windows

package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

/* -------------------- Exported Functions -------------------- */

func DnsServers() []string {
	switch runtime.GOOS {
	case "linux":
		return dnsLinux()
	case "darwin":
		return dnsMacOS()
	default:
		return []string{runtime.GOOS}
	}
}

/* -------------------- Unexported Functions -------------------- */

func dnsLinux() []string {
	// This may be very Ubuntu specific
	cmd := exec.Command("nmcli", "device", "show")
	out := wtf.ExecuteCommand(cmd)

	lines := strings.Split(out, "\n")

	dns := []string{}

	for _, l := range lines {
		if strings.HasPrefix(l, "IP4.DNS") {
			parts := strings.Split(l, ":")
			dns = append(dns, strings.TrimSpace(parts[1]))
		}
	}
	return dns
}

func dnsMacOS() []string {
	cmdString := `scutil --dns | head -n 7 | grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}'`
	cmd := exec.Command("sh", "-c", cmdString)
	out := wtf.ExecuteCommand(cmd)

	lines := strings.Split(out, "\n")

	if len(lines) > 0 {
		return lines
	} else {
		return []string{}
	}
}
