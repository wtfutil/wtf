package lunarphase

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

// --- helpers ---

const (
	minimalYAML = `
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	globalYAML = `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
)

func parseSettings(t *testing.T, yml string) *Settings {
	t.Helper()
	ymlCfg, err := config.ParseYaml(yml)
	if err != nil {
		t.Fatalf("parse yaml: %v", err)
	}
	globalCfg, err := config.ParseYaml(globalYAML)
	if err != nil {
		t.Fatalf("parse global yaml: %v", err)
	}
	return NewSettingsFromYAML("lunarphase", ymlCfg, globalCfg)
}

func newTestWidget(t *testing.T, settings *Settings) *Widget {
	t.Helper()
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	pages := tview.NewPages()
	w := NewWidget(app, redrawChan, pages, settings)
	return w
}

// --- Settings tests ---

func TestNewSettingsFromYAML_Defaults(t *testing.T) {
	settings := parseSettings(t, minimalYAML)

	if settings.language != "en" {
		t.Errorf("expected default language 'en', got %q", settings.language)
	}
	if settings.requestTimeout != 30 {
		t.Errorf("expected default timeout 30, got %d", settings.requestTimeout)
	}
	if !strings.Contains(settings.Title, "Phase of the Moon") {
		t.Errorf("expected title to contain 'Phase of the Moon', got %q", settings.Title)
	}
}

func TestNewSettingsFromYAML_Custom(t *testing.T) {
	yml := `
language: "fr"
timeout: 10
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	settings := parseSettings(t, yml)

	if settings.language != "fr" {
		t.Errorf("expected language 'fr', got %q", settings.language)
	}
	if settings.requestTimeout != 10 {
		t.Errorf("expected timeout 10, got %d", settings.requestTimeout)
	}
}

// --- Date format tests ---

func TestDateFormat(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		format   string
		expected string
	}{
		{"dateFormat standard", time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), dateFormat, "2024-03-15"},
		{"dateFormat new year", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), dateFormat, "2025-01-01"},
		{"phaseFormat standard", time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), phaseFormat, "03-15-2024"},
		{"phaseFormat end of year", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), phaseFormat, "12-31-2024"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.date.Format(tc.format)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

// --- lunarPhase HTTP tests ---

func TestLunarPhase_Success(t *testing.T) {
	responseBody := "  🌕 Full Moon\n  Illumination: 100%  \n"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/Moon@") {
			t.Errorf("expected path to contain /Moon@, got %s", r.URL.Path)
		}
		if r.Header.Get("User-Agent") != "curl" {
			t.Errorf("expected User-Agent 'curl', got %q", r.Header.Get("User-Agent"))
		}
		if !strings.Contains(r.URL.RawQuery, "lang=en") {
			t.Errorf("expected query to contain lang=en, got %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseBody))
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.day = "2024-03-15"

	w.lunarPhase()

	if w.result == "" {
		t.Fatal("expected non-empty result")
	}
	if w.last != "2024-03-15" {
		t.Errorf("expected last to be set to '2024-03-15', got %q", w.last)
	}
}

func TestLunarPhase_LanguageHeader(t *testing.T) {
	tests := []struct {
		name     string
		language string
	}{
		{"english", "en"},
		{"french", "fr"},
		{"german", "de"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				got := r.Header.Get("Accept-Language")
				if got != tc.language {
					t.Errorf("expected Accept-Language %q, got %q", tc.language, got)
				}
				if !strings.Contains(r.URL.RawQuery, "lang="+tc.language) {
					t.Errorf("expected query to contain lang=%s, got %s", tc.language, r.URL.RawQuery)
				}
				_, _ = w.Write([]byte("moon data"))
			}))
			defer srv.Close()

			yml := `
language: "` + tc.language + `"
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
			settings := parseSettings(t, yml)
			widget := newTestWidget(t, settings)
			widget.baseURL = srv.URL
			widget.day = "2024-01-01"

			widget.lunarPhase()

			if widget.result == "" {
				t.Error("expected non-empty result")
			}
		})
	}
}

func TestLunarPhase_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.day = "2024-03-15"

	w.lunarPhase()

	// The widget stores whatever the server returns (no status check in source)
	if w.last != "2024-03-15" {
		t.Errorf("expected last='2024-03-15', got %q", w.last)
	}
}

func TestLunarPhase_ConnectionRefused(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1" // unreachable port
	w.day = "2024-03-15"

	w.lunarPhase()

	if w.result == "" {
		t.Fatal("expected error message in result")
	}
	if w.last == "2024-03-15" {
		t.Error("last should not be set on connection error")
	}
}

func TestLunarPhase_Timeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		_, _ = w.Write([]byte("delayed"))
	}))
	defer srv.Close()

	yml := `
timeout: 1
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	settings := parseSettings(t, yml)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.day = "2024-03-15"
	w.timeout = 100 * time.Millisecond

	w.lunarPhase()

	if w.result == "" {
		t.Fatal("expected error in result due to timeout")
	}
	if w.last == "2024-03-15" {
		t.Error("last should not be set on timeout error")
	}
}

func TestLunarPhase_EmptyBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.day = "2024-03-15"

	w.lunarPhase()

	if w.last != "2024-03-15" {
		t.Errorf("expected last='2024-03-15', got %q", w.last)
	}
	if w.result != "" {
		t.Errorf("expected empty result for empty body, got %q", w.result)
	}
}

func TestLunarPhase_ResponseWithANSI(t *testing.T) {
	ansiContent := "\033[1m\033[38;5;220mFull Moon\033[0m"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(ansiContent))
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.day = "2024-03-15"

	w.lunarPhase()

	// ASCIItoTviewColors should convert ANSI codes
	if strings.Contains(w.result, "\033[") {
		t.Error("expected ANSI codes to be converted")
	}
	if !strings.Contains(w.result, "Full Moon") {
		t.Errorf("expected result to contain 'Full Moon', got %q", w.result)
	}
}

func TestLunarPhase_CachesResult(t *testing.T) {
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		_, _ = w.Write([]byte("moon data"))
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.day = "2024-03-15"
	w.current = false

	// First call should fetch
	w.Refresh()
	if callCount != 1 {
		t.Errorf("expected 1 call, got %d", callCount)
	}

	// Second call with same day should not fetch again
	w.Refresh()
	if callCount != 1 {
		t.Errorf("expected still 1 call (cached), got %d", callCount)
	}

	// Change day, should fetch again
	w.day = "2024-03-16"
	w.Refresh()
	if callCount != 2 {
		t.Errorf("expected 2 calls after day change, got %d", callCount)
	}
}

// --- Navigation tests ---

func TestNextDay(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1" // won't actually connect in setDay path
	w.date = time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	w.day = w.date.Format(dateFormat)

	w.NextDay()

	if w.current {
		t.Error("expected current to be false after NextDay")
	}
	expectedDay := "2024-03-16"
	if w.day != expectedDay {
		t.Errorf("expected day %q, got %q", expectedDay, w.day)
	}
}

func TestPrevDay(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1"
	w.date = time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	w.day = w.date.Format(dateFormat)

	w.PrevDay()

	if w.current {
		t.Error("expected current to be false after PrevDay")
	}
	expectedDay := "2024-03-14"
	if w.day != expectedDay {
		t.Errorf("expected day %q, got %q", expectedDay, w.day)
	}
}

func TestNextWeek(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1"
	w.date = time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	w.day = w.date.Format(dateFormat)

	w.NextWeek()

	if w.current {
		t.Error("expected current to be false after NextWeek")
	}
	expectedDay := "2024-03-22"
	if w.day != expectedDay {
		t.Errorf("expected day %q, got %q", expectedDay, w.day)
	}
}

func TestPrevWeek(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1"
	w.date = time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	w.day = w.date.Format(dateFormat)

	w.PrevWeek()

	if w.current {
		t.Error("expected current to be false after PrevWeek")
	}
	expectedDay := "2024-03-08"
	if w.day != expectedDay {
		t.Errorf("expected day %q, got %q", expectedDay, w.day)
	}
}

func TestNavigationTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		startDate   time.Time
		action      string
		expectedDay string
	}{
		{"next day from month end", time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC), "next", "2024-02-01"},
		{"prev day from month start", time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), "prev", "2024-02-29"},
		{"next week across month", time.Date(2024, 2, 25, 0, 0, 0, 0, time.UTC), "nextweek", "2024-03-03"},
		{"prev week across month", time.Date(2024, 3, 5, 0, 0, 0, 0, time.UTC), "prevweek", "2024-02-27"},
		{"next day year boundary", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), "next", "2025-01-01"},
		{"prev day year boundary", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "prev", "2024-12-31"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			settings := parseSettings(t, minimalYAML)
			w := newTestWidget(t, settings)
			w.baseURL = "http://127.0.0.1:1"
			w.date = tc.startDate
			w.day = w.date.Format(dateFormat)

			switch tc.action {
			case "next":
				w.NextDay()
			case "prev":
				w.PrevDay()
			case "nextweek":
				w.NextWeek()
			case "prevweek":
				w.PrevWeek()
			}

			if w.day != tc.expectedDay {
				t.Errorf("expected day %q, got %q", tc.expectedDay, w.day)
			}
			if w.current {
				t.Error("expected current=false after navigation")
			}
		})
	}
}

// --- Disable/Enable tests ---

func TestDisableWidget(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1"

	if !w.settings.Enabled {
		t.Fatal("expected widget to start enabled")
	}

	w.DisableWidget()
	if w.settings.Enabled {
		t.Error("expected widget to be disabled after DisableWidget")
	}

	w.DisableWidget()
	if !w.settings.Enabled {
		t.Error("expected widget to be re-enabled after second DisableWidget")
	}
}

// --- URL construction test ---

func TestLunarPhase_URLConstruction(t *testing.T) {
	tests := []struct {
		name     string
		day      string
		language string
		wantPath string
		wantLang string
	}{
		{"standard", "2024-03-15", "en", "/Moon@2024-03-15", "lang=en"},
		{"french", "2024-12-25", "fr", "/Moon@2024-12-25", "lang=fr"},
		{"german", "2025-01-01", "de", "/Moon@2025-01-01", "lang=de"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var gotPath, gotQuery string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.Path
				gotQuery = r.URL.RawQuery
				_, _ = w.Write([]byte("ok"))
			}))
			defer srv.Close()

			yml := `
language: "` + tc.language + `"
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
			settings := parseSettings(t, yml)
			widget := newTestWidget(t, settings)
			widget.baseURL = srv.URL
			widget.day = tc.day

			widget.lunarPhase()

			if gotPath != tc.wantPath {
				t.Errorf("expected path %q, got %q", tc.wantPath, gotPath)
			}
			if !strings.Contains(gotQuery, tc.wantLang) {
				t.Errorf("expected query to contain %q, got %q", tc.wantLang, gotQuery)
			}
		})
	}
}

// --- OpenMoonPhase URL format test ---

func TestOpenMoonPhase_URLFormat(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{"march 2024", time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), "03-15-2024"},
		{"january 2025", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "01-01-2025"},
		{"december 2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), "12-31-2024"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.date.Format(phaseFormat)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

// --- Widget initialization tests ---

func TestNewWidget_Initialization(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)

	if !w.current {
		t.Error("expected current=true on init")
	}
	if w.day == "" {
		t.Error("expected day to be set on init")
	}
	if w.last != "" {
		t.Error("expected last to be empty on init")
	}
	if w.timeout != 30*time.Second {
		t.Errorf("expected timeout=30s, got %v", w.timeout)
	}
	if w.baseURL != "https://wttr.in" {
		t.Errorf("expected baseURL='https://wttr.in', got %q", w.baseURL)
	}
}

func TestNewWidget_CustomTimeout(t *testing.T) {
	yml := `
timeout: 60
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	settings := parseSettings(t, yml)
	w := newTestWidget(t, settings)

	if w.timeout != 60*time.Second {
		t.Errorf("expected timeout=60s, got %v", w.timeout)
	}
}

// --- Today test ---

func TestToday(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("today moon"))
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.current = false
	w.date = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	w.day = "2020-01-01"

	w.Today()

	if !w.current {
		t.Error("expected current=true after Today()")
	}
	// date should be updated to today
	today := time.Now().Format(dateFormat)
	if w.day != today {
		t.Errorf("expected day=%q, got %q", today, w.day)
	}
}

// --- Refresh when disabled ---

func TestRefresh_Disabled(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	settings.Enabled = false
	w := newTestWidget(t, settings)
	w.baseURL = "http://127.0.0.1:1"
	w.current = false
	w.day = "2024-03-15"

	w.Refresh()

	if !strings.Contains(w.settings.Title, "Disabled") {
		t.Errorf("expected title to contain 'Disabled', got %q", w.settings.Title)
	}
}

func TestRefresh_CurrentUpdatesDay(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("moon"))
	}))
	defer srv.Close()

	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)
	w.baseURL = srv.URL
	w.current = true
	w.day = "2020-01-01" // old date
	w.last = ""

	w.Refresh()

	today := time.Now().Format(dateFormat)
	if w.day != today {
		t.Errorf("expected day to be updated to today %q, got %q", today, w.day)
	}
}

// --- ConfigText test ---

func TestConfigText(t *testing.T) {
	settings := parseSettings(t, minimalYAML)
	w := newTestWidget(t, settings)

	text := w.ConfigText()
	if text == "" {
		t.Error("expected non-empty config text")
	}
}
