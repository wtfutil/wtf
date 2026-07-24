package mercurial

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func newTestWidget(t *testing.T, settings *Settings) *Widget {
	t.Helper()
	// Use a buffered redraw channel so display()/Redraw() calls made during
	// tests don't block waiting for a consumer (no tview event loop is running).
	redrawChan := make(chan bool, 10)
	return NewWidget(tview.NewApplication(), redrawChan, tview.NewPages(), settings)
}

// fullSettings builds Settings via the real YAML loader (rather than a bare
// struct literal) so that Common.Config is populated; this is required by
// view.NewMultiSourceWidget, which reads repositories from it directly.
func fullSettings(t *testing.T, repoPaths []string) *Settings {
	t.Helper()

	yaml := "title: Mercurial\n"
	if len(repoPaths) > 0 {
		yaml += "repositories:\n"
		for _, p := range repoPaths {
			yaml += "  - " + p + "\n"
		}
	}

	ymlConfig, err := config.ParseYaml(yaml)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml(`wtf: {}`)
	assert.NoError(t, err)

	return NewSettingsFromYAML("mercurial", ymlConfig, globalConfig)
}

func Test_NewWidget(t *testing.T) {
	settings := fullSettings(t, nil)

	widget := newTestWidget(t, settings)

	assert.NotNil(t, widget)
	assert.NotNil(t, widget.View)
}

func Test_currentData(t *testing.T) {
	widget := &Widget{}
	assert.Nil(t, widget.currentData())

	repo := &MercurialRepo{Repository: "repo-a"}
	widget.Data = []*MercurialRepo{repo}

	widget.Idx = 0
	assert.Equal(t, repo, widget.currentData())

	widget.Idx = -1
	assert.Nil(t, widget.currentData())

	widget.Idx = 5
	assert.Nil(t, widget.currentData())
}

func Test_mercurialRepos_WithFakeHg(t *testing.T) {
	fakeHg(t)

	settings := testSettings()
	settings.commitCount = 1
	settings.commitFormat = "{rev}"
	widget := &Widget{settings: settings}

	repoPath := t.TempDir()
	repos := widget.mercurialRepos([]string{repoPath})

	assert.Len(t, repos, 1)
	assert.Equal(t, "feature-branch", repos[0].Branch)
}

func Test_mercurialRepos_Empty(t *testing.T) {
	widget := &Widget{settings: testSettings()}

	repos := widget.mercurialRepos([]string{})

	assert.Empty(t, repos)
}

func Test_Refresh_WithFakeHg(t *testing.T) {
	fakeHg(t)

	repoPath := t.TempDir()
	settings := fullSettings(t, []string{repoPath})
	widget := newTestWidget(t, settings)

	widget.Refresh()

	assert.Len(t, widget.Data, 1)
	assert.Equal(t, "feature-branch", widget.Data[0].Branch)
}

func Test_Pull_WithFakeHg(t *testing.T) {
	fakeHg(t)

	settings := fullSettings(t, nil)
	widget := newTestWidget(t, settings)
	widget.Data = []*MercurialRepo{{Path: t.TempDir()}}
	widget.Idx = 0

	assert.NotPanics(t, func() {
		widget.Pull()
	})
}
