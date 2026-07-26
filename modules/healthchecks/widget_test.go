package healthchecks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

func TestTimeSincePing(t *testing.T) {
	tests := []struct {
		name     string
		ts       time.Time
		wantPos  bool // result should be a positive duration string
		wantZero bool // result should be "0s" or similar
	}{
		{
			name:    "recent ping",
			ts:      time.Now().Add(-5 * time.Second),
			wantPos: true,
		},
		{
			name:    "old ping",
			ts:      time.Now().Add(-2 * time.Hour),
			wantPos: true,
		},
		{
			name:     "zero time",
			ts:       time.Time{},
			wantPos:  true,
			wantZero: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timeSincePing(tt.ts)
			if result == "" {
				t.Error("expected non-empty string")
			}
			if tt.wantPos && !strings.ContainsAny(result, "0123456789") {
				t.Errorf("expected numeric duration, got %q", result)
			}
		})
	}
}

func TestMakeURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		path    string
		tags    []string
		want    string
		wantErr bool
	}{
		{
			name:    "simple URL no tags",
			baseURL: "https://healthchecks.io",
			path:    "/api/v1/checks/",
			tags:    nil,
			want:    "https://healthchecks.io/api/v1/checks/",
		},
		{
			name:    "URL with single tag",
			baseURL: "https://healthchecks.io",
			path:    "/api/v1/checks/",
			tags:    []string{"prod"},
			want:    "https://healthchecks.io/api/v1/checks/?tag=prod",
		},
		{
			name:    "URL with multiple tags",
			baseURL: "https://healthchecks.io",
			path:    "/api/v1/checks/",
			tags:    []string{"prod", "web"},
			want:    "https://healthchecks.io/api/v1/checks/?tag=prod&tag=web",
		},
		{
			name:    "empty tags slice",
			baseURL: "https://healthchecks.io",
			path:    "/api/v1/checks/",
			tags:    []string{},
			want:    "https://healthchecks.io/api/v1/checks/",
		},
		{
			name:    "invalid base URL",
			baseURL: "://invalid",
			path:    "/api/v1/checks/",
			tags:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeURL(tt.baseURL, tt.path, tt.tags)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestContentFrom(t *testing.T) {
	tests := []struct {
		name   string
		checks []Checks
		want   []string // substrings expected in output
	}{
		{
			name: "up status uses green prefix",
			checks: []Checks{
				{Name: "DB Backup", Status: "up", NPings: 42, LastPing: time.Now().Add(-1 * time.Minute)},
			},
			want: []string{"[green] + ", "DB Backup", "42"},
		},
		{
			name: "down status uses red prefix",
			checks: []Checks{
				{Name: "Web Server", Status: "down", NPings: 10, LastPing: time.Now().Add(-5 * time.Minute)},
			},
			want: []string{"[red] - ", "Web Server", "10"},
		},
		{
			name: "paused status uses lightgray prefix",
			checks: []Checks{
				{Name: "Cron Job", Status: "paused", NPings: 5, LastPing: time.Now().Add(-1 * time.Hour)},
			},
			want: []string{"[lightgray] × ", "Cron Job", "5"},
		},
		{
			name: "new status uses lightgray prefix",
			checks: []Checks{
				{Name: "New Check", Status: "new", NPings: 0, LastPing: time.Now()},
			},
			want: []string{"[lightgray] × ", "New Check", "0"},
		},
		{
			name: "grace status uses yellow prefix",
			checks: []Checks{
				{Name: "Grace Check", Status: "grace", NPings: 3, LastPing: time.Now().Add(-30 * time.Second)},
			},
			want: []string{"[yellow] ~ ", "Grace Check", "3"},
		},
		{
			name: "unknown status uses yellow prefix",
			checks: []Checks{
				{Name: "Unknown", Status: "something_else", NPings: 1, LastPing: time.Now()},
			},
			want: []string{"[yellow] ~ ", "Unknown"},
		},
		{
			name: "multiple checks",
			checks: []Checks{
				{Name: "Check1", Status: "up", NPings: 1, LastPing: time.Now()},
				{Name: "Check2", Status: "down", NPings: 2, LastPing: time.Now()},
			},
			want: []string{"Check1", "Check2", "[green]", "[red]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{}
			got := widget.contentFrom(tt.checks)
			for _, substr := range tt.want {
				if !strings.Contains(got, substr) {
					t.Errorf("output %q does not contain %q", got, substr)
				}
			}
		})
	}
}

func TestGetExistingChecks_Success(t *testing.T) {
	checks := Health{
		Checks: []Checks{
			{
				Name:    "Test Check",
				Status:  "up",
				NPings:  100,
				Grace:   60,
				PingURL: "https://hc-ping.com/abc123",
			},
			{
				Name:   "Another Check",
				Status: "down",
				NPings: 50,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != "test-api-key" {
			t.Errorf("expected X-Api-Key header 'test-api-key', got %q", r.Header.Get("X-Api-Key"))
		}
		if r.Header.Get("User-Agent") != "WTFUtil" {
			t.Errorf("expected User-Agent 'WTFUtil', got %q", r.Header.Get("User-Agent"))
		}
		if r.URL.Path != "/api/v1/checks/" {
			t.Errorf("expected path /api/v1/checks/, got %q", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(checks)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{
			apiKey: "test-api-key",
			apiURL: server.URL,
			tags:   nil,
		},
	}

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 checks, got %d", len(result))
	}
	if result[0].Name != "Test Check" {
		t.Errorf("expected first check name 'Test Check', got %q", result[0].Name)
	}
	if result[0].Status != "up" {
		t.Errorf("expected first check status 'up', got %q", result[0].Status)
	}
	if result[1].Name != "Another Check" {
		t.Errorf("expected second check name 'Another Check', got %q", result[1].Name)
	}
	if result[1].NPings != 50 {
		t.Errorf("expected second check NPings 50, got %d", result[1].NPings)
	}
}

func TestGetExistingChecks_WithTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tags := r.URL.Query()["tag"]
		if len(tags) != 2 || tags[0] != "prod" || tags[1] != "web" {
			t.Errorf("expected tags [prod, web], got %v", tags)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"checks":[]}`)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{
			apiKey: "key",
			apiURL: server.URL,
			tags:   []string{"prod", "web"},
		},
	}

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected 0 checks, got %d", len(result))
	}
}

func TestGetExistingChecks_Non200Status(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		wantErr    string
	}{
		{"401 unauthorized", 401, "401 Unauthorized"},
		{"403 forbidden", 403, "403 Forbidden"},
		{"500 internal server error", 500, "500 Internal Server Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "error", tt.statusCode)
			}))
			defer server.Close()

			widget := &Widget{
				settings: &Settings{
					apiKey: "bad-key",
					apiURL: server.URL,
					tags:   nil,
				},
			}

			_, err := widget.getExistingChecks()
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("expected error containing %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}

func TestGetExistingChecks_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"checks": invalid}`)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{
			apiKey: "key",
			apiURL: server.URL,
			tags:   nil,
		},
	}

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestGetExistingChecks_ConnectionError(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			apiKey: "key",
			apiURL: "http://127.0.0.1:1", // should fail to connect
			tags:   nil,
		},
	}

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestGetExistingChecks_InvalidBaseURL(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			apiKey: "key",
			apiURL: "://invalid-url",
			tags:   nil,
		},
	}

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("expected error for invalid base URL, got nil")
	}
}

func TestContentFrom_EmptyChecks(t *testing.T) {
	widget := &Widget{}
	got := widget.contentFrom([]Checks{})
	if got != "" {
		t.Errorf("expected empty string for no checks, got %q", got)
	}
}

func newTestWidget() *Widget {
	common := &cfg.Common{
		Module: cfg.Module{Name: "healthchecks"},
		Title:  "Healthchecks",
	}
	base := view.NewBase(nil, nil, nil, common)
	w := &Widget{}
	w.ScrollableWidget.TextWidget.Base = base
	return w
}

func TestContent_WithError(t *testing.T) {
	widget := newTestWidget()
	widget.err = fmt.Errorf("something went wrong")

	_, body, wrap := widget.content()
	if body != "something went wrong" {
		t.Errorf("expected error message in body, got %q", body)
	}
	if !wrap {
		t.Error("expected wrap=true when error is set")
	}
}

func TestContent_NilChecks(t *testing.T) {
	widget := newTestWidget()

	_, body, wrap := widget.content()
	if body != "No checks to display" {
		t.Errorf("expected 'No checks to display', got %q", body)
	}
	if wrap {
		t.Error("expected wrap=false when no checks")
	}
}

func TestContent_WithChecks(t *testing.T) {
	widget := newTestWidget()
	widget.checks = []Checks{
		{Name: "Up1", Status: "up", NPings: 1, LastPing: time.Now()},
		{Name: "Up2", Status: "up", NPings: 2, LastPing: time.Now()},
		{Name: "Down1", Status: "down", NPings: 3, LastPing: time.Now()},
	}

	title, body, wrap := widget.content()
	if !strings.Contains(title, "2/3") {
		t.Errorf("expected title to contain '2/3' (up/total), got %q", title)
	}
	if !strings.Contains(body, "Up1") || !strings.Contains(body, "Down1") {
		t.Errorf("expected body to contain check names, got %q", body)
	}
	if wrap {
		t.Error("expected wrap=false when checks are present")
	}
}

func TestGetExistingChecks_FullResponse(t *testing.T) {
	// Test parsing of all JSON fields
	responseJSON := `{
		"checks": [{
			"name": "Full Check",
			"tags": "prod web",
			"desc": "A description",
			"grace": 120,
			"n_pings": 500,
			"status": "up",
			"last_ping": "2025-01-15T10:30:00Z",
			"next_ping": "2025-01-15T11:30:00Z",
			"manual_resume": true,
			"methods": "POST",
			"ping_url": "https://hc-ping.com/abc",
			"update_url": "https://healthchecks.io/api/v1/checks/abc",
			"pause_url": "https://healthchecks.io/api/v1/checks/abc/pause",
			"channels": "*",
			"timeout": 3600,
			"schedule": "0 * * * *",
			"tz": "UTC"
		}]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, responseJSON)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{
			apiKey: "key",
			apiURL: server.URL,
			tags:   nil,
		},
	}

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 check, got %d", len(result))
	}

	check := result[0]
	if check.Name != "Full Check" {
		t.Errorf("Name = %q, want 'Full Check'", check.Name)
	}
	if check.Tags != "prod web" {
		t.Errorf("Tags = %q, want 'prod web'", check.Tags)
	}
	if check.Desc != "A description" {
		t.Errorf("Desc = %q, want 'A description'", check.Desc)
	}
	if check.Grace != 120 {
		t.Errorf("Grace = %d, want 120", check.Grace)
	}
	if check.NPings != 500 {
		t.Errorf("NPings = %d, want 500", check.NPings)
	}
	if check.Status != "up" {
		t.Errorf("Status = %q, want 'up'", check.Status)
	}
	if check.ManualResume != true {
		t.Error("ManualResume = false, want true")
	}
	if check.Methods != "POST" {
		t.Errorf("Methods = %q, want 'POST'", check.Methods)
	}
	if check.Timeout != 3600 {
		t.Errorf("Timeout = %d, want 3600", check.Timeout)
	}
	if check.Schedule != "0 * * * *" {
		t.Errorf("Schedule = %q, want '0 * * * *'", check.Schedule)
	}
	if check.Tz != "UTC" {
		t.Errorf("Tz = %q, want 'UTC'", check.Tz)
	}
}
