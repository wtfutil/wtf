package prettyweather

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

func testSettings(overrides map[string]string) *Settings {
	city := "London"
	unit := "m"
	view := "0"
	language := "en"

	if v, ok := overrides["city"]; ok {
		city = v
	}
	if v, ok := overrides["unit"]; ok {
		unit = v
	}
	if v, ok := overrides["view"]; ok {
		view = v
	}
	if v, ok := overrides["language"]; ok {
		language = v
	}

	ymlContent := "prettyweather:\n" +
		"  enabled: true\n" +
		"  position:\n" +
		"    top: 0\n" +
		"    left: 0\n" +
		"    height: 3\n" +
		"    width: 3\n" +
		"  city: " + city + "\n" +
		"  unit: " + unit + "\n" +
		"  view: " + view + "\n" +
		"  language: " + language + "\n"

	ymlConfig, _ := config.ParseYaml(ymlContent)
	globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    border:\n      focusable: darkslateblue\n      focused: orange\n      normal: gray\n")

	moduleConfig, _ := ymlConfig.Get("prettyweather")

	return &Settings{
		Common:   cfg.NewCommonSettingsFromModule("prettyweather", defaultTitle, defaultFocusable, moduleConfig, globalConfig),
		city:     city,
		unit:     unit,
		view:     view,
		language: language,
	}
}

func newTestWidget(serverURL string, overrides map[string]string) *Widget {
	settings := testSettings(overrides)
	widget := &Widget{
		settings: settings,
		baseURL:  serverURL + "/",
	}
	return widget
}

func TestPrettyWeather_SuccessfulResponse(t *testing.T) {
	tests := []struct {
		name         string
		responseBody string
		overrides    map[string]string
		wantContains string
	}{
		{
			name:         "basic weather response",
			responseBody: "Sunny 25°C\nWind: 10 km/h",
			overrides:    nil,
			wantContains: "Sunny 25°C",
		},
		{
			name:         "response with leading/trailing whitespace is trimmed",
			responseBody: "  \n  Cloudy 18°C  \n  ",
			overrides:    nil,
			wantContains: "Cloudy 18°C",
		},
		{
			name:         "empty response",
			responseBody: "",
			overrides:    nil,
			wantContains: "",
		},
		{
			name:         "multiline forecast",
			responseBody: "Monday: Rain 12°C\nTuesday: Sun 20°C\nWednesday: Cloud 15°C",
			overrides:    nil,
			wantContains: "Tuesday: Sun 20°C",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			overrides := tt.overrides
			if overrides == nil {
				overrides = map[string]string{}
			}
			widget := newTestWidget(server.URL, overrides)
			widget.prettyWeather()

			if tt.wantContains != "" && !strings.Contains(widget.result, tt.wantContains) {
				t.Errorf("result = %q, want it to contain %q", widget.result, tt.wantContains)
			}
		})
	}
}

func TestPrettyWeather_RequestParameters(t *testing.T) {
	tests := []struct {
		name           string
		overrides      map[string]string
		wantPath       string
		wantLangHeader string
		wantUAHeader   string
	}{
		{
			name:           "default settings build correct URL",
			overrides:      map[string]string{},
			wantPath:       "/London?0?m",
			wantLangHeader: "en",
			wantUAHeader:   "curl",
		},
		{
			name:           "custom city in URL",
			overrides:      map[string]string{"city": "Paris"},
			wantPath:       "/Paris?0?m",
			wantLangHeader: "en",
			wantUAHeader:   "curl",
		},
		{
			name:           "custom unit in URL",
			overrides:      map[string]string{"unit": "u"},
			wantPath:       "/London?0?u",
			wantLangHeader: "en",
			wantUAHeader:   "curl",
		},
		{
			name:           "custom view in URL",
			overrides:      map[string]string{"view": "1"},
			wantPath:       "/London?1?m",
			wantLangHeader: "en",
			wantUAHeader:   "curl",
		},
		{
			name:           "custom language header",
			overrides:      map[string]string{"language": "fr"},
			wantPath:       "/London?0?m",
			wantLangHeader: "fr",
			wantUAHeader:   "curl",
		},
		{
			name:           "all custom settings",
			overrides:      map[string]string{"city": "Tokyo", "unit": "u", "view": "2", "language": "ja"},
			wantPath:       "/Tokyo?2?u",
			wantLangHeader: "ja",
			wantUAHeader:   "curl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotPath string
			var gotLang string
			var gotUA string

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.RequestURI()
				gotLang = r.Header.Get("Accept-Language")
				gotUA = r.Header.Get("User-Agent")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("OK"))
			}))
			defer server.Close()

			widget := newTestWidget(server.URL, tt.overrides)
			widget.prettyWeather()

			if gotPath != tt.wantPath {
				t.Errorf("request path = %q, want %q", gotPath, tt.wantPath)
			}
			if gotLang != tt.wantLangHeader {
				t.Errorf("Accept-Language = %q, want %q", gotLang, tt.wantLangHeader)
			}
			if gotUA != tt.wantUAHeader {
				t.Errorf("User-Agent = %q, want %q", gotUA, tt.wantUAHeader)
			}
		})
	}
}

func TestPrettyWeather_ErrorPaths(t *testing.T) {
	tests := []struct {
		name        string
		setupServer func() *httptest.Server
		baseURL     string
		wantErr     bool
	}{
		{
			name: "server returns 500",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write([]byte("Internal Server Error"))
				}))
			},
			wantErr: false, // still reads body, no Go-level error
		},
		{
			name:        "connection refused (invalid URL)",
			setupServer: nil,
			baseURL:     "http://127.0.0.1:1/",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var baseURL string
			if tt.setupServer != nil {
				server := tt.setupServer()
				defer server.Close()
				baseURL = server.URL + "/"
			} else {
				baseURL = tt.baseURL
			}

			widget := newTestWidget("placeholder", map[string]string{})
			widget.baseURL = baseURL
			widget.prettyWeather()

			if tt.wantErr && widget.result == "" {
				t.Error("expected non-empty error result, got empty string")
			}
			if tt.wantErr && widget.result == "OK" {
				t.Error("expected error in result, got success response")
			}
		})
	}
}

func TestPrettyWeather_InvalidBaseURL(t *testing.T) {
	widget := newTestWidget("placeholder", map[string]string{})
	widget.baseURL = "://invalid-url"
	widget.prettyWeather()

	if widget.result == "" {
		t.Error("expected error result for invalid URL, got empty string")
	}
}

func TestPrettyWeather_ResponseWithANSICodes(t *testing.T) {
	// ANSI escape codes that wttr.in typically returns
	ansiResponse := "\033[38;5;226m☀\033[0m Sunny +25°C"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(ansiResponse))
	}))
	defer server.Close()

	widget := newTestWidget(server.URL, map[string]string{})
	widget.prettyWeather()

	// The result should have been processed through ASCIItoTviewColors
	// and should not contain raw ANSI escape sequences
	if strings.Contains(widget.result, "\033[") {
		t.Error("result still contains raw ANSI escape codes after processing")
	}
}

func TestPrettyWeather_ServerClosedBeforeRead(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Close connection immediately by hijacking
		hj, ok := w.(http.Hijacker)
		if !ok {
			w.WriteHeader(http.StatusOK)
			return
		}
		conn, _, _ := hj.Hijack()
		_ = conn.Close()
	}))
	defer server.Close()

	widget := newTestWidget(server.URL, map[string]string{})
	widget.prettyWeather()

	// Should get an error in result since the connection was forcefully closed
	if widget.result == "" {
		t.Error("expected error result when server closes connection, got empty string")
	}
}

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name         string
		yml          string
		wantCity     string
		wantUnit     string
		wantView     string
		wantLanguage string
	}{
		{
			name: "all values specified",
			yml: "prettyweather:\n  enabled: true\n  position:\n    top: 0\n    left: 0\n    height: 3\n    width: 3\n" +
				"  city: Berlin\n  unit: u\n  view: 2\n  language: de\n",
			wantCity:     "Berlin",
			wantUnit:     "u",
			wantView:     "2",
			wantLanguage: "de",
		},
		{
			name:         "defaults when nothing specified",
			yml:          "prettyweather:\n  enabled: true\n  position:\n    top: 0\n    left: 0\n    height: 3\n    width: 3\n",
			wantCity:     "Barcelona",
			wantUnit:     "m",
			wantView:     "0",
			wantLanguage: "en",
		},
		{
			name: "partial overrides",
			yml: "prettyweather:\n  enabled: true\n  position:\n    top: 0\n    left: 0\n    height: 3\n    width: 3\n" +
				"  city: NYC\n  language: es\n",
			wantCity:     "NYC",
			wantUnit:     "m",
			wantView:     "0",
			wantLanguage: "es",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yml)
			if err != nil {
				t.Fatalf("failed to parse YAML: %v", err)
			}

			globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    border:\n      focusable: darkslateblue\n      focused: orange\n      normal: gray\n")

			moduleConfig, _ := ymlConfig.Get("prettyweather")
			settings := NewSettingsFromYAML("prettyweather", moduleConfig, globalConfig)

			if settings.city != tt.wantCity {
				t.Errorf("city = %q, want %q", settings.city, tt.wantCity)
			}
			if settings.unit != tt.wantUnit {
				t.Errorf("unit = %q, want %q", settings.unit, tt.wantUnit)
			}
			if settings.view != tt.wantView {
				t.Errorf("view = %q, want %q", settings.view, tt.wantView)
			}
			if settings.language != tt.wantLanguage {
				t.Errorf("language = %q, want %q", settings.language, tt.wantLanguage)
			}
		})
	}
}
