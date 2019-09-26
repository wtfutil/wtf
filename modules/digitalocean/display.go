package digitalocean

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	str := " [red]Droplets[white]\n"

	for idx, droplet := range widget.droplets {
		dropletName := droplet.Name
		ipV4, _ := droplet.PublicIPv4()

		row := fmt.Sprintf(
			"[%s]%3d %s %s %s[white]",
			widget.RowColor(idx),
			(idx + 1),
			tview.Escape(dropletName),
			droplet.Region.Slug,
			ipV4,
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(dropletName))
	}

	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
