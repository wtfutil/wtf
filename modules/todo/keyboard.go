package todo

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar(" ", widget.toggleChecked)
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.displayNext)
	widget.SetKeyboardChar("k", widget.displayPrev)
	widget.SetKeyboardChar("n", widget.newItem)
	widget.SetKeyboardChar("o", widget.openFile)

	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelected)
	widget.SetKeyboardKey(tcell.KeyCtrlJ, widget.demoteSelected)
	widget.SetKeyboardKey(tcell.KeyCtrlK, widget.promoteSelected)
	widget.SetKeyboardKey(tcell.KeyDown, widget.displayNext)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.editSelected)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.displayPrev)
}

func (widget *Widget) deleteSelected() {
	widget.list.Delete()
	widget.persist()
	widget.display()
}

func (widget *Widget) demoteSelected() {
	widget.list.Demote()
	widget.persist()
	widget.display()
}

func (widget *Widget) displayNext() {
	widget.list.Next()
	widget.display()
}

func (widget *Widget) displayPrev() {
	widget.list.Prev()
	widget.display()
}

func (widget *Widget) openFile() {
	confDir, _ := cfg.ConfigDir()
	wtf.OpenFile(fmt.Sprintf("%s/%s", confDir, widget.filePath))
}

func (widget *Widget) promoteSelected() {
	widget.list.Promote()
	widget.persist()
	widget.display()
}

func (widget *Widget) toggleChecked() {
	widget.list.Toggle()
	widget.persist()
	widget.display()
}

func (widget *Widget) unselect() {
	widget.list.Unselect()
	widget.display()
}
