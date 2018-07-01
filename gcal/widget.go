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
		TextWidget: wtf.NewTextWidget(" Calendar ", "gcal", true),
		app:        app,
		ch:         make(chan struct{}),
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
	calEvents, err := Fetch(CreateCodeInputDialog(" Calendar ", widget))
	if err != nil {
		widget.calEvents = []*CalEvent{}
	} else {
		widget.calEvents = calEvents
	}

	widget.UpdateRefreshedAt()
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

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
