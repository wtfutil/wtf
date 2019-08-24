package jenkins

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next item")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("o", widget.openJob, "Open job in browser")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openJob, "Open job in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
