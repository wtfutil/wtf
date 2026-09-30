package clocks

import "fmt"

// labelWidth computes the column width needed for clock labels,
// using a minimum of 12.
func labelWidth(clocks []Clock) int {
	width := 12
	for _, clock := range clocks {
		if len(clock.Label) > width {
			width = len(clock.Label) + 2
		}
	}
	return width
}

// formatClocks builds the display string for a set of clocks.
// rowColor is called with the row index to determine the color tag.
func formatClocks(clocks []Clock, dateFormat string, timeFormat string, rowColor func(int) string) string {
	if len(clocks) == 0 {
		return fmt.Sprintf("\n%s", " no timezone data available")
	}

	locWidth := labelWidth(clocks)
	str := ""
	for idx, clock := range clocks {
		str += fmt.Sprintf(
			" [%s]%-*s %-10s %7s[white]\n",
			rowColor(idx),
			locWidth,
			clock.Label,
			clock.Time(timeFormat),
			clock.Date(dateFormat),
		)
	}
	return str
}

func (widget *Widget) display(clocks []Clock, dateFormat string, timeFormat string) {
	str := formatClocks(clocks, dateFormat, timeFormat, widget.CommonSettings().RowColor)
	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, str, false })
}
