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
	hidden := 0
	if widget.settings.checkedPos == "last" {
		str, hidden = widget.sortListByChecked(widget.list.UncheckedItems(), widget.list.CheckedItems())
	} else if widget.settings.checkedPos == "first" {
		str, hidden = widget.sortListByChecked(widget.list.CheckedItems(), widget.list.UncheckedItems())
	} else {
		str, hidden = widget.sortListByChecked(widget.list.Items, []*checklist.ChecklistItem{})
	}

	if widget.Error != "" {
		str = widget.Error
	}

	title := widget.CommonSettings().Title
	if widget.showTagPrefix != "" {
		title += " #" + widget.showTagPrefix
	}
	if widget.showFilter != "" {
		title += fmt.Sprintf(" /%s", widget.showFilter)
	}
	if widget.settings.hiddenNumInTitle {
		title += fmt.Sprintf(" (%d hidden)", hidden)
	}

	return title, str, false
}

func (widget *Widget) sortListByChecked(firstGroup []*checklist.ChecklistItem, secondGroup []*checklist.ChecklistItem) (string, int) {
	str := ""
	hidden := 0
	newList := checklist.NewChecklist(
		widget.settings.Sigils.Checkbox.Checked,
		widget.settings.Sigils.Checkbox.Unchecked,
	)

	offset := 0
	selectedItem := widget.SelectedItem()
	for idx, item := range firstGroup {
		if widget.shouldShowItem(item) {
			str += widget.formattedItemLine(idx, hidden, item)
		} else {
			hidden = hidden + 1
		}
		newList.Items = append(newList.Items, item)
		offset++
	}

	for idx, item := range secondGroup {
		if widget.shouldShowItem(item) {
			str += widget.formattedItemLine(idx+offset, hidden, item)
		} else {
			hidden = hidden + 1
		}
		newList.Items = append(newList.Items, item)
	}
	if idx, ok := newList.IndexByItem(selectedItem); ok {
		widget.Selected = idx
	}

	widget.SetList(newList)
	return str, hidden
}

func (widget *Widget) shouldShowItem(item *checklist.ChecklistItem) bool {
	if widget.showFilter != "" && !strings.Contains(strings.ToLower(item.Text), widget.showFilter) {
		return false
	}

	if !widget.settings.parseTags {
		return true
	}

	if len(item.Tags) == 0 {
		return widget.showTagPrefix == ""
	}

	for _, tag := range item.Tags {
		for _, hideTag := range widget.settings.hideTags {
			if widget.showTagPrefix == "" && tag == hideTag {
				return false
			}
		}
		if widget.showTagPrefix == "" || strings.HasPrefix(tag, widget.showTagPrefix) {
			return true
		}
	}

	return false
}

func (widget *Widget) RowColor(idx int, hidden int, checked bool) string {
	if widget.View.HasFocus() && (idx == widget.Selected) {
		foreground := widget.CommonSettings().Colors.RowTheme.HighlightedForeground
		if checked {
			foreground = widget.settings.Colors.CheckboxTheme.Checked
		}
		return fmt.Sprintf(
			"%s:%s",
			foreground,
			widget.CommonSettings().Colors.RowTheme.HighlightedBackground,
		)
	}

	if checked {
		return widget.settings.Colors.CheckboxTheme.Checked
	} else {
		return widget.CommonSettings().RowColor(idx - hidden)
	}
}

func (widget *Widget) formattedItemLine(idx int, hidden int, currItem *checklist.ChecklistItem) string {
	rowColor := widget.RowColor(idx, hidden, currItem.Checked)

	todoDate := currItem.Date
	row := fmt.Sprintf(
		` [%s]|%s| `,
		rowColor,
		currItem.CheckMark(),
	)

	if widget.settings.parseDates && todoDate != nil {
		row += fmt.Sprintf(
			`[%s]%s `,
			widget.settings.dateColor,
			widget.getDateString(todoDate),
		)
	}

	tagsPart := ""
	if len(currItem.Tags) > 0 {
		tagsPart = fmt.Sprintf(
			`[%s]%s[white]`,
			widget.settings.tagColor,
			currItem.TagString(),
		)
	}

	textPart := fmt.Sprintf(
		`[%s]%s[white]`,
		rowColor,
		tview.Escape(currItem.Text),
	)

	if widget.settings.parseTags && widget.settings.tagsAtEnd {
		row += textPart + " " + tagsPart
	} else if widget.settings.parseTags {
		row += tagsPart + textPart
	} else {
		row += textPart
	}

	return utils.HighlightableHelper(widget.View, row, idx-hidden, len(currItem.Text))
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
