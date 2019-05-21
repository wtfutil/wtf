package todoist

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")
	widget.SetKeyboardChar("j", widget.Down, "Select next item")
	widget.SetKeyboardChar("k", widget.Up, "Select previous item")
	widget.SetKeyboardChar("c", widget.Close, "Close item")
	widget.SetKeyboardChar("d", widget.Delete, "Delete item")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous project")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next project")

	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next project")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyDown, widget.Down, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Up, "Select previous item")
}
