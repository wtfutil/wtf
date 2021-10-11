package security

import (
	"bytes"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

/* -------------------- Exported Functions -------------------- */

func FirewallState() string {
	switch runtime.GOOS {
	case "darwin":
		return firewallStateMacOS()
	case "linux":
		return firewallStateLinux()
	case "windows":
		return firewallStateWindows()
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
	case "windows":
		return firewallStealthStateWindows()
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

		if strings.Contains(o.String(), "inactive") {
			return "[red]Disabled[white]"
		} else {
			return "[green]Enabled[white]"
		}
	} else {
		return "[red]N/A[white]"
	}
}

func firewallStateMacOS() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")
	str := utils.ExecuteCommand(cmd)

	return statusLabel(str)
}

func firewallStateWindows() string {
	// The raw way to do this in PS, not using netsh, nor registry, is the following:
	//   if (((Get-NetFirewallProfile | select name,enabled)
	//                                | where { $_.Enabled -eq $True } | measure ).Count -eq 3)
	//   { Write-Host "OK" -ForegroundColor Green} else { Write-Host "OFF" -ForegroundColor Red }

	cmd := exec.Command("powershell.exe", "-NoProfile",
		"-Command", "& { ((Get-NetFirewallProfile | select name,enabled) | where { $_.Enabled -eq $True } | measure ).Count }")

	fwStat := utils.ExecuteCommand(cmd)
	fwStat = strings.TrimSpace(fwStat) // Always sanitize PowerShell output:  "3\r\n"

	switch fwStat {
	case "3":
		return "[green]Good[white] (3/3)"
	case "2":
		return "[orange]Poor[white] (2/3)"
	case "1":
		return "[yellow]Bad[white] (1/3)"
	case "0":
		return "[red]Disabled[white]"
	default:
		return "[white]N/A[white]"
	}
}

/* -------------------- Getting Stealth State ------------------- */
// "Stealth": Not responding to pings from unauthorized devices

func firewallStealthStateLinux() string {
	return "[white]N/A[white]"
}

func firewallStealthStateMacOS() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")
	str := utils.ExecuteCommand(cmd)

	return statusLabel(str)
}

func firewallStealthStateWindows() string {
	return "[white]N/A[white]"
}

func statusLabel(str string) string {
	label := "off"

	if strings.Contains(str, "enabled") {
		label = "on"
	}

	return label
}
