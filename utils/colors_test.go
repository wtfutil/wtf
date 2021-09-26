package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ColorizePercent(t *testing.T) {
	tests := []struct {
		name     string
		percent  float64
		expected string
	}{
		{
			name:     "with high percent",
			percent:  70,
			expected: "[green]70[white]",
		},
		{
			name:     "with medium percent",
			percent:  35,
			expected: "[yellow]35[white]",
		},
		{
			name:     "with low percent",
			percent:  1,
			expected: "[red]1[white]",
		},
		{
			name:     "with negative percent",
			percent:  -5,
			expected: "[grey]-5[white]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ColorizePercent(tt.percent)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
