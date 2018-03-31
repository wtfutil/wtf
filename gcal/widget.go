package gcal

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"google.golang.org/api/calendar/v3"
)

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:            "Calendar",
			RefreshedAt:     time.Now(),
			RefreshInterval: 60,
		},
	}

	widget.addView()
	go widget.refresher()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	events := Fetch()

	widget.View.SetTitle(" 🐸 Calendar ")
	widget.RefreshedAt = time.Now()

	fmt.Fprintf(widget.View, "%s", widget.contentFrom(events))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(events *calendar.Events) string {
	str := "\n"

	for _, item := range events.Items {
		ts, _ := time.Parse(time.RFC3339, item.Start.DateTime)
		timestamp := ts.Format("Mon, Jan 2 - 15:04")

		str = str + fmt.Sprintf(" [%s]%s[white]\n [%s]%s[white]\n\n", titleColor(item), item.Summary, descriptionColor(item), timestamp)
	}

	return str
}

func titleColor(item *calendar.Event) string {
	ts, _ := time.Parse(time.RFC3339, item.Start.DateTime)

	color := "red"
	if strings.Contains(item.Summary, "1on1") {
		color = "green"
	}

	if ts.Before(time.Now()) {
		color = "grey"
	}

	return color
}

func descriptionColor(item *calendar.Event) string {
	ts, _ := time.Parse(time.RFC3339, item.Start.DateTime)

	color := "white"
	if ts.Before(time.Now()) {
		color = "grey"
	}

	return color
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
