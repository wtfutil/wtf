package newrelic

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous application")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next application")
}
