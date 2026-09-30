package uptimekuma

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name        string
		rawURL      string
		wantBase    string
		wantSlug    string
		wantErrMsg  string
	}{
		{
			name:     "valid URL",
			rawURL:   "https://uptime.example.com/status/overview",
			wantBase: "https://uptime.example.com",
			wantSlug: "overview",
		},
		{
			name:     "valid URL with trailing slash",
			rawURL:   "https://uptime.example.com/status/mypage/",
			wantBase: "https://uptime.example.com",
			wantSlug: "mypage",
		},
		{
			name:     "valid URL with port",
			rawURL:   "http://localhost:3001/status/test",
			wantBase: "http://localhost:3001",
			wantSlug: "test",
		},
		{
			name:       "empty URL",
			rawURL:     "",
			wantErrMsg: "URL is not defined",
		},
		{
			name:       "missing status prefix",
			rawURL:     "https://example.com/page/overview",
			wantErrMsg: "invalid status page URL format",
		},
		{
			name:       "only one path segment",
			rawURL:     "https://example.com/status",
			wantErrMsg: "invalid status page URL format",
		},
		{
			name:       "no path",
			rawURL:     "https://example.com",
			wantErrMsg: "invalid status page URL format",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			base, slug, err := parseURL(tc.rawURL)
			if tc.wantErrMsg != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErrMsg)
				}
				if !strings.Contains(err.Error(), tc.wantErrMsg) {
					t.Fatalf("expected error containing %q, got %q", tc.wantErrMsg, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if base != tc.wantBase {
				t.Errorf("baseURL: got %q, want %q", base, tc.wantBase)
			}
			if slug != tc.wantSlug {
				t.Errorf("slug: got %q, want %q", slug, tc.wantSlug)
			}
		})
	}
}

func TestFetchStatusData_Success(t *testing.T) {
	data := StatusPageData{
		Incident: &Incident{CreatedDate: "2023-10-27 10:30:00.123"},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/status-page/myslug" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}))
	defer srv.Close()

	widget := &Widget{}
	result, err := widget.fetchStatusData(srv.URL, "myslug")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Incident == nil {
		t.Fatal("expected incident, got nil")
	}
	if result.Incident.CreatedDate != "2023-10-27 10:30:00.123" {
		t.Errorf("unexpected createdDate: %s", result.Incident.CreatedDate)
	}
}

func TestFetchStatusData_NoIncident(t *testing.T) {
	data := StatusPageData{Incident: nil}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}))
	defer srv.Close()

	widget := &Widget{}
	result, err := widget.fetchStatusData(srv.URL, "slug")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Incident != nil {
		t.Errorf("expected nil incident, got %+v", result.Incident)
	}
}

func TestFetchStatusData_Non200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	widget := &Widget{}
	_, err := widget.fetchStatusData(srv.URL, "slug")
	if err == nil {
		t.Fatal("expected error for 404")
	}
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("expected error to contain '404', got %q", err.Error())
	}
}

func TestFetchStatusData_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json"))
	}))
	defer srv.Close()

	widget := &Widget{}
	_, err := widget.fetchStatusData(srv.URL, "slug")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestFetchStatusData_ConnectionError(t *testing.T) {
	widget := &Widget{}
	_, err := widget.fetchStatusData("http://127.0.0.1:1", "slug")
	if err == nil {
		t.Fatal("expected error for connection refused")
	}
}

func TestFetchHeartbeatData_Success(t *testing.T) {
	data := HeartbeatData{
		HeartbeatList: map[string][]*Heartbeat{
			"1": {{Status: 1}, {Status: 1}},
			"2": {{Status: 0}},
		},
		UptimeList: map[string]float64{
			"1": 0.999,
			"2": 0.95,
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/status-page/heartbeat/myslug" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(data)
	}))
	defer srv.Close()

	widget := &Widget{}
	result, err := widget.fetchHeartbeatData(srv.URL, "myslug")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.HeartbeatList) != 2 {
		t.Errorf("expected 2 heartbeat lists, got %d", len(result.HeartbeatList))
	}
	if len(result.UptimeList) != 2 {
		t.Errorf("expected 2 uptime entries, got %d", len(result.UptimeList))
	}
}

func TestFetchHeartbeatData_Non200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	widget := &Widget{}
	_, err := widget.fetchHeartbeatData(srv.URL, "slug")
	if err == nil {
		t.Fatal("expected error for 500")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected error to contain '500', got %q", err.Error())
	}
}

func TestFetchHeartbeatData_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{invalid"))
	}))
	defer srv.Close()

	widget := &Widget{}
	_, err := widget.fetchHeartbeatData(srv.URL, "slug")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestFetchHeartbeatData_ConnectionError(t *testing.T) {
	widget := &Widget{}
	_, err := widget.fetchHeartbeatData("http://127.0.0.1:1", "slug")
	if err == nil {
		t.Fatal("expected error for connection refused")
	}
}

func TestContent_Error(t *testing.T) {
	widget := &Widget{
		err: fmt.Errorf("something went wrong"),
	}
	got := widget.content()
	if !strings.Contains(got, "Error") {
		t.Errorf("expected 'Error' in output, got %q", got)
	}
	if !strings.Contains(got, "something went wrong") {
		t.Errorf("expected error message in output, got %q", got)
	}
}

func TestContent_Loading(t *testing.T) {
	widget := &Widget{}
	got := widget.content()
	if got != "Loading..." {
		t.Errorf("expected 'Loading...', got %q", got)
	}

	// Only statusData set, heartbeatData nil
	widget.statusData = &StatusPageData{}
	got = widget.content()
	if got != "Loading..." {
		t.Errorf("expected 'Loading...' when heartbeatData nil, got %q", got)
	}
}

func TestContent_AllUp(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{
				"1": {{Status: 1}},
				"2": {{Status: 1}},
				"3": {{Status: 1}},
			},
			UptimeList: map[string]float64{
				"1": 1.0,
				"2": 0.99,
				"3": 0.98,
			},
		},
	}
	got := widget.content()
	if !strings.Contains(got, "Up: [green]3") {
		t.Errorf("expected 'Up: [green]3' in output, got %q", got)
	}
	if !strings.Contains(got, "Down: [green]0") {
		t.Errorf("expected green 0 for down when none down, got %q", got)
	}
	// Average uptime: (1.0+0.99+0.98)/3 * 100 = 99.0%
	if !strings.Contains(got, "99.0%") {
		t.Errorf("expected '99.0%%' in output, got %q", got)
	}
	// Should not show maintenance or pending
	if strings.Contains(got, "Maint") {
		t.Errorf("unexpected 'Maint' in output: %q", got)
	}
	if strings.Contains(got, "Pend") {
		t.Errorf("unexpected 'Pend' in output: %q", got)
	}
}

func TestContent_MixedStatuses(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{
				"1": {{Status: 1}},           // UP
				"2": {{Status: 0}},           // DOWN
				"3": {{Status: 3}},           // MAINTENANCE
				"4": {{Status: 2}},           // PENDING
				"5": {{Status: 1}, {Status: 0}}, // last is DOWN
			},
			UptimeList: map[string]float64{
				"1": 1.0,
				"2": 0.5,
				"3": 0.99,
				"4": 0.8,
				"5": 0.7,
			},
		},
	}
	got := widget.content()
	if !strings.Contains(got, "Up: [green]1") {
		t.Errorf("expected 'Up: [green]1', got %q", got)
	}
	if !strings.Contains(got, "Down: [red]2") {
		t.Errorf("expected 'Down: [red]2', got %q", got)
	}
	if !strings.Contains(got, "Maint:") {
		t.Errorf("expected 'Maint:' in output, got %q", got)
	}
	if !strings.Contains(got, "Pend:") {
		t.Errorf("expected 'Pend:' in output, got %q", got)
	}
}

func TestContent_WithIncident(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{
			Incident: &Incident{CreatedDate: "2023-10-27 10:30:00.123"},
		},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{
				"1": {{Status: 1}},
			},
			UptimeList: map[string]float64{
				"1": 1.0,
			},
		},
	}
	got := widget.content()
	if !strings.Contains(got, "Incident:") {
		t.Errorf("expected 'Incident:' in output, got %q", got)
	}
	if !strings.Contains(got, "h ago") {
		t.Errorf("expected 'h ago' in output, got %q", got)
	}
}

func TestContent_WithIncidentUnparsableDate(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{
			Incident: &Incident{CreatedDate: "not-a-date"},
		},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{},
			UptimeList:    map[string]float64{},
		},
	}
	got := widget.content()
	if !strings.Contains(got, "unparsable date") {
		t.Errorf("expected 'unparsable date' in output, got %q", got)
	}
}

func TestContent_EmptyHeartbeatList(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{},
			UptimeList:    map[string]float64{},
		},
	}
	got := widget.content()
	if !strings.Contains(got, "Up: [green]0") {
		t.Errorf("expected 'Up: [green]0', got %q", got)
	}
	if !strings.Contains(got, "Down: [green]0") {
		t.Errorf("expected 'Down: [green]0', got %q", got)
	}
	// 0 monitors means 0.0% uptime
	if !strings.Contains(got, "0.0%") {
		t.Errorf("expected '0.0%%', got %q", got)
	}
}

func TestContent_EmptyHeartbeatSlice(t *testing.T) {
	// A monitor key exists but has no heartbeats
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{
				"1": {},
			},
			UptimeList: map[string]float64{
				"1": 0.5,
			},
		},
	}
	got := widget.content()
	// Should not panic and should show 0 for all counts
	if !strings.Contains(got, "Up: [green]0") {
		t.Errorf("expected 'Up: [green]0', got %q", got)
	}
}

func TestHeartbeatStatusConstants(t *testing.T) {
	tests := []struct {
		status HeartbeatStatus
		want   int
	}{
		{DOWN, 0},
		{UP, 1},
		{PENDING, 2},
		{MAINTENANCE, 3},
	}
	for _, tc := range tests {
		if int(tc.status) != tc.want {
			t.Errorf("HeartbeatStatus %d: got %d, want %d", tc.status, int(tc.status), tc.want)
		}
	}
}

func TestConfigText(t *testing.T) {
	widget := &Widget{}
	got := widget.ConfigText()
	if !strings.Contains(got, "url") {
		t.Errorf("expected 'url' in config text, got %q", got)
	}
}

func TestRefresh_Success(t *testing.T) {
	statusData := StatusPageData{Incident: nil}
	heartbeatData := HeartbeatData{
		HeartbeatList: map[string][]*Heartbeat{
			"1": {{Status: 1}},
		},
		UptimeList: map[string]float64{"1": 0.99},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "/api/status-page/heartbeat/"):
			_ = json.NewEncoder(w).Encode(heartbeatData)
		case strings.Contains(r.URL.Path, "/api/status-page/"):
			_ = json.NewEncoder(w).Encode(statusData)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
			url:    srv.URL + "/status/test",
		},
	}

	// Refresh calls display() which calls Redraw(). Without a real TextWidget,
	// Redraw will panic. We just test the fetch logic by calling the fetch
	// methods directly (already tested above). But let's verify parseURL works
	// with the test server URL.
	baseURL, slug, err := parseURL(widget.settings.url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if slug != "test" {
		t.Errorf("expected slug 'test', got %q", slug)
	}

	sd, err := widget.fetchStatusData(baseURL, slug)
	if err != nil {
		t.Fatalf("fetchStatusData error: %v", err)
	}
	widget.statusData = sd

	hd, err := widget.fetchHeartbeatData(baseURL, slug)
	if err != nil {
		t.Fatalf("fetchHeartbeatData error: %v", err)
	}
	widget.heartbeatData = hd

	// Verify content renders correctly
	got := widget.content()
	if !strings.Contains(got, "Up: [green]1") {
		t.Errorf("expected 'Up: [green]1' in content, got %q", got)
	}
}

func TestRefresh_InvalidURL(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
			url:    "",
		},
	}

	baseURL, slug, err := parseURL(widget.settings.url)
	if err == nil {
		t.Fatalf("expected error for empty URL, got base=%q slug=%q", baseURL, slug)
	}
	widget.err = err

	got := widget.content()
	if !strings.Contains(got, "URL is not defined") {
		t.Errorf("expected URL error in content, got %q", got)
	}
}

func TestContent_OutOfRangeStatus(t *testing.T) {
	// Status value outside 0-3 range should be ignored
	widget := &Widget{
		settings: &Settings{
			common: makeTestCommon(),
		},
		statusData: &StatusPageData{},
		heartbeatData: &HeartbeatData{
			HeartbeatList: map[string][]*Heartbeat{
				"1": {{Status: 99}},
				"2": {{Status: -1}},
			},
			UptimeList: map[string]float64{
				"1": 1.0,
				"2": 1.0,
			},
		},
	}
	got := widget.content()
	// All counts should be 0 since both statuses are out of range
	if !strings.Contains(got, "Up: [green]0") {
		t.Errorf("expected 'Up: [green]0', got %q", got)
	}
	if !strings.Contains(got, "Down: [green]0") {
		t.Errorf("expected 'Down: [green]0', got %q", got)
	}
}

func TestNewWidget(t *testing.T) {
	widget := makeTestWidget("https://example.com/status/test")
	if widget == nil {
		t.Fatal("expected non-nil widget")
	}
	if widget.settings == nil {
		t.Fatal("expected settings to be set")
	}
}

func TestRefresh_FullIntegration_Success(t *testing.T) {
	statusData := StatusPageData{Incident: nil}
	heartbeatData := HeartbeatData{
		HeartbeatList: map[string][]*Heartbeat{
			"1": {{Status: 1}},
		},
		UptimeList: map[string]float64{"1": 0.99},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "/api/status-page/heartbeat/"):
			_ = json.NewEncoder(w).Encode(heartbeatData)
		case strings.Contains(r.URL.Path, "/api/status-page/"):
			_ = json.NewEncoder(w).Encode(statusData)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	widget := makeTestWidget(srv.URL + "/status/test")
	widget.Refresh()

	if widget.err != nil {
		t.Fatalf("unexpected error after Refresh: %v", widget.err)
	}
	if widget.statusData == nil {
		t.Fatal("expected statusData to be set")
	}
	if widget.heartbeatData == nil {
		t.Fatal("expected heartbeatData to be set")
	}
}

func TestRefresh_FullIntegration_InvalidURL(t *testing.T) {
	widget := makeTestWidget("")
	widget.Refresh()

	if widget.err == nil {
		t.Fatal("expected error for empty URL")
	}
	if !strings.Contains(widget.err.Error(), "URL is not defined") {
		t.Errorf("unexpected error: %v", widget.err)
	}
}

func TestRefresh_FullIntegration_StatusFetchError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	widget := makeTestWidget(srv.URL + "/status/test")
	widget.Refresh()

	if widget.err == nil {
		t.Fatal("expected error for status fetch failure")
	}
}

func TestRefresh_FullIntegration_HeartbeatFetchError(t *testing.T) {
	statusData := StatusPageData{Incident: nil}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/api/status-page/heartbeat/"):
			w.WriteHeader(http.StatusInternalServerError)
		case strings.Contains(r.URL.Path, "/api/status-page/"):
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(statusData)
		}
	}))
	defer srv.Close()

	widget := makeTestWidget(srv.URL + "/status/test")
	widget.Refresh()

	if widget.err == nil {
		t.Fatal("expected error for heartbeat fetch failure")
	}
}
