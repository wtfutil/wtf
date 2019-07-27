package wtf

// Schedulable is the interface that enforces scheduling capabilities on a module
type Schedulable interface {
	Refresh()
	Refreshing() bool
	RefreshInterval() int
}
