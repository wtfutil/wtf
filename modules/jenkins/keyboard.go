package jenkins

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("o", widget.openJob, "Open job in browser")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openJob, "Open job in browser")
}
