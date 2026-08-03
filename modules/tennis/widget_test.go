package tennis

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

// createTestWidget builds a widget wired to an httptest server. No real API
// key is ever used.
func createTestWidget(apiKey string, handler http.HandlerFunc) (*Widget, *httptest.Server) {
	srv := httptest.NewServer(handler)

	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{
			Title:   "Tennis",
			Enabled: true,
		},
		apiKey:     apiKey,
		status:     defaultStatus,
		matchLimit: defaultMatchLimit,
	}

	widget := NewWidget(tviewApp, redrawChan, nil, settings)
	widget.client = NewClient(apiKey, srv.Client(), srv.URL)

	return widget, srv
}

func TestContent_NoAPIKey(t *testing.T) {
	widget, srv := createTestWidget("", func(w http.ResponseWriter, r *http.Request) {
		t.Error("server should not be called without an API key")
	})
	defer srv.Close()

	widget.Refresh()

	_, body, wrap := widget.content()
	if !wrap {
		t.Error("expected wrap=true for setup hint")
	}
	if !strings.Contains(body, "WTF_TENNIS_API_KEY") {
		t.Errorf("expected env var hint in body, got %q", body)
	}
	if !strings.Contains(body, FreeKeyURL) {
		t.Errorf("expected free key URL in body, got %q", body)
	}
}

func TestContent_Unauthorized(t *testing.T) {
	widget, srv := createTestWidget("bad-key", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
	})
	defer srv.Close()

	widget.Refresh()

	_, body, wrap := widget.content()
	if !wrap {
		t.Error("expected wrap=true for error content")
	}
	if !strings.Contains(body, "401") {
		t.Errorf("expected 401 in body, got %q", body)
	}
	if !strings.Contains(body, FreeKeyURL) {
		t.Errorf("expected free key URL in body, got %q", body)
	}
}

func TestContent_RateLimited(t *testing.T) {
	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	})
	defer srv.Close()

	widget.Refresh()

	_, body, _ := widget.content()
	if !strings.Contains(body, "429") {
		t.Errorf("expected 429 in body, got %q", body)
	}
	if !strings.Contains(body, "refreshInterval") {
		t.Errorf("expected refreshInterval hint in body, got %q", body)
	}
}

func TestContent_Empty(t *testing.T) {
	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": [], "meta": {}}`))
	})
	defer srv.Close()

	widget.Refresh()

	_, body, wrap := widget.content()
	if wrap {
		t.Error("expected wrap=false for empty state")
	}
	if body != "No live matches" {
		t.Errorf("expected 'No live matches', got %q", body)
	}
}

func TestContent_LiveMatches(t *testing.T) {
	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(liveFixture))
	})
	defer srv.Close()

	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error: %v", widget.err)
	}

	title, body, wrap := widget.content()
	if wrap {
		t.Error("expected wrap=false for match content")
	}
	if title != "Tennis (live)" {
		t.Errorf("unexpected title %q", title)
	}

	lines := strings.Split(body, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), body)
	}

	// Live match: score, green serving marker on p1's side, points, location
	want := "Sinner (1) 6-3 4-6 2-1[green]*[-] (40-AD) vs Alcaraz (2) • Tampere QF"
	if lines[0] != want {
		t.Errorf("live line mismatch:\n got %q\nwant %q", lines[0], want)
	}

	// Upcoming match (null score): no score block, scheduled time shown
	if !strings.Contains(lines[1], "Djokovic (7) vs Musetti (10)") {
		t.Errorf("expected upcoming players line, got %q", lines[1])
	}
	if !strings.Contains(lines[1], "2026-07-24 18:30:00Z") {
		t.Errorf("expected scheduled time, got %q", lines[1])
	}
}

func TestContent_MatchLimit(t *testing.T) {
	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(liveFixture))
	})
	defer srv.Close()

	widget.settings.matchLimit = 1
	widget.Refresh()

	if len(widget.matches) != 1 {
		t.Fatalf("expected 1 match after limit, got %d", len(widget.matches))
	}
}

func TestRefresh_Disabled(t *testing.T) {
	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {
		t.Error("server should not be called when widget is disabled")
	})
	defer srv.Close()

	widget.Disable()
	widget.Refresh()

	if widget.matches != nil {
		t.Error("expected nil matches when disabled")
	}
}

func TestWidgetTitle(t *testing.T) {
	tests := []struct {
		name   string
		tour   string
		status string
		want   string
	}{
		{"no tour", "", "live", "Tennis (live)"},
		{"with tour", "wta", "upcoming", "Tennis WTA (upcoming)"},
		{"completed", "atp", "completed", "Tennis ATP (completed)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {})
			defer srv.Close()

			widget.settings.tour = tt.tour
			widget.settings.status = tt.status

			got := widget.title()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			// tview parses square brackets in titles as style tags and
			// silently swallows them, so the title must not contain any.
			if strings.ContainsAny(got, "[]") {
				t.Errorf("title must not contain square brackets: %q", got)
			}
		})
	}
}

func TestConfigText(t *testing.T) {
	widget, srv := createTestWidget("key", func(w http.ResponseWriter, r *http.Request) {})
	defer srv.Close()

	if widget.ConfigText() == "" {
		t.Error("expected non-empty config text")
	}
}
