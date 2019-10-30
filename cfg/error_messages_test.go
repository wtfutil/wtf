package cfg

import (
	"errors"
)

var (
	errTest = errors.New("Busted")
)

func ExampleDisplayError() {
	displayError(errTest)
	// Output: [31mError:[0m Busted
}
