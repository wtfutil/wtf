// +build linux

package power

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

var batteryState string

type Battery struct {
	args   []string
	cmd    string
	result string

	Charge    string
	Remaining string
}

func NewBattery() *Battery {
	return &Battery{}
}

/* -------------------- Exported Functions -------------------- */

func (battery *Battery) Refresh() {
	data := battery.execute()
	battery.result = battery.parse(data)
}

func (battery *Battery) String() string {
	return battery.result
}

/* -------------------- Unexported Functions -------------------- */

func (battery *Battery) execute() string {
	cmd := exec.Command("upower", "-e")
	lines := strings.Split(wtf.ExecuteCommand(cmd), "\n")
	var target string
	for _, l := range lines {
		if strings.Contains(l, "/battery") {
			target = l
			break
		}
	}
	cmd = exec.Command("upower", "-i", target)
	return wtf.ExecuteCommand(cmd)
}

func (battery *Battery) parse(data string) string {
	lines := strings.Split(data, "\n")
	if len(lines) < 2 {
		return "unknown"
	}
	table := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}
		table[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	if s := table["time to empty"]; s == "" {
		table["time to empty"] = "∞"
	}
	str := ""
	str = str + fmt.Sprintf(" %10s: %s\n", "Charge", battery.formatCharge(table["percentage"]))
	str = str + fmt.Sprintf(" %10s: %s\n", "Remaining", table["time to empty"])
	str = str + fmt.Sprintf(" %10s: %s\n", "State", battery.formatState(table["state"]))
	if s := table["time to full"]; s != "" {
		str = str + fmt.Sprintf(" %10s: %s\n", "TimeToFull", table["time to full"])
	}
	batteryState = table["state"]
	return str
}

func (battery *Battery) formatCharge(data string) string {
	percent, _ := strconv.ParseFloat(strings.Replace(data, "%", "", -1), 32)
	color := ""

	switch {
	case percent >= 70:
		color = "[green]"
	case percent >= 35:
		color = "[yellow]"
	default:
		color = "[red]"
	}

	return color + data + "[white]"
}

func (battery *Battery) formatState(data string) string {
	color := ""

	switch data {
	case "charging":
		color = "[green]"
	case "discharging":
		color = "[yellow]"
	default:
		color = "[white]"
	}

	return color + data + "[white]"
}
