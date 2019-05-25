package weather

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh Widget")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous city")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next city")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous city")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next city")
}
