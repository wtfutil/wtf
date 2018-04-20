package todo

import (
	"fmt"
)

func (widget *Widget) display() {
	widget.View.Clear()

	title := fmt.Sprintf(" 📝 %s ", widget.FilePath)
	widget.View.SetTitle(title)

	str := ""
	for idx, item := range widget.list.Items {
		foreColor, backColor := "white", "black"

		if widget.View.HasFocus() && idx == widget.list.selected {
			foreColor, backColor = "black", "olive"
		}

		str = str + fmt.Sprintf(
			"[%s:%s]|%s| %s[white]\n",
			foreColor,
			backColor,
			item.CheckMark(),
			item.Text,
		)
	}

	fmt.Fprintf(widget.View, "%s", str)
}
