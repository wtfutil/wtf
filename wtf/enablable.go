package wtf

// Enablable is the interface that enforces enable/disable capabilities on a module
type Enablable interface {
	Disable()
	Disabled() bool
	Enabled() bool
}
