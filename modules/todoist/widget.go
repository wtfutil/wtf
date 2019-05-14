package todoist

import (
	"github.com/darkSasori/todoist"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// A Widget represents a Todoist widget
type Widget struct {
	wtf.KeyboardWidget
	wtf.TextWidget
	wtf.MultiSourceWidget

	projects []*Project
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    wtf.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:        wtf.NewTextWidget(app, settings.common, true),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "project", "projects"),

		settings: settings,
	}

	widget.loadAPICredentials()
	widget.loadProjects()

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

func (w *Widget) Refresh() {
	if w.Disabled() || w.CurrentProject() == nil {
		return
	}

	w.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Keyboard Movement -------------------- */

// Down selects the next item in the list
func (w *Widget) Down() {
	w.CurrentProject().down()
	w.display()
}

// Up selects the previous item in the list
func (w *Widget) Up() {
	w.CurrentProject().up()
	w.display()
}

// Close closes the currently-selected task in the currently-selected project
func (w *Widget) Close() {
	w.CurrentProject().closeSelectedTask()

	if w.CurrentProject().isLast() {
		w.Up()
		return
	}

	w.Down()
}

// Delete deletes the currently-selected task in the currently-selected project
func (w *Widget) Delete() {
	w.CurrentProject().deleteSelectedTask()

	if w.CurrentProject().isLast() {
		w.Up()
		return
	}

	w.Down()
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

func (w *Widget) vimBindings(event *tcell.EventKey) tcell.Key {
	switch string(event.Rune()) {
	case "h":
		return tcell.KeyLeft
	case "l":
		return tcell.KeyRight
	case "k":
		return tcell.KeyUp
	case "j":
		return tcell.KeyDown
	}
	return event.Key()
}
