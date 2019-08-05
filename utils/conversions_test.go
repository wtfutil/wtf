package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
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

	Equal(t, expected, MapToStrs(source))
}

func Test_ToInts(t *testing.T) {
	expected := []int{1, 2, 3}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	Equal(t, expected, ToInts(source))
}

func Test_ToStrs(t *testing.T) {
	expected := []string{"cat", "dog", "rat"}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	Equal(t, expected, ToStrs(source))
}
