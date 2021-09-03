package todo

import (
	"fmt"
	"strings"
	"time"

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
	now := getNowDate()
	diff := int(date.Sub(now).Hours() / 24)
	if diff == 0 {
		return "today"
	} else if diff == 1 {
		return "tomorrow"
	} else if diff <= widget.settings.switchToInDaysIn {
		return fmt.Sprintf("in %d days", diff)
	} else {
		dateStr := ""
		y, m, d := date.Year(), date.Month(), date.Day()
		switch widget.settings.dateFormat {
		case "yyyy-mm-dd":
			dateStr = fmt.Sprintf("%d-%02d-%02d", y, m, d)
		case "yy-mm-dd":
			dateStr = fmt.Sprintf("%d-%02d-%02d", y-2000, m, d)
		case "dd-mm-yyyy":
			dateStr = fmt.Sprintf("%02d-%02d-%d", d, m, y)
		case "dd-mm-yy":
			dateStr = fmt.Sprintf("%02d-%02d-%d", d, m, y-2000)
		case "dd M yyyy":
			dateStr = fmt.Sprintf("%02d %s %d", d, date.Month().String()[:3], y)
			// date
		case "dd M yy":
			dateStr = fmt.Sprintf("%02d %s %d", d, date.Month().String()[:3], y-2000)
			// dateStr = "aaasdada"
		default:
			dateStr = fmt.Sprintf("%d-%02d-%02d", y, m, d)
			// dateStr = fmt.Sprintf("%d-%02d-%02d", y, m, d)
		}
		if widget.settings.hideYearIfCurrent && date.Year() == now.Year() {
			if widget.settings.dateFormat[:1] == "y" {
				dateStr = dateStr[strings.Index(dateStr, "-")+1:]
			} else if widget.settings.dateFormat[3:4] == "-" {
				dateStr = dateStr[:5]
			} else {
				parts := strings.Split(dateStr, " ")
				dateStr = parts[0] + " " + parts[1]
			}
		}
		return dateStr
	}
}

func getNowDate() time.Time {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Now().Location())
	return now
}
