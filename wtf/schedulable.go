package wtf

import "time"

// Schedulable is the interface that enforces scheduling capabilities on a module
type Schedulable interface {
	Refresh()
	Refreshing() bool
	RefreshInterval() time.Duration
}
