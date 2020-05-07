package checklist

// ChecklistItem is a module for creating generic checklist implementations
// See 'Todo' for an implementation example
type ChecklistItem struct {
	Checked       bool
	CheckedIcon   string
	Text          string
	UncheckedIcon string
}

func NewChecklistItem(checked bool, text string, checkedIcon, uncheckedIcon string) *ChecklistItem {
	item := &ChecklistItem{
		Checked:       checked,
		CheckedIcon:   checkedIcon,
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
