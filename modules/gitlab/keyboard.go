package gitlab

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")
	widget.SetKeyboardChar("h", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("l", widget.Next, "Select next item")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next, "Select next item")
}
