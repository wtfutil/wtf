package digitalocean

import (
	"fmt"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	str := fmt.Sprintf(
		" [%s]Droplets\n\n",
		widget.settings.common.Colors.Subheading,
	)

	for idx, droplet := range widget.droplets {
		dropletName := droplet.Name

		row := fmt.Sprintf(
			"[%s] %-8s %-24s %s",
			widget.RowColor(idx),
			droplet.Status,
			dropletName,
			utils.Truncate(strings.Join(droplet.Tags, ","), 24, true),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, 33)
	}

	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
