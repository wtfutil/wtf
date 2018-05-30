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
	checked := []*Item{}
	uncheckedLen := 0
	var selected *Item
	var newList List
	for idx, item := range widget.list.Items {
		foreColor, backColor := "white", "black"

		// save the selected one
		if idx == widget.list.selected {
			selected = item
		}

		if item.Checked {
			checked = append(checked, item)
			continue
		}

		uncheckedLen++

		if widget.View.HasFocus() && item == selected {
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

		newList.Items = append(newList.Items, item)
	}
	for _, item := range checked {
		foreColor, backColor := Config.UString("wtf.mods.todo.colors.checked", "white"), "black"

		if widget.View.HasFocus() && item == selected {
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

		newList.Items = append(newList.Items, item)
	}

	// update new index of selected item
	for idx, item := range newList.Items {
		if item == selected {
			newList.selected = idx
		}
	}

	// update list with new Items and selected item index
	widget.list = &newList

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
