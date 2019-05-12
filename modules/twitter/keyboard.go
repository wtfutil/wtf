package twitter

import (
	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("r", widget.ShowHelp, "Refresh widget")
	widget.SetKeyboardChar("l", widget.Next, "Select next item")
	widget.SetKeyboardChar("h", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("o", widget.openFile, "Open item")

	widget.SetKeyboardKey(tcell.KeyRight, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openFile, "Open item")
}

func (widget *Widget) openFile() {
	src := widget.currentSourceURI()
	wtf.OpenFile(src)
}
