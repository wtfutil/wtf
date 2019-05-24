package todoist

import (
	"github.com/darkSasori/todoist"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// A Widget represents a Todoist widget
type Widget struct {
	wtf.KeyboardWidget
	wtf.MultiSourceWidget
	wtf.ScrollableWidget

	projects []*Project
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    wtf.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "project", "projects"),
		ScrollableWidget:  wtf.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.loadAPICredentials()
	widget.loadProjects()

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.SetDisplayFunction(widget.display)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) CurrentProject() *Project {
	return widget.ProjectAt(widget.Idx)
}

func (widget *Widget) ProjectAt(idx int) *Project {
	if len(widget.projects) == 0 {
		return nil
	}

	return widget.projects[idx]
}

func (widget *Widget) Refresh() {
	if widget.Disabled() || widget.CurrentProject() == nil {
		widget.SetItemCount(0)
		return
	}

	widget.loadProjects()

	widget.SetItemCount(len(widget.CurrentProject().tasks))
	widget.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

func (widget *Widget) NextSource() {
	widget.MultiSourceWidget.NextSource()
	widget.Selected = widget.CurrentProject().index
	widget.SetItemCount(len(widget.CurrentProject().tasks))
	widget.RenderFunction()
}

func (widget *Widget) PrevSource() {
	widget.MultiSourceWidget.PrevSource()
	widget.Selected = widget.CurrentProject().index
	widget.SetItemCount(len(widget.CurrentProject().tasks))
	widget.RenderFunction()
}

func (widget *Widget) Prev() {
	widget.ScrollableWidget.Prev()
	widget.CurrentProject().index = widget.Selected
}

func (widget *Widget) Next() {
	widget.ScrollableWidget.Next()
	widget.CurrentProject().index = widget.Selected
}

func (widget *Widget) Unselect() {
	widget.ScrollableWidget.Unselect()
	widget.CurrentProject().index = -1
	widget.RenderFunction()
}

/* -------------------- Keyboard Movement -------------------- */

// Close closes the currently-selected task in the currently-selected project
func (w *Widget) Close() {
	w.CurrentProject().closeSelectedTask()

	if w.CurrentProject().isLast() {
		w.Prev()
		return
	}

	w.Next()
}

// Delete deletes the currently-selected task in the currently-selected project
func (w *Widget) Delete() {
	w.CurrentProject().deleteSelectedTask()

	if w.CurrentProject().isLast() {
		w.Prev()
		return
	}

	w.Next()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) loadAPICredentials() {
	todoist.Token = widget.settings.apiKey
}

func (widget *Widget) loadProjects() {
	projects := []*Project{}

	for _, id := range widget.settings.projects {
		proj := NewProject(id.(int))
		projects = append(projects, proj)
	}

	widget.projects = projects
}
