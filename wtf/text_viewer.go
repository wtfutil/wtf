package wtf

import (
	"github.com/rivo/tview"
)

// TODO: I really need to come up with a better name for this now
type TextViewer interface {
	Enabler
	Scheduler

	BorderColor() string
	Focusable() bool
	TextView() *tview.TextView

	Top() int
	Left() int
	Width() int
	Height() int
}
