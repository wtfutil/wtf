package todo

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	str := ""
	newList := List{selected: -1}

	selectedItem := widget.list.Selected()
	maxLineLen := widget.list.LongestLine()

	for _, item := range widget.list.UncheckedItems() {
		str = str + widget.formattedItemLine(item, selectedItem, maxLineLen)
		newList.Items = append(newList.Items, item)
	}

	for _, item := range widget.list.CheckedItems() {
		str = str + widget.formattedItemLine(item, selectedItem, maxLineLen)
		newList.Items = append(newList.Items, item)
	}

	newList.SetSelectedByItem(widget.list.Selected())
	widget.SetList(&newList)

	widget.View.Clear()
	widget.View.SetText(str)
}

func (widget *Widget) formattedItemLine(item *Item, selectedItem *Item, maxLen int) string {
	foreColor, backColor := "white", wtf.Config.UString("wtf.colors.background", "black")

	if item.Checked {
		foreColor = wtf.Config.UString("wtf.mods.todo.colors.checked", "white")
	}

	if widget.View.HasFocus() && (item == selectedItem) {
		foreColor = wtf.Config.UString("wtf.mods.todo.colors.highlight.fore", "black")
		backColor = wtf.Config.UString("wtf.mods.todo.colors.highlight.back", "white")
	}

	str := fmt.Sprintf(
		"[%s:%s]|%s| %s[white]",
		foreColor,
		backColor,
		item.CheckMark(),
		item.Text,
	)

	str = str + wtf.PadRow((checkWidth+len(item.Text)), (checkWidth+maxLen)) + "\n"

	return str
}
