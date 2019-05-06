package git

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("h", widget.Prev)
	widget.SetKeyboardChar("l", widget.Next)
	widget.SetKeyboardChar("p", widget.Pull)
	widget.SetKeyboardChar("c", widget.Checkout)

	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev)
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next)
}
