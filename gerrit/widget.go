package gerrit

import (
	"crypto/tls"
	"fmt"
	glb "github.com/andygrunwald/go-gerrit"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"net/http"
	"os"
	"regexp"
)

const HelpText = `
  Keyboard commands for Gerrit:

    /: Show/hide this help window
    h: Previous project
    l: Next project
    r: Refresh the data

    arrow left:  Previous project
    arrow right: Next project
`

type Widget struct {
	wtf.TextWidget

	app   *tview.Application
	pages *tview.Pages

	gerrit *glb.Client

	GerritProjects []*GerritProject
	Idx            int
}

var (
	GerritURLPattern = regexp.MustCompile(`^(http|https)://(.*)$`)
)

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	baseURL := wtf.Config.UString("wtf.mods.gerrit.domain")
	username := wtf.Config.UString("wtf.mods.gerrit.username")
	password := os.Getenv("WTF_GERRIT_PASSWORD")
	verifyServerCertificate := wtf.Config.UBool("wtf.mods.gerrit.verifyServerCertificate", true)

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !verifyServerCertificate,
		},
	},
	}

	gerritUrl := baseURL
	submatches := GerritURLPattern.FindAllStringSubmatch(baseURL, -1)

	if len(submatches) > 0 && len(submatches[0]) > 2 {
		submatch := submatches[0]
		gerritUrl = fmt.Sprintf(
			"%s://%s:%s@%s", submatch[1], username, password, submatch[2])
	}

	gerrit, err := glb.NewClient(gerritUrl, httpClient)
	if err != nil {
		panic(err)
	}

	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Gerrit ", "gerrit", true),

		app:   app,
		pages: pages,

		gerrit: gerrit,

		Idx: 0,
	}

	widget.GerritProjects = widget.buildProjectCollection(wtf.Config.UList("wtf.mods.gerrit.projects"))

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, project := range widget.GerritProjects {
		project.Refresh()
	}

	widget.UpdateRefreshedAt()
	widget.display()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.GerritProjects) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.GerritProjects) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

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

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.showHelp()
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

func (widget *Widget) showHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.View)
	}

	modal := wtf.NewBillboardModal(HelpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
}
