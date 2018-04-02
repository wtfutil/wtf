package wtf

import (
	"time"
)

type Refresher interface {
	Refresh()
	RefreshInterval() int
}

func Refresh(widget Refresher) {
	tick := time.NewTicker(time.Duration(widget.RefreshInterval()) * time.Second)
	quit := make(chan struct{})

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
