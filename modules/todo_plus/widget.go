package todo_plus

import (
	"log"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/modules/todo_plus/backend"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Todoist widget
type Widget struct {
	view.MultiSourceWidget
	view.ScrollableWidget

	projects []*backend.Project
	settings *Settings
	backend  backend.Backend
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "project", "projects"),
		ScrollableWidget:  view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.backend = getBackend(settings.backendType)
	widget.backend.Setup(settings.backendSettings)
	widget.CommonSettings().Title = widget.backend.Title()

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()
	widget.SetDisplayFunction(widget.display)

	return &widget
}

func getBackend(backendType string) backend.Backend {
	switch backendType {
	case "trello":
		backend := &backend.Trello{}
		return backend
	case "todoist":
		backend := &backend.Todoist{}
		return backend
	default:
		log.Fatal(backendType + " is not a supported backend")
		return nil
	}

}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) CurrentProject() *backend.Project {
	return widget.ProjectAt(widget.Idx)
}

func (widget *Widget) ProjectAt(idx int) *backend.Project {
	if len(widget.projects) == 0 {
		return nil
	}

	return widget.projects[idx]
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.projects = widget.backend.BuildProjects()
	widget.Sources = widget.backend.Sources()
	widget.SetItemCount(len(widget.CurrentProject().Tasks))
	widget.display()
}

func (widget *Widget) NextSource() {
	widget.MultiSourceWidget.NextSource()
	widget.Selected = widget.CurrentProject().Index
	widget.SetItemCount(len(widget.CurrentProject().Tasks))
	widget.RenderFunction()
}

func (widget *Widget) PrevSource() {
	widget.MultiSourceWidget.PrevSource()
	widget.Selected = widget.CurrentProject().Index
	widget.SetItemCount(len(widget.CurrentProject().Tasks))
	widget.RenderFunction()
}

func (widget *Widget) Prev() {
	widget.ScrollableWidget.Prev()
	widget.CurrentProject().Index = widget.Selected
}

func (widget *Widget) Next() {
	widget.ScrollableWidget.Next()
	widget.CurrentProject().Index = widget.Selected
}

func (widget *Widget) Unselect() {
	widget.ScrollableWidget.Unselect()
	widget.CurrentProject().Index = -1
	widget.RenderFunction()
}

/* -------------------- Keyboard Movement -------------------- */

// Close closes the currently-selected task in the currently-selected project
func (w *Widget) Close() {
	w.CurrentProject().CloseSelectedTask()
	w.SetItemCount(len(w.CurrentProject().Tasks))

	if w.CurrentProject().IsLast() {
		w.Prev()
		return
	}
	w.CurrentProject().Index = w.Selected
	w.RenderFunction()
}

// Delete deletes the currently-selected task in the currently-selected project
func (w *Widget) Delete() {
	w.CurrentProject().DeleteSelectedTask()
	w.SetItemCount(len(w.CurrentProject().Tasks))

	if w.CurrentProject().IsLast() {
		w.Prev()
	}
	w.CurrentProject().Index = w.Selected
	w.RenderFunction()
}
