package nbascore

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("h", widget.prev, "Select previous item")
	widget.SetKeyboardChar("l", widget.next, "Select next item")
	widget.SetKeyboardChar("c", widget.center, "???")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.next, "Select next item")
}

func (widget *Widget) center() {
	offset = 0
	widget.Refresh()
}

func (widget *Widget) next() {
	offset++
	widget.Refresh()
}

func (widget *Widget) prev() {
	offset--
	widget.Refresh()
}
