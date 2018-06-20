package todo

import (
	"github.com/senorprogrammer/wtf/wtf"
)

type Item struct {
	Checked bool
	Text    string
}

func (item *Item) CheckMark() string {
	if item.Checked {
		return wtf.Config.UString("wtf.mods.todo.checkedIcon", "x")
	} else {
		return " "
	}
}

func (item *Item) Toggle() {
	item.Checked = !item.Checked
}
