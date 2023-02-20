package gitlab

import (
	"strconv"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type ContentItem struct {
	Type string
	ID   int
}

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	GitlabProjects []*GitlabProject

	context  *context
	settings *Settings
	Selected int
	maxItems int
	Items    []ContentItem

	configError error
}

// NewWidget creates a new instance of the widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	context, err := newContext(settings)

	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "project", "projects"),
		TextWidget:        view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		context:  context,
		settings: settings,

		configError: err,
	}

	widget.GitlabProjects = widget.buildProjectCollection(context, settings.projects)

	widget.initializeKeyboardControls()
	widget.View.SetRegions(true)
	widget.SetDisplayFunction(widget.display)

	widget.Unselect()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.context == nil || widget.configError != nil {
		widget.displayError()
		return
	}

	for _, project := range widget.GitlabProjects {
		project.Refresh()
	}

	widget.display()
}

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
	widget.View.Highlight(strconv.Itoa(widget.Selected))
	widget.View.ScrollToHighlight()
}

// Prev cycles the currently highlighted text up
func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = widget.maxItems - 1
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected))
	widget.View.ScrollToHighlight()
}

// Unselect stops highlighting the text and jumps the scroll position to the top
func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.Highlight()
	widget.View.ScrollToBeginning()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildProjectCollection(context *context, projectData []string) []*GitlabProject {
	gitlabProjects := []*GitlabProject{}

	for _, projectPath := range projectData {
		project := NewGitlabProject(context, projectPath)
		gitlabProjects = append(gitlabProjects, project)
	}

	return gitlabProjects
}

func (widget *Widget) currentGitlabProject() *GitlabProject {
	if len(widget.GitlabProjects) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.GitlabProjects) {
		return nil
	}

	return widget.GitlabProjects[widget.Idx]
}

func (widget *Widget) openItemInBrowser() {
	currentSelection := widget.View.GetHighlights()
	if widget.Selected >= 0 && currentSelection[0] != "" {

		item := widget.Items[widget.Selected]
		url := ""

		project := widget.currentGitlabProject()
		if project == nil {
			// This is a problem. We will just bail out for now
			return
		}

		switch item.Type {
		case "MR":
			url = (project.RemoteProject.WebURL + "/merge_requests/" + strconv.Itoa(item.ID))
		case "ISSUE":
			url = (project.RemoteProject.WebURL + "/issues/" + strconv.Itoa(item.ID))
		}

		utils.OpenFile(url)
	}
}

func (widget *Widget) openRepo() {
	project := widget.currentGitlabProject()
	if project == nil {
		return
	}
	url := project.RemoteProject.WebURL
	utils.OpenFile(url)
}

func (widget *Widget) openPulls() {
	project := widget.currentGitlabProject()
	if project == nil {
		return
	}
	url := project.RemoteProject.WebURL + "/merge_requests/"
	utils.OpenFile(url)
}

func (widget *Widget) openIssues() {
	project := widget.currentGitlabProject()
	if project == nil {
		return
	}
	url := project.RemoteProject.WebURL + "/issues/"
	utils.OpenFile(url)
}
