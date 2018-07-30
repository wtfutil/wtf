package gcal

import (
	"sync"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	calEvents []*CalEvent
	ch        chan struct{}
	mutex     sync.Mutex
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Calendar", "gcal", false),
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
	calEvents, err := Fetch()
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
