package app

import (
	"time"

	"github.com/wtfutil/wtf/wtf"
)

// Schedule kicks off the first refresh of a module's data and then queues the rest of the
// data refreshes on a timer
func Schedule(widget wtf.Wtfable) {
	widget.Refresh()

	interval := widget.CommonSettings().RefreshInterval

	if interval <= 0 {
		return
	}

	timer := time.NewTicker(interval)

	for {
		select {
		case <-timer.C:
			if widget.Enabled() {
				widget.Refresh()
			} else {
				timer.Stop()
				return
			}
		case quit := <-widget.QuitChan():
			if quit {
				timer.Stop()
				return
			}
		}
	}
}
