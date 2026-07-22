//go:build windows

package security

import (
	"os/exec"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

/* -------------------- Exported Functions -------------------- */

func FirewallState() string {
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

func FirewallStealthState() string {
	return "[white]N/A[white]"
}
