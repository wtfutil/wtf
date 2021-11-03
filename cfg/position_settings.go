package cfg

import (
	"github.com/olebedev/config"
)

const (
	positionPath = "position"
)

// PositionSettings represents the onscreen location of a widget
type PositionSettings struct {
	Validations *Validations

	Height int
	Left   int
	Top    int
	Width  int
}

// NewPositionSettingsFromYAML creates and returns a new instance of cfg.Position
func NewPositionSettingsFromYAML(moduleConfig *config.Config) PositionSettings {
	var currVal int
	var err error

	validations := NewValidations()

	// Parse the positional data from the config data
	currVal, err = moduleConfig.Int(positionPath + ".top")
	validations.append("top", newPositionValidation("top", currVal, err))

	currVal, err = moduleConfig.Int(positionPath + ".left")
	validations.append("left", newPositionValidation("left", currVal, err))

	currVal, err = moduleConfig.Int(positionPath + ".width")
	validations.append("width", newPositionValidation("width", currVal, err))

	currVal, err = moduleConfig.Int(positionPath + ".height")
	validations.append("height", newPositionValidation("height", currVal, err))

	pos := PositionSettings{
		Validations: validations,

		Top:    validations.intValueFor("top"),
		Left:   validations.intValueFor("left"),
		Width:  validations.intValueFor("width"),
		Height: validations.intValueFor("height"),
	}

	return pos
}
