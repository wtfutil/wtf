package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_padSlice(t *testing.T) {
	tests := []struct {
		name    string
		padLeft bool
		values  []string
		want    []string
	}{
		{
			name:    "pad right (left-aligned) to longest value",
			padLeft: false,
			values:  []string{"a", "bb", "ccc"},
			want:    []string{"a  ", "bb ", "ccc"},
		},
		{
			name:    "pad left (right-aligned) to longest value",
			padLeft: true,
			values:  []string{"a", "bb", "ccc"},
			want:    []string{"  a", " bb", "ccc"},
		},
		{
			name:    "all equal length is a no-op",
			padLeft: false,
			values:  []string{"aa", "bb"},
			want:    []string{"aa", "bb"},
		},
		{
			name:    "empty slice",
			padLeft: false,
			values:  []string{},
			want:    []string{},
		},
		{
			name:    "single element",
			padLeft: true,
			values:  []string{"solo"},
			want:    []string{"solo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := append([]string{}, tt.values...)

			padSlice(tt.padLeft, values, func(i int) string {
				return values[i]
			}, func(i int, newVal string) {
				values[i] = newVal
			})

			assert.Equal(t, tt.want, values)
		})
	}
}
