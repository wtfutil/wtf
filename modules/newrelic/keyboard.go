package newrelic

import "github.com/gdamore/tcell/v2"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous application")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next application")
}
