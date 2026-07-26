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

func testWidget(t *testing.T) *Widget {
	t.Helper()
	settings := testSettings(t)
	widget := &Widget{
		TextWidget: view.NewTextWidget(nil, nil, nil, settings.Common),
		settings:   settings,
	}
	return widget
}

// buildESPNResponse builds a minimal ESPN scoreboard JSON response.
func buildESPNResponse(events []espnEvent) []byte {
	data := espnResponse{Events: events}
	b, _ := json.Marshal(data)
	return b
}

// buildESPNResponseWithDay builds an ESPN response with a day.date field.
func buildESPNResponseWithDay(events []espnEvent, dayDate string) []byte {
	data := espnResponse{
		Day:    espnDay{Date: dayDate},
		Events: events,
	}
	b, _ := json.Marshal(data)
	return b
}

func makeEvent(awayAbbr, homeAbbr, awayScore, homeScore string, period int, state string) espnEvent {
	return espnEvent{
		Competitions: []espnCompetition{
			{
				Competitors: []espnCompetitor{
					{
						HomeAway: "home",
						Team:     espnTeam{Abbreviation: homeAbbr},
						Score:    homeScore,
					},
					{
						HomeAway: "away",
						Team:     espnTeam{Abbreviation: awayAbbr},
						Score:    awayScore,
					},
				},
			},
		},
		Status: espnStatus{
			Period: period,
			Type:   espnStatusType{State: state},
		},
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
	events := []espnEvent{
		makeEvent("BOS", "LAL", "110", "105", 4, "post"),
		makeEvent("GSW", "MIA", "98", "102", 3, "in"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buildESPNResponse(events))
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
		_, _ = w.Write(buildESPNResponseWithDay([]espnEvent{}, "2026-10-05"))
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
	today := time.Now().Format(utils.FriendlyDateFormat)
	if !strings.Contains(content, today) {
		t.Errorf("expected content to contain date %q", today)
	}
	if !strings.Contains(content, "No games scheduled") {
		t.Errorf("expected 'No games scheduled' message, got %q", content)
	}
	if !strings.Contains(content, "Next game:") {
		t.Errorf("expected 'Next game:' message, got %q", content)
	}
}

func TestNbascore_EmptyGamesNoNextDate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buildESPNResponse([]espnEvent{}))
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
	if !strings.Contains(content, "No games scheduled") {
		t.Errorf("expected 'No games scheduled' message, got %q", content)
	}
	if strings.Contains(content, "Next game:") {
		t.Errorf("should not show 'Next game:' when no date available, got %q", content)
	}
}

func TestNbascore_Non200Status(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("server error"))
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
		_, _ = w.Write([]byte("not json"))
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
	nbaBaseURL = "http://127.0.0.1:1"
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
		name     string
		offset   int
		wantDate string
	}{
		{"today", 0, time.Now().Format("20060102")},
		{"yesterday", -1, time.Now().AddDate(0, 0, -1).Format("20060102")},
		{"tomorrow", 1, time.Now().AddDate(0, 0, 1).Format("20060102")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var receivedQuery string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				receivedQuery = r.URL.RawQuery
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(buildESPNResponse([]espnEvent{}))
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

			expectedQuery := "dates=" + tc.wantDate
			if receivedQuery != expectedQuery {
				t.Errorf("expected query %q, got %q", expectedQuery, receivedQuery)
			}
		})
	}
}

// --- Score display / color formatting tests ---

func TestNbascore_ScoreColorFormatting(t *testing.T) {
	tests := []struct {
		name      string
		vScore    string
		hScore    string
		period    int
		state     string
		wantInOut []string
	}{
		{
			name:      "visitor winning",
			vScore:    "100",
			hScore:    "90",
			period:    4,
			state:     "post",
			wantInOut: []string{"[orange]BOS", "100"},
		},
		{
			name:      "home winning",
			vScore:    "90",
			hScore:    "100",
			period:    4,
			state:     "post",
			wantInOut: []string{"[orange]", "LAL"},
		},
		{
			name:      "tied score",
			vScore:    "95",
			hScore:    "95",
			period:    3,
			state:     "in",
			wantInOut: []string{"[orange]BOS", "[orange]", "[sandybrown]"},
		},
		{
			name:      "game not started (period 0)",
			vScore:    "0",
			hScore:    "0",
			period:    0,
			state:     "pre",
			wantInOut: []string{"BOS", "LAL"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			events := []espnEvent{
				makeEvent("BOS", "LAL", tc.vScore, tc.hScore, tc.period, tc.state),
			}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(buildESPNResponse(events))
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
	events := []espnEvent{
		makeEvent("BOS", "LAL", "50", "48", 2, "in"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buildESPNResponse(events))
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
	events := []espnEvent{
		makeEvent("BOS", "LAL", "110", "105", 4, "post"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buildESPNResponse(events))
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

	if strings.Contains(content, "[sandybrown]") {
		t.Errorf("inactive game should not have [sandybrown] color, got: %s", content)
	}
}

// --- Multiple games test ---

func TestNbascore_MultipleGames(t *testing.T) {
	events := []espnEvent{
		makeEvent("BOS", "LAL", "110", "105", 4, "post"),
		makeEvent("GSW", "MIA", "98", "102", 3, "in"),
		makeEvent("NYK", "CHI", "0", "0", 0, "pre"),
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buildESPNResponse(events))
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
	for _, tm := range []string{"BOS", "LAL", "GSW", "MIA", "NYK", "CHI"} {
		if !strings.Contains(content, tm) {
			t.Errorf("expected content to contain %q", tm)
		}
	}
}

// --- Request header tests ---

func TestNbascore_RequestHeaders(t *testing.T) {
	var userAgent string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buildESPNResponse([]espnEvent{}))
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

	if userAgent != "WTFUtil (+https://wtfutil.com)" {
		t.Errorf("expected User-Agent 'WTFUtil (+https://wtfutil.com)', got %q", userAgent)
	}
}

// --- ConfigText test ---

func TestConfigText(t *testing.T) {
	widget := testWidget(t)
	_ = widget.ConfigText() // just verify no panic
}

// --- Keyboard controls offset tests ---

func TestOffsetControls(t *testing.T) {
	origOffset := offset
	defer func() { offset = origOffset }()

	offset = 0
	offset++
	if offset != 1 {
		t.Errorf("expected offset=1 after next, got %d", offset)
	}

	offset--
	if offset != 0 {
		t.Errorf("expected offset=0 after prev, got %d", offset)
	}

	offset = 5
	offset = 0
	if offset != 0 {
		t.Errorf("expected offset=0 after center, got %d", offset)
	}
}
