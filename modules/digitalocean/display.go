package digitalocean

import (
	"fmt"

	"github.com/wtfutil/wtf/utils"
)

const maxColWidth = 12

func (widget *Widget) content() (string, string, bool) {
	columnSet := widget.settings.columns

	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(columnSet) < 1 {
		return title, " no columns defined", false
	}

	str := fmt.Sprintf(" [::b][%s]", widget.settings.Colors.Subheading)

	for _, colName := range columnSet {
		truncName := utils.Truncate(colName, maxColWidth, false)

		str += fmt.Sprintf("%-12s", truncName)
	}

	str += "\n"

	for idx, droplet := range widget.droplets {
		// This defines the formatting for the row, one tab-separated string for each defined column
		fmtStr := " [%s]"

		for range columnSet {
			fmtStr += "%-12s"
		}

		vals := []interface{}{
			widget.RowColor(idx),
		}

		// Dynamically access the droplet to get the requested columns values
		for _, colName := range columnSet {
			val, err := droplet.StringValueForProperty(colName)
			if err != nil {
				val = "???"
			}

			truncVal := utils.Truncate(val, maxColWidth, false)

			vals = append(vals, truncVal)
		}

		// And format, print, and color the row
		row := fmt.Sprintf(fmtStr, vals...)
		str += utils.HighlightableHelper(widget.View, row, idx, 33)
	}

	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
