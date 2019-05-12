package git

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help window")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next source")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous source")
	widget.SetKeyboardChar("p", widget.Pull, "Pull repo")
	widget.SetKeyboardChar("c", widget.Checkout, "Checkout branch")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous source")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next source")
}
