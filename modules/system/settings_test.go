package system

import (
	"testing"

	"github.com/olebedev/config"
	"gotest.tools/assert"
)

func Test_NewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name          string
		yaml          string
		expectedTitle string
	}{
		{
			name:          "with no title set",
			yaml:          "{}",
			expectedTitle: defaultTitle,
		},
		{
			name:          "with custom title",
			yaml:          "title: \"Custom System\"",
			expectedTitle: "Custom System",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NilError(t, err)

			globalConfig, err := config.ParseYaml("{}")
			assert.NilError(t, err)

			settings := NewSettingsFromYAML("system", ymlConfig, globalConfig)

			assert.Assert(t, settings != nil)
			assert.Assert(t, settings.Common != nil)
			assert.Equal(t, tt.expectedTitle, settings.Title)
			assert.Equal(t, defaultFocusable, settings.Focusable)
		})
	}
}
