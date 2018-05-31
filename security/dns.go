package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

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
	cmd := exec.Command("networksetup", "-getdnsservers", "Wi-Fi")
	out := wtf.ExecuteCommand(cmd)
	records := strings.Split(out, "\n")

	if len(records) > 0 {
		return records
	} else {
		return []string{}
	}
}

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
