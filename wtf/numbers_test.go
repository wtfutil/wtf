package wtf

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Round(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "negative",
			input:    -3,
			expected: -3,
		},
		{
			name:     "integer",
			input:    3,
			expected: 3,
		},
		{
			name:     "float down",
			input:    3.123456,
			expected: 3,
		},
		{
			name:     "float up",
			input:    3.998786,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Round(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_TruncateFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		precision int
		expected  float64
	}{
		{
			name:      "negative precision",
			input:     23.234567,
			precision: -2,
			expected:  0,
		},
		{
			name:      "zero precision",
			input:     23.234567,
			precision: 0,
			expected:  23,
		},
		{
			name:      "positive precision",
			input:     23.234567,
			precision: 2,
			expected:  23.23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := TruncateFloat64(tt.input, tt.precision)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
