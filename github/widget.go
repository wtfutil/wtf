package github

import (
	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

const helpText = `
  Keyboard commands for Github:

    /: Show/hide this help window
    h: Previous git repository
    l: Next git repository
    r: Refresh the data

    arrow left:  Previous git repository
    arrow right: Next git repository
`

type Widget struct {
	wtf.TextWidget

	app   *tview.Application
	Data  []*GithubRepo
	Idx   int
	pages *tview.Pages
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Github ", "github", true),

		app:   app,
		Idx:   0,
		pages: pages,
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Data = widget.buildRepoCollection(Config.UMap("wtf.mods.github.repositories"))

	widget.UpdateRefreshedAt()
	widget.display()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Data) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Data) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildRepoCollection(repoData map[string]interface{}) []*GithubRepo {
	githubRepos := []*GithubRepo{}

	for name, owner := range repoData {
		repo := NewGithubRepo(name, owner.(string))
		githubRepos = append(githubRepos, repo)
	}

	return githubRepos
}

func (widget *Widget) currentData() *GithubRepo {
	if len(widget.Data) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Data) {
		return nil
	}

	return widget.Data[widget.Idx]
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

	modal := wtf.NewBillboardModal(helpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
}
