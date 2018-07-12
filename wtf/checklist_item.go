package wtf

type ChecklistItem struct {
	Checked bool
	Text    string
}

func (item *ChecklistItem) CheckMark() string {
	if item.Checked {
		return Config.UString("wtf.mods.todo.checkedIcon", "x")
	} else {
		return " "
	}
}

func (item *ChecklistItem) Toggle() {
	item.Checked = !item.Checked
}
