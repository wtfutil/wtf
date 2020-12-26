package github

import (
	"strconv"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	GithubRepos []*Repo

	settings *Settings
	Selected int
	maxItems int
	Items    []int
}

// NewWidget creates a new instance of the widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "repository", "repositories"),
		TextWidget:        view.NewTextWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.GithubRepos = widget.buildRepoCollection(widget.settings.repositories)

	widget.initializeKeyboardControls()

	widget.View.SetRegions(true)
	widget.SetDisplayFunction(widget.display)

	widget.Unselect()

	widget.Sources = widget.settings.repositories

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// SetItemCount sets the amount of PRs RRs and other PRs throughout the widgets display creation
func (widget *Widget) SetItemCount(items int) {
	widget.maxItems = items
}

// GetItemCount returns the amount of PRs RRs and other PRs calculated so far as an int
func (widget *Widget) GetItemCount() int {
	return widget.maxItems
}

// GetSelected returns the index of the currently highlighted item as an int
func (widget *Widget) GetSelected() int {
	if widget.Selected < 0 {
		return 0
	}
	return widget.Selected
}

// Next cycles the currently highlighted text down
func (widget *Widget) Next() {
	widget.Selected++
	if widget.Selected >= widget.maxItems {
		widget.Selected = 0
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

// Prev cycles the currently highlighted text up
func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = widget.maxItems - 1
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected)).ScrollToHighlight()
}

// Unselect stops highlighting the text and jumps the scroll position to the top
func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.Highlight()
	widget.View.ScrollToBeginning()
}

// Refresh reloads the github data via the Github API and reruns the display
func (widget *Widget) Refresh() {
	for _, repo := range widget.GithubRepos {
		repo.Refresh()
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildRepoCollection(repoData []string) []*Repo {
	githubRepos := []*Repo{}

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

func (widget *Widget) currentGithubRepo() *Repo {
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
		url := (*widget.currentGithubRepo().RemoteRepo.HTMLURL + "/pull/" + strconv.Itoa(widget.Items[widget.Selected]))
		utils.OpenFile(url)
	}
}

func (widget *Widget) openRepo() {
	repo := widget.currentGithubRepo()

	if repo != nil {
		repo.Open()
	}
}

func (widget *Widget) openPulls() {
	repo := widget.currentGithubRepo()

	if repo != nil {
		repo.OpenPulls()
	}
}

func (widget *Widget) openIssues() {
	repo := widget.currentGithubRepo()

	if repo != nil {
		repo.OpenIssues()
	}
}
