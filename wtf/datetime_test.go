package wtf

import (
	"testing"
	"time"
)

func Test_IsToday(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected bool
	}{
		{
			name:     "when yesterday",
			date:     time.Now().Local().AddDate(0, 0, -1),
			expected: false,
		},
		{
			name:     "when today",
			date:     time.Now().Local(),
			expected: true,
		},
		{
			name:     "when tomorrow",
			date:     time.Now().Local().AddDate(0, 0, +1),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsToday(tt.date)

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_PrettyDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected string
	}{
		{
			name:     "with empty date",
			date:     "",
			expected: "",
		},
		{
			name:     "with invalid date",
			date:     "10-21-1999",
			expected: "10-21-1999",
		},
		{
			name:     "with valid date",
			date:     "1999-10-21",
			expected: "Oct 21, 1999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PrettyDate(tt.date)

			if tt.expected != actual {
				t.Errorf("\nexpected: %s\n     got: %s", tt.expected, actual)
			}
		})
	}
}
func Test_UnixTime(t *testing.T) {
	tests := []struct {
		name     string
		unixVal  int64
		expected string
	}{
		{
			name:     "with 0 time",
			unixVal:  0,
			expected: "1970-01-01 00:00:00 +0000 UTC",
		},
		{
			name:     "with explicit time",
			unixVal:  1564883266,
			expected: "2019-08-04 01:47:46 +0000 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := UnixTime(tt.unixVal).UTC()

			if tt.expected != actual.String() {
				t.Errorf("\nexpected: %s\n     got: %s", tt.expected, actual)
			}
		})
	}
}
