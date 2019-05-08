package todo

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/checklist"
	"github.com/wtfutil/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	str := ""
	newList := checklist.NewChecklist(
		widget.settings.common.Sigils.Checkbox.Checked,
		widget.settings.common.Sigils.Checkbox.Unchecked,
	)

	offset := 0

	for idx, item := range widget.list.UncheckedItems() {
		str = str + widget.formattedItemLine(idx, item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
		offset++
	}

	for idx, item := range widget.list.CheckedItems() {
		str = str + widget.formattedItemLine(idx+offset, item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
	}

	newList.SetSelectedByItem(widget.list.SelectedItem())
	widget.SetList(newList)

	widget.Redraw(widget.CommonSettings.Title, str, false)
}

func (widget *Widget) formattedItemLine(idx int, item *checklist.ChecklistItem, selectedItem *checklist.ChecklistItem, maxLen int) string {
	foreColor, backColor := widget.settings.common.Colors.Text, widget.settings.common.Colors.Background

	if item.Checked {
		foreColor = widget.settings.common.Colors.Checked
	}

	if widget.View.HasFocus() && (item == selectedItem) {
		foreColor = widget.settings.common.Colors.HighlightFore
		backColor = widget.settings.common.Colors.HighlightBack
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
