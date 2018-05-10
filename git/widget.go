package git

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
		h: Previous weather location
		l: Next weather location

    arrow left:  Previous weather location
    arrow right: Next weather location
`

type Widget struct {
	wtf.TextWidget

	Data []*GitRepo
	Idx  int
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Git ", "git", true),
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

	repoPaths := wtf.ToStrs(Config.UList("wtf.mods.git.repositories"))
	widget.Data = widget.gitRepos(repoPaths)

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

func (widget *Widget) currentData() *GitRepo {
	if len(widget.Data) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Data) {
		return nil
	}

	return widget.Data[widget.Idx]
}

func (widget *Widget) gitRepos(repoPaths []string) []*GitRepo {
	repos := []*GitRepo{}

	for _, repoPath := range repoPaths {
		repo := NewGitRepo(repoPath)
		repos = append(repos, repo)
	}

	return repos
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
