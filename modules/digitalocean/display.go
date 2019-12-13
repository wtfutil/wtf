package digitalocean

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	str := fmt.Sprintf(
		" [%s]Droplets[-:-:-]\n\n",
		widget.settings.common.Colors.Subheading,
	)

	for idx, droplet := range widget.droplets {
		dropletName := droplet.Name

		row := fmt.Sprintf(
			"[%s] %-24s %-8s %s[-:-:-]",
			widget.RowColor(idx),
			utils.Truncate(tview.Escape(dropletName), 24, true),
			widget.statusColor(droplet.Status, widget.RowColor(idx)),
			utils.Truncate(strings.Join(droplet.Tags, ","), 24, true),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(dropletName))
	}

	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}

func (widget *Widget) statusColor(status string, rowColor string) string {
	color := rowColor

	switch status {
	case "active":
		color = "green"
	case "archive":
		color = "gray"
	case "new":
		color = "yellow"
	case "off":
		color = "gray"
	}

	return fmt.Sprintf("[%s]%s[%s]", color, status, rowColor)
}
