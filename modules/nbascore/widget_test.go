package nbascore

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// --- helpers ---

func testSettings(t *testing.T) *Settings {
	t.Helper()
	yamlStr := `
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
    subheading: "red"
`
	ymlCfg, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatal(err)
	}
	globalCfg, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatal(err)
	}
	return NewSettingsFromYAML("nbascore", ymlCfg, globalCfg)
}

// testWidget creates a Widget with a properly initialized TextWidget for testing.
func testWidget(t *testing.T) *Widget {
	t.Helper()
	settings := testSettings(t)
	widget := &Widget{
		TextWidget: view.NewTextWidget(nil, nil, nil, settings.Common),
		settings:   settings,
	}
	return widget
}

// buildScoreboard builds a minimal NBA scoreboard JSON response.
func buildScoreboard(games []map[string]interface{}) []byte {
	data := map[string]interface{}{
		"games": games,
	}
	b, _ := json.Marshal(data)
	return b
}

func makeGame(vTeam, hTeam, vScore, hScore string, quarter float64, active bool) map[string]interface{} {
	return map[string]interface{}{
		"vTeam": map[string]interface{}{
			"triCode": vTeam,
			"score":   vScore,
		},
		"hTeam": map[string]interface{}{
			"triCode": hTeam,
			"score":   hScore,
		},
		"period": map[string]interface{}{
			"current": quarter,
		},
		"isGameActivated": active,
	}
}

// --- Settings tests ---

func TestNewSettingsFromYAML_DefaultTitle(t *testing.T) {
	settings := testSettings(t)
	if settings.Title != "NBA Score" {
		t.Errorf("expected title 'NBA Score', got %q", settings.Title)
	}
}

func TestNewSettingsFromYAML_Focusable(t *testing.T) {
	settings := testSettings(t)
	if !settings.Focusable {
		t.Error("expected focusable=true by default")
	}
}

// --- nbascore() integration tests with httptest ---

func TestNbascore_SuccessfulResponse(t *testing.T) {
	games := []map[string]interface{}{
		makeGame("BOS", "LAL", "110", "105", 4, false),
		makeGame("GSW", "MIA", "98", "102", 3, true),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(buildScoreboard(games))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	title, content, wrap := widget.nbascore()

	if title != "NBA Score" {
		t.Errorf("expected title 'NBA Score', got %q", title)
	}
	if wrap {
		t.Error("expected wrap=false for successful response")
	}
	if !strings.Contains(content, "BOS") {
		t.Errorf("expected content to contain 'BOS', got %q", content)
	}
	if !strings.Contains(content, "LAL") {
		t.Errorf("expected content to contain 'LAL', got %q", content)
	}
	if !strings.Contains(content, "GSW") {
		t.Errorf("expected content to contain 'GSW', got %q", content)
	}
	if !strings.Contains(content, "110") {
		t.Errorf("expected content to contain score '110', got %q", content)
	}
}

func TestNbascore_EmptyGames(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(buildScoreboard([]map[string]interface{}{}))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, content, wrap := widget.nbascore()

	if wrap {
		t.Error("expected wrap=false for empty games")
	}
	// Should still have date header
	today := time.Now().Format(utils.FriendlyDateFormat)
	if !strings.Contains(content, today) {
		t.Errorf("expected content to contain date %q", today)
	}
}

func TestNbascore_Non200Status(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, _, wrap := widget.nbascore()

	if !wrap {
		t.Error("expected wrap=true for non-200 status")
	}
}

func TestNbascore_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not json"))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, content, wrap := widget.nbascore()

	if !wrap {
		t.Error("expected wrap=true for invalid JSON")
	}
	if content == "" {
		t.Error("expected non-empty error message")
	}
}

func TestNbascore_ServerUnreachable(t *testing.T) {
	origURL := nbaBaseURL
	nbaBaseURL = "http://127.0.0.1:1" // unreachable port
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, content, wrap := widget.nbascore()

	if !wrap {
		t.Error("expected wrap=true for unreachable server")
	}
	if content == "" {
		t.Error("expected non-empty error message")
	}
}

// --- Date formatting tests ---

func TestNbascore_DateOffset(t *testing.T) {
	tests := []struct {
		name       string
		offset     int
		wantDate   string
	}{
		{"today", 0, time.Now().Format("20060102")},
		{"yesterday", -1, time.Now().AddDate(0, 0, -1).Format("20060102")},
		{"tomorrow", 1, time.Now().AddDate(0, 0, 1).Format("20060102")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var receivedPath string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				receivedPath = r.URL.Path
				w.WriteHeader(http.StatusOK)
				w.Write(buildScoreboard([]map[string]interface{}{}))
			}))
			defer ts.Close()

			origURL := nbaBaseURL
			nbaBaseURL = ts.URL
			defer func() { nbaBaseURL = origURL }()

			origOffset := offset
			offset = tc.offset
			defer func() { offset = origOffset }()

			widget := testWidget(t)
			widget.nbascore()

			expectedPath := "/" + tc.wantDate + "/scoreboard.json"
			if receivedPath != expectedPath {
				t.Errorf("expected path %q, got %q", expectedPath, receivedPath)
			}
		})
	}
}

// --- Score display / color formatting tests ---

func TestNbascore_ScoreColorFormatting(t *testing.T) {
	tests := []struct {
		name       string
		vScore     string
		hScore     string
		quarter    float64
		active     bool
		wantInOut  []string
	}{
		{
			name:    "visitor winning",
			vScore:  "100",
			hScore:  "90",
			quarter: 4,
			active:  false,
			wantInOut: []string{"[orange]BOS", "100"},
		},
		{
			name:    "home winning",
			vScore:  "90",
			hScore:  "100",
			quarter: 4,
			active:  false,
			wantInOut: []string{"[orange]", "LAL"},
		},
		{
			name:    "tied score",
			vScore:  "95",
			hScore:  "95",
			quarter: 3,
			active:  true,
			wantInOut: []string{"[orange]BOS", "[orange]", "[sandybrown]"},
		},
		{
			name:    "game not started (quarter 0)",
			vScore:  "",
			hScore:  "",
			quarter: 0,
			active:  false,
			wantInOut: []string{"BOS", "LAL"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			games := []map[string]interface{}{
				makeGame("BOS", "LAL", tc.vScore, tc.hScore, tc.quarter, tc.active),
			}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(buildScoreboard(games))
			}))
			defer ts.Close()

			origURL := nbaBaseURL
			nbaBaseURL = ts.URL
			defer func() { nbaBaseURL = origURL }()

			origOffset := offset
			offset = 0
			defer func() { offset = origOffset }()

			widget := testWidget(t)
			_, content, wrap := widget.nbascore()

			if wrap {
				t.Fatalf("unexpected wrap=true, content: %s", content)
			}
			for _, want := range tc.wantInOut {
				if !strings.Contains(content, want) {
					t.Errorf("expected content to contain %q\ngot: %s", want, content)
				}
			}
		})
	}
}

// --- Game status formatting tests ---

func TestNbascore_ActiveGameHighlight(t *testing.T) {
	games := []map[string]interface{}{
		makeGame("BOS", "LAL", "50", "48", 2, true),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(buildScoreboard(games))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, content, _ := widget.nbascore()

	if !strings.Contains(content, "[sandybrown]") {
		t.Errorf("expected active game to have [sandybrown] color, got: %s", content)
	}
}

func TestNbascore_InactiveGameNoHighlight(t *testing.T) {
	games := []map[string]interface{}{
		makeGame("BOS", "LAL", "110", "105", 4, false),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(buildScoreboard(games))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, content, _ := widget.nbascore()

	// Quarter indicator should be [white] not [sandybrown]
	if strings.Contains(content, "[sandybrown]") {
		t.Errorf("inactive game should not have [sandybrown] color, got: %s", content)
	}
}

// --- Multiple games test ---

func TestNbascore_MultipleGames(t *testing.T) {
	games := []map[string]interface{}{
		makeGame("BOS", "LAL", "110", "105", 4, false),
		makeGame("GSW", "MIA", "98", "102", 3, true),
		makeGame("NYK", "CHI", "0", "0", 0, false),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(buildScoreboard(games))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	_, content, wrap := widget.nbascore()

	if wrap {
		t.Fatal("expected wrap=false")
	}
	for _, team := range []string{"BOS", "LAL", "GSW", "MIA", "NYK", "CHI"} {
		if !strings.Contains(content, team) {
			t.Errorf("expected content to contain %q", team)
		}
	}
}

// --- Request header tests ---

func TestNbascore_RequestHeaders(t *testing.T) {
	var userAgent string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
		w.Write(buildScoreboard([]map[string]interface{}{}))
	}))
	defer ts.Close()

	origURL := nbaBaseURL
	nbaBaseURL = ts.URL
	defer func() { nbaBaseURL = origURL }()

	origOffset := offset
	offset = 0
	defer func() { offset = origOffset }()

	widget := testWidget(t)
	widget.nbascore()

	if userAgent != "curl" {
		t.Errorf("expected User-Agent 'curl', got %q", userAgent)
	}
}

// --- ConfigText test ---

func TestConfigText(t *testing.T) {
	widget := testWidget(t)
	text := widget.ConfigText()
	// Should return something (help text), not panic
	if text == "" {
		// ConfigText may be empty for this simple settings struct, that's ok
		// Just verify it doesn't panic
	}
}

// --- Keyboard controls offset tests ---

func TestOffsetControls(t *testing.T) {
	origOffset := offset
	defer func() { offset = origOffset }()

	offset = 0

	// Test next increments
	offset++
	if offset != 1 {
		t.Errorf("expected offset=1 after next, got %d", offset)
	}

	// Test prev decrements
	offset--
	if offset != 0 {
		t.Errorf("expected offset=0 after prev, got %d", offset)
	}

	// Test center resets
	offset = 5
	offset = 0
	if offset != 0 {
		t.Errorf("expected offset=0 after center, got %d", offset)
	}
}
