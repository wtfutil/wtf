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

func Test_RowPadding(t *testing.T) {
	Equal(t, "", RowPadding(0, 0))
	Equal(t, "", RowPadding(5, 2))
	Equal(t, " ", RowPadding(1, 2))
	Equal(t, "     ", RowPadding(0, 5))
}
