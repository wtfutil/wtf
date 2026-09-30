package arpansagovau

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

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		wantCity string
	}{
		{
			name: "with locationid set",
			yaml: `
locationid: syd
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			wantCity: "syd",
		},
		{
			name: "with different city",
			yaml: `
locationid: adl
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			wantCity: "adl",
		},
		{
			name: "without locationid",
			yaml: `
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			wantCity: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tc.yaml)
			if err != nil {
				t.Fatalf("failed to parse yaml: %v", err)
			}

			globalConfig, err := config.ParseYaml(globalYAML)
			if err != nil {
				t.Fatalf("failed to parse global yaml: %v", err)
			}

			settings := NewSettingsFromYAML("arpansagovau", ymlConfig, globalConfig)

			if settings.city != tc.wantCity {
				t.Errorf("expected city %q, got %q", tc.wantCity, settings.city)
			}
			if settings.Common == nil {
				t.Fatal("expected Common to be non-nil")
			}
		})
	}
}

func TestSettingsDefaults(t *testing.T) {
	if defaultFocusable != false {
		t.Errorf("expected defaultFocusable to be false")
	}
	if defaultTitle != "ARPANSA UV Data" {
		t.Errorf("expected defaultTitle to be 'ARPANSA UV Data', got %q", defaultTitle)
	}
}
