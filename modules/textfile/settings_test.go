package textfile

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

const testGlobalConfig = "wtf:\n  colors:\n    theme: default\n"

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name                string
		yaml                string
		expectedTitle       string
		expectedFormat      bool
		expectedFormatStyle string
		expectedWrapText    bool
		expectedFilePaths   int
	}{
		{
			name:                "defaults when nothing configured",
			yaml:                "{}",
			expectedTitle:       defaultTitle,
			expectedFormat:      false,
			expectedFormatStyle: "vim",
			expectedWrapText:    true,
			expectedFilePaths:   0,
		},
		{
			name: "explicit values are honored",
			yaml: "" +
				"title: \"My Text Files\"\n" +
				"format: true\n" +
				"formatStyle: monokai\n" +
				"wrapText: false\n" +
				"filePaths:\n" +
				"  - /tmp/a.txt\n" +
				"  - /tmp/b.txt\n",
			expectedTitle:       "My Text Files",
			expectedFormat:      true,
			expectedFormatStyle: "monokai",
			expectedWrapText:    false,
			expectedFilePaths:   2,
		},
		{
			name: "format explicitly false",
			yaml: "" +
				"format: false\n" +
				"formatStyle: dracula\n",
			expectedTitle:       defaultTitle,
			expectedFormat:      false,
			expectedFormatStyle: "dracula",
			expectedWrapText:    true,
			expectedFilePaths:   0,
		},
		{
			name:                "single filePath is not counted in filePaths list",
			yaml:                "filePath: /tmp/only.txt\n",
			expectedTitle:       defaultTitle,
			expectedFormat:      false,
			expectedFormatStyle: "vim",
			expectedWrapText:    true,
			expectedFilePaths:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NoError(t, err)

			globalConfig, err := config.ParseYaml(testGlobalConfig)
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("textfile", ymlConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.expectedTitle, settings.Title)
			assert.Equal(t, tt.expectedFormat, settings.format)
			assert.Equal(t, tt.expectedFormatStyle, settings.formatStyle)
			assert.Equal(t, tt.expectedWrapText, settings.wrapText)
			assert.Len(t, settings.filePaths, tt.expectedFilePaths)
		})
	}
}

func TestConfigText(t *testing.T) {
	widget := &Widget{}
	text := widget.ConfigText()
	assert.NotEmpty(t, text)
}
