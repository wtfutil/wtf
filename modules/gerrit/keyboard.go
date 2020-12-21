package gerrit

import (
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("h", widget.prevProject, "Select previous project")
	widget.SetKeyboardChar("l", widget.nextProject, "Select next project")
	widget.SetKeyboardChar("j", widget.nextReview, "Select next review")
	widget.SetKeyboardChar("k", widget.prevReview, "Select previous review")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.prevProject, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyRight, widget.nextProject, "Select next project")
	widget.SetKeyboardKey(tcell.KeyDown, widget.nextReview, "Select next review")
	widget.SetKeyboardKey(tcell.KeyUp, widget.prevReview, "Select previous review")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openReview, "Open review in browser")
}
