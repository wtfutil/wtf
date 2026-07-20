//go:build windows

package power

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

type Battery struct {
	result string

	Charge    string
	Remaining string
}

func NewBattery() *Battery {
	return &Battery{}
}

/* -------------------- Exported Functions -------------------- */

func (battery *Battery) Refresh() {
	battery.result = battery.queryWindows()
}

func (battery *Battery) String() string {
	return battery.result
}

/* -------------------- Unexported Functions -------------------- */

func (battery *Battery) queryWindows() string {
	// Use WMI to get battery info
	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command",
		"Get-CimInstance Win32_Battery | Select-Object EstimatedChargeRemaining, BatteryStatus, EstimatedRunTime | Format-List")
	out := utils.ExecuteCommand(cmd)

	if strings.TrimSpace(out) == "" {
		return " no battery found"
	}

	lines := strings.Split(out, "\n")
	info := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			info[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	charge := info["EstimatedChargeRemaining"]
	status := battery.formatStatus(info["BatteryStatus"])
	remaining := info["EstimatedRunTime"]

	str := ""
	if percent, err := strconv.ParseFloat(charge, 64); err == nil {
		str += fmt.Sprintf(" %14s: %s\n", "Charge", utils.ColorizePercent(percent))
	} else {
		str += fmt.Sprintf(" %14s: %s\n", "Charge", charge+"%")
	}

	if remaining == "" || remaining == "0" {
		remaining = "-"
	}
	str += fmt.Sprintf(" %14s: %s min\n", "Remaining", remaining)
	str += fmt.Sprintf(" %14s: %s\n", "State", status)

	return str
}

func (battery *Battery) formatStatus(code string) string {
	// Win32_Battery BatteryStatus codes
	switch strings.TrimSpace(code) {
	case "1":
		return "[yellow]discharging[white]"
	case "2":
		return "[white]AC connected[white]"
	case "3":
		return "[white]fully charged[white]"
	case "4":
		return "[yellow]low[white]"
	case "5":
		return "[red]critical[white]"
	case "6":
		return "[green]charging[white]"
	case "7":
		return "[green]charging (high)[white]"
	case "8":
		return "[green]charging (low)[white]"
	case "9":
		return "[red]critical (charging)[white]"
	default:
		return "[white]unknown[white]"
	}
}
