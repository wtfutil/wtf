package progress

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name          string
		yaml          string
		expectedTitle string
		expectedMax   float64
		expectedShow  string
	}{
		{
			name:          "defaults when no config provided",
			yaml:          "{}",
			expectedTitle: defaultTitle,
			expectedMax:   0,
			expectedShow:  "right",
		},
		{
			name: "custom values override defaults",
			yaml: `
title: "Custom Progress"
showPercentage: "left"
padding: 2
minimum: 10
maximum: 200
current: 50
colors:
  gradientA: "#111111"
  gradientB: "#222222"
  solid: "#333333"
`,
			expectedTitle: "Custom Progress",
			expectedMax:   200,
			expectedShow:  "left",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NoError(t, err)

			globalConfig, err := config.ParseYaml("wtf: {}")
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("progress", ymlConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.Equal(t, tt.expectedTitle, settings.common.Title)
			assert.Equal(t, tt.expectedMax, settings.maximum)
			assert.Equal(t, tt.expectedShow, settings.showPercentage)
		})
	}
}

func TestNewSettingsFromYAML_Colors(t *testing.T) {
	yaml := `
colors:
  gradientA: "#111111"
  gradientB: "#222222"
  solid: "#333333"
`
	ymlConfig, err := config.ParseYaml(yaml)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("progress", ymlConfig, globalConfig)

	assert.Equal(t, "#111111", settings.gradientA)
	assert.Equal(t, "#222222", settings.gradientB)
	assert.Equal(t, "#333333", settings.solid)
}

func TestNewSettingsFromYAML_ColorDefaults(t *testing.T) {
	ymlConfig, err := config.ParseYaml("{}")
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("progress", ymlConfig, globalConfig)

	assert.Equal(t, "#56ab2f", settings.gradientA)
	assert.Equal(t, "#a8e063", settings.gradientB)
	assert.Equal(t, "", settings.solid)
}

func TestConfigText(t *testing.T) {
	widget := createTestWidget(0, 100, 0, "right", 1)

	text := widget.ConfigText()

	assert.NotEmpty(t, text)
}
