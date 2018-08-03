package gitlab

import (
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
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

	gitlab *glb.Client

	GitlabProjects []*GitlabProject
	Idx            int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	baseURL := wtf.Config.UString("wtf.mods.gitlab.domain")
	gitlab := glb.NewClient(nil, apiKey())

	if baseURL != "" {
		gitlab.SetBaseURL(baseURL)
	}

	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Gitlab", "gitlab", true),

		gitlab: gitlab,

		Idx: 0,
	}

	widget.GitlabProjects = widget.buildProjectCollection(wtf.Config.UMap("wtf.mods.gitlab.projects"))

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, project := range widget.GitlabProjects {
		project.Refresh()
	}

	widget.UpdateRefreshedAt()
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

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.gitlab.apiKey",
		os.Getenv("WTF_GITLAB_TOKEN"),
	)
}

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
