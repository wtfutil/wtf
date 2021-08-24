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
	if widget.settings.checkedPos == "last" {
		str += widget.sortListByChecked(widget.list.UncheckedItems(), widget.list.CheckedItems())
	} else if widget.settings.checkedPos == "first" {
		str += widget.sortListByChecked(widget.list.CheckedItems(), widget.list.UncheckedItems())
	} else {
		str += widget.sortListByChecked(widget.list.Items, []*checklist.ChecklistItem{})
	}
	return widget.CommonSettings().Title, str, false
}

func (widget *Widget) sortListByChecked(firstGroup []*checklist.ChecklistItem, secondGroup []*checklist.ChecklistItem) string {
	str := ""
	newList := checklist.NewChecklist(
		widget.settings.Sigils.Checkbox.Checked,
		widget.settings.Sigils.Checkbox.Unchecked,
	)

	offset := 0
	selectedItem := widget.SelectedItem()
	for idx, item := range firstGroup {
		str += widget.formattedItemLine(idx, item, selectedItem, widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
		offset++
	}

	for idx, item := range secondGroup {
		str += widget.formattedItemLine(idx+offset, item, selectedItem, widget.list.LongestLine())
		newList.Items = append(newList.Items, item)
	}
	if idx, ok := newList.IndexByItem(selectedItem); ok {
		widget.Selected = idx
	}

	widget.SetList(newList)
	return str
}

func (widget *Widget) formattedItemLine(idx int, currItem *checklist.ChecklistItem, selectedItem *checklist.ChecklistItem, maxLen int) string {
	rowColor := widget.RowColor(idx)

	if currItem.Checked {
		rowColor = widget.settings.Colors.CheckboxTheme.Checked
	}

	if widget.View.HasFocus() && (currItem == selectedItem) {
		rowColor = widget.RowColor(idx)
	}

	row := fmt.Sprintf(
		` [%s]|%s| %s[white]`,
		rowColor,
		currItem.CheckMark(),
		tview.Escape(currItem.Text),
	)

	return utils.HighlightableHelper(widget.View, row, idx, len(currItem.Text))
}
