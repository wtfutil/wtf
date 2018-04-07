package gcal

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"google.golang.org/api/calendar/v3"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Calendar", "gcal"),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Enabled == false {
		return
	}

	events, _ := Fetch()

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
	if events == nil {
		return ""
	}

	var prevEvent *calendar.Event

	str := ""
	for _, event := range events.Items {
		str = str + fmt.Sprintf(
			"%s [%s]%s[white]\n [%s]%s %s[white]\n\n",
			widget.dayDivider(event, prevEvent),
			widget.titleColor(event),
			widget.eventSummary(event),
			widget.descriptionColor(event),
			widget.eventTimestamp(event),
			widget.until(event),
		)

		prevEvent = event
	}

	return str
}

func (widget *Widget) dayDivider(event, prevEvent *calendar.Event) string {
	if prevEvent != nil {
		prevStartTime, _ := time.Parse(time.RFC3339, prevEvent.Start.DateTime)
		currStartTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)

		if currStartTime.Day() != prevStartTime.Day() {
			return "\n"
		}
	}

	return ""
}

func (widget *Widget) eventSummary(event *calendar.Event) string {
	if widget.eventIsNow(event) {
		return fmt.Sprintf("%s %s", Config.UString("wtf.gcal.currentIcon", "🔸"), event.Summary)
	}

	return event.Summary
}

func (widget *Widget) eventTimestamp(event *calendar.Event) string {
	startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	return startTime.Format("Mon, Jan 2, 15:04")
}

// eventIsNow returns true if the event is happening now, false if it not
func (widget *Widget) eventIsNow(event *calendar.Event) bool {
	startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, event.End.DateTime)

	return time.Now().After(startTime) && time.Now().Before(endTime)
}

func (widget *Widget) descriptionColor(event *calendar.Event) string {
	ts, _ := time.Parse(time.RFC3339, event.Start.DateTime)

	color := "white"
	if (widget.eventIsNow(event) == false) && ts.Before(time.Now()) {
		color = "grey"
	}

	return color
}

func (widget *Widget) titleColor(event *calendar.Event) string {
	ts, _ := time.Parse(time.RFC3339, event.Start.DateTime)

	color := "red"
	if strings.Contains(event.Summary, "1on1") {
		color = "green"
	}

	if (widget.eventIsNow(event) == false) && ts.Before(time.Now()) {
		color = "grey"
	}

	return color
}

// until returns the number of hours or days until the event
// If the event is in the past, returns nil
func (widget *Widget) until(event *calendar.Event) string {
	startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	duration := time.Until(startTime)

	duration = duration.Round(time.Minute)

	if duration < 0 {
		return ""
	}

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
