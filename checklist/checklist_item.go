package checklist

import (
	"fmt"
	"time"
)

// ChecklistItem is a module for creating generic checklist implementations
// See 'Todo' for an implementation example
type ChecklistItem struct {
	Checked       bool
	CheckedIcon   string `yaml:"-"`
	Date          *time.Time
	Tags          []string
	Text          string
	UncheckedIcon string `yaml:"-"`
}

func NewChecklistItem(checked bool, date *time.Time, tags []string, text string, checkedIcon, uncheckedIcon string) *ChecklistItem {
	item := &ChecklistItem{
		Checked:       checked,
		CheckedIcon:   checkedIcon,
		Date:          date,
		Tags:          tags,
		Text:          text,
		UncheckedIcon: uncheckedIcon,
	}

	return item
}

// CheckMark returns the string used to indicate a ChecklistItem is checked or unchecked
func (item *ChecklistItem) CheckMark() string {
	item.ensureItemIcons()

	if item.Checked {
		return item.CheckedIcon
	}

	return item.UncheckedIcon
}

// EditText returns the content of the edit todo form, so includes formatted date and tags
func (item *ChecklistItem) EditText() string {
	datePrefix := ""
	if item.Date != nil {
		datePrefix = fmt.Sprintf("%d-%02d-%02d", item.Date.Year(), item.Date.Month(), item.Date.Day()) + " "
	}

	tagsPrefix := item.TagString()

	return datePrefix + tagsPrefix + item.Text
}

func (item *ChecklistItem) TagString() string {
	if len(item.Tags) == 0 {
		return ""
	}

	s := ""
	for _, tag := range item.Tags {
		s += "#" + tag + " "
	}

	return s
}

// Toggle changes the checked state of the ChecklistItem
// If checked, it is unchecked. If unchecked, it is checked
func (item *ChecklistItem) Toggle() {
	item.Checked = !item.Checked
}

/* -------------------- Unexported Functions -------------------- */

func (item *ChecklistItem) ensureItemIcons() {
	if item.CheckedIcon == "" {
		item.CheckedIcon = "x"
	}

	if item.UncheckedIcon == "" {
		item.UncheckedIcon = " "
	}
}
