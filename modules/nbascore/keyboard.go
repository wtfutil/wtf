package nbascore

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("h", widget.prev)
	widget.SetKeyboardChar("l", widget.next)
	widget.SetKeyboardChar("c", widget.center)

	widget.SetKeyboardKey(tcell.KeyLeft, widget.prev)
	widget.SetKeyboardKey(tcell.KeyRight, widget.next)
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
