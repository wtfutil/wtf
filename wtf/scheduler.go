package wtf

import (
	"time"
)

type Scheduler interface {
	Refresh()
	RefreshInterval() int
}

func Schedule(widget Wtfable) {
	// Kick off the first refresh and then leave the rest to the timer
	widget.Refresh()

	interval := time.Duration(widget.RefreshInterval()) * time.Second

	if interval <= 0 {
		return
	}

	tick := time.NewTicker(interval)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			if widget.Enabled() {
				widget.Refresh()
			} else {
				tick.Stop()
				return
			}
		case <-quit:
			tick.Stop()
			return
		}
	}
}
