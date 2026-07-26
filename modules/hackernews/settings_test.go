package hackernews

import (
	"testing"

	"github.com/olebedev/config"
	"gotest.tools/assert"
)

func TestNewSettingsFromYAML_Defaults(t *testing.T) {
	ymlConfig, err := config.ParseYaml("{}")
	assert.NilError(t, err)

	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NilError(t, err)

	settings := NewSettingsFromYAML("hackernews", ymlConfig, globalConfig)

	assert.Equal(t, 10, settings.numberOfStories)
	assert.Equal(t, "top", settings.storyType)
	assert.Equal(t, "HackerNews", settings.Title)
}

func TestNewSettingsFromYAML_CustomValues(t *testing.T) {
	tests := []struct {
		name        string
		yaml        string
		wantStories int
		wantType    string
		wantTitle   string
	}{
		{
			name:        "custom number of stories",
			yaml:        "numberOfStories: 5",
			wantStories: 5,
			wantType:    "top",
			wantTitle:   "HackerNews",
		},
		{
			name:        "custom story type",
			yaml:        "storyType: new",
			wantStories: 10,
			wantType:    "new",
			wantTitle:   "HackerNews",
		},
		{
			name:        "custom title",
			yaml:        "title: My HN",
			wantStories: 10,
			wantType:    "top",
			wantTitle:   "My HN",
		},
		{
			name:        "all custom values",
			yaml:        "numberOfStories: 20\nstoryType: ask\ntitle: Ask Stories",
			wantStories: 20,
			wantType:    "ask",
			wantTitle:   "Ask Stories",
		},
		{
			name:        "job stories type",
			yaml:        "storyType: job\nnumberOfStories: 3",
			wantStories: 3,
			wantType:    "job",
			wantTitle:   "HackerNews",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NilError(t, err)

			globalConfig, err := config.ParseYaml("wtf: {}")
			assert.NilError(t, err)

			settings := NewSettingsFromYAML("hackernews", ymlConfig, globalConfig)

			assert.Equal(t, tt.wantStories, settings.numberOfStories)
			assert.Equal(t, tt.wantType, settings.storyType)
			assert.Equal(t, tt.wantTitle, settings.Title)
		})
	}
}

func TestDefaultConstants(t *testing.T) {
	assert.Equal(t, true, defaultFocusable)
	assert.Equal(t, "HackerNews", defaultTitle)
}
