package todo

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next item")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous item")
	widget.SetKeyboardChar(" ", widget.toggleChecked, "Toggle checkmark")
	widget.SetKeyboardChar("n", widget.newItem, "Create new item")
	widget.SetKeyboardChar("o", widget.openFile, "Open file")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelected, "Delete item")
	widget.SetKeyboardKey(tcell.KeyCtrlJ, widget.demoteSelected, "Demote item")
	widget.SetKeyboardKey(tcell.KeyCtrlK, widget.promoteSelected, "Promote item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.updateSelected, "Edit item")

}

func (widget *Widget) deleteSelected() {

	if !widget.isItemSelected() {
		return
	}

	widget.list.Delete(widget.Selected)
	widget.ScrollableWidget.SetItemCount(len(widget.list.Items))
	widget.Prev()
	widget.persist()
	widget.display()
}

func (widget *Widget) demoteSelected() {
	if !widget.isItemSelected() {
		return
	}

	j := widget.Selected + 1
	if j >= len(widget.list.Items) {
		j = 0
	}

	widget.list.Swap(widget.Selected, j)
	widget.Selected = j

	widget.persist()
	widget.display()
}

func (widget *Widget) openFile() {
	confDir, _ := cfg.WtfConfigDir()
	utils.OpenFile(fmt.Sprintf("%s/%s", confDir, widget.filePath))
}

func (widget *Widget) promoteSelected() {
	if !widget.isItemSelected() {
		return
	}

	k := widget.Selected - 1
	if k < 0 {
		k = len(widget.list.Items) - 1
	}

	widget.list.Swap(widget.Selected, k)
	widget.Selected = k
	widget.persist()
	widget.display()
}

func (widget *Widget) toggleChecked() {
	selectedItem := widget.SelectedItem()
	if selectedItem == nil {
		return
	}

	selectedItem.Toggle()
	widget.persist()
	widget.display()
}

func (widget *Widget) unselect() {
	widget.Selected = -1
	widget.display()
}
