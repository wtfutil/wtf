package gerrit

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("h", widget.prevProject)
	widget.SetKeyboardChar("l", widget.nextProject)
	widget.SetKeyboardChar("j", widget.nextReview)
	widget.SetKeyboardChar("k", widget.prevReview)
	widget.SetKeyboardChar("r", widget.Refresh)

	widget.SetKeyboardKey(tcell.KeyLeft, widget.prevProject)
	widget.SetKeyboardKey(tcell.KeyRight, widget.nextProject)
	widget.SetKeyboardKey(tcell.KeyDown, widget.nextReview)
	widget.SetKeyboardKey(tcell.KeyUp, widget.prevReview)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openReview)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
}
