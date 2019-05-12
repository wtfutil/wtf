package git

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help window")
	widget.SetKeyboardChar("h", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("l", widget.Next, "Select next item")
	widget.SetKeyboardChar("p", widget.Pull, "Pull repo")
	widget.SetKeyboardChar("c", widget.Checkout, "Checkout branch")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next, "Select next item")
}
