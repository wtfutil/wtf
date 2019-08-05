package view

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

type Base struct {
	app             *tview.Application
	bordered        bool
	commonSettings  *cfg.Common
	enabled         bool
	focusChar       string
	focusable       bool
	key             string
	name            string
	quitChan        chan bool
	refreshing      bool
	refreshInterval int
}
