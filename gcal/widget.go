package gcal

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
	"google.golang.org/api/calendar/v3"
)

type Widget struct {
	wtf.TextWidget

	events *calendar.Events
	ch     chan struct{}
	mutex  sync.Mutex
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Calendar ", "gcal", false),
		ch:         make(chan struct{}),
	}

	go updateLoop(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	events, _ := Fetch()
	widget.events = events

	widget.UpdateRefreshedAt()

	widget.display()
}

func (widget *Widget) Disable() {
	close(widget.ch)
	widget.TextWidget.Disable()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.events == nil || len(widget.events.Items) == 0 {
		return
	}

	widget.mutex.Lock()
	defer widget.mutex.Unlock()
	widget.View.SetText(widget.contentFrom(widget.events))
}

// conflicts returns TRUE if this event conflicts with another, FALSE if it does not
func (widget *Widget) conflicts(event *calendar.Event, events *calendar.Events) bool {
	conflict := false

	for _, otherEvent := range events.Items {
		if event == otherEvent {
			continue
		}

		eventStart, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		eventEnd, _ := time.Parse(time.RFC3339, event.End.DateTime)

		otherEnd, _ := time.Parse(time.RFC3339, otherEvent.End.DateTime)
		otherStart, _ := time.Parse(time.RFC3339, otherEvent.Start.DateTime)

		if eventStart.Before(otherEnd) && eventEnd.After(otherStart) {
			conflict = true
			break
		}
	}

	return conflict
}

func (widget *Widget) contentFrom(events *calendar.Events) string {
	if events == nil {
		return ""
	}

	var prevEvent *calendar.Event

	str := ""
	for _, event := range events.Items {
		conflict := widget.conflicts(event, events)

		dayDivider := widget.dayDivider(event, prevEvent)
		responseIcon := widget.responseIcon(event)
		timestamp := fmt.Sprintf("[%s]%s",
			widget.descriptionColor(event),
			widget.eventTimestamp(event))
		title := fmt.Sprintf("[%s]%s",
			widget.titleColor(event),
			widget.eventSummary(event, conflict))
		until := widget.until(event)

		lineOne := fmt.Sprintf(
			"%s %s %s %s %s[white]",
			dayDivider,
			responseIcon,
			timestamp,
			title,
			until,
		)
		str = str + fmt.Sprintf("%s%s\n\n",
			lineOne,
			widget.location(event), // prefixes newline if non-empty
		)

		prevEvent = event
	}

	return str
}

func (widget *Widget) dayDivider(event, prevEvent *calendar.Event) string {
	var prevStartTime time.Time
	if prevEvent != nil {
		prevStartTime, _ = time.Parse(time.RFC3339, prevEvent.Start.DateTime)
	}
	currStartTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)

	if currStartTime.Day() != prevStartTime.Day() {
		_, _, width, _ := widget.View.GetInnerRect()
		return fmt.Sprintf("[%s]", wtf.Config.UString("wtf.mods.gcal.colors.day", "forestgreen")) +
			wtf.CenterText(currStartTime.Format(wtf.FullDateFormat), width) +
			"\n"
	}

	return ""
}

func (widget *Widget) descriptionColor(event *calendar.Event) string {
	color := wtf.Config.UString("wtf.mods.gcal.colors.description", "white")

	if widget.eventIsPast(event) {
		color = wtf.Config.UString("wtf.mods.gcal.colors.past", "gray")
	}

	return color
}

func (widget *Widget) eventSummary(event *calendar.Event, conflict bool) string {
	summary := event.Summary

	if widget.eventIsNow(event) {
		summary = fmt.Sprintf(
			"%s %s",
			wtf.Config.UString("wtf.mods.gcal.currentIcon", "🔸"),
			event.Summary,
		)
	}

	if conflict {
		return fmt.Sprintf("%s %s", wtf.Config.UString("wtf.mods.gcal.conflictIcon", "🚨"), summary)
	} else {
		return summary
	}
}

func (widget *Widget) eventTimestamp(event *calendar.Event) string {
	if len(event.Start.Date) > 0 {
		startTime, _ := time.Parse("2006-01-02", event.Start.Date)
		return startTime.Format(wtf.FriendlyDateFormat)
	} else {
		startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		return startTime.Format(wtf.MinimumTimeFormat)
	}
}

// eventIsNow returns true if the event is happening now, false if it not
func (widget *Widget) eventIsNow(event *calendar.Event) bool {
	startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, event.End.DateTime)

	return time.Now().After(startTime) && time.Now().Before(endTime)
}

func (widget *Widget) eventIsPast(event *calendar.Event) bool {
	ts, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	return (widget.eventIsNow(event) == false) && ts.Before(time.Now())
}

func (widget *Widget) titleColor(event *calendar.Event) string {
	color := wtf.Config.UString("wtf.mods.gcal.colors.title", "white")

	for _, untypedArr := range wtf.Config.UList("wtf.mods.gcal.colors.highlights") {
		highlightElements := wtf.ToStrs(untypedArr.([]interface{}))

		match, _ := regexp.MatchString(
			strings.ToLower(highlightElements[0]),
			strings.ToLower(event.Summary),
		)

		if match == true {
			color = highlightElements[1]
		}
	}

	if widget.eventIsPast(event) {
		color = wtf.Config.UString("wtf.mods.gcal.colors.past", "gray")
	}

	return color
}

func (widget *Widget) location(event *calendar.Event) string {
	if wtf.Config.UBool("wtf.mods.gcal.displayLocation", true) == false {
		return ""
	}

	if event.Location == "" {
		return ""
	}

	return fmt.Sprintf(
		"\n [%s]%s",
		widget.descriptionColor(event),
		event.Location,
	)
}

func (widget *Widget) responseIcon(event *calendar.Event) string {
	if false == wtf.Config.UBool("wtf.mods.gcal.displayResponseStatus", true) {
		return ""
	}

	response := ""

	for _, attendee := range event.Attendees {
		if attendee.Email == wtf.Config.UString("wtf.mods.gcal.email") {
			response = attendee.ResponseStatus
			break
		}
	}

	icon := "[gray]"

	switch response {
	case "accepted":
		icon = icon + "✔︎"
	case "declined":
		icon = icon + "✘"
	case "needsAction":
		icon = icon + "?"
	case "tentative":
		icon = icon + "~"
	default:
		icon = icon + " "
	}

	return icon
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

	var color = "[lightblue]"
	if days > 0 {
		untilStr = fmt.Sprintf("%dd", days)
	} else if hours > 0 {
		untilStr = fmt.Sprintf("%dh", hours)
	} else {
		untilStr = fmt.Sprintf("%dm", mins)
		if mins < 30 {
			color = "[red]"
		}
	}

	return color + untilStr + "[white]"
}

func updateLoop(widget *Widget) {
	interval := wtf.Config.UInt("wtf.mods.gcal.textInterval", 30)
	if interval == 0 {
		return
	}

	tick := time.NewTicker(time.Duration(interval) * time.Second)
	defer tick.Stop()
outer:
	for {
		select {
		case <-tick.C:
			widget.display()
		case <-widget.ch:
			break outer
		}
	}
}
