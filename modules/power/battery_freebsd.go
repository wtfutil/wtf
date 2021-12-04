//go:build freebsd

package power

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/wtfutil/wtf/utils"
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

// returns 3 numbers
//   1/0   = AC/battery
//   c     = battery charge percentage
//   -1/s  = charging / seconds to empty
func (battery *Battery) execute() string {
	cmd := exec.Command("apm", "-alt")
	return utils.ExecuteCommand(cmd)
}

func (battery *Battery) parse(data string) string {
	lines := strings.Split(data, "\n")
	if len(lines) < 3 {
		return "unknown"
	}
	batteryState = strings.TrimSpace(lines[0])
	charge := strings.TrimSpace(lines[1])
	timeToEmpty := "∞"
	seconds, err := strconv.Atoi(strings.TrimSpace(lines[2]))
	if err == nil && seconds >= 0 {
		h := seconds / 3600
		m := seconds % 3600 / 60
		s := seconds % 60
		timeToEmpty = fmt.Sprintf("%2d:%02d:%02d", h, m, s)
	}

	str := fmt.Sprintf(" %14s: %s%%\n", "Charge", battery.formatCharge(charge))
	str += fmt.Sprintf(" %14s: %s\n", "Remaining", timeToEmpty)
	str += fmt.Sprintf(" %14s: %s\n", "State", battery.formatState(batteryState))

	return str
}

func (battery *Battery) formatCharge(data string) string {
	percent, _ := strconv.ParseFloat(strings.Replace(data, "%", "", -1), 32)
	return utils.ColorizePercent(percent)
}

func (battery *Battery) formatState(data string) string {
	color := ""

	switch data {
	case "1":
		color = "[green]charging"
	case "0":
		color = "[yellow]discharging"
	default:
		color = "[white]unknown"
	}

	return color + "[white]"
}
