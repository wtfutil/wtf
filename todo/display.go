package todo

import (
	"fmt"
	"strings"
	//"github.com/gdamore/tcell"
)

const checkWidth = 4

func (widget *Widget) display() {
	widget.View.Clear()

	str := ""
	for idx, item := range widget.list.Items {
		foreColor, backColor := "white", "black"

		if item.Checked {
			foreColor = Config.UString("wtf.mods.todo.colors.checked", "white")
		}

		if widget.View.HasFocus() && idx == widget.list.selected {
			foreColor = Config.UString("wtf.mods.todo.colors.highlight.fore", "black")
			backColor = Config.UString("wtf.mods.todo.colors.highlight.back", "white")
		}

		str = str + fmt.Sprintf(
			"[%s:%s]|%s| %s[white]",
			foreColor,
			backColor,
			item.CheckMark(),
			item.Text,
		)

		str = widget.padLine(str, item) + "\n"
	}

	fmt.Fprintf(widget.View, "%s", str)
}

// Pad with spaces to get full-line highlighting
func (widget *Widget) padLine(str string, item *Item) string {
	_, _, w, _ := widget.View.GetInnerRect()

	padSize := w - checkWidth - len(item.Text)
	if padSize < 0 {
		padSize = 0
	}

	return str + strings.Repeat(" ", padSize)
}
