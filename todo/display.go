package todo

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	widget.View.Clear()

	maxLen := widget.longestLineLen(widget.list.Items)

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

		str = str + wtf.PadRow((4+len(item.Text)), (4+maxLen)) + "\n"
	}

	fmt.Fprintf(widget.View, "%s", str)
}

// longestLineLen returns the length of the longest todo item line
func (widget *Widget) longestLineLen(items []*Item) int {
	maxLen := 0

	for _, item := range items {
		if len(item.Text) > maxLen {
			maxLen = len(item.Text)
		}
	}

	return maxLen
}
