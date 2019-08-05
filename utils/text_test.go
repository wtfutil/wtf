package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_CenterText(t *testing.T) {
	Equal(t, "cat", CenterText("cat", -9))
	Equal(t, "cat", CenterText("cat", 0))
	Equal(t, "   cat   ", CenterText("cat", 9))
}

func Test_PadRow(t *testing.T) {
	Equal(t, "", PadRow(0, 0))
	Equal(t, "", PadRow(5, 2))
	Equal(t, " ", PadRow(1, 2))
}
