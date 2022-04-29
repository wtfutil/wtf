package gcal

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	calEvents []*CalEvent
	err       error
	settings  *Settings
	tviewApp  *tview.Application
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		tviewApp: tviewApp,
		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Disable() {
	widget.TextWidget.Disable()
}

func (widget *Widget) Refresh() {
	if isAuthenticated(widget.settings.email) {
		widget.fetchAndDisplayEvents()
		return
	}

	widget.tviewApp.Suspend(widget.authenticate)
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
