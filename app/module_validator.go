package app

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

// ModuleValidator is responsible for validating the state of a module's configuration
type ModuleValidator struct{}

type widgetError struct {
	name             string
	validationErrors []cfg.Validatable
}

// NewModuleValidator creates and returns an instance of ModuleValidator
func NewModuleValidator() *ModuleValidator {
	return &ModuleValidator{}
}

// Validate rolls through all the enabled widgets and looks for configuration errors.
// If it finds any it stringifies them, writes them to the console, and kills the app gracefully
func (val *ModuleValidator) Validate(widgets []wtf.Wtfable) {
	validationErrors := validate(widgets)

	if len(validationErrors) > 0 {
		fmt.Println()
		for _, error := range validationErrors {
			for _, message := range error.errorMessages() {
				fmt.Println(message)
			}
		}
		fmt.Println()

		os.Exit(1)
	}
}

func validate(widgets []wtf.Wtfable) (widgetErrors []widgetError) {
	for _, widget := range widgets {
		err := widgetError{name: widget.Name()}

		for _, val := range widget.CommonSettings().Validations() {
			if val.HasError() {
				err.validationErrors = append(err.validationErrors, val)
			}
		}

		if len(err.validationErrors) > 0 {
			widgetErrors = append(widgetErrors, err)
		}
	}

	return widgetErrors
}

func (err widgetError) errorMessages() (messages []string) {
	widgetMessage := fmt.Sprintf(
		"%s in %s configuration",
		aurora.Red("Errors"),
		aurora.Yellow(
			fmt.Sprintf(
				"%s.position",
				err.name,
			),
		),
	)
	messages = append(messages, widgetMessage)

	for _, e := range err.validationErrors {
		configMessage := fmt.Sprintf(" - %s\t%s %v", e.String(), aurora.Red("Error:"), e.Error())

		messages = append(messages, configMessage)
	}

	return messages
}
