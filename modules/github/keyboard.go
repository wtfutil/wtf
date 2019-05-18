package github

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous item")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next item")
	widget.SetKeyboardChar("o", widget.openRepo, "Open item in browser")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openRepo, "Open item in browser")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next item")
}
