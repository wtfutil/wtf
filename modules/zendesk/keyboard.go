package zendesk

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("o", widget.openTicket, "Open item")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openTicket, "Open item")
}
