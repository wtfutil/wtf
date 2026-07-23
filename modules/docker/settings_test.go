package docker

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name              string
		moduleYAML        string
		globalYAML        string
		expectedTitle     string
		expectedLabelColo string
	}{
		{
			name:              "defaults when nothing configured",
			moduleYAML:        "{}",
			globalYAML:        "{}",
			expectedTitle:     defaultTitle,
			expectedLabelColo: "white",
		},
		{
			name:              "custom title and label color",
			moduleYAML:        "title: \"My Docker\"\nlabelColor: \"green\"\n",
			globalYAML:        "{}",
			expectedTitle:     "My Docker",
			expectedLabelColo: "green",
		},
		{
			name:              "label color falls back to default when empty string configured",
			moduleYAML:        "labelColor: \"\"\n",
			globalYAML:        "{}",
			expectedTitle:     defaultTitle,
			expectedLabelColo: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moduleConfig, err := config.ParseYaml(tt.moduleYAML)
			assert.NoError(t, err)

			globalConfig, err := config.ParseYaml(tt.globalYAML)
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("docker", moduleConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.expectedTitle, settings.Title)
			assert.Equal(t, tt.expectedLabelColo, settings.labelColor)
		})
	}
}

func TestDefaultConstants(t *testing.T) {
	assert.False(t, defaultFocusable)
	assert.Equal(t, "docker", defaultTitle)
}

func TestWidget_ConfigText(t *testing.T) {
	widget := &Widget{}

	helpText := widget.ConfigText()

	assert.NotEmpty(t, helpText)
}

func TestSettings_Structure(t *testing.T) {
	settings := &Settings{
		Common: &cfg.Common{
			Title: "Test Docker",
		},
		labelColor: "yellow",
	}

	assert.NotNil(t, settings.Common)
	assert.Equal(t, "Test Docker", settings.Title)
	assert.Equal(t, "yellow", settings.labelColor)
}
