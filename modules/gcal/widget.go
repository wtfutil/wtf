package gcal

import (
	"sync"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app       *tview.Application
	calEvents []*CalEvent
	ch        chan struct{}
	mutex     sync.Mutex
	settings  *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		ch:       make(chan struct{}),
		settings: settings,
	}

	go updateLoop(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Disable() {
	close(widget.ch)
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

func updateLoop(widget *Widget) {
	if widget.settings.textInterval == 0 {
		return
	}

	tick := time.NewTicker(time.Duration(widget.settings.textInterval) * time.Second)
	defer tick.Stop()
outer:
	for {
		select {
		case <-tick.C:
			widget.display()
		case <-widget.ch:
			break outer
		}
	}
}
