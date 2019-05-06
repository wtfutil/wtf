package github

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("h", widget.Prev)
	widget.SetKeyboardChar("l", widget.Next)
	widget.SetKeyboardChar("o", widget.openRepo)
	widget.SetKeyboardChar("r", widget.Refresh)

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openRepo)
	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev)
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next)
}
