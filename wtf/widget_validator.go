package wtf

import (
	"fmt"
	"log"
)

// Check that all the loaded widgets are valid for display
func ValidateWidgets(widgets []Wtfable) (bool, error) {
	result := true
	var err error

	for _, widget := range widgets {
		if widget.Enabled() && !widget.IsPositionable() {
			errStr := fmt.Sprintf("Widget config has invalid values: %s", widget.Key())
			log.Fatalln(errStr)
		}
	}

	return result, err
}
