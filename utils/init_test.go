package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	Init("cats")
	Equal(t, OpenFileUtil, "cats")
}
