//go:build !linux && !freebsd

package power

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

const (
	timeRegExp = "^(?:\\d|[01]\\d|2[0-3]):[0-5]\\d"
)

type Battery struct {
	args   []string
	cmd    string
	result string

	Charge    string
	Remaining string
}

func NewBattery() *Battery {
	battery := &Battery{
		args: []string{"-g", "batt"},
		cmd:  "pmset",
	}

	return battery
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
	return utils.ExecuteCommand(cmd)
}

func (battery *Battery) parse(data string) string {
	lines := strings.Split(data, "\n")
	if len(lines) < 2 {
		return msgNoBattery
	}

	stats := strings.Split(lines[1], "\t")
	if len(stats) < 2 {
		return msgNoBattery
	}

	details := strings.Split(stats[1], "; ")
	if len(details) < 3 {
		return msgNoBattery
	}

	str := ""
	str = str + fmt.Sprintf(" %14s: %s\n", "Charge", battery.formatCharge(details[0]))
	str = str + fmt.Sprintf(" %14s: %s\n", "Remaining", battery.formatRemaining(details[2]))
	str = str + fmt.Sprintf(" %14s: %s\n", "State", battery.formatState(details[1]))

	return str
}

func (battery *Battery) formatCharge(data string) string {
	percent, _ := strconv.ParseFloat(strings.Replace(data, "%", "", -1), 32)
	return utils.ColorizePercent(percent)
}

func (battery *Battery) formatRemaining(data string) string {
	r, _ := regexp.Compile(timeRegExp)

	result := r.FindString(data)
	if result == "" || result == "0:00" {
		result = "-"
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
