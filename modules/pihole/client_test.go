package pihole

import (
	"errors"
	"testing"
)

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil error", nil, "unknown error"},
		{"error without token", errors.New("connection refused"), "connection refused"},
		{"error with token redacted", errors.New("request failed: auth=abc123XYZ"), "request failed: auth=<token>"},
		{"error with token in url query", errors.New("Get \"http://host/api.php?auth=secrettoken123&summary\": timeout"), "Get \"http://host/api.php?auth=<token>&summary\": timeout"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseError(tt.err)
			if got != tt.want {
				t.Errorf("parseError(%v) = %q, want %q", tt.err, got, tt.want)
			}
		})
	}
}
