package jira

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.selectNext)
	widget.SetKeyboardChar("k", widget.selectPrev)
	widget.SetKeyboardChar("o", widget.openItem)

	widget.SetKeyboardKey(tcell.KeyDown, widget.selectNext)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openItem)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.selectPrev)
}

func (widget *Widget) selectNext() {
	widget.next()
	widget.display()
}

func (widget *Widget) selectPrev() {
	widget.prev()
	widget.display()
}
