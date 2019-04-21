package gitlab

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
	glb "github.com/xanzy/go-gitlab"
)

const HelpText = `
  Keyboard commands for Gitlab:

    /: Show/hide this help window
    h: Previous project
    l: Next project
    r: Refresh the data

    arrow left:  Previous project
    arrow right: Next project
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	GitlabProjects []*GitlabProject
	Idx            int
	gitlab         *glb.Client
	settings       *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	baseURL := settings.domain
	gitlab := glb.NewClient(nil, settings.apiKey)

	if baseURL != "" {
		gitlab.SetBaseURL(baseURL)
	}

	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, settings.common, true),

		Idx:      0,
		gitlab:   gitlab,
		settings: settings,
	}

	widget.GitlabProjects = widget.buildProjectCollection(settings.projects)

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, project := range widget.GitlabProjects {
		project.Refresh()
	}

	widget.display()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.GitlabProjects) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.GitlabProjects) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildProjectCollection(projectData map[string]interface{}) []*GitlabProject {
	gitlabProjects := []*GitlabProject{}

	for name, namespace := range projectData {
		project := NewGitlabProject(name, namespace.(string), widget.gitlab)
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
