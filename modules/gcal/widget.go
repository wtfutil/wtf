package gcal

import (
	"sync"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app       *tview.Application
	calEvents []*CalEvent
	mutex     sync.Mutex
	settings  *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, true),

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
