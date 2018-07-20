package todo

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/checklist"
	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	str := ""
	newList := checklist.NewChecklist()

	for _, item := range widget.list.UncheckedItems() {
		str = str + widget.formattedItemLine(item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
	}

	for _, item := range widget.list.CheckedItems() {
		str = str + widget.formattedItemLine(item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
	}

	newList.SetSelectedByItem(widget.list.SelectedItem())
	widget.SetList(newList)

	widget.View.Clear()
	widget.View.SetText(str)
}

func (widget *Widget) formattedItemLine(item *checklist.ChecklistItem, selectedItem *checklist.ChecklistItem, maxLen int) string {
	foreColor, backColor := "white", wtf.Config.UString("wtf.colors.background", "black")

	if item.Checked {
		foreColor = wtf.Config.UString("wtf.colors.checked", "white")
	}

	if widget.View.HasFocus() && (item == selectedItem) {
		foreColor = wtf.Config.UString("wtf.colors.highlight.fore", "black")
		backColor = wtf.Config.UString("wtf.colors.highlight.back", "orange")
	}

	str := fmt.Sprintf(
		"[%s:%s]|%s| %s[white]",
		foreColor,
		backColor,
		item.CheckMark(),
		tview.Escape(item.Text),
	)

	_, _, w, _ := widget.View.GetInnerRect()
	if w > maxLen {
		maxLen = w
	}

	return str + wtf.PadRow((checkWidth+len(item.Text)), (checkWidth+maxLen+1)) + "\n"
}
