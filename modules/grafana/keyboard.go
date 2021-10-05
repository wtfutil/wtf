package grafana

import "github.com/gdamore/tcell/v2"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous alert")
	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next alert")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openAlert, "Open alert in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
