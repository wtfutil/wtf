package gcal

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	app       *tview.Application
	calEvents []*CalEvent
	settings  *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Disable() {
	widget.TextWidget.Disable()
}

func (widget *Widget) Refresh() {
	if isAuthenticated() {
		widget.fetchAndDisplayEvents()
		return
	}

	widget.app.Suspend(widget.authenticate)
	widget.Refresh()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) fetchAndDisplayEvents() {
	calEvents, err := widget.Fetch()
	if err != nil {
		widget.calEvents = []*CalEvent{}
	} else {
		widget.calEvents = calEvents
	}

	widget.display()
}
