package clocks

import (
	"fmt"
)

func (widget *Widget) display(clocks []Clock) {
	if len(clocks) == 0 {
		fmt.Fprintf(widget.View, "\n%s", " no timezone data available")
		return
	}

	str := "\n"
	for idx, clock := range clocks {
		str = str + fmt.Sprintf(
			" [%s]%-12s %-10s %7s[white]\n",
			widget.rowColor(idx),
			clock.Label,
			clock.LocalTime.Format(TimeFormat),
			clock.LocalTime.Format(DateFormat),
		)
	}

	fmt.Fprintf(widget.View, "%s", str)
}

func (widget *Widget) rowColor(idx int) string {
	rowCol := Config.UString("wtf.mods.clocks.colors.row.even", "lightblue")

	if idx%2 == 0 {
		rowCol = Config.UString("wtf.mods.clocks.colors.row.odd", "white")
	}

	return rowCol
}
