package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	Init("cats", []string{"dogs"})

	assert.Equal(t, OpenFileUtil, "cats")
	assert.Equal(t, OpenUrlUtil, []string{"dogs"})
}
