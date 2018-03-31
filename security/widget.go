package security

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:            "Weather",
			RefreshedAt:     time.Now(),
			RefreshInterval: 5,
		},
	}

	widget.addView()
	go widget.refresher()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	//data := Fetch()

	widget.View.SetTitle(" 🐼 Security")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()

	fmt.Fprintf(widget.View, "%s", "cats and dogs")
	//fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)

	widget.View = view
}

func (widget *Widget) refresher() {
	tick := time.NewTicker(time.Duration(widget.RefreshInterval) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			widget.Refresh()
		case <-quit:
			tick.Stop()
			return
		}
	}
}
