package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumInts(t *testing.T) {
	s := []int{1, 3, 2}
	assert.Equal(t, 6, SumInts(s))
	s = []int{4, 6, 7, 23, 6}
	assert.Equal(t, 46, SumInts(s))
	s = []int{4}
	assert.Equal(t, 4, SumInts(s))
}
