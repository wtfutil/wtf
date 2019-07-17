package wtf

import (
	"time"
)

// Schedulable is the interface that enforces scheduling capabilities on a module
type Schedulable interface {
	Refresh()
	Refreshing() bool
	RefreshInterval() int
}

// Schedule kicks off the first refresh of a module's data and then queues the rest of the
// data refreshes on a timer
func Schedule(widget Wtfable) {
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
