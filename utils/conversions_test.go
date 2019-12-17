package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapToStrs(t *testing.T) {
	expected := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	source := make(map[string]interface{})
	for _, val := range expected {
		source[val] = val
	}

	assert.Equal(t, expected, MapToStrs(source))
}

func Test_IntsToUints(t *testing.T) {
	tests := []struct {
		name     string
		src      []int
		expected []uint
	}{
		{
			name:     "empty set",
			src:      []int{},
			expected: []uint{},
		},
		{
			name:     "full set",
			src:      []int{1, 2, 3},
			expected: []uint{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IntsToUints(tt.src)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToInts(t *testing.T) {
	expected := []int{1, 2, 3}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	assert.Equal(t, expected, ToInts(source))
}

func Test_ToStrs(t *testing.T) {
	expectedInts := []int{1, 2, 3}
	expectedStrs := []string{"1", "2", "3"}

	fromInts := make([]interface{}, 3)
	for idx, val := range expectedInts {
		fromInts[idx] = val
	}

	fromStrs := make([]interface{}, 3)
	for idx, val := range expectedStrs {
		fromStrs[idx] = val
	}

	assert.Equal(t, expectedStrs, ToStrs(fromInts))
	assert.Equal(t, expectedStrs, ToStrs(fromStrs))
}

func Test_ToUints(t *testing.T) {
	expected := []uint{1, 2, 3}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	assert.Equal(t, expected, ToUints(source))
}
