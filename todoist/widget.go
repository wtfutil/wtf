package todoist

import (
	"os"

	"github.com/darkSasori/todoist"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app   *tview.Application
	pages *tview.Pages
	list  []*List
	idx   int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Todoist ", "todoist", true),

		app:   app,
		pages: pages,
	}

	todoist.Token = os.Getenv("WTF_TODOIST_TOKEN")
	widget.list = loadProjects()
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

func (w *Widget) Refresh() {
	if w.Disabled() || len(w.list) == 0 {
		return
	}

	w.UpdateRefreshedAt()
	w.display()
}

func (w *Widget) Next() {
	w.idx = w.idx + 1
	if w.idx == len(w.list) {
		w.idx = 0
	}

	w.display()
}

func (w *Widget) Prev() {
	w.idx = w.idx - 1
	if w.idx < 0 {
		w.idx = len(w.list) - 1
	}

	w.display()
}

func (w *Widget) Down() {
	w.list[w.idx].down()
	w.display()
}

func (w *Widget) UP() {
	w.list[w.idx].up()
	w.display()
}

func (w *Widget) Close() {
	w.list[w.idx].close()
	if w.list[w.idx].isLast() {
		w.UP()
		return
	}
	w.Down()
}

func (w *Widget) Delete() {
	w.list[w.idx].close()
	if w.list[w.idx].isLast() {
		w.UP()
		return
	}
	w.Down()
}

func loadProjects() []*List {
	lists := []*List{}
	for _, id := range wtf.Config.UList("wtf.mods.todoist.projects") {
		list := NewList(id.(int))
		lists = append(lists, list)
	}

	return lists
}

func fromVim(event *tcell.EventKey) tcell.Key {
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
