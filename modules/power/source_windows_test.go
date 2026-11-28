//go:build windows

package power

import "testing"

func Test_mapBatteryStatusToSource(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		expected string
	}{
		{name: "status 1 is battery", status: "1", expected: "Battery Power"},
		{name: "status 2 is AC", status: "2", expected: "AC Power"},
		{name: "status 3 is AC", status: "3", expected: "AC Power"},
		{name: "status 4 is battery", status: "4", expected: "Battery Power"},
		{name: "status 5 is battery", status: "5", expected: "Battery Power"},
		{name: "status 6 is AC", status: "6", expected: "AC Power"},
		{name: "status 7 is AC", status: "7", expected: "AC Power"},
		{name: "status 8 is AC", status: "8", expected: "AC Power"},
		{name: "status 9 is AC", status: "9", expected: "AC Power"},
		{name: "empty defaults to AC", status: "", expected: "AC Power"},
		{name: "unknown code defaults to AC", status: "42", expected: "AC Power"},
		{name: "non-numeric defaults to AC", status: "abc", expected: "AC Power"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mapBatteryStatusToSource(tt.status)
			if actual != tt.expected {
				t.Errorf("mapBatteryStatusToSource(%q) = %q, want %q", tt.status, actual, tt.expected)
			}
		})
	}
}
