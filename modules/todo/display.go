package todo

import (
	"fmt"
	"time"
	"sort"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/checklist"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

type byDate struct {
	items []*checklist.ChecklistItem
	undatedAsDays int
}
func (s byDate) Len() int {
	return len(s.items)
}
func (s byDate) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}
func (s byDate) Less(i, j int) bool {
	defaultVal := time.Now().AddDate(0,0,s.undatedAsDays)
	d1 := s.items[i].Date
	if d1 == nil {
		d1 = &defaultVal
	}
	d2 := s.items[j].Date
	if d2 == nil {
		d2 = &defaultVal
	}
	if d1.Equal(*d2) {
		return i < j
	}
	return d1.Before(*d2)
}

func (widget *Widget) content() (string, string, bool) {
	str := ""
	if widget.settings.parseDates {
		sort.Sort(byDate{items: widget.list.Items, undatedAsDays: widget.settings.undatedAsDays})
	}
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

	todoDate := getTodoDate(currItem.Text)
	row := ""

	if todoDate == nil {
		row += fmt.Sprintf(
			` [%s]|%s| %s[white]`,
			rowColor,
			currItem.CheckMark(),
			tview.Escape(currItem.Text),
		)
	} else {
		row += fmt.Sprintf(
			` [%s]|%s| [%s]%s [%s]%s[white]`,
			rowColor,
			currItem.CheckMark(),
			widget.settings.dateColor,
			widget.getDateString(todoDate),
			rowColor,
			tview.Escape(currItem.Text[13:]),
		)
	}

	return utils.HighlightableHelper(widget.View, row, idx, len(currItem.Text))
}

func getTodoDate(text string, defaultVal ...time.Time) *time.Time {
	if len(text) < 12 {
		if len(defaultVal) > 0 {
			return &defaultVal[0]
		} else {
			return nil
		}
	}
	date, err := time.Parse("2006-01-02", text[1:11])
	if err != nil {
		if len(defaultVal) > 0 {
			return &defaultVal[0]
		} else {
			return nil
		}
	}
	return &date
}

func (widget *Widget) getDateString(date *time.Time) string {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(),0,0,0,0,time.Now().Location())
	diff := int(date.Sub(now).Hours() / 24)
	if diff == 0 {
		return "today"
	} else if diff == 1 {
		return "tomorrow"
	} else if diff <= widget.settings.switchToInDaysIn {
		return fmt.Sprintf("in %d days", diff)
	} else {
		return widget._textWithDate(*date,"")
	}
}

