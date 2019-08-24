package twitter

import (
	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("l", widget.NextSource, "Select next source")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous source")
	widget.SetKeyboardChar("o", widget.openFile, "Open source")

	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next source")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous source")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openFile, "Open source")
}

func (widget *Widget) openFile() {
	src := widget.currentSourceURI()
	utils.OpenFile(src)
}
