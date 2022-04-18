package asana

import "github.com/gdamore/tcell/v2"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next task")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous task")
	widget.SetKeyboardChar("q", widget.Unselect, "Unselect task")
	widget.SetKeyboardChar("o", widget.openTask, "Open task in browser")
	widget.SetKeyboardChar("x", widget.toggleTaskCompletion, "Toggles the task's completion state")
	widget.SetKeyboardChar("?", widget.ShowHelp, "Shows help")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next task")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous task")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Unselect task")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openTask, "Open task in browser")
}
