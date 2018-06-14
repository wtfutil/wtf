// +build !linux

package power

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

const TimeRegExp = "^(?:\\d|[01]\\d|2[0-3]):[0-5]\\d"

type Battery struct {
	args   []string
	cmd    string
	result string

	Charge    string
	Remaining string
}

func NewBattery() *Battery {
	battery := Battery{
		args: []string{"-g", "batt"},
		cmd:  "pmset",
	}

	return &battery
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
	cmd := exec.Command(battery.cmd, battery.args...)
	return wtf.ExecuteCommand(cmd)
}

func (battery *Battery) parse(data string) string {
	lines := strings.Split(data, "\n")
	if len(lines) < 2 {
		return "unknown (1)"
	}

	stats := strings.Split(lines[1], "\t")
	if len(stats) < 2 {
		return "unknown (2)"
	}

	details := strings.Split(stats[1], "; ")
	if len(details) < 3 {
		return "unknown (3)"
	}

	str := ""
	str = str + fmt.Sprintf(" %10s: %s\n", "Charge", battery.formatCharge(details[0]))
	str = str + fmt.Sprintf(" %10s: %s\n", "Remaining", battery.formatRemaining(details[2]))
	str = str + fmt.Sprintf(" %10s: %s\n", "State", battery.formatState(details[1]))

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

func (battery *Battery) formatRemaining(data string) string {
	r, _ := regexp.Compile(TimeRegExp)

	result := r.FindString(data)
	if result == "" {
		result = "∞"
	}

	return result
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
