package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	Init("cats", []string{"dogs"})

	Equal(t, OpenFileUtil, "cats")
	Equal(t, OpenUrlUtil, []string{"dogs"})
}
