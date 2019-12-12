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
		" [%s]Droplets[white]\n",
		widget.settings.common.Colors.Subheading,
	)

	for idx, droplet := range widget.droplets {
		dropletName := droplet.Name

		row := fmt.Sprintf(
			"[%s] %-24s %-8s %s[white]",
			widget.RowColor(idx),
			utils.Truncate(tview.Escape(dropletName), 24, true),
			utils.Truncate(droplet.Status, 8, true),
			utils.Truncate(strings.Join(droplet.Tags, ","), 24, true),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(dropletName))
	}

	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
