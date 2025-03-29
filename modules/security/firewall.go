package security

import (
	"os/exec"
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

func firewallStateLinux() string {
	// Check UFW first
	if hasUfw := checkUfw(); hasUfw != "" {
		return hasUfw
	}

	// Check nftables
	if hasNft := checkNftables(); hasNft != "" {
		return hasNft
	}

	// Check iptables as last resort
	if hasIpt := checkIptables(); hasIpt != "" {
		return hasIpt
	}

	return "[red]No firewall[white]"
}

func checkUfw() string {
	// First check if UFW is installed
	checkInstalled := exec.Command("which", "ufw")
	if err := checkInstalled.Run(); err != nil {
		return ""
	}

	// Then check if service is running
	cmd := exec.Command("systemctl", "is-active", "ufw")
	out := utils.ExecuteCommand(cmd)

	if strings.Contains(out, "active") {
		return "[green]Enabled (ufw)[white]"
	} else if out != "" {
		return "[red]Disabled (ufw)[white]"
	}
	return ""
}

func checkNftables() string {
	// First check if nftables is installed
	checkInstalled := exec.Command("which", "nft")
	if err := checkInstalled.Run(); err != nil {
		return ""
	}

	// Then check if service is running
	cmd := exec.Command("systemctl", "is-active", "nftables")
	out := utils.ExecuteCommand(cmd)
	if strings.Contains(out, "active") {
		return "[green]Enabled (nftables)[white]"
	} else if out != "" {
		return "[red]Disabled (nftables)[white]"
	}
	return ""
}

func checkIptables() string {
	// First check if iptables is installed
	checkInstalled := exec.Command("which", "iptables")
	if strings.Contains(utils.ExecuteCommand(checkInstalled), "not found") {
		return ""
	}

	// Check if iptables module is loaded
	cmd := exec.Command("lsmod")
	out := utils.ExecuteCommand(cmd)

	if strings.Contains(out, "ip_tables") {
		// Check for any active rules
		cmd := exec.Command("iptables", "-L")
		out := utils.ExecuteCommand(cmd)
		if strings.Contains(out, "Chain") && !strings.Contains(out, "0 references") {
			return "[green]Enabled (iptables)[white]"
		}
		return "[yellow]Loaded but unable to check rules (iptables)[white]"
	}
	return ""
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
