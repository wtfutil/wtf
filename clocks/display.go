package clocks

import (
	"fmt"
	"strings"
)

func (widget *Widget) display() {
	locs := widget.locations(Config.UMap("wtf.mods.clocks.locations"))

	if len(locs) == 0 {
		fmt.Fprintf(widget.View, "\n%s", " no timezone data available")
		return
	}

	labels := widget.sortedLabels(locs)

	tzs := []string{}
	for idx, label := range labels {
		zoneStr := fmt.Sprintf(
			" [%s]%-12s %-10s %7s[white]",
			widget.colorFor(idx),
			label,
			locs[label].Format(TimeFormat),
			locs[label].Format(DateFormat),
		)

		tzs = append(tzs, zoneStr)
	}

	fmt.Fprintf(widget.View, "\n%s", strings.Join(tzs, "\n"))
}
