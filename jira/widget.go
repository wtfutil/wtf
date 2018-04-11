package jira

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("JIRA", "jira"),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	issues, err := IssuesFor(Config.UString("wtf.mods.jira.username"))

	widget.View.SetTitle(fmt.Sprintf(" %s ", widget.Name))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()

	if err != nil {
		widget.View.SetWrap(true)
		fmt.Fprintf(widget.View, "%v", err)
	} else {
		widget.View.SetWrap(false)
		fmt.Fprintf(widget.View, "%v", issues)
	}
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)

	widget.View = view
}
