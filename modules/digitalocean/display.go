package digitalocean

import (
	"fmt"

	"github.com/wtfutil/wtf/utils"
)

const maxColWidth = 10

// defaultColumns defines the default set of columns to display in the widget
// This can be over-ridden in the cofig by explicitly defining a set of columns
var defaultColumns = []string{
	"Name",
	"Status",
	"Vcpus",
	"Disk",
	"Memory",
	"Region.Slug",
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(defaultColumns) < 1 {
		return title, " no columns defined", false
	}

	str := fmt.Sprintf(" [::b][%s]", widget.settings.common.Colors.Subheading)

	for _, colName := range defaultColumns {
		truncName := utils.Truncate(colName, maxColWidth, false)

		str += fmt.Sprintf("%-10s", truncName)
	}

	str += "\n"

	for idx, droplet := range widget.droplets {
		// This defines the formatting for the row, one tab-seperated string
		// for each defined column
		fmtStr := " [%s]"
		for range defaultColumns {
			fmtStr += "%-10s"
		}

		vals := []interface{}{
			widget.RowColor(idx),
		}

		// Dynamically access the droplet to get the requested columns values
		for _, colName := range defaultColumns {
			val := droplet.ValueForColumn(colName)
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
