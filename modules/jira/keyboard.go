package jira

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.Next)
	widget.SetKeyboardChar("k", widget.Prev)
	widget.SetKeyboardChar("o", widget.openItem)

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openItem)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev)
}
