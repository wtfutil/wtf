package bamboohr

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

type Widget struct {
	RefreshedAt     time.Time
	RefreshInterval int
	View            *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		RefreshedAt:     time.Now(),
		RefreshInterval: 3600,
	}

	widget.addView()
	go widget.refresher()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	items := Fetch()

	widget.View.SetTitle(fmt.Sprintf(" 🐨 Away (%d) ", len(items)))
	widget.RefreshedAt = time.Now()

	fmt.Fprintf(widget.View, "%s", widget.contentFrom(items))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(" BambooHR ")

	widget.View = view
}

func (widget *Widget) contentFrom(items []Item) string {
	str := ""

	for _, item := range items {
		str = str + widget.display(item)
	}

	return str
}

func (widget *Widget) display(item Item) string {
	var str string

	if item.IsOneDay() {
		str = fmt.Sprintf(" [green]%s[white]\n %s\n\n", item.Name(), item.PrettyEnd())
	} else {
		str = fmt.Sprintf(" [green]%s[white]\n %s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
	}

	return str
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
