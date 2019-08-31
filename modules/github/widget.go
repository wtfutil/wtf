package github

import (
	"strings"
	"strconv"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.MultiSourceWidget
	view.KeyboardWidget
	view.TextWidget

	GithubRepos []*GithubRepo

	settings *Settings
	Selected int
	maxItems int
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    view.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: view.NewMultiSourceWidget(settings.common, "repository", "repositories"),
		TextWidget:        view.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.GithubRepos = widget.buildRepoCollection(widget.settings.repositories)

	widget.initializeKeyboardControls()
	widget.View.SetRegions(true)
	widget.View.SetInputCapture(widget.InputCapture)
	widget.SetDisplayFunction(widget.display)

	widget.Unselect()

	widget.Sources = widget.settings.repositories

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */
func (widget *Widget) SetItemCount(items int) {
	widget.maxItems = items
}

func (widget *Widget) GetItemCount() int {
	return widget.maxItems
}

func (widget *Widget) GetSelected() int {
	if widget.Selected < 0 {
		return 0
	}
	return widget.Selected
}

func (widget *Widget) Next() {
	widget.Selected++
	if widget.Selected >= widget.maxItems {
		widget.Selected = 0
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = widget.maxItems - 1
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.ScrollToBeginning()
}

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

func (widget *Widget) openPr() {
	currentSelection := widget.View.GetHighlights()
	if widget.Selected >= 0 && currentSelection[0] != "" {
		url := (*widget.currentGithubRepo().RemoteRepo.HTMLURL + "/pull/" +  widget.View.GetRegionText(currentSelection[0]))
		utils.OpenFile(url)
	}
}

func (widget *Widget) openRepo() {
	repo := widget.currentGithubRepo()

	if repo != nil {
		repo.Open()
	}
}
