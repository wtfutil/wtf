package weather

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous city")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next city")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous city")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next city")
}
