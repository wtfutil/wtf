package digitalocean

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(nil)

	widget.SetKeyboardChar("j", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("k", widget.Next, "Select next item")
	widget.SetKeyboardChar("u", widget.Unselect, "Clear selection")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
}
