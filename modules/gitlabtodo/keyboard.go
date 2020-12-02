package gitlabtodo

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next item")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("o", widget.openTodo, "Open todo in browser")
	widget.SetKeyboardChar("x", widget.markAsDone, "Mark todo as done")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openTodo, "Open todo in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
