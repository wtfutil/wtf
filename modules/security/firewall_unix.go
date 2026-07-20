//go:build !windows

package security

import (
	"fmt"
	"os/exec"
	"syscall"
)

func firewallStateLinux() string {
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

func firewallStealthStateLinux() string {
	return "[white]N/A[white]"
}
