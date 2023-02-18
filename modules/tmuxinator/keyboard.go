package tmuxinator

import (
	"github.com/gdamore/tcell/v2"
	tc "github.com/wtfutil/wtf/modules/tmuxinator/client"
	"strings"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next project")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous project")
	widget.SetKeyboardChar("e", widget.editProject, "Edit project")
	widget.SetKeyboardChar("n", widget.newProject, "Create new project")
	widget.SetKeyboardChar("c", widget.cloneProject, "Clone existing project")

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

		tc.StartProject(projectName)
	}
}

func (widget *Widget) editProject() {
	if widget.GetSelected() >= 0 && len(widget.Items) > 0 {
		projectName := widget.Items[widget.GetSelected()]

		tc.EditProject(projectName)
	}
}

func (widget *Widget) newProject() {
	widget.processFormInput("New Project:", "", func(t string) {
		projectName := strings.ReplaceAll(t, " ", "_")

		tc.EditProject(projectName)
		widget.Base.RedrawChan <- true
	})
}

func (widget *Widget) cloneProject() {
	if widget.GetSelected() >= 0 && len(widget.Items) > 0 {
		currentProjectName := widget.Items[widget.GetSelected()]

		widget.processFormInput("Copy Project:", currentProjectName, func(t string) {
			newProjectName := strings.ReplaceAll(t, " ", "_")

			tc.CopyProject(currentProjectName, newProjectName)
			widget.Base.RedrawChan <- true
		})
	}
}
