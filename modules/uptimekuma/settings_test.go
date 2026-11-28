package uptimekuma

import (
	"testing"

	"github.com/olebedev/config"
)

func TestNewSettingsFromYAML(t *testing.T) {
	yamlStr := `
url: "https://uptime.example.com/status/overview"
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	ymlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}

	globalYaml := `
wtf:
  colors:
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray
`
	globalConfig, err := config.ParseYaml(globalYaml)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	settings := NewSettingsFromYAML("uptimekuma", ymlConfig, globalConfig)
	if settings.url != "https://uptime.example.com/status/overview" {
		t.Errorf("expected url to be set, got %q", settings.url)
	}
	if settings.common == nil {
		t.Fatal("expected common settings to be non-nil")
	}
}

func TestNewSettingsFromYAML_EmptyURL(t *testing.T) {
	yamlStr := `
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	ymlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}

	globalYaml := `
wtf:
  colors:
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray
`
	globalConfig, err := config.ParseYaml(globalYaml)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	settings := NewSettingsFromYAML("uptimekuma", ymlConfig, globalConfig)
	if settings.url != "" {
		t.Errorf("expected empty url, got %q", settings.url)
	}
}

func TestDefaultConstants(t *testing.T) {
	if defaultFocusable != true {
		t.Errorf("expected defaultFocusable=true, got %v", defaultFocusable)
	}
	if defaultTitle != "Uptime Kuma" {
		t.Errorf("expected defaultTitle='Uptime Kuma', got %q", defaultTitle)
	}
}
