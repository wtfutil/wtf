package tennis

import (
	"net/http"
	"testing"
)

func TestShiftStatus(t *testing.T) {
	tests := []struct {
		name    string
		current string
		offset  int
		want    string
	}{
		{"forward", "live", 1, "upcoming"},
		{"forward wraps", "upcoming", 1, "live"},
		{"backward", "upcoming", -1, "live"},
		{"backward wraps", "live", -1, "upcoming"},
		{"unknown status starts at the beginning", "bogus", 1, "upcoming"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftStatus(tt.current, tt.offset); got != tt.want {
				t.Errorf("shiftStatus(%q, %d) = %q, want %q", tt.current, tt.offset, got, tt.want)
			}
		})
	}
}

func TestStatusKeysRequeryTheAPI(t *testing.T) {
	var seen []string

	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.URL.Query().Get("status"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": []}`))
	})
	defer srv.Close()

	widget.nextStatus()
	widget.nextStatus()
	widget.prevStatus()

	want := []string{"upcoming", "live", "upcoming"}
	if len(seen) != len(want) {
		t.Fatalf("expected %d requests, got %d (%v)", len(want), len(seen), seen)
	}
	for i, status := range want {
		if seen[i] != status {
			t.Errorf("request %d: got status %q, want %q", i, seen[i], status)
		}
	}

	if widget.settings.status != "upcoming" {
		t.Errorf("expected settings.status 'upcoming', got %q", widget.settings.status)
	}
}
