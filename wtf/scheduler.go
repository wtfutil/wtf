package wtf

import (
	"time"
)

type Scheduler interface {
	Refresh()
	RefreshInterval() int
}

func Schedule(widget Scheduler) {
	tick := time.NewTicker(time.Duration(widget.RefreshInterval()) * time.Second)
	quit := make(chan struct{})

	// Kick off the first refresh and then leave the rest to the timer
	widget.Refresh()

	for {
		select {
		case <-tick.C:
			widget.Refresh()
		case <-quit:
			tick.Stop()
			return
		}
	}
}
