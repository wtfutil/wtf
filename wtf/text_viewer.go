package wtf

import (
	"github.com/rivo/tview"
)

type TextViewer interface {
	Enabler

	Refresh()
	TextView() *tview.TextView

	Top() int
	Left() int
	Width() int
	Height() int
}
