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

// Position represents the onscreen location of a widget
type Position struct {
	Height int
	Left   int
	Top    int
	Width  int
}

// NewPositionFromYAML creates and returns a new instance of Position
func NewPositionFromYAML(moduleName string, moduleConfig *config.Config) Position {
	errs := make(map[string]error)

	// Parse the positional data from the config data
	top, err := moduleConfig.Int(positionPath + ".top")
	errs["top"] = err

	left, err := moduleConfig.Int(positionPath + ".left")
	errs["left"] = err

	width, err := moduleConfig.Int(positionPath + ".width")
	errs["width"] = err

	height, err := moduleConfig.Int(positionPath + ".height")
	errs["height"] = err

	validatePositions(moduleName, errs)

	pos := Position{
		Top:    top,
		Left:   left,
		Width:  width,
		Height: height,
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
func validatePositions(moduleName string, errs map[string]error) {
	var errStr string

	for pos, err := range errs {
		if err != nil {
			errStr += fmt.Sprintf("  - Invalid value for %s\n", pos)
		}
	}

	if errStr != "" {
		fmt.Println()
		fmt.Printf("\033[0;31mErrors in %s configuration\033[0m\n", strings.Title(moduleName))
		fmt.Println(errStr)
		fmt.Println("Please check your config.yml file.")
		fmt.Println()

		os.Exit(1)
	}
}
