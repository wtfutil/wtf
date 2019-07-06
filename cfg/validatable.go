package cfg

// Validatable is implemented by any value that validates a configuration setting
type Validatable interface {
	Error() error
	HasError() bool
	String() string
	IntValue() int
}
