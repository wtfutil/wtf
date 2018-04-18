package git

import (
	"fmt"
	"strings"
	"time"
	

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	Data []*GitRepo
	Idx  int
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Git ", "git"),
		Idx:        0,
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Fetch(repoPaths []string) []*GitRepo {
	data := []*GitRepo{}

	for _, repoPath := range repoPaths {
		repo := NewGitRepo(repoPath)
		data = append(data, repo)
	}

	return data
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	repoPaths := wtf.ToStrs(Config.UList("wtf.mods.git.repositories"))
	widget.Data = widget.Fetch(repoPaths)

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
  return widget.Data[widget.Idx]
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
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

	return event
}

func (widget *Widget) tickMarks(data []*GitRepo) string {
	str := ""

	if len(data) > 1 {
		marks := strings.Repeat("*", len(data))
		marks = marks[:widget.Idx] + "_" + marks[widget.Idx+1:]

		str = "[lightblue]" + fmt.Sprintf(wtf.RightAlignFormat(widget.View), marks) + "[white]"
	}

	return str
}
