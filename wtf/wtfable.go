package wtf

import (
	"github.com/rivo/tview"
)

type Wtfable interface {
	Enabler
	Scheduler

	BorderColor() string
	Focusable() bool
	FocusChar() string
	Key() string
	Name() string
	SetFocusChar(string)
	TextView() *tview.TextView

	Top() int
	Left() int
	Width() int
	Height() int
}
