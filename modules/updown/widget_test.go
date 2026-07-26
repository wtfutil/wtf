package updown

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeSincePing(t *testing.T) {
	tests := []struct {
		name     string
		ts       time.Time
		wantMin  time.Duration
		wantFmt  bool // just check it's non-empty and formatted
	}{
		{
			name:    "recent timestamp",
			ts:      time.Now().Add(-5 * time.Second),
			wantMin: 5 * time.Second,
			wantFmt: true,
		},
		{
			name:    "older timestamp",
			ts:      time.Now().Add(-2 * time.Minute),
			wantMin: 2 * time.Minute,
			wantFmt: true,
		},
		{
			name:    "zero time",
			ts:      time.Time{},
			wantFmt: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timeSincePing(tt.ts)
			if result == "" {
				t.Error("expected non-empty string")
			}
		})
	}
}

func TestMakeURL(t *testing.T) {
	tests := []struct {
		name    string
		baseurl string
		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "valid URL with path",
			baseurl: "https://updown.io",
			path:    "/api/checks",
			want:    "https://updown.io/api/checks",
		},
		{
			name:    "valid URL with different path",
			baseurl: "http://localhost:8080",
			path:    "/health",
			want:    "http://localhost:8080/health",
		},
		{
			name:    "invalid base URL",
			baseurl: "://invalid",
			path:    "/test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeURL(tt.baseurl, tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("makeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterChecks(t *testing.T) {
	checks := []Check{
		{Token: "abc", Alias: "Site A"},
		{Token: "def", Alias: "Site B"},
		{Token: "ghi", Alias: "Site C"},
		{Token: "jkl", Alias: "Site D"},
	}

	tests := []struct {
		name     string
		tokenSet map[string]struct{}
		wantLen  int
		wantAlis []string
	}{
		{
			name:     "filter to single token",
			tokenSet: map[string]struct{}{"abc": {}},
			wantLen:  1,
			wantAlis: []string{"Site A"},
		},
		{
			name:     "filter to multiple tokens",
			tokenSet: map[string]struct{}{"abc": {}, "ghi": {}},
			wantLen:  2,
			wantAlis: []string{"Site A", "Site C"},
		},
		{
			name:     "no matching tokens",
			tokenSet: map[string]struct{}{"zzz": {}},
			wantLen:  0,
			wantAlis: []string{},
		},
		{
			name:     "all tokens match",
			tokenSet: map[string]struct{}{"abc": {}, "def": {}, "ghi": {}, "jkl": {}},
			wantLen:  4,
			wantAlis: []string{"Site A", "Site B", "Site C", "Site D"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid mutation across tests
			input := make([]Check, len(checks))
			copy(input, checks)

			result := filterChecks(input, tt.tokenSet)
			if len(result) != tt.wantLen {
				t.Errorf("filterChecks() returned %d checks, want %d", len(result), tt.wantLen)
				return
			}
			for i, alias := range tt.wantAlis {
				if result[i].Alias != alias {
					t.Errorf("filterChecks()[%d].Alias = %v, want %v", i, result[i].Alias, alias)
				}
			}
		})
	}
}

func TestFilterChecks_EmptyInput(t *testing.T) {
	result := filterChecks([]Check{}, map[string]struct{}{"abc": {}})
	if len(result) != 0 {
		t.Errorf("filterChecks() with empty input returned %d checks, want 0", len(result))
	}
}

func TestContentFrom(t *testing.T) {
	widget := &Widget{}

	tests := []struct {
		name       string
		checks     []Check
		wantPrefix string
	}{
		{
			name: "up and enabled check",
			checks: []Check{
				{Alias: "MyApp", Down: false, Enabled: true, Uptime: 99.95, LastCheckAt: time.Now()},
			},
			wantPrefix: "[green] + ",
		},
		{
			name: "down check",
			checks: []Check{
				{Alias: "MyApp", Down: true, Enabled: true, Uptime: 80.00, LastCheckAt: time.Now()},
			},
			wantPrefix: "[red] - ",
		},
		{
			name: "disabled check",
			checks: []Check{
				{Alias: "MyApp", Down: false, Enabled: false, Uptime: 99.00, LastCheckAt: time.Now()},
			},
			wantPrefix: "[yellow] ~ ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := widget.contentFrom(tt.checks)
			if len(result) == 0 {
				t.Fatal("contentFrom() returned empty string")
			}
			if result[:len(tt.wantPrefix)] != tt.wantPrefix {
				t.Errorf("contentFrom() prefix = %q, want %q", result[:len(tt.wantPrefix)], tt.wantPrefix)
			}
			// Should contain alias
			if !containsStr(result, tt.checks[0].Alias) {
				t.Errorf("contentFrom() does not contain alias %q", tt.checks[0].Alias)
			}
		})
	}
}

func TestContentFrom_MultipleChecks(t *testing.T) {
	widget := &Widget{}
	checks := []Check{
		{Alias: "App1", Down: false, Enabled: true, Uptime: 99.9, LastCheckAt: time.Now()},
		{Alias: "App2", Down: true, Enabled: true, Uptime: 50.0, LastCheckAt: time.Now()},
		{Alias: "App3", Down: false, Enabled: false, Uptime: 88.5, LastCheckAt: time.Now()},
	}

	result := widget.contentFrom(checks)
	if !containsStr(result, "App1") || !containsStr(result, "App2") || !containsStr(result, "App3") {
		t.Error("contentFrom() should contain all check aliases")
	}
	if !containsStr(result, "[green]") || !containsStr(result, "[red]") || !containsStr(result, "[yellow]") {
		t.Error("contentFrom() should contain all status colors")
	}
}

func TestContent_WithError(t *testing.T) {
	widget := &Widget{
		err: fmt.Errorf("connection refused"),
	}

	title, body, wrap := widget.content()
	if title != "Updown (0/0)" {
		t.Errorf("content() title = %q, want %q", title, "Updown (0/0)")
	}
	if body != "connection refused" {
		t.Errorf("content() body = %q, want %q", body, "connection refused")
	}
	if !wrap {
		t.Error("content() wrap should be true on error")
	}
}

func TestContent_NilChecks(t *testing.T) {
	widget := &Widget{
		checks: nil,
	}

	title, body, wrap := widget.content()
	if title != "Updown (0/0)" {
		t.Errorf("content() title = %q, want %q", title, "Updown (0/0)")
	}
	if body != "No checks to display" {
		t.Errorf("content() body = %q, want %q", body, "No checks to display")
	}
	if wrap {
		t.Error("content() wrap should be false for nil checks")
	}
}

func TestContent_WithChecks(t *testing.T) {
	widget := &Widget{
		checks: []Check{
			{Alias: "Up1", Down: false, Enabled: true, Uptime: 99.9, LastCheckAt: time.Now()},
			{Alias: "Up2", Down: false, Enabled: true, Uptime: 99.5, LastCheckAt: time.Now()},
			{Alias: "Down1", Down: true, Enabled: true, Uptime: 50.0, LastCheckAt: time.Now()},
		},
	}

	title, body, wrap := widget.content()
	if title != "Updown (2/3)" {
		t.Errorf("content() title = %q, want %q", title, "Updown (2/3)")
	}
	if body == "" {
		t.Error("content() body should not be empty")
	}
	if wrap {
		t.Error("content() wrap should be false for valid checks")
	}
}

func TestGetExistingChecks_Success(t *testing.T) {
	checks := []Check{
		{Token: "tok1", Alias: "Test Site", URL: "https://example.com", Down: false, Enabled: true, Uptime: 99.99},
		{Token: "tok2", Alias: "Other Site", URL: "https://other.com", Down: true, Enabled: true, Uptime: 80.0},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/checks" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("X-API-KEY") != "test-api-key" {
			t.Errorf("unexpected API key: %s", r.Header.Get("X-API-KEY"))
		}
		if r.Header.Get("User-Agent") != userAgent {
			t.Errorf("unexpected User-Agent: %s", r.Header.Get("User-Agent"))
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(checks)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "test-api-key"},
		tokenSet: map[string]struct{}{},
	}

	// Override the API base to use our test server
	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("getExistingChecks() error = %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("getExistingChecks() returned %d checks, want 2", len(result))
	}
	if result[0].Alias != "Test Site" {
		t.Errorf("result[0].Alias = %q, want %q", result[0].Alias, "Test Site")
	}
	if result[1].Down != true {
		t.Errorf("result[1].Down = %v, want true", result[1].Down)
	}
}

func TestGetExistingChecks_WithTokenFilter(t *testing.T) {
	checks := []Check{
		{Token: "tok1", Alias: "Wanted"},
		{Token: "tok2", Alias: "Unwanted"},
		{Token: "tok3", Alias: "Also Wanted"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(checks)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "key"},
		tokenSet: map[string]struct{}{"tok1": {}, "tok3": {}},
	}

	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("getExistingChecks() error = %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("got %d checks, want 2", len(result))
	}
	if result[0].Alias != "Wanted" || result[1].Alias != "Also Wanted" {
		t.Errorf("unexpected filtered results: %+v", result)
	}
}

func TestGetExistingChecks_Non200Status(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "bad-key"},
		tokenSet: map[string]struct{}{},
	}

	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("getExistingChecks() expected error for 401 response")
	}
	if !containsStr(err.Error(), "401") {
		t.Errorf("error should mention 401, got: %v", err)
	}
}

func TestGetExistingChecks_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`not valid json`))
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "key"},
		tokenSet: map[string]struct{}{},
	}

	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("getExistingChecks() expected error for invalid JSON")
	}
}

func TestGetExistingChecks_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "key"},
		tokenSet: map[string]struct{}{},
	}

	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("getExistingChecks() expected error for 500 response")
	}
	if !containsStr(err.Error(), "500") {
		t.Errorf("error should mention 500, got: %v", err)
	}
}

func TestGetExistingChecks_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "key"},
		tokenSet: map[string]struct{}{},
	}

	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("getExistingChecks() error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("got %d checks, want 0", len(result))
	}
}

func TestCheckJSONParsing(t *testing.T) {
	raw := `[{"token":"abc","url":"https://example.com","alias":"Example","last_status":200,"uptime":99.98,"down":false,"enabled":true,"period":30}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(raw))
	}))
	defer server.Close()

	widget := &Widget{
		settings: &Settings{apiKey: "key"},
		tokenSet: map[string]struct{}{},
	}

	origBase := apiURLBase
	setAPIURLBase(server.URL)
	defer setAPIURLBase(origBase)

	result, err := widget.getExistingChecks()
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("got %d checks, want 1", len(result))
	}
	c := result[0]
	if c.Token != "abc" {
		t.Errorf("Token = %q, want %q", c.Token, "abc")
	}
	if c.URL != "https://example.com" {
		t.Errorf("URL = %q, want %q", c.URL, "https://example.com")
	}
	if c.LastStatus != 200 {
		t.Errorf("LastStatus = %d, want 200", c.LastStatus)
	}
	if c.Uptime != 99.98 {
		t.Errorf("Uptime = %f, want 99.98", c.Uptime)
	}
	if c.Down != false {
		t.Error("Down should be false")
	}
	if c.Enabled != true {
		t.Error("Enabled should be true")
	}
	if c.Period != 30 {
		t.Errorf("Period = %d, want 30", c.Period)
	}
}

func TestGetExistingChecks_ConnectionError(t *testing.T) {
	widget := &Widget{
		settings: &Settings{apiKey: "key"},
		tokenSet: map[string]struct{}{},
	}

	origBase := apiURLBase
	setAPIURLBase("http://127.0.0.1:1") // port 1 should refuse connections
	defer setAPIURLBase(origBase)

	_, err := widget.getExistingChecks()
	if err == nil {
		t.Fatal("getExistingChecks() expected error for connection failure")
	}
}

// helper
func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && stringContains(s, substr))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
