package devto

import (
	"strings"
	"testing"

	"github.com/olebedev/config"
)

func TestNewSettingsFromYAML(t *testing.T) {
	yamlStr := `
numberOfArticles: 5
contentTag: "golang"
contentUsername: "testuser"
contentState: "fresh"
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
`

	yamlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}

	globalConfig, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	settings := NewSettingsFromYAML("devto", yamlConfig, globalConfig)

	if settings.numberOfArticles != 5 {
		t.Errorf("expected numberOfArticles=5, got %d", settings.numberOfArticles)
	}
	if settings.contentTag != "golang" {
		t.Errorf("expected contentTag='golang', got %q", settings.contentTag)
	}
	if settings.contentUsername != "testuser" {
		t.Errorf("expected contentUsername='testuser', got %q", settings.contentUsername)
	}
	if settings.contentState != "fresh" {
		t.Errorf("expected contentState='fresh', got %q", settings.contentState)
	}
}

func TestNewSettingsFromYAML_Defaults(t *testing.T) {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`

	yamlConfig, _ := config.ParseYaml(yamlStr)
	globalConfig, _ := config.ParseYaml(globalStr)

	settings := NewSettingsFromYAML("devto", yamlConfig, globalConfig)

	if settings.numberOfArticles != 10 {
		t.Errorf("expected default numberOfArticles=10, got %d", settings.numberOfArticles)
	}
	if settings.contentTag != "" {
		t.Errorf("expected empty default contentTag, got %q", settings.contentTag)
	}
	if settings.contentState != "" {
		t.Errorf("expected empty default contentState, got %q", settings.contentState)
	}
	if !strings.Contains(settings.Title, "dev.to") {
		t.Errorf("expected title to contain 'dev.to', got %q", settings.Title)
	}
}
