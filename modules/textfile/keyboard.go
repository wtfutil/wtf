package textfile

import (
	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous item")
	widget.SetKeyboardChar("l", widget.NextSource, "Select next item")
	widget.SetKeyboardChar("o", widget.openFile, "Open item")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next item")
}

func (widget *Widget) openFile() {
	src := widget.CurrentSource()
	wtf.OpenFile(src)
}
