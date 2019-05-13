package rollbar

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("o", widget.openBuild, "Open item in browser")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openBuild, "Open item in browser")
}
