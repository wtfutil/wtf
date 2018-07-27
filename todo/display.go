package todo

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/checklist"
	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	str := ""
	newList := checklist.NewChecklist()

  offset := 0

	for idx, item := range widget.list.UncheckedItems() {
		str = str + widget.formattedItemLine(idx, item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
		offset++
	}

	for idx, item := range widget.list.CheckedItems() {
		str = str + widget.formattedItemLine(idx + offset, item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
	}

	newList.SetSelectedByItem(widget.list.SelectedItem())
	widget.SetList(newList)

	widget.View.Clear()
	widget.View.SetText(str)
	widget.View.Highlight(strconv.Itoa(widget.list.Selected)).ScrollToHighlight()
}

func (widget *Widget) formattedItemLine(idx int, item *checklist.ChecklistItem, selectedItem *checklist.ChecklistItem, maxLen int) string {
	foreColor, backColor := "white", wtf.Config.UString("wtf.colors.background", "black")

	if item.Checked {
		foreColor = wtf.Config.UString("wtf.colors.checked", "white")
	}

	if widget.View.HasFocus() && (item == selectedItem) {
		foreColor = wtf.Config.UString("wtf.colors.highlight.fore", "black")
		backColor = wtf.Config.UString("wtf.colors.highlight.back", "orange")
	}

	str := fmt.Sprintf(
		`["%d"][""][%s:%s]|%s| %s[white]`,
		idx,
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
