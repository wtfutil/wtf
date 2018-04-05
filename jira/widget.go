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
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "JIRA",
			RefreshedAt: time.Now(),
			RefreshInt:  Config.UInt("wtf.jira.refreshInterval", 900),
		},
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.View.SetTitle(fmt.Sprintf(" %s ", widget.Name))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", "jira")
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
