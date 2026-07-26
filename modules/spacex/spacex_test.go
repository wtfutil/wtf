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

func sampleLaunch() Launch {
	return Launch{
		FlightNumber: 100,
		MissionName:  "Starlink-15",
		LaunchDate:   1609459200, // 2021-01-01 00:00:00 UTC
		IsTentative:  false,
		Rocket:       Rocket{Name: "Falcon 9"},
		LaunchSite:   LaunchSite{Name: "KSC LC 39A"},
		Links:        Links{RedditLink: "https://reddit.com/r/spacex", YouTubeLink: "https://youtube.com/watch?v=abc"},
		Details:      "Launching satellites",
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
	launch := sampleLaunch()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(launch)
	})
	defer cleanup()

	result, err := NextLaunch()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.MissionName != "Starlink-15" {
		t.Errorf("expected MissionName 'Starlink-15', got %q", result.MissionName)
	}
	if result.FlightNumber != 100 {
		t.Errorf("expected FlightNumber 100, got %d", result.FlightNumber)
	}
	if result.Rocket.Name != "Falcon 9" {
		t.Errorf("expected Rocket.Name 'Falcon 9', got %q", result.Rocket.Name)
	}
	if result.LaunchSite.Name != "KSC LC 39A" {
		t.Errorf("expected LaunchSite.Name 'KSC LC 39A', got %q", result.LaunchSite.Name)
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
	if result.Links.RedditLink != "https://reddit.com/r/spacex" {
		t.Errorf("expected RedditLink, got %q", result.Links.RedditLink)
	}
	if result.Details != "Launching satellites" {
		t.Errorf("expected Details, got %q", result.Details)
	}
}

func TestNextLaunch_InvalidJSON(t *testing.T) {
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "not valid json{{{")
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
		fmt.Fprint(w, "")
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
		fmt.Fprint(w, "{}")
	})
	defer cleanup()

	result, err := NextLaunch()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.MissionName != "" {
		t.Errorf("expected empty MissionName, got %q", result.MissionName)
	}
	if result.FlightNumber != 0 {
		t.Errorf("expected FlightNumber 0, got %d", result.FlightNumber)
	}
}

// --- JSON Parsing Table Tests ---

func TestLaunchParsing(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		wantErr     bool
		checkFlight int
		checkName   string
	}{
		{
			name:        "full payload",
			json:        `{"flight_number":42,"mission_name":"CRS-21","launch_date_unix":1700000000,"tentative":true,"rocket":{"rocket_name":"Falcon Heavy"},"launch_site":{"site_name_long":"VAFB"},"links":{"reddit_campaign":"https://r.com","video_link":"https://yt.com"},"details":"cargo"}`,
			wantErr:     false,
			checkFlight: 42,
			checkName:   "CRS-21",
		},
		{
			name:        "minimal payload",
			json:        `{"flight_number":1}`,
			wantErr:     false,
			checkFlight: 1,
			checkName:   "",
		},
		{
			name:    "invalid JSON",
			json:    `{invalid`,
			wantErr: true,
		},
		{
			name:        "extra fields ignored",
			json:        `{"flight_number":7,"mission_name":"Test","unknown_field":"hello"}`,
			wantErr:     false,
			checkFlight: 7,
			checkName:   "Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, tt.json)
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
			if result.FlightNumber != tt.checkFlight {
				t.Errorf("FlightNumber: got %d, want %d", result.FlightNumber, tt.checkFlight)
			}
			if result.MissionName != tt.checkName {
				t.Errorf("MissionName: got %q, want %q", result.MissionName, tt.checkName)
			}
		})
	}
}

// --- Display / Content Formatting Tests ---

func TestContentFormatting(t *testing.T) {
	launch := sampleLaunch()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(launch)
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
		"Name: Starlink-15",
		fmt.Sprintf("Date: %s", expectedDate),
		"Site: KSC LC 39A",
		"YouTube: https://youtube.com/watch?v=abc",
		"Reddit: https://reddit.com/r/spacex",
		"RocketName: Falcon 9",
		"Details: Launching satellites",
	}
	for _, check := range checks {
		if !strings.Contains(content, check) {
			t.Errorf("content missing %q", check)
		}
	}
}

func TestContentFormattingSmallHeight(t *testing.T) {
	launch := sampleLaunch()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(launch)
	})
	defer cleanup()

	widget := buildTestWidget(1) // height < 2
	_, content, _ := widget.content()

	// Should NOT contain details section with height < 2
	if strings.Contains(content, "RocketName:") {
		t.Error("details section should not appear when height < 2")
	}
	if strings.Contains(content, "Details: Launching") {
		t.Error("details should not appear when height < 2")
	}
	// Should still have mission info
	if !strings.Contains(content, "Name: Starlink-15") {
		t.Error("should contain mission name regardless of height")
	}
}

func TestContentOnError(t *testing.T) {
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "broken json!!!!")
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
	launch := sampleLaunch()
	_, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(launch)
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
			if settings.Common.Name != tt.expectedName {
				t.Errorf("Name: got %q, want %q", settings.Common.Name, tt.expectedName)
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
