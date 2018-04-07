package wtf

import (
	"github.com/rivo/tview"
)

type TextViewer interface {
	Enabler
	TextView() *tview.TextView
}
