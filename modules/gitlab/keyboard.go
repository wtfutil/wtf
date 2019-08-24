package gitlab

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls()

	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous project")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next project")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next project")
}
