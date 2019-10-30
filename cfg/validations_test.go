package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	vals = NewValidations()
)

func Test_intValueFor(t *testing.T) {
	vals.append("left", newPositionValidation("left", 3, nil))

	tests := []struct {
		name     string
		key      string
		expected int
	}{
		{
			name:     "with valid key",
			key:      "left",
			expected: 3,
		},
		{
			name:     "with invalid key",
			key:      "cat",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, vals.intValueFor(tt.key))
		})
	}
}
