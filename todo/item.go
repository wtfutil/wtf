package todo

import (
)

type Item struct {
	Checked bool
	Index   int
	Text    string
}

func (item *Item) CheckMark() string {
	if item.Checked {
		return "x"
	} else {
		return " "
	}
}

func (item *Item) Toggle() {
	item.Checked = !item.Checked
}
