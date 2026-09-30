package pihole

import (
	"testing"

	"github.com/olebedev/config"

	"github.com/wtfutil/wtf/cfg"
)

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name               string
		yamlStr            string
		wantTitle          string
		wantAPIURL         string
		wantShowSummary    bool
		wantShowTopItems   int
		wantShowTopClients int
	}{
		{
			name:               "defaults",
			yamlStr:            "{}",
			wantTitle:          defaultTitle,
			wantAPIURL:         "",
			wantShowSummary:    true,
			wantShowTopItems:   5,
			wantShowTopClients: 5,
		},
		{
			name: "custom values",
			yamlStr: `
apiUrl: "http://pi.hole/admin/api.php"
showSummary: false
showTopItems: 10
showTopClients: 3
`,
			wantTitle:          defaultTitle,
			wantAPIURL:         "http://pi.hole/admin/api.php",
			wantShowSummary:    false,
			wantShowTopItems:   10,
			wantShowTopClients: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yamlStr)
			if err != nil {
				t.Fatalf("failed to parse test YAML: %v", err)
			}

			globalConfig, err := config.ParseYaml("wtf: {}")
			if err != nil {
				t.Fatalf("failed to parse global YAML: %v", err)
			}

			settings := NewSettingsFromYAML("pihole", ymlConfig, globalConfig)

			if settings == nil {
				t.Fatal("NewSettingsFromYAML() returned nil")
			}

			if settings.Title != tt.wantTitle {
				t.Errorf("Title = %q, want %q", settings.Title, tt.wantTitle)
			}

			if settings.apiUrl != tt.wantAPIURL {
				t.Errorf("apiUrl = %q, want %q", settings.apiUrl, tt.wantAPIURL)
			}

			if settings.showSummary != tt.wantShowSummary {
				t.Errorf("showSummary = %v, want %v", settings.showSummary, tt.wantShowSummary)
			}

			if settings.showTopItems != tt.wantShowTopItems {
				t.Errorf("showTopItems = %d, want %d", settings.showTopItems, tt.wantShowTopItems)
			}

			if settings.showTopClients != tt.wantShowTopClients {
				t.Errorf("showTopClients = %d, want %d", settings.showTopClients, tt.wantShowTopClients)
			}
		})
	}
}

func TestConfigText(t *testing.T) {
	widget := &Widget{
		settings: &Settings{Common: &cfg.Common{Title: "Test"}},
	}

	got := widget.ConfigText()
	if got == "" {
		t.Error("ConfigText() returned empty string, want help text")
	}
}
