package gitlab

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
	glb "github.com/xanzy/go-gitlab"
)

type Widget struct {
	wtf.KeyboardWidget
	wtf.MultiSourceWidget
	wtf.TextWidget

	GitlabProjects []*GitlabProject

	gitlab   *glb.Client
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	baseURL := settings.domain
	gitlab := glb.NewClient(nil, settings.apiKey)

	if baseURL != "" {
		gitlab.SetBaseURL(baseURL)
	}

	widget := Widget{
		KeyboardWidget:    wtf.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "repository", "repositories"),
		TextWidget:        wtf.NewTextWidget(app, settings.common, true),

		gitlab:   gitlab,
		settings: settings,
	}

	widget.GitlabProjects = widget.buildProjectCollection(settings.projects)

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.SetDisplayFunction(widget.display)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, project := range widget.GitlabProjects {
		project.Refresh()
	}

	widget.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
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
