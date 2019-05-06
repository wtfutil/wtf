package textfile

import (
	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("h", widget.Prev)
	widget.SetKeyboardChar("l", widget.Next)
	widget.SetKeyboardChar("o", widget.openFile)

	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev)
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next)
}

func (widget *Widget) openFile() {
	src := widget.CurrentSource()
	wtf.OpenFile(src)
}
