package todoist

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("c", widget.Close)
	widget.SetKeyboardChar("d", widget.Delete)
	widget.SetKeyboardChar("h", widget.PreviousProject)
	widget.SetKeyboardChar("j", widget.Up)
	widget.SetKeyboardChar("k", widget.Down)
	widget.SetKeyboardChar("l", widget.NextProject)
	widget.SetKeyboardChar("r", widget.Refresh)

	widget.SetKeyboardKey(tcell.KeyDown, widget.Down)
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PreviousProject)
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextProject)
	widget.SetKeyboardKey(tcell.KeyUp, widget.Up)
}
