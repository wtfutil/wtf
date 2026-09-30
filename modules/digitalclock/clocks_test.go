package digitalclock

import "testing"

func TestIntStrConv(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"single digit", 5, "05"},
		{"zero", 0, "00"},
		{"double digit", 12, "12"},
		{"nine", 9, "09"},
		{"ten", 10, "10"},
		{"large number", 123, "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intStrConv(tt.input)
			if result != tt.expected {
				t.Errorf("intStrConv(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetColon(t *testing.T) {
	// getColon returns ":" or " " based on current second being even/odd
	result := getColon()
	if result != ":" && result != " " {
		t.Errorf("getColon() = %q, want \":\" or \" \"", result)
	}
}
