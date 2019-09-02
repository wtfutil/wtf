package gcal

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) sortedEvents() ([]*CalEvent, []*CalEvent) {
	allDayEvents := []*CalEvent{}
	timedEvents := []*CalEvent{}

	for _, calEvent := range widget.calEvents {
		if calEvent.AllDay() {
			allDayEvents = append(allDayEvents, calEvent)
		} else {
			timedEvents = append(timedEvents, calEvent)
		}
	}

	return allDayEvents, timedEvents
}

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.settings.common.Title
	calEvents := widget.calEvents

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if (calEvents == nil) || (len(calEvents) == 0) {
		return title, "No calendar events", false
	}

	var str string
	var prevEvent *CalEvent

	if !widget.settings.showDeclined {
		calEvents = widget.removeDeclined(calEvents)
	}

	for _, calEvent := range calEvents {
		timestamp := fmt.Sprintf("[%s]%s", widget.descriptionColor(calEvent), calEvent.Timestamp())
		if calEvent.AllDay() {
			timestamp = ""
		}

		title := fmt.Sprintf("[%s]%s",
			widget.titleColor(calEvent),
			widget.eventSummary(calEvent, calEvent.ConflictsWith(calEvents)),
		)

		lineOne := fmt.Sprintf(
			"%s %s %s %s[white]\n",
			widget.dayDivider(calEvent, prevEvent),
			widget.responseIcon(calEvent),
			timestamp,
			title,
		)

		str += fmt.Sprintf("%s   %s%s\n",
			lineOne,
			widget.location(calEvent),
			widget.timeUntil(calEvent),
		)

		if (widget.location(calEvent) != "") || (widget.timeUntil(calEvent) != "") {
			str += "\n"
		}

		prevEvent = calEvent
	}

	return title, str, false
}

func (widget *Widget) dayDivider(event, prevEvent *CalEvent) string {
	var prevStartTime time.Time

	if prevEvent != nil {
		prevStartTime = prevEvent.Start()
	}

	// round times to midnight for comparison
	toMidnight := func(t time.Time) time.Time {
		t = t.Local()
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	}
	prevStartDay := toMidnight(prevStartTime)
	eventStartDay := toMidnight(event.Start())

	if !eventStartDay.Equal(prevStartDay) {

		return fmt.Sprintf("[%s::b]",
			widget.settings.colors.day) +
			event.Start().Format(utils.FullDateFormat) +
			"\n"
	}

	return ""
}

func (widget *Widget) descriptionColor(calEvent *CalEvent) string {
	if calEvent.Past() {
		return widget.settings.colors.past
	}

	return widget.settings.colors.description
}

func (widget *Widget) eventSummary(calEvent *CalEvent, conflict bool) string {
	summary := calEvent.event.Summary

	if calEvent.Now() {
		summary = fmt.Sprintf(
			"%s %s",
			widget.settings.currentIcon,
			summary,
		)
	}

	if conflict {
		return fmt.Sprintf("%s %s", widget.settings.conflictIcon, summary)
	}

	return summary
}

// timeUntil returns the number of hours or days until the event
// If the event is in the past, returns nil
func (widget *Widget) timeUntil(calEvent *CalEvent) string {
	duration := time.Until(calEvent.Start()).Round(time.Minute)

	if duration < 0 {
		return ""
	}

	days := duration / (24 * time.Hour)
	duration -= days * (24 * time.Hour)

	hours := duration / time.Hour
	duration -= hours * time.Hour

	mins := duration / time.Minute

	untilStr := ""

	color := "[lightblue]"
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

func (widget *Widget) titleColor(calEvent *CalEvent) string {
	color := widget.settings.colors.title

	for _, untypedArr := range widget.settings.highlights {
		highlightElements := utils.ToStrs(untypedArr.([]interface{}))

		match, _ := regexp.MatchString(
			strings.ToLower(highlightElements[0]),
			strings.ToLower(calEvent.event.Summary),
		)

		if match == true {
			color = highlightElements[1]
		}
	}

	if calEvent.Past() {
		color = widget.settings.colors.past
	}

	return color
}

func (widget *Widget) location(calEvent *CalEvent) string {
	if widget.settings.withLocation == false {
		return ""
	}

	if calEvent.event.Location == "" {
		return ""
	}

	return fmt.Sprintf(
		"[%s]%s ",
		widget.descriptionColor(calEvent),
		calEvent.event.Location,
	)
}

func (widget *Widget) responseIcon(calEvent *CalEvent) string {
	if widget.settings.displayResponseStatus == false {
		return ""
	}

	icon := "[gray]"

	switch calEvent.ResponseFor(widget.settings.email) {
	case "accepted":
		return icon + "✔"
	case "declined":
		return icon + "✘"
	case "needsAction":
		return icon + "?"
	case "tentative":
		return icon + "~"
	default:
		return icon + " "
	}
}

func (widget *Widget) removeDeclined(events []*CalEvent) []*CalEvent {
	var ret []*CalEvent
	for _, e := range events {
		if e.ResponseFor(widget.settings.email) != "declined" {
			ret = append(ret, e)
		}
	}
	return ret
}
