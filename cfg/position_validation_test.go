package cfg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	posVal = &positionValidation{
		err:    errors.New("Busted"),
		name:   "top",
		intVal: -3,
	}
)

func Test_Attributes(t *testing.T) {
	assert.EqualError(t, posVal.Error(), "Busted")
	assert.Equal(t, true, posVal.HasError())
	assert.Equal(t, -3, posVal.IntValue())

	assert.Contains(t, posVal.String(), "Invalid")
	assert.Contains(t, posVal.String(), "top")
	assert.Contains(t, posVal.String(), "-3")
}
