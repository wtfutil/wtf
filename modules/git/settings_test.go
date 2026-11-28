package git

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseYAML(t *testing.T, yml string) *config.Config {
	t.Helper()

	cfg, err := config.ParseYaml(yml)
	require.NoError(t, err)

	return cfg
}

func Test_NewSettingsFromYAML(t *testing.T) {
	globalConfig := parseYAML(t, "wtf:\n  colors: {}\n")

	tests := []struct {
		name                 string
		yml                  string
		wantCommitCount      int
		wantSections         []interface{}
		wantShowModuleName   bool
		wantBranchInTitle    bool
		wantShowFilesIfEmpty bool
		wantLastFolderTitle  bool
		wantCommitFormat     string
		wantDateFormat       string
		wantRepositories     []interface{}
	}{
		{
			name:                 "defaults when no config provided",
			yml:                  "{}",
			wantCommitCount:      10,
			wantSections:         []interface{}{"branch", "files", "commits"},
			wantShowModuleName:   true,
			wantBranchInTitle:    false,
			wantShowFilesIfEmpty: true,
			wantLastFolderTitle:  false,
			wantCommitFormat:     "[forestgreen]%h [white]%s [grey]%an on %cd[white]",
			wantDateFormat:       "%b %d, %Y",
			wantRepositories:     []interface{}{},
		},
		{
			name: "explicit values override defaults",
			yml: `
commitCount: 3
sections: ["branch", "commits"]
showModuleName: false
branchInTitle: true
showFilesIfEmpty: false
lastFolderTitle: true
commitFormat: "%h %s"
dateFormat: "%Y-%m-%d"
repositories: ["/home/user/project"]
`,
			wantCommitCount:      3,
			wantSections:         []interface{}{"branch", "commits"},
			wantShowModuleName:   false,
			wantBranchInTitle:    true,
			wantShowFilesIfEmpty: false,
			wantLastFolderTitle:  true,
			wantCommitFormat:     "%h %s",
			wantDateFormat:       "%Y-%m-%d",
			wantRepositories:     []interface{}{"/home/user/project"},
		},
		{
			name:                 "empty sections list falls back to default sections",
			yml:                  "sections: []\n",
			wantCommitCount:      10,
			wantSections:         []interface{}{"branch", "files", "commits"},
			wantShowModuleName:   true,
			wantShowFilesIfEmpty: true,
			wantCommitFormat:     "[forestgreen]%h [white]%s [grey]%an on %cd[white]",
			wantDateFormat:       "%b %d, %Y",
			wantRepositories:     []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig := parseYAML(t, tt.yml)

			settings := NewSettingsFromYAML("git", ymlConfig, globalConfig)

			require.NotNil(t, settings)
			require.NotNil(t, settings.Common)
			assert.Equal(t, tt.wantCommitCount, settings.commitCount)
			assert.Equal(t, tt.wantSections, settings.sections)
			assert.Equal(t, tt.wantShowModuleName, settings.showModuleName)
			assert.Equal(t, tt.wantBranchInTitle, settings.branchInTitle)
			assert.Equal(t, tt.wantShowFilesIfEmpty, settings.showFilesIfEmpty)
			assert.Equal(t, tt.wantLastFolderTitle, settings.lastFolderTitle)
			assert.Equal(t, tt.wantCommitFormat, settings.commitFormat)
			assert.Equal(t, tt.wantDateFormat, settings.dateFormat)
			assert.Equal(t, tt.wantRepositories, settings.repositories)
		})
	}
}

func Test_ConfigText(t *testing.T) {
	widget := &Widget{}

	text := widget.ConfigText()

	assert.NotEmpty(t, text)
	assert.Contains(t, text, "commitCount")
	assert.Contains(t, text, "repositories")
}
