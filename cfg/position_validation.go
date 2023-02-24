package cfg

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

// Common examples of invalid position configuration are:
//
//	position:
//	  top: -3
//	  left: 2
//	  width: 0
//	  height: 1
//
//	position:
//	  top: 3
//	  width: 2
//	  height: 1
//
//	position:
//	  top: 3
//	  # left: 2
//	  width: 2
//	  height: 1
//
//	position:
//	top: 3
//	left: 2
//	width: 2
//	height: 1
type positionValidation struct {
	err    error
	name   string
	intVal int
}

func (posVal *positionValidation) Error() error {
	return posVal.err
}

func (posVal *positionValidation) HasError() bool {
	return posVal.err != nil
}

func (posVal *positionValidation) IntValue() int {
	return posVal.intVal
}

// String returns the Stringer representation of the positionValidation
func (posVal *positionValidation) String() string {
	return fmt.Sprintf("Invalid value for %s:\t%d", aurora.Yellow(posVal.name), posVal.intVal)
}

/* -------------------- Unexported Functions -------------------- */

func newPositionValidation(name string, intVal int, err error) *positionValidation {
	posVal := &positionValidation{
		err:    err,
		name:   name,
		intVal: intVal,
	}

	return posVal
}
