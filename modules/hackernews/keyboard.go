
package hackernews

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.next)
	widget.SetKeyboardChar("k", widget.prev)
	widget.SetKeyboardChar("o", widget.openStory)
	widget.SetKeyboardChar("r", widget.Refresh)
	widget.SetKeyboardChar("c", widget.openComments)

	widget.SetKeyboardKey(tcell.KeyDown, widget.next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openStory)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.prev)
}