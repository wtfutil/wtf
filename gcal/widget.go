package gcal

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell"
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
			Name:        "Calendar",
			RefreshedAt: time.Now(),
			RefreshInt:  300,
		},
	}

	widget.addView()
	go wtf.Refresh(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	events := Fetch()

	widget.View.SetTitle(" 🍿 Calendar ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(events))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGrey)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(events *calendar.Events) string {
	str := "\n"

	for _, event := range events.Items {
		startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		timestamp := startTime.Format("Mon, Jan 2, 15:04")
		until := widget.until(startTime)

		summary := event.Summary
		if widget.eventIsNow(event) {
			summary = "🔥 " + summary
		}

		str = str + fmt.Sprintf(" [%s]%s[white]\n [%s]%s %s[white]\n\n", titleColor(event), summary, descriptionColor(event), timestamp, until)
	}

	return str
}

// eventIsNow returns true if the event is happening now, false if it not
func (widget *Widget) eventIsNow(event *calendar.Event) bool {
	startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, event.End.DateTime)

	return time.Now().After(startTime) && time.Now().Before(endTime)
}

func descriptionColor(item *calendar.Event) string {
	ts, _ := time.Parse(time.RFC3339, item.Start.DateTime)

	color := "white"
	if ts.Before(time.Now()) {
		color = "grey"
	}

	return color
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

// until returns the number of hours or days until the event
// If the event is in the past, returns nil
func (widget *Widget) until(start time.Time) string {
	duration := time.Until(start)

	duration = duration.Round(time.Minute)

	days := duration / (24 * time.Hour)
	duration -= days * (24 * time.Hour)

	hours := duration / time.Hour
	duration -= hours * time.Hour

	mins := duration / time.Minute

	untilStr := ""

	if days > 0 {
		untilStr = fmt.Sprintf("%dd", days)
	} else if hours > 0 {
		untilStr = fmt.Sprintf("%dh", hours)
	} else {
		untilStr = fmt.Sprintf("%dm", mins)
	}

	return "[grey]" + untilStr + "[white]"
}
