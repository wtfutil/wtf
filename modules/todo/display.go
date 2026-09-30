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

	switch widget.settings.checkedPos {
	case "last":
		str, hidden = widget.sortListByChecked(widget.list.UncheckedItems(), widget.list.CheckedItems())
	case "first":
		str, hidden = widget.sortListByChecked(widget.list.CheckedItems(), widget.list.UncheckedItems())
	default:
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
		widget.settings.Checkbox.Checked,
		widget.settings.Checkbox.Unchecked,
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

// rowColorResult computes the color string for a row based on state.
// This is the pure logic extracted for testability.
func rowColorResult(hasFocus bool, idx int, selected int, hidden int, checked bool, highlightedFg, highlightedBg, checkedColor string, rowColorFn func(int) string) string {
	if hasFocus && (idx == selected) {
		foreground := highlightedFg
		if checked {
			foreground = checkedColor
		}
		return fmt.Sprintf("%s:%s", foreground, highlightedBg)
	}

	if checked {
		return checkedColor
	}
	return rowColorFn(idx - hidden)
}

func (widget *Widget) RowColor(idx int, hidden int, checked bool) string {
	return rowColorResult(
		widget.View.HasFocus(),
		idx,
		widget.Selected,
		hidden,
		checked,
		widget.CommonSettings().Colors.HighlightedForeground,
		widget.CommonSettings().Colors.HighlightedBackground,
		widget.settings.Colors.Checked,
		widget.CommonSettings().RowColor,
	)
}

// formatItemRow builds the display string for a checklist item without any
// view-dependent operations (no HighlightableHelper wrapping).
// This is the pure logic extracted for testability.
func formatItemRow(rowColor string, checkMark string, parseDates bool, todoDate *time.Time, dateColor string, dateStr string, parseTags bool, tagsAtEnd bool, tagColor string, tagString string, text string) string {
	row := fmt.Sprintf(
		` [%s]|%s| `,
		rowColor,
		checkMark,
	)

	if parseDates && todoDate != nil {
		row += fmt.Sprintf(
			`[%s]%s `,
			dateColor,
			dateStr,
		)
	}

	tagsPart := ""
	if len(tagString) > 0 {
		tagsPart = fmt.Sprintf(
			`[%s]%s[white]`,
			tagColor,
			tagString,
		)
	}

	textPart := fmt.Sprintf(
		`[%s]%s[white]`,
		rowColor,
		tview.Escape(text),
	)

	if parseTags && tagsAtEnd {
		row += textPart + " " + tagsPart
	} else if parseTags {
		row += tagsPart + textPart
	} else {
		row += textPart
	}

	return row
}

func (widget *Widget) formattedItemLine(idx int, hidden int, currItem *checklist.ChecklistItem) string {
	rowColor := widget.RowColor(idx, hidden, currItem.Checked)

	dateStr := ""
	if widget.settings.parseDates && currItem.Date != nil {
		dateStr = widget.getDateString(currItem.Date)
	}

	row := formatItemRow(
		rowColor,
		currItem.CheckMark(),
		widget.settings.parseDates,
		currItem.Date,
		widget.settings.dateColor,
		dateStr,
		widget.settings.parseTags,
		widget.settings.tagsAtEnd,
		widget.settings.tagColor,
		currItem.TagString(),
		currItem.Text,
	)

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
			} else if widget.settings.dateFormat[2:3] == "-" {
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
