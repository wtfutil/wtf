package clocks

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

const TimeFormat = "15:04 MST"
const DateFormat = "Jan 2"

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 🕗 World Clocks ", "clocks"),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.View.Clear()

	fmt.Fprintf(widget.View, "\n%s", widget.locations())

	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) locations() string {
	timezones := Timezones(Config.UMap("wtf.mods.clocks.locations"))

	if len(timezones) == 0 {
		return ""
	}

	// All this is to display the time entries in alphabetical order
	labels := []string{}
	for label, _ := range timezones {
		labels = append(labels, label)
	}

	sort.Strings(labels)

	tzs := []string{}
	for idx, label := range labels {
		rowColor := Config.UString("wtf.mods.clocks.rowcolors.even", "lightblue")
		if idx%2 == 0 {
			rowColor = Config.UString("wtf.mods.clocks.rowcolors.odd", "white")
		}

		zoneStr := fmt.Sprintf(
			" [%s]%-12s %-10s %7s[white]",
			rowColor, label,
			timezones[label].Format(TimeFormat),
			timezones[label].Format(DateFormat),
		)
		tzs = append(tzs, zoneStr)
	}

	return strings.Join(tzs, "\n")
}
