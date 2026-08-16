package tennis

import (
	"testing"

	"github.com/olebedev/config"
)

const globalYAML = `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`

func parseConfigs(t *testing.T, moduleYAML string) (*config.Config, *config.Config) {
	t.Helper()

	moduleConfig, err := config.ParseYaml(moduleYAML)
	if err != nil {
		t.Fatalf("failed to parse module yaml: %v", err)
	}

	globalConfig, err := config.ParseYaml(globalYAML)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	return moduleConfig, globalConfig
}

func TestNewSettingsFromYAML(t *testing.T) {
	t.Setenv("WTF_TENNIS_API_KEY", "")

	moduleConfig, globalConfig := parseConfigs(t, `
apiKey: "yaml-key"
baseURL: "https://tennis.example.test/v1"
tour: "wta"
status: "upcoming"
matchLimit: 3
enabled: true
refreshInterval: 60s
position:
  top: 0
  left: 0
  height: 1
  width: 1
`)

	settings := NewSettingsFromYAML("tennis", moduleConfig, globalConfig)

	if settings.apiKey != "yaml-key" {
		t.Errorf("expected apiKey='yaml-key', got %q", settings.apiKey)
	}
	if settings.tour != "wta" {
		t.Errorf("expected tour='wta', got %q", settings.tour)
	}
	if settings.status != "upcoming" {
		t.Errorf("expected status='upcoming', got %q", settings.status)
	}
	if settings.baseURL != "https://tennis.example.test/v1" {
		t.Errorf("expected overridden baseURL, got %q", settings.baseURL)
	}
	if settings.matchLimit != 3 {
		t.Errorf("expected matchLimit=3, got %d", settings.matchLimit)
	}
	if settings.RefreshInterval.Seconds() != 60 {
		t.Errorf("expected refreshInterval=60s, got %v", settings.RefreshInterval)
	}
	if settings.DocPath != "sports/tennis" {
		t.Errorf("expected DocPath='sports/tennis', got %q", settings.DocPath)
	}
}

func TestNewSettingsFromYAML_Defaults(t *testing.T) {
	t.Setenv("WTF_TENNIS_API_KEY", "")

	moduleConfig, globalConfig := parseConfigs(t, `
apiKey: "yaml-key"
position:
  top: 0
  left: 0
  height: 1
  width: 1
`)

	settings := NewSettingsFromYAML("tennis", moduleConfig, globalConfig)

	if settings.tour != "" {
		t.Errorf("expected empty default tour, got %q", settings.tour)
	}
	if settings.baseURL != defaultBaseURL {
		t.Errorf("expected default baseURL %q, got %q", defaultBaseURL, settings.baseURL)
	}
	if settings.status != defaultStatus {
		t.Errorf("expected default status %q, got %q", defaultStatus, settings.status)
	}
	if settings.matchLimit != defaultMatchLimit {
		t.Errorf("expected default matchLimit=%d, got %d", defaultMatchLimit, settings.matchLimit)
	}
	if settings.Title != defaultTitle {
		t.Errorf("expected default title %q, got %q", defaultTitle, settings.Title)
	}
}

func TestNewSettingsFromYAML_APIKeySources(t *testing.T) {
	tests := []struct {
		name       string
		envValue   string
		moduleYAML string
		want       string
	}{
		{
			name:       "environment variable",
			envValue:   "env-key",
			moduleYAML: "enabled: true\n",
			want:       "env-key",
		},
		{
			name:       "yaml overrides the environment",
			envValue:   "env-key",
			moduleYAML: "apiKey: \"yaml-key\"\n",
			want:       "yaml-key",
		},
		{
			name:       "lowercase apikey is accepted",
			envValue:   "env-key",
			moduleYAML: "apikey: \"lower-key\"\n",
			want:       "lower-key",
		},
		{
			name:       "unset everywhere",
			envValue:   "",
			moduleYAML: "enabled: true\n",
			want:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("WTF_TENNIS_API_KEY", tt.envValue)

			moduleConfig, globalConfig := parseConfigs(t, tt.moduleYAML)
			settings := NewSettingsFromYAML("tennis", moduleConfig, globalConfig)

			if settings.apiKey != tt.want {
				t.Errorf("got apiKey %q, want %q", settings.apiKey, tt.want)
			}
		})
	}
}

func TestNormalizeStatus(t *testing.T) {
	tests := []struct{ in, want string }{
		{"live", "live"},
		{"upcoming", "upcoming"},
		{"completed", "live"}, // completed requires a paid plan; clamp to the free-tier surface
		{"bogus", "live"},
		{"LIVE", "live"},
		{"", "live"},
	}

	for _, tt := range tests {
		if got := normalizeStatus(tt.in); got != tt.want {
			t.Errorf("normalizeStatus(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
