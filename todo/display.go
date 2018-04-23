package todo

import (
	"fmt"
	//"strings"
	//"github.com/gdamore/tcell"

	"github.com/senorprogrammer/wtf/wtf"
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

		str = str + wtf.PadRow((4+len(item.Text)), widget.View) + "\n"
	}

	fmt.Fprintf(widget.View, "%s", str)
}
