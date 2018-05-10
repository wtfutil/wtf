package github

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

const helpText = `
  Keyboard commands for Git:

    /: Show/hide this help window
		h: Previous git repository
		l: Next git repository

    arrow left:  Previous git repository
    arrow right: Next git repository
`

type Widget struct {
	wtf.TextWidget

	Data []*GithubRepo
	Idx  int
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Github ", "github", true),
		Idx:        0,
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

	widget.display()
	widget.RefreshedAt = time.Now()
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
		githubColl = append(githubRepos, repo)
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
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
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
