package cdsstatus

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next line")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous line")
	widget.SetKeyboardChar("o", widget.openWorkflow, "Open status in browser")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next line")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous line")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openWorkflow, "Open status in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
