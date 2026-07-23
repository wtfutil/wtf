package mercurial

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name             string
		yaml             string
		wantTitle        string
		wantCommitCount  int
		wantCommitFormat string
		wantRepoCount    int
	}{
		{
			name:             "defaults",
			yaml:             `{}`,
			wantTitle:        defaultTitle,
			wantCommitCount:  10,
			wantCommitFormat: "[forestgreen]{rev}:{phase} [white]{desc|firstline|strip} [grey]{author|person} {date|age}[white]",
			wantRepoCount:    0,
		},
		{
			name: "custom values",
			yaml: `
title: "My Hg Repos"
commitCount: 5
commitFormat: "{rev}:{desc}"
repositories:
  - /repo/one
  - /repo/two
`,
			wantTitle:        "My Hg Repos",
			wantCommitCount:  5,
			wantCommitFormat: "{rev}:{desc}",
			wantRepoCount:    2,
		},
	}

	globalConfig, err := config.ParseYaml(`wtf: {}`)
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("mercurial", ymlConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.wantTitle, settings.Title)
			assert.Equal(t, tt.wantCommitCount, settings.commitCount)
			assert.Equal(t, tt.wantCommitFormat, settings.commitFormat)
			assert.Len(t, settings.repositories, tt.wantRepoCount)
		})
	}
}

func Test_ConfigText(t *testing.T) {
	widget := &Widget{}
	text := widget.ConfigText()

	assert.NotEmpty(t, text)
}

func Test_defaultConstants(t *testing.T) {
	assert.True(t, defaultFocusable)
	assert.Equal(t, "Mercurial", defaultTitle)
}
