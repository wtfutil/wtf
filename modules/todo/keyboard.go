package todo

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("j", widget.displayNext, "Select next item")
	widget.SetKeyboardChar("k", widget.displayPrev, "Select previous item")
	widget.SetKeyboardChar(" ", widget.toggleChecked, "Toggle checkmark")
	widget.SetKeyboardChar("n", widget.newItem, "Create new item")
	widget.SetKeyboardChar("o", widget.openFile, "Open file")

	widget.SetKeyboardKey(tcell.KeyDown, widget.displayNext, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.displayPrev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelected, "Delete item")
	widget.SetKeyboardKey(tcell.KeyCtrlJ, widget.demoteSelected, "Demote item")
	widget.SetKeyboardKey(tcell.KeyCtrlK, widget.promoteSelected, "Promote item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.editSelected, "Edit item")

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
	confDir, _ := cfg.WtfConfigDir()
	utils.OpenFile(fmt.Sprintf("%s/%s", confDir, widget.filePath))
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
