package system

import (
	"strings"
	"testing"

	"gotest.tools/assert"
)

func Test_prettyDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected string
		wantErr  bool
	}{
		{
			name:     "valid timestamp",
			date:     "2021-03-05T14:30:00-0700",
			expected: "Mar  5, 14:30",
		},
		{
			name:     "valid timestamp, single digit day",
			date:     "2020-01-02T09:05:00+0000",
			expected: "Jan  2, 09:05",
		},
		{
			name:     "valid timestamp, double digit day",
			date:     "2020-12-25T23:59:00+0000",
			expected: "Dec 25, 23:59",
		},
		{
			name:    "empty date",
			date:    "",
			wantErr: true,
		},
		{
			name:    "malformed date",
			date:    "not-a-date",
			wantErr: true,
		},
		{
			name:    "wrong format",
			date:    "2021-03-05",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				Date: tt.date,
			}

			actual := widget.prettyDate()

			if tt.wantErr {
				// On error, prettyDate returns the error message itself,
				// which will not match the "Jan _2, 15:04" pretty format.
				assert.Assert(t, actual != "")
				assert.Assert(t, !strings.Contains(actual, ","))
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
