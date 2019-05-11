package hackernews

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.Next)
	widget.SetKeyboardChar("k", widget.Prev)
	widget.SetKeyboardChar("o", widget.openStory)
	widget.SetKeyboardChar("r", widget.Refresh)
	widget.SetKeyboardChar("c", widget.openComments)

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openStory)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev)
}
