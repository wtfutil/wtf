package wtf

import (
	"github.com/rivo/tview"
)

type TextViewer interface {
	Disabled() bool
	Enabled() bool
	TextView() *tview.TextView
}
