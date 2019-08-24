package github

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("l", widget.NextSource, "Select next source")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous source")
	widget.SetKeyboardChar("o", widget.openRepo, "Open item in browser")

	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next source")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous source")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openRepo, "Open item in browser")
}
