package status

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

	Current int
	View    *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "Status",
			RefreshedAt: time.Now(),
			RefreshInt:  Config.UInt("wtf.status.refreshInterval", 1),
		},
		Current: 0,
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.View.SetTitle(" 🎉 Status ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, " %s", widget.contentFrom())
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

func (widget *Widget) contentFrom() string {
	icons := []string{"👍", "🤜", "🤙", "🤜", "🤘", "🤜", "✊", "🤜", "👌", "🤜"}
	next := icons[widget.Current]

	widget.Current = widget.Current + 1
	if widget.Current == len(icons) {
		widget.Current = 0
	}

	return next
}
