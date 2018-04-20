package todo

import(
	"time"
)

type Item struct {
	Checked bool
	Index   int
	Text    string

	createdAt time.Time
	updatedAt time.Time
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
	item.updatedAt = time.Now()
}


