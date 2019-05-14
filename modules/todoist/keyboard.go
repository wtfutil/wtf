package todoist

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("c", widget.Close, "Close item")
	widget.SetKeyboardChar("d", widget.Delete, "Delete item")
	widget.SetKeyboardChar("h", widget.Prev, "Select previous project")
	widget.SetKeyboardChar("j", widget.Up, "Select previous item")
	widget.SetKeyboardChar("k", widget.Down, "Select next item")
	widget.SetKeyboardChar("l", widget.Next, "Select next project")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Down, "Select next item")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next, "Select next project")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Up, "Select previous item")
}
