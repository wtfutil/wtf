package ipapi

import (
	"testing"

	"github.com/olebedev/config"
)

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
	yamlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}
	globalConfig, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	settings := NewSettingsFromYAML("ipapi", yamlConfig, globalConfig)

	if settings.name != "red" {
		t.Errorf("expected default name color 'red', got %q", settings.name)
	}
	if settings.value != "white" {
		t.Errorf("expected default value color 'white', got %q", settings.value)
	}
	if settings.Common == nil {
		t.Fatal("expected Common to be set")
	}
	if len(settings.args) != 0 {
		t.Errorf("expected empty args by default, got %v", settings.args)
	}
}

func TestNewSettingsFromYAML_CustomColors(t *testing.T) {
	yamlStr := `
colors:
  name: "blue"
  value: "green"
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

	settings := NewSettingsFromYAML("ipapi", yamlConfig, globalConfig)

	if settings.name != "blue" {
		t.Errorf("expected name color 'blue', got %q", settings.name)
	}
	if settings.value != "green" {
		t.Errorf("expected value color 'green', got %q", settings.value)
	}
}

func TestNewSettingsFromYAML_WithArgs(t *testing.T) {
	yamlStr := `
args:
  - ip
  - city
  - country
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

	settings := NewSettingsFromYAML("ipapi", yamlConfig, globalConfig)

	if len(settings.args) != 3 {
		t.Fatalf("expected 3 args, got %d", len(settings.args))
	}
}

func TestNewSettingsFromYAML_Constants(t *testing.T) {
	if defaultFocusable != false {
		t.Error("expected defaultFocusable to be false")
	}
	if defaultTitle != "IP API" {
		t.Errorf("expected defaultTitle 'IP API', got %q", defaultTitle)
	}
}
