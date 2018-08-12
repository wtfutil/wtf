package gerrit

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"

	glb "github.com/andygrunwald/go-gerrit"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const HelpText = `
  Keyboard commands for Gerrit:

    /: Show/hide this help window
    h: Show the previous project
    l: Show the next project
    j: Select the next review in the list
    k: Select the previous review in the list
    r: Refresh the data

    arrow left:  Show the previous project
    arrow right: Show the next project
    arrow down:  Select the next review in the list
	arrow up:    Select the previous review in the list
    
	return: Open the selected review in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	gerrit *glb.Client

	GerritProjects []*GerritProject
	Idx            int
	selected       int
}

var (
	GerritURLPattern = regexp.MustCompile(`^(http|https)://(.*)$`)
)

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Gerrit", "gerrit", true),

		Idx: 0,
	}

	widget.HelpfulWidget.SetView(widget.View)

	widget.View.SetInputCapture(widget.keyboardIntercept)
	widget.unselect()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	baseURL := wtf.Config.UString("wtf.mods.gerrit.domain")
	username := wtf.Config.UString("wtf.mods.gerrit.username")

	password := wtf.Config.UString(
		"wtf.mods.gerrit.password",
		os.Getenv("WTF_GERRIT_PASSWORD"),
	)

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
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		widget.View.SetText(err.Error())
		return
	}
	widget.gerrit = gerrit
	widget.GerritProjects = widget.buildProjectCollection(wtf.Config.UList("wtf.mods.gerrit.projects"))

	for _, project := range widget.GerritProjects {
		project.Refresh()
	}

	widget.UpdateRefreshedAt()
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) nextProject() {
	widget.Idx = widget.Idx + 1
	widget.unselect()
	if widget.Idx == len(widget.GerritProjects) {
		widget.Idx = 0
	}

	widget.unselect()
}

func (widget *Widget) prevProject() {
	widget.Idx = widget.Idx - 1
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
		wtf.OpenFile(fmt.Sprintf("%s/%s/%d", wtf.Config.UString("wtf.mods.gerrit.domain"), "#/c", change.Number))
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

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "h":
		widget.prevProject()
		return nil
	case "l":
		widget.nextProject()
		return nil
	case "j":
		widget.nextReview()
		return nil
	case "k":
		widget.prevReview()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyLeft:
		widget.prevProject()
		return nil
	case tcell.KeyRight:
		widget.nextProject()
		return nil
	case tcell.KeyDown:
		widget.nextReview()
		return nil
	case tcell.KeyUp:
		widget.prevReview()
		return nil
	case tcell.KeyEnter:
		widget.openReview()
		return nil
	case tcell.KeyEsc:
		widget.unselect()
		return event
	default:
		return event
	}
}
