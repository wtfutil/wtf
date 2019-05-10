package rollbar

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.Next)
	widget.SetKeyboardChar("k", widget.Prev)
	widget.SetKeyboardChar("o", widget.openBuild)
	widget.SetKeyboardChar("r", widget.Refresh)
	widget.SetKeyboardChar("u", widget.Unselect)

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openBuild)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev)
}
