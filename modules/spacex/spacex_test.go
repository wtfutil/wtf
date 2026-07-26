package spacex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

func sampleLL2Response() ll2Response {
	return ll2Response{
		Count: 1,
		Results: []ll2Launch{
			{
				Name:   "Falcon 9 Block 5 | Starlink Group 17-52",
				Net:    "2021-01-01T00:00:00Z",
				Status: ll2Status{Name: "Go for Launch"},
				Rocket: ll2Rocket{Configuration: ll2RocketConfig{FullName: "Falcon 9 Block 5"}},
				Mission: &ll2Mission{
					Name:        "Starlink Group 17-52",
					Description: "A batch of satellites for Starlink",
				},
				Pad:     &ll2Pad{Name: "Space Launch Complex 4E"},
				VidURLs: []ll2VidURL{{URL: "https://youtube.com/watch?v=abc"}},
			},
		},
	}
}

func setupTestServer(handler http.HandlerFunc) (*httptest.Server, func()) {
	server := httptest.NewServer(handler)
	original := spacexLaunchAPI
	spacexLaunchAPI = server.URL
	cleanup := func() {
		spacexLaunchAPI = original
		server.Close()
	}
	return server, cleanup
}

// --- Client Tests ---

func TestNextLaunch_Success(t *testing.T) {
	ll2Resp := sampleLL2Response()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ll2Resp)
	})
	defer cleanup()

	result, err := NextLaunch()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.MissionName != "Starlink Group 17-52" {
		t.Errorf("expected MissionName 'Starlink Group 17-52', got %q", result.MissionName)
	}
	if result.Rocket.Name != "Falcon 9 Block 5" {
		t.Errorf("expected Rocket.Name 'Falcon 9 Block 5', got %q", result.Rocket.Name)
	}
	if result.LaunchSite.Name != "Space Launch Complex 4E" {
		t.Errorf("expected LaunchSite.Name 'Space Launch Complex 4E', got %q", result.LaunchSite.Name)
	}
	if result.LaunchDate != 1609459200 {
		t.Errorf("expected LaunchDate 1609459200, got %d", result.LaunchDate)
	}
	if result.IsTentative != false {
		t.Errorf("expected IsTentative false, got true")
	}
	if result.Links.YouTubeLink != "https://youtube.com/watch?v=abc" {
		t.Errorf("expected YouTubeLink, got %q", result.Links.YouTubeLink)
	}
	if result.Details != "A batch of satellites for Starlink" {
		t.Errorf("expected Details, got %q", result.Details)
	}
}

func TestNextLaunch_InvalidJSON(t *testing.T) {
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "not valid json{{{")
	})
	defer cleanup()

	result, err := NextLaunch()
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
	if result != nil {
		t.Errorf("expected nil result on error, got %+v", result)
	}
}

func TestNextLaunch_ServerError(t *testing.T) {
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "")
	})
	defer cleanup()

	// Server returns 500 with empty body — JSON decode will fail
	_, err := NextLaunch()
	if err == nil {
		t.Fatal("expected error for empty response body, got nil")
	}
}

func TestNextLaunch_NetworkError(t *testing.T) {
	original := spacexLaunchAPI
	spacexLaunchAPI = "http://127.0.0.1:1" // unreachable port
	defer func() { spacexLaunchAPI = original }()

	_, err := NextLaunch()
	if err == nil {
		t.Fatal("expected network error, got nil")
	}
}

func TestNextLaunch_EmptyObject(t *testing.T) {
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"count":0,"results":[]}`)
	})
	defer cleanup()

	_, err := NextLaunch()
	if err == nil {
		t.Fatal("expected error for empty results, got nil")
	}
	if err.Error() != "no upcoming SpaceX launches found" {
		t.Errorf("unexpected error: %v", err)
	}
}

// --- JSON Parsing Table Tests ---

func TestLaunchParsing(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		wantErr     bool
		checkName   string
		checkRocket string
	}{
		{
			name:        "full payload",
			json:        `{"count":1,"results":[{"name":"Falcon Heavy | Europa Clipper","net":"2026-10-10T12:00:00Z","status":{"name":"Go for Launch"},"rocket":{"configuration":{"full_name":"Falcon Heavy"}},"mission":{"name":"Europa Clipper","description":"NASA mission"},"pad":{"name":"LC-39A"},"vidURLs":[{"url":"https://yt.com"}]}]}`,
			wantErr:     false,
			checkName:   "Europa Clipper",
			checkRocket: "Falcon Heavy",
		},
		{
			name:        "minimal payload",
			json:        `{"count":1,"results":[{"name":"Falcon 9 | Test","net":"2026-01-01T00:00:00Z","status":{"name":"TBD"},"rocket":{"configuration":{"full_name":"Falcon 9"}}}]}`,
			wantErr:     false,
			checkName:   "Falcon 9 | Test",
			checkRocket: "Falcon 9",
		},
		{
			name:    "invalid JSON",
			json:    `{invalid`,
			wantErr: true,
		},
		{
			name:    "empty results",
			json:    `{"count":0,"results":[]}`,
			wantErr: true,
		},
		{
			name:        "TBD is tentative",
			json:        `{"count":1,"results":[{"name":"Test","net":"2026-01-01T00:00:00Z","status":{"name":"TBD"},"rocket":{"configuration":{"full_name":"F9"}}}]}`,
			wantErr:     false,
			checkName:   "Test",
			checkRocket: "F9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = fmt.Fprint(w, tt.json)
			})
			defer cleanup()

			result, err := NextLaunch()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.MissionName != tt.checkName {
				t.Errorf("MissionName: got %q, want %q", result.MissionName, tt.checkName)
			}
			if result.Rocket.Name != tt.checkRocket {
				t.Errorf("Rocket.Name: got %q, want %q", result.Rocket.Name, tt.checkRocket)
			}
		})
	}
}

// --- Display / Content Formatting Tests ---

func TestContentFormatting(t *testing.T) {
	ll2Resp := sampleLL2Response()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ll2Resp)
	})
	defer cleanup()

	widget := buildTestWidget(10)
	title, content, wrap := widget.content()

	if title != "Next SpaceX 🚀" {
		t.Errorf("expected default title, got %q", title)
	}
	if !wrap {
		t.Error("expected wrap=true")
	}

	expectedDate := wtf.UnixTime(1609459200).Format(time.RFC822)
	checks := []string{
		"Name: Starlink Group 17-52",
		fmt.Sprintf("Date: %s", expectedDate),
		"Site: Space Launch Complex 4E",
		"YouTube: https://youtube.com/watch?v=abc",
		"RocketName: Falcon 9 Block 5",
		"Details: A batch of satellites for Starlink",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("content missing %q", check)
		}
	}
}

func TestContentFormattingSmallHeight(t *testing.T) {
	ll2Resp := sampleLL2Response()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ll2Resp)
	})
	defer cleanup()

	widget := buildTestWidget(1) // height < 2
	_, content, _ := widget.content()

	// Should NOT contain details section with height < 2
	if strings.Contains(content, "RocketName:") {
		t.Error("details section should not appear when height < 2")
	}
	if strings.Contains(content, "Details: A batch") {
		t.Error("details should not appear when height < 2")
	}
	// Should still have mission info
	if !strings.Contains(content, "Name: Starlink Group 17-52") {
		t.Error("should contain mission name regardless of height")
	}
}

func TestContentOnError(t *testing.T) {
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "broken json!!!!")
	})
	defer cleanup()

	widget := buildTestWidget(10)
	_, content, _ := widget.content()

	// On error, content should be empty string
	if content != "" {
		t.Errorf("expected empty content on error, got %q", content)
	}
	if widget.err == nil {
		t.Error("expected widget.err to be set")
	}
}

func TestContentCustomTitle(t *testing.T) {
	ll2Resp := sampleLL2Response()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ll2Resp)
	})
	defer cleanup()

	widget := buildTestWidgetWithTitle(10, "My Custom Title")
	title, _, _ := widget.content()

	if title != "My Custom Title" {
		t.Errorf("expected custom title, got %q", title)
	}
}

// --- Settings Tests ---

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name         string
		yaml         string
		expectedName string
	}{
		{
			name:         "basic settings",
			yaml:         "enabled: true\nposition:\n  top: 0\n  left: 0\n  height: 3\n  width: 2\n",
			expectedName: "spacex",
		},
		{
			name:         "with title override",
			yaml:         "enabled: true\ntitle: My Rockets\nposition:\n  top: 0\n  left: 0\n  height: 3\n  width: 2\n",
			expectedName: "spacex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			if err != nil {
				t.Fatalf("failed to parse yaml: %v", err)
			}

			globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    subheading: red\n")
			settings := NewSettingsFromYAML("spacex", ymlConfig, globalConfig)

			if settings.Common == nil {
				t.Fatal("expected Common to be set")
			}
			if settings.Name != tt.expectedName {
				t.Errorf("Name: got %q, want %q", settings.Name, tt.expectedName)
			}
		})
	}
}

func TestDefaultFocusable(t *testing.T) {
	if defaultFocusable != false {
		t.Errorf("expected defaultFocusable to be false")
	}
}

// --- handleError Tests ---

func TestHandleError(t *testing.T) {
	widget := buildTestWidget(10)
	if widget.err != nil {
		t.Fatal("err should be nil initially")
	}

	testErr := fmt.Errorf("test error")
	handleError(widget, testErr)

	if widget.err == nil {
		t.Fatal("err should be set after handleError")
	}
	if widget.err.Error() != "test error" {
		t.Errorf("expected 'test error', got %q", widget.err.Error())
	}
}

// --- Helpers ---

func buildTestWidget(height int) *Widget {
	return buildTestWidgetWithTitle(height, "")
}

func buildTestWidgetWithTitle(height int, title string) *Widget {
	yaml := fmt.Sprintf("enabled: true\nposition:\n  top: 0\n  left: 0\n  height: %d\n  width: 2\n", height)
	if title != "" {
		yaml += fmt.Sprintf("title: %s\n", title)
	}
	ymlConfig, _ := config.ParseYaml(yaml)
	globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    subheading: red\n")
	settings := NewSettingsFromYAML("spacex", ymlConfig, globalConfig)

	widget := &Widget{
		settings: settings,
	}
	widget.TextWidget = newTestTextWidget(settings, height)
	return widget
}

// newTestTextWidget creates a minimal TextWidget for testing without tview.
func newTestTextWidget(settings *Settings, height int) view.TextWidget {
	return view.NewTextWidget(nil, nil, nil, settings.Common)
}
