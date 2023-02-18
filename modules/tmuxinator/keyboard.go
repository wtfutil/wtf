package tmuxinator

import (
	"github.com/gdamore/tcell/v2"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next project")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous project")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next project")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.startProject, "Start or go to project")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}

func (widget *Widget) Next() {
	widget.Selected++

	if widget.Selected >= widget.MaxItems() {
		widget.Selected = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Selected--

	if widget.Selected < 0 {
		widget.Selected = widget.MaxItems() - 1
	}

	widget.display()
}

func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.display()
}

func (widget *Widget) startProject() {
	if widget.GetSelected() >= 0 && len(widget.Items) > 0 {
		projectName := widget.Items[widget.GetSelected()]

		startTmuxProject(projectName)
	}
}
