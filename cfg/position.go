package cfg

import (
	"fmt"
	"os"
	"strings"

	"github.com/olebedev/config"
)

const (
	positionPath = "position"
)

/* -------------------- Position Validation -------------------- */

type positionValidation struct {
	err  error
	name string
	val  int
}

func (posVal *positionValidation) hasError() bool {
	return posVal.err != nil
}

// String returns the Stringer representation of the positionValidation
func (posVal *positionValidation) String() string {
	return fmt.Sprintf("Invalid value for %s:\t%d", posVal.name, posVal.val)
}

func newPositionValidation(name string, val int, err error) *positionValidation {
	posVal := &positionValidation{
		err:  err,
		name: name,
		val:  val,
	}

	return posVal
}

/* -------------------- Position -------------------- */

// Position represents the onscreen location of a widget
type Position struct {
	Height int
	Left   int
	Top    int
	Width  int
}

// NewPositionFromYAML creates and returns a new instance of Position
func NewPositionFromYAML(moduleName string, moduleConfig *config.Config) Position {
	var val int
	var err error
	validations := make(map[string]*positionValidation)

	// Parse the positional data from the config data
	val, err = moduleConfig.Int(positionPath + ".top")
	validations["top"] = newPositionValidation("top", val, err)

	val, err = moduleConfig.Int(positionPath + ".left")
	validations["left"] = newPositionValidation("left", val, err)

	val, err = moduleConfig.Int(positionPath + ".width")
	validations["width"] = newPositionValidation("width", val, err)

	val, err = moduleConfig.Int(positionPath + ".height")
	validations["height"] = newPositionValidation("height", val, err)

	validatePositions(moduleName, validations)

	pos := Position{
		Top:    validations["top"].val,
		Left:   validations["left"].val,
		Width:  validations["width"].val,
		Height: validations["height"].val,
	}

	return pos
}

/* -------------------- Unexported Functions -------------------- */

// If any of the position values have an error then we inform the user and exit the app
// Common examples of invalid position configuration are:
//
//    position:
//      top: 3
//      width: 2
//      height: 1
//
//    position:
//      top: 3
//      # left: 2
//      width: 2
//      height: 1
//
//    position:
//    top: 3
//    left: 2
//    width: 2
//    height: 1
//
func validatePositions(moduleName string, validations map[string]*positionValidation) {
	var errStr string

	for _, posVal := range validations {
		if posVal.hasError() {
			errStr += fmt.Sprintf("  - %s.\t\033[0;31mError\033[0m %v\n", posVal, posVal.err)
		}
	}

	if errStr != "" {
		fmt.Println()
		fmt.Printf("\033[0;1mErrors in %s position configuration\033[0m\n", strings.Title(moduleName))
		fmt.Println(errStr)
		fmt.Println("Please check your config.yml file.")
		fmt.Println()

		os.Exit(1)
	}
}
