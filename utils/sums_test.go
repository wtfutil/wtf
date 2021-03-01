package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumInts(t *testing.T) {
	expected := 6
	result := SumInts([]int{1, 3, 2})

	assert.Equal(t, expected, result)

	expected = 46
	result = SumInts([]int{4, 6, 7, 23, 6})

	assert.Equal(t, expected, result)

	expected = 4
	result = SumInts([]int{4})

	assert.Equal(t, expected, result)
}
