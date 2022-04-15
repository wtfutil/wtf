package todo

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.NextTodo, "Select next item")
	widget.SetKeyboardChar("k", widget.PrevTodo, "Select previous item")
	widget.SetKeyboardChar(" ", widget.toggleChecked, "Toggle checkmark")
	widget.SetKeyboardChar("n", widget.newItem, "Create new item")
	widget.SetKeyboardChar("o", widget.openFile, "Open file")
	widget.SetKeyboardChar("#", widget.setTag, "Set tag(s) to show")
	widget.SetKeyboardChar("f", widget.setFilter, "Filter shown items")

	widget.SetKeyboardKey(tcell.KeyDown, widget.NextTodo, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.PrevTodo, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelected, "Delete item")
	widget.SetKeyboardKey(tcell.KeyCtrlJ, widget.demoteSelected, "Demote item")
	widget.SetKeyboardKey(tcell.KeyCtrlL, widget.makeSelectedLast, "Make item last")
	widget.SetKeyboardKey(tcell.KeyCtrlK, widget.promoteSelected, "Promote item")
	widget.SetKeyboardKey(tcell.KeyCtrlF, widget.makeSelectedFirst, "Make item first")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.updateSelected, "Edit item")

}

func (widget *Widget) NextTodo() {
	newIndex := widget.Selected + 1
	for newIndex < len(widget.list.Items) && !widget.shouldShowItem(widget.list.Items[newIndex]) {
		newIndex = newIndex + 1
	}
	if newIndex < len(widget.list.Items) {
		widget.Selected = newIndex
	}
	widget.display()
}

func (widget *Widget) PrevTodo() {
	newIndex := widget.Selected - 1
	for newIndex >= 0 && !widget.shouldShowItem(widget.list.Items[newIndex]) {
		newIndex = newIndex - 1
	}
	if newIndex >= 0 {
		widget.Selected = newIndex
	}
	widget.display()
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

func (widget *Widget) makeSelectedLast() {
	if !widget.isItemSelected() {
		return
	}

	j := widget.Selected + 1
	if j >= len(widget.list.Items) {
		return
	}

	for j < len(widget.list.Items) {
		widget.list.Swap(widget.Selected, j)
		widget.Selected = j
		j = j + 1
	}

	if widget.settings.parseDates {
		widget.Selected = widget.placeItemBasedOnDate(widget.Selected)
	}

	widget.persist()
	widget.display()
}

func (widget *Widget) openFile() {
	confDir, _ := cfg.WtfConfigDir()
	utils.OpenFile(fmt.Sprintf("%s/%s", confDir, widget.filePath))
}

func (widget *Widget) setTag() {
	if !widget.settings.parseTags {
		return
	}

	widget.processFormInput("Tag prefix:", "", func(filter string) {
		widget.showTagPrefix = filter
	})
}

func (widget *Widget) setFilter() {
	widget.processFormInput("Filter:", "", func(filter string) {
		widget.showFilter = strings.ToLower(filter)
	})
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

func (widget *Widget) makeSelectedFirst() {
	if !widget.isItemSelected() {
		return
	}

	j := widget.Selected - 1
	if j < 0 {
		return
	}

	for j >= 0 {
		widget.list.Swap(widget.Selected, j)
		widget.Selected = j
		j = j - 1
	}

	if widget.settings.parseDates {
		widget.Selected = widget.placeItemBasedOnDate(widget.Selected)
	}

	widget.persist()
	widget.display()
}

func (widget *Widget) toggleChecked() {
	selectedItem := widget.SelectedItem()
	if selectedItem == nil {
		return
	}

	selectedItem.Toggle()

	if !selectedItem.Checked {
		widget.Selected = widget.placeItemBasedOnDate(widget.Selected)
	}

	widget.persist()
	widget.display()
}

func (widget *Widget) unselect() {
	if widget.showFilter != "" {
		widget.showFilter = ""
	} else {
		widget.Selected = -1
	}
	widget.display()
}
