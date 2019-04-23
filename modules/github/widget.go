package github

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
  Keyboard commands for GitHub:

    /: Show/hide this help window
    h: Previous git repository
    l: Next git repository
    r: Refresh the data

    arrow left:  Previous git repository
    arrow right: Next git repository

    return: Open the selected repository in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	GithubRepos []*GithubRepo
	Idx         int
	settings    *Settings
}

func NewWidget(app *tview.Application, refreshChan chan<- string, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(refreshChan, settings.common, true),

		Idx:      0,
		settings: settings,
	}

	widget.GithubRepos = widget.buildRepoCollection(widget.settings.repositories)

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, repo := range widget.GithubRepos {
		repo.Refresh()
	}

	widget.display()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.GithubRepos) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.GithubRepos) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildRepoCollection(repoData map[string]interface{}) []*GithubRepo {
	githubRepos := []*GithubRepo{}

	for name, owner := range repoData {
		repo := NewGithubRepo(
			name,
			owner.(string),
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

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyEnter:
		widget.openRepo()
		return nil
	case tcell.KeyLeft:
		widget.Prev()
		return nil
	case tcell.KeyRight:
		widget.Next()
		return nil
	default:
		return event
	}
}

func (widget *Widget) openRepo() {
	repo := widget.currentGithubRepo()

	if repo != nil {
		repo.Open()
	}
}
