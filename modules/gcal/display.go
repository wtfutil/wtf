package gcal

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.settings.Title
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
		if calEvent.AllDay() && !widget.settings.showAllDay {
			continue
		}

		ts := calEvent.Timestamp(widget.settings.hourFormat, widget.settings.showEndTime)
		timestamp := fmt.Sprintf("[%s]%s", widget.eventTimeColor(), ts)
		if calEvent.AllDay() {
			timestamp = ""
		}

		eventTitle := fmt.Sprintf("[%s]%s",
			widget.titleColor(calEvent),
			widget.eventSummary(calEvent, calEvent.ConflictsWith(calEvents)),
		)

		lineOne := fmt.Sprintf(
			"%s %s %s %s[white]\n",
			widget.dayDivider(calEvent, prevEvent),
			widget.responseIcon(calEvent),
			timestamp,
			eventTitle,
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
		return fmt.Sprintf("[%s]",
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

func (widget *Widget) eventTimeColor() string {
	return widget.settings.colors.eventTime
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
	switch {
	case days > 0:
		untilStr = fmt.Sprintf("%dd", days)
	case hours > 0:
		untilStr = fmt.Sprintf("%dh", hours)
	default:
		untilStr = fmt.Sprintf("%dm", mins)
		if mins < 30 {
			color = "[red]"
		}
	}

	return color + untilStr + "[white]"
}

func (widget *Widget) titleColor(calEvent *CalEvent) string {
	color := widget.settings.colors.title

	for _, untypedArr := range widget.settings.colors.highlights {
		highlightElements := utils.ToStrs(untypedArr.([]interface{}))

		match, _ := regexp.MatchString(
			strings.ToLower(highlightElements[0]),
			strings.ToLower(calEvent.event.Summary),
		)

		if match {
			color = highlightElements[1]
		}
	}

	if calEvent.Past() {
		color = widget.settings.colors.past
	}

	return color
}

func (widget *Widget) location(calEvent *CalEvent) string {
	if !widget.settings.withLocation {
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
	if !widget.settings.displayResponseStatus {
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
