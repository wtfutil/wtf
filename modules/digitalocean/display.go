package digitalocean

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	str := fmt.Sprintf(
		" [%s]Droplets [grey](%d)[white]\n",
		widget.settings.common.Colors.Subheading,
		len(widget.droplets),
	)

	for idx, droplet := range widget.droplets {
		dropletName := droplet.Name

		row := fmt.Sprintf(
			"[%s] %14s %28s %12s[white]",
			widget.RowColor(idx),
			wtf.PrettyDate(strings.Split(droplet.Created, "T")[0]),
			tview.Escape(dropletName),
			strings.Join(droplet.Tags, ","),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(dropletName))
	}

	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
