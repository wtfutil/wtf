package cmdrunner

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSettingsFromYAML(t *testing.T) {
	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	tests := []struct {
		name               string
		yaml               string
		expectedCmd        string
		expectedArgs       []string
		expectedTail       bool
		expectedPty        bool
		expectedPtyErrors  bool
		expectedMaxLines   int
		expectedWorkingDir string
	}{
		{
			name:               "defaults",
			yaml:               `{}`,
			expectedCmd:        "",
			expectedArgs:       []string{},
			expectedTail:       false,
			expectedPty:        false,
			expectedPtyErrors:  false,
			expectedMaxLines:   256,
			expectedWorkingDir: ".",
		},
		{
			name: "fully specified",
			yaml: `
cmd: "echo"
args: ["hello", "world"]
tail: true
pty: true
ptySuppressErrors: true
maxLines: 42
workingDir: "/tmp"
`,
			expectedCmd:        "echo",
			expectedArgs:       []string{"hello", "world"},
			expectedTail:       true,
			expectedPty:        true,
			expectedPtyErrors:  true,
			expectedMaxLines:   42,
			expectedWorkingDir: "/tmp",
		},
		{
			name: "cmd without args",
			yaml: `
cmd: "whoami"
`,
			expectedCmd:        "whoami",
			expectedArgs:       []string{},
			expectedTail:       false,
			expectedPty:        false,
			expectedPtyErrors:  false,
			expectedMaxLines:   256,
			expectedWorkingDir: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moduleConfig, err := config.ParseYaml(tt.yaml)
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("cmdrunner", moduleConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.expectedCmd, settings.cmd)
			assert.Equal(t, tt.expectedArgs, settings.args)
			assert.Equal(t, tt.expectedTail, settings.tail)
			assert.Equal(t, tt.expectedPty, settings.pty)
			assert.Equal(t, tt.expectedPtyErrors, settings.ptySuppressErrors)
			assert.Equal(t, tt.expectedMaxLines, settings.maxLines)
			assert.Equal(t, tt.expectedWorkingDir, settings.workingDir)
		})
	}
}

func TestNewSettingsFromYAML_DefaultTitle(t *testing.T) {
	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	moduleConfig, err := config.ParseYaml("{}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("cmdrunner", moduleConfig, globalConfig)

	assert.Equal(t, defaultTitle, settings.Title)
	assert.True(t, defaultFocusable)
}

func TestConfigText(t *testing.T) {
	widget := &Widget{}

	text := widget.ConfigText()

	assert.NotEmpty(t, text)
	assert.Contains(t, text, "cmd")
	assert.Contains(t, text, "args")
}
