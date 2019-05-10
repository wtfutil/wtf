package datadog

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.next)
	widget.SetKeyboardChar("k", widget.prev)
	widget.SetKeyboardChar("o", widget.openItem)

	widget.SetKeyboardKey(tcell.KeyDown, widget.next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openItem)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.prev)
}
