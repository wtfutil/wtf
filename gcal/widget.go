package gcal

import (
	"sync"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
	"google.golang.org/api/calendar/v3"
)

type Widget struct {
	wtf.TextWidget

	events *calendar.Events
	ch     chan struct{}
	mutex  sync.Mutex
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Calendar ", "gcal", false),
		ch:         make(chan struct{}),
	}

	go updateLoop(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	events, _ := Fetch()
	widget.events = events

	widget.UpdateRefreshedAt()

	widget.display()
}

func (widget *Widget) Disable() {
	close(widget.ch)
	widget.TextWidget.Disable()
}

/* -------------------- Unexported Functions -------------------- */

// conflicts returns TRUE if this event conflicts with another, FALSE if it does not
func (widget *Widget) conflicts(event *calendar.Event, events *calendar.Events) bool {
	conflict := false

	for _, otherEvent := range events.Items {
		if event == otherEvent {
			continue
		}

		eventStart, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		eventEnd, _ := time.Parse(time.RFC3339, event.End.DateTime)

		otherEnd, _ := time.Parse(time.RFC3339, otherEvent.End.DateTime)
		otherStart, _ := time.Parse(time.RFC3339, otherEvent.Start.DateTime)

		if eventStart.Before(otherEnd) && eventEnd.After(otherStart) {
			conflict = true
			break
		}
	}

	return conflict
}

func (widget *Widget) eventIsAllDay(event *calendar.Event) bool {
	return len(event.Start.Date) > 0
}

// eventIsNow returns true if the event is happening now, false if it not
func (widget *Widget) eventIsNow(event *calendar.Event) bool {
	startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, event.End.DateTime)

	return time.Now().After(startTime) && time.Now().Before(endTime)
}

func (widget *Widget) eventIsPast(event *calendar.Event) bool {
	if widget.eventIsAllDay(event) {
		return false
	} else {
		ts, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		return (widget.eventIsNow(event) == false) && ts.Before(time.Now())
	}
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
