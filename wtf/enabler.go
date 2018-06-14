package wtf

type Enabler interface {
	Disabled() bool
	Enabled() bool
	Disable()
}
