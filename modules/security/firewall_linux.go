//go:build linux

package security

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/wtfutil/wtf/utils"
)

func FirewallState() string {
	// Check UFW first
	if hasUfw := checkUfw(); hasUfw != "" {
		return hasUfw
	}

	// Check Firewalld
	if hasFirewalld := checkFirewalld(); hasFirewalld != "" {
		return hasFirewalld
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

func FirewallStealthState() string {
	return "[white]N/A[white]"
}

func checkFirewalld() string {
	checkInstalled := exec.Command("which", "firewall-cmd")
	if err := checkInstalled.Run(); err != nil {
		return ""
	}

	cmd := exec.Command("firewall-cmd", "--state")
	err := cmd.Start()
	if err != nil {
		return "[red]Failed to start status check (firewalld)[white]"
	}

	err = cmd.Wait()
	if err == nil {
		return "[green]Active (firewalld)[white]"
	}

	if exitError, ok := err.(*exec.ExitError); ok {
		sc := exitError.Sys().(syscall.WaitStatus).ExitStatus()
		switch sc {
		case 251:
			return "[yellow]Running but failed (firewalld)[white]"
		case 252:
			return "[red]Not running (firewalld)[white]"
		default:
			return fmt.Sprintf("[red]Unexpected state (%d) assume not running (firewalld)[white]", sc)
		}
	} else {
		return fmt.Sprintf("[red] Error waiting for command: %v (firewalld)[white]", err)
	}
}

func checkUfw() string {
	checkInstalled := exec.Command("which", "ufw")
	if err := checkInstalled.Run(); err != nil {
		return ""
	}

	cmd := exec.Command("systemctl", "is-active", "ufw")
	err := cmd.Run()
	if err == nil {
		return "[green]Enabled (ufw)[white]"
	}
	return "[red]Disabled (ufw)[white]"
}

func checkNftables() string {
	checkInstalled := exec.Command("which", "nft")
	if err := checkInstalled.Run(); err != nil {
		return ""
	}

	cmd := exec.Command("systemctl", "is-active", "nftables")
	err := cmd.Run()
	if err == nil {
		return "[green]Enabled (nftables)[white]"
	}
	return "[red]Disabled (nftables)[white]"
}

func checkIptables() string {
	checkInstalled := exec.Command("which", "iptables")
	if strings.Contains(utils.ExecuteCommand(checkInstalled), "not found") {
		return ""
	}

	cmd := exec.Command("lsmod")
	out := utils.ExecuteCommand(cmd)

	if strings.Contains(out, "ip_tables") {
		cmd := exec.Command("iptables", "-L")
		out := utils.ExecuteCommand(cmd)
		if strings.Contains(out, "Chain") && !strings.Contains(out, "0 references") {
			return "[green]Enabled (iptables)[white]"
		}
		return "[yellow]Loaded but unable to check rules (iptables)[white]"
	}
	return ""
}
