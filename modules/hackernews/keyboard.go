package hackernews

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("o", widget.openStory, "Open story in browser")
	widget.SetKeyboardChar("c", widget.openComments, "Open comments in browser")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openStory, "Open story in browser")
}
