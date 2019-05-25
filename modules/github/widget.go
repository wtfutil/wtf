package github

import (
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.MultiSourceWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	GithubRepos []*GithubRepo

	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    wtf.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "repository", "repositories"),
		TextWidget:        wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.GithubRepos = widget.buildRepoCollection(widget.settings.repositories)

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.SetDisplayFunction(widget.display)

	widget.Sources = widget.settings.repositories

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, repo := range widget.GithubRepos {
		repo.Refresh()
	}

	widget.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildRepoCollection(repoData []string) []*GithubRepo {
	githubRepos := []*GithubRepo{}

	for _, repo := range repoData {
		split := strings.Split(repo, "/")
		owner, name := split[0], split[1]
		repo := NewGithubRepo(
			name,
			owner,
			widget.settings.apiKey,
			widget.settings.baseURL,
			widget.settings.uploadURL,
		)

		githubRepos = append(githubRepos, repo)
	}

	return githubRepos
}

func (widget *Widget) currentGithubRepo() *GithubRepo {
	if len(widget.GithubRepos) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.GithubRepos) {
		return nil
	}

	return widget.GithubRepos[widget.Idx]
}

func (widget *Widget) openRepo() {
	repo := widget.currentGithubRepo()

	if repo != nil {
		repo.Open()
	}
}
