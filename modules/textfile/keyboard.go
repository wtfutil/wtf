package textfile

import (
	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("h", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("l", widget.Next, "Select next item")
	widget.SetKeyboardChar("o", widget.openFile, "Open item")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next, "Select next item")
}

func (widget *Widget) openFile() {
	src := widget.CurrentSource()
	wtf.OpenFile(src)
}
