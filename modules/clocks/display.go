package clocks

import (
	"fmt"
)

func (widget *Widget) display(clocks []Clock, dateFormat string, timeFormat string) {
	str := ""
	if len(clocks) == 0 {
		str = fmt.Sprintf("\n%s", " no timezone data available")
	} else {

		for idx, clock := range clocks {
			rowColor := widget.settings.colors.rows.odd

			if idx%2 == 0 {
				rowColor = widget.settings.colors.rows.even
			}

			str += fmt.Sprintf(
				" [%s]%-12s %-10s %7s[white]\n",
				rowColor,
				clock.Label,
				clock.Time(timeFormat),
				clock.Date(dateFormat),
			)
		}
	}

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, str, false })
}
