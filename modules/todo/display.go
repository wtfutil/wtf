package todo

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/checklist"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	str := ""
	newList := checklist.NewChecklist(
		widget.settings.common.Sigils.Checkbox.Checked,
		widget.settings.common.Sigils.Checkbox.Unchecked,
	)

	offset := 0

	for idx, item := range widget.list.UncheckedItems() {
		str += widget.formattedItemLine(idx, item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
		offset++
	}

	for idx, item := range widget.list.CheckedItems() {
		str += widget.formattedItemLine(idx+offset, item, widget.list.SelectedItem(), widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
	}

	newList.SetSelectedByItem(widget.list.SelectedItem())
	widget.SetList(newList)

	return widget.CommonSettings().Title, str, false
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

	row := fmt.Sprintf(
		` [%s:%s]|%s| %s[white]`,
		foreColor,
		backColor,
		item.CheckMark(),
		tview.Escape(item.Text),
	)

	return utils.HighlightableHelper(widget.View, row, idx, len(item.Text))
}
