package zendesk

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("j", widget.selectNext)
	widget.SetKeyboardChar("k", widget.selectPrev)
	widget.SetKeyboardChar("o", widget.openTicket)

	widget.SetKeyboardKey(tcell.KeyDown, widget.selectNext)
	widget.SetKeyboardKey(tcell.KeyUp, widget.selectPrev)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openTicket)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
}

func (widget *Widget) selectNext() {
	widget.next()
	widget.display()
}

func (widget *Widget) selectPrev() {
	widget.prev()
	widget.display()
}
