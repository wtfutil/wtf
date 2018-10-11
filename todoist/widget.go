package todoist

import (
	"os"

	"github.com/darkSasori/todoist"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Todoist:

   /: Show/hide this help window
   c: Close the selected item
   d: Delete the selected item
   h: Previous Todoist list
   j: Select the next item in the list
   k: Select the previous item in the list
   l: Next Todoist list
   r: Refresh the todo list data

   arrow down: Select the next item in the list
   arrow left: Previous Todoist list
   arrow right: Next Todoist list
   arrow up: Select the previous item in the list
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	projects []*Project
	idx      int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, "Todoist", "todoist", true),
	}

	widget.loadAPICredentials()
	widget.projects = loadProjects()

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) CurrentProject() *Project {
	return widget.ProjectAt(widget.idx)
}

func (widget *Widget) ProjectAt(idx int) *Project {
	if len(widget.projects) == 0 {
		return nil
	}

	return widget.projects[idx]
}

func (w *Widget) NextProject() {
	w.idx = w.idx + 1
	if w.idx == len(w.projects) {
		w.idx = 0
	}

	w.display()
}

func (w *Widget) PreviousProject() {
	w.idx = w.idx - 1
	if w.idx < 0 {
		w.idx = len(w.projects) - 1
	}

	w.display()
}

func (w *Widget) Refresh() {
	if w.Disabled() || w.CurrentProject() == nil {
		return
	}

	w.display()
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

func (w *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	if len(w.projects) == 0 {
		return event
	}

	switch string(event.Rune()) {
	case "/":
		w.ShowHelp()
		return nil
	case "r":
		w.Refresh()
		return nil
	case "d":
		w.Delete()
		return nil
	case "c":
		w.Close()
		return nil
	}

	switch w.vimBindings(event) {
	case tcell.KeyLeft:
		w.PreviousProject()
		return nil
	case tcell.KeyRight:
		w.NextProject()
		return nil
	case tcell.KeyUp:
		w.Up()
		return nil
	case tcell.KeyDown:
		w.Down()
		return nil
	}

	return event
}

func (widget *Widget) loadAPICredentials() {
	todoist.Token = wtf.Config.UString(
		"wtf.mods.todoist.apiKey",
		os.Getenv("WTF_TODOIST_TOKEN"),
	)
}

func loadProjects() []*Project {
	projects := []*Project{}

	for _, id := range wtf.Config.UList("wtf.mods.todoist.projects") {
		proj := NewProject(id.(int))
		projects = append(projects, proj)
	}

	return projects
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
