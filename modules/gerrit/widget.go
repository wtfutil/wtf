package gerrit

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"

	glb "github.com/andygrunwald/go-gerrit"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.TextWidget

	gerrit *glb.Client

	GerritProjects []*GerritProject
	Idx            int

	selected int
	settings *Settings
	err      error
}

var (
	GerritURLPattern = regexp.MustCompile(`^(http|https)://(.*)$`)
)

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: view.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     view.NewTextWidget(app, settings.common, true),

		Idx: 0,

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	widget.unselect()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !widget.settings.verifyServerCertificate,
			},
			Proxy: http.ProxyFromEnvironment,
		},
	}

	gerritUrl := widget.settings.domain
	submatches := GerritURLPattern.FindAllStringSubmatch(widget.settings.domain, -1)

	if len(submatches) > 0 && len(submatches[0]) > 2 {
		submatch := submatches[0]
		gerritUrl = fmt.Sprintf(
			"%s://%s:%s@%s",
			submatch[1],
			widget.settings.username,
			widget.settings.password,
			submatch[2],
		)
	}
	gerrit, err := glb.NewClient(gerritUrl, httpClient)
	if err != nil {
		widget.err = err
		widget.gerrit = nil
		widget.GerritProjects = nil
	} else {
		widget.err = nil
		widget.gerrit = gerrit
		widget.GerritProjects = widget.buildProjectCollection(widget.settings.projects)
		for _, project := range widget.GerritProjects {
			project.Refresh(widget.settings.username)
		}
	}

	widget.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) nextProject() {
	widget.Idx++
	widget.unselect()
	if widget.Idx == len(widget.GerritProjects) {
		widget.Idx = 0
	}

	widget.unselect()
}

func (widget *Widget) prevProject() {
	widget.Idx--
	if widget.Idx < 0 {
		widget.Idx = len(widget.GerritProjects) - 1
	}

	widget.unselect()
}

func (widget *Widget) nextReview() {
	widget.selected++
	project := widget.GerritProjects[widget.Idx]
	if widget.selected >= project.ReviewCount {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prevReview() {
	widget.selected--
	project := widget.GerritProjects[widget.Idx]
	if widget.selected < 0 {
		widget.selected = project.ReviewCount - 1
	}

	widget.display()
}

func (widget *Widget) openReview() {
	sel := widget.selected
	project := widget.GerritProjects[widget.Idx]
	if sel >= 0 && sel < project.ReviewCount {
		change := glb.ChangeInfo{}
		if sel < len(project.IncomingReviews) {
			change = project.IncomingReviews[sel]
		} else {
			change = project.OutgoingReviews[sel-len(project.IncomingReviews)]
		}
		utils.OpenFile(fmt.Sprintf("%s/%s/%d", widget.settings.domain, "#/c", change.Number))
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) buildProjectCollection(projectData []interface{}) []*GerritProject {
	gerritProjects := []*GerritProject{}

	for _, name := range projectData {
		project := NewGerritProject(name.(string), widget.gerrit)
		gerritProjects = append(gerritProjects, project)
	}

	return gerritProjects
}

func (widget *Widget) currentGerritProject() *GerritProject {
	if len(widget.GerritProjects) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.GerritProjects) {
		return nil
	}

	return widget.GerritProjects[widget.Idx]
}
