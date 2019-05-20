package mercurial

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous item")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next item")
	widget.SetKeyboardChar("p", widget.Pull, "Pull repo")
	widget.SetKeyboardChar("c", widget.Checkout, "Checkout branch")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next item")
}
