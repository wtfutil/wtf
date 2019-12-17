package app

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/wtfutil/wtf/wtf"
)

type ModuleValidator struct{}

func NewModuleValidator() *ModuleValidator {
	val := &ModuleValidator{}
	return val
}

// Validate rolls through all the enabled widgets and looks for configuration errors.
// If it finds any it stringifies them, writes them to the console, and kills the app gracefully
func (val *ModuleValidator) Validate(widgets []wtf.Wtfable) {
	var errStr string
	hasErrors := false

	for _, widget := range widgets {
		var widgetErrStr string

		for _, val := range widget.CommonSettings().Validations() {
			if val.HasError() {
				hasErrors = true
				widgetErrStr += fmt.Sprintf(" - %s\t%s %v\n", val, aurora.Red("Error:"), val.Error())
			}
		}

		if widgetErrStr != "" {
			errStr += fmt.Sprintf(
				"%s\n",
				fmt.Sprintf(
					"%s in %s configuration",
					aurora.Red("Errors"),
					aurora.Yellow(
						fmt.Sprintf(
							"%s.position",
							widget.Name(),
						),
					),
				),
			)

			errStr += widgetErrStr + "\n"
		}
	}

	if hasErrors {
		fmt.Println()
		fmt.Println(errStr)

		os.Exit(1)
	}
}
