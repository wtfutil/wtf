package wtf

import (
	"github.com/wtfutil/wtf/cfg"

	"github.com/rivo/tview"
)

type Wtfable interface {
	Enabler
	Scheduler

	BorderColor() string
	ConfigText() string
	FocusChar() string
	Focusable() bool
	HelpText() string
	Name() string
	SetFocusChar(string)
	TextView() *tview.TextView

	CommonSettings() *cfg.Common
}
