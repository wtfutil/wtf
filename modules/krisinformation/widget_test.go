package krisinformation

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

// newTestWidget creates a minimal Widget suitable for testing content()
func newTestWidget(client *Client, settings *Settings) *Widget {
	app := tview.NewApplication()
	common := &cfg.Common{
		Module: cfg.Module{Name: "krisinformation"},
		Title:  "",
	}
	settings.common = common

	redrawChan := make(chan bool, 1)
	return &Widget{
		TextWidget: view.NewTextWidget(app, redrawChan, nil, common),
		app:        app,
		settings:   settings,
		client:     client,
	}
}

func TestContent_AgeFiltering(t *testing.T) {
	now := time.Now()
	old := now.Add(-800 * time.Hour)  // older than default 720h
	recent := now.Add(-100 * time.Hour) // within default 720h

	data := Krisinformation{
		{
			Headline:    "Old Alert",
			PushMessage: "old",
			SenderName:  "MSB",
			Updated:     old,
			Published:   old,
			Area: []struct {
				Type                string      `json:"Type"`
				Description         string      `json:"Description"`
				Coordinate          string      `json:"Coordinate"`
				GeometryInformation interface{} `json:"GeometryInformation"`
			}{{Type: "Country", Description: "Sverige"}},
		},
		{
			Headline:    "Recent Alert",
			PushMessage: "recent",
			SenderName:  "MSB",
			Updated:     recent,
			Published:   recent,
			Area: []struct {
				Type                string      `json:"Type"`
				Description         string      `json:"Description"`
				Coordinate          string      `json:"Coordinate"`
				GeometryInformation interface{} `json:"GeometryInformation"`
			}{{Type: "Country", Description: "Sverige"}},
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}))
	defer srv.Close()

	tests := []struct {
		name        string
		maxage      int
		wantInBody  string
		wantMissing string
	}{
		{
			name:        "default maxage filters old items",
			maxage:      720,
			wantInBody:  "Recent Alert",
			wantMissing: "Old Alert",
		},
		{
			name:        "maxage -1 shows all items",
			maxage:      -1,
			wantInBody:  "Old Alert",
			wantMissing: "",
		},
		{
			name:        "very small maxage filters everything",
			maxage:      1,
			wantInBody:  "",
			wantMissing: "Recent Alert",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient(0, 0, -1, "", true)
			client.apiURL = srv.URL

			w := newTestWidget(client, &Settings{
				maxage:   tc.maxage,
				maxitems: -1,
				country:  true,
			})

			_, body, _ := w.content()

			if tc.wantInBody != "" && !strings.Contains(body, tc.wantInBody) {
				t.Errorf("expected body to contain %q, got: %s", tc.wantInBody, body)
			}
			if tc.wantMissing != "" && strings.Contains(body, tc.wantMissing) {
				t.Errorf("expected body NOT to contain %q, got: %s", tc.wantMissing, body)
			}
		})
	}
}

func TestContent_MaxItemsLimiting(t *testing.T) {
	now := time.Now()
	data := Krisinformation{
		{
			Headline: "Alert 1", PushMessage: "msg1", SenderName: "MSB",
			Updated: now, Published: now,
			Area: []struct {
				Type                string      `json:"Type"`
				Description         string      `json:"Description"`
				Coordinate          string      `json:"Coordinate"`
				GeometryInformation interface{} `json:"GeometryInformation"`
			}{{Type: "Country", Description: "Sverige"}},
		},
		{
			Headline: "Alert 2", PushMessage: "msg2", SenderName: "MSB",
			Updated: now, Published: now,
			Area: []struct {
				Type                string      `json:"Type"`
				Description         string      `json:"Description"`
				Coordinate          string      `json:"Coordinate"`
				GeometryInformation interface{} `json:"GeometryInformation"`
			}{{Type: "Country", Description: "Sverige"}},
		},
		{
			Headline: "Alert 3", PushMessage: "msg3", SenderName: "MSB",
			Updated: now, Published: now,
			Area: []struct {
				Type                string      `json:"Type"`
				Description         string      `json:"Description"`
				Coordinate          string      `json:"Coordinate"`
				GeometryInformation interface{} `json:"GeometryInformation"`
			}{{Type: "Country", Description: "Sverige"}},
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}))
	defer srv.Close()

	tests := []struct {
		name     string
		maxitems int
		wantN    int
	}{
		{"maxitems -1 shows all", -1, 3},
		{"maxitems 2 limits to 2", 2, 2},
		{"maxitems 1 limits to 1", 1, 1},
		{"maxitems 10 shows all when fewer exist", 10, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient(0, 0, -1, "", true)
			client.apiURL = srv.URL

			w := newTestWidget(client, &Settings{
				maxage:   -1,
				maxitems: tc.maxitems,
				country:  true,
			})

			_, body, _ := w.content()
			lines := strings.Split(strings.TrimSpace(body), "\n")
			if body == "" {
				lines = []string{}
			}
			if len(lines) != tc.wantN {
				t.Errorf("expected %d items, got %d: %q", tc.wantN, len(lines), body)
			}
		})
	}
}

func TestContent_Title(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	}))
	defer srv.Close()

	client := NewClient(0, 0, -1, "", true)
	client.apiURL = srv.URL

	w := newTestWidget(client, &Settings{
		maxage:   -1,
		maxitems: -1,
		country:  true,
	})

	title, _, _ := w.content()
	if title != defaultTitle {
		t.Errorf("expected default title %q, got %q", defaultTitle, title)
	}
}

func TestContent_FormatOutput(t *testing.T) {
	now := time.Now()
	data := Krisinformation{
		{
			Headline: "Test Headline", PushMessage: "push", SenderName: "Sender",
			Updated: now, Published: now,
			Area: []struct {
				Type                string      `json:"Type"`
				Description         string      `json:"Description"`
				Coordinate          string      `json:"Coordinate"`
				GeometryInformation interface{} `json:"GeometryInformation"`
			}{{Type: "Country", Description: "Sverige"}},
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}))
	defer srv.Close()

	client := NewClient(0, 0, -1, "", true)
	client.apiURL = srv.URL

	w := newTestWidget(client, &Settings{
		maxage:   -1,
		maxitems: -1,
		country:  true,
	})

	_, body, wrap := w.content()
	if !wrap {
		t.Error("expected wrap to be true")
	}
	expected := "- Test Headline\n"
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}

func TestContent_ErrorHandling(t *testing.T) {
	client := NewClient(0, 0, -1, "", true)
	client.apiURL = "http://127.0.0.1:1" // connection refused

	w := newTestWidget(client, &Settings{
		maxage:   -1,
		maxitems: -1,
		country:  true,
	})

	_, body, _ := w.content()
	// Even on error, content returns (it calls handleError, doesn't return early)
	if body != "" {
		t.Errorf("expected empty body on error, got %q", body)
	}
	if w.err == nil {
		t.Error("expected widget.err to be set")
	}
}
