package zendesk

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("j", widget.Next)
	widget.SetKeyboardChar("k", widget.Prev)
	widget.SetKeyboardChar("o", widget.openTicket)

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next)
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openTicket)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect)
}
