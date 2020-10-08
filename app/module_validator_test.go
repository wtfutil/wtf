package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewModuleValidator(t *testing.T) {
	assert.IsType(t, &ModuleValidator{}, NewModuleValidator())
}
