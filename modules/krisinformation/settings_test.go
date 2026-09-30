package krisinformation

import (
	"testing"
)

func TestDefaultSettings(t *testing.T) {
	tests := []struct {
		name     string
		constant interface{}
		expected interface{}
	}{
		{"defaultFocusable", defaultFocusable, false},
		{"defaultTitle", defaultTitle, "Krisinformation"},
		{"defaultRadius", defaultRadius, -1},
		{"defaultCountry", defaultCountry, true},
		{"defaultCounty", defaultCounty, ""},
		{"defaultMaxItems", defaultMaxItems, -1},
		{"defaultMaxAge", defaultMaxAge, 720},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.constant != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, tc.constant)
			}
		})
	}
}
