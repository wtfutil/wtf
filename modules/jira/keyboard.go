package jira

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("o", widget.openItem, "Open item in browser")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openItem, "Open item in browser")
}
