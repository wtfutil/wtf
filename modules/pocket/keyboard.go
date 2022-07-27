package pocket

import "github.com/gdamore/tcell/v2"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("a", widget.toggleLink, "Toggle Link")
	widget.SetKeyboardChar("t", widget.toggleView, "Toggle view (links ,archived links)")
	widget.SetKeyboardChar("j", widget.Next, "Select Next Link")
	widget.SetKeyboardChar("k", widget.Prev, "Select Previous Link")
	widget.SetKeyboardChar("o", widget.openLink, "Open Link in the browser")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select Next Link")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select Previous Link")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openLink, "Open Link in the browser")
}
