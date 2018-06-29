package wtftests

import (
	"testing"

	"github.com/gdamore/tcell"
	. "github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

func TestColorFor(t *testing.T) {
	Equal(t, tcell.ColorRed, ColorFor("red"))
	Equal(t, tcell.ColorGreen, ColorFor("cat"))
}
