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
	err       error
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
		widget.err = err
		widget.calEvents = []*CalEvent{}
	} else {
		widget.err = nil
		widget.calEvents = calEvents
	}

	widget.display()
}
