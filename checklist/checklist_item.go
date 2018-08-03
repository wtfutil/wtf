package checklist

import (
	"github.com/senorprogrammer/wtf/wtf"
)

// ChecklistItem is a module for creating generic checklist implementations
// See 'Todo' for an implementation example
type ChecklistItem struct {
	Checked bool
	Text    string
}

// CheckMark returns the string used to indicate a ChecklistItem is checked or unchecked
func (item *ChecklistItem) CheckMark() string {
	if item.Checked {
		return wtf.Config.UString("wtf.mods.todo.checkedIcon", "x")
	}

	return " "
}

// Toggle changes the checked state of the ChecklistItem
// If checked, it is unchecked. If unchecked, it is checked
func (item *ChecklistItem) Toggle() {
	item.Checked = !item.Checked
}
