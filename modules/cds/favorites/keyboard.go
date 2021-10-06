package cdsfavorites

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next workflow")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous workflow")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next source")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous source")
	widget.SetKeyboardChar("o", widget.openWorkflow, "Open workflow in browser")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next workflow")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous workflow")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next source")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous source")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openWorkflow, "Open workflow in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
