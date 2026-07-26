//go:build windows

package power

import (
	"os/exec"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

// powerSource returns the name of the current power source on Windows
func powerSource() string {
	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command",
		"(Get-CimInstance Win32_Battery).BatteryStatus")
	out := strings.TrimSpace(utils.ExecuteCommand(cmd))

	return mapBatteryStatusToSource(out)
}

// mapBatteryStatusToSource maps a WMI BatteryStatus code to a human-readable source name
func mapBatteryStatusToSource(status string) string {
	switch status {
	case "2", "3", "6", "7", "8", "9":
		return "AC Power"
	case "1", "4", "5":
		return "Battery Power"
	default:
		return "AC Power"
	}
}
