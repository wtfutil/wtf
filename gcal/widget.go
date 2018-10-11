package gcal

import (
	"sync"
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	calEvents []*CalEvent
	ch        chan struct{}
	mutex     sync.Mutex
	app       *tview.Application
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "Calendar", "gcal", true),
		ch:         make(chan struct{}),
		app:        app,
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
	widget.app.Suspend(authenticate)
	widget.Refresh()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) fetchAndDisplayEvents() {
	calEvents, err := Fetch()
	if err != nil {
		widget.calEvents = []*CalEvent{}
	} else {
		widget.calEvents = calEvents
	}
	widget.display()
}

func updateLoop(widget *Widget) {
	interval := wtf.Config.UInt("wtf.mods.gcal.textInterval", 30)
	if interval == 0 {
		return
	}

	tick := time.NewTicker(time.Duration(interval) * time.Second)
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
