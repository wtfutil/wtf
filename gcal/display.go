package gcal

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
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
	if widget.calEvents == nil || len(widget.calEvents) == 0 {
		return
	}

	widget.mutex.Lock()
	defer widget.mutex.Unlock()

	widget.View.SetTitle(widget.ContextualTitle(widget.Name))
	widget.View.SetText(widget.contentFrom(widget.calEvents))
}

func (widget *Widget) contentFrom(calEvents []*CalEvent) string {
	if (calEvents == nil) || (len(calEvents) == 0) {
		return ""
	}

	var str string
	var prevEvent *CalEvent

	if !wtf.Config.UBool("wtf.mods.gcal.showDeclined", false) {
		calEvents = removeDeclined(calEvents)
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

		str = str + fmt.Sprintf("%s   %s%s\n",
			lineOne,
			widget.location(calEvent),
			widget.timeUntil(calEvent),
		)

		if (widget.location(calEvent) != "") || (widget.timeUntil(calEvent) != "") {
			str = str + "\n"
		}

		prevEvent = calEvent
	}

	return str
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
			wtf.Config.UString("wtf.mods.gcal.colors.day", "forestgreen")) +
			event.Start().Format(wtf.FullDateFormat) +
			"\n"
	}

	return ""
}

func (widget *Widget) descriptionColor(calEvent *CalEvent) string {
	if calEvent.Past() {
		return wtf.Config.UString("wtf.mods.gcal.colors.past", "gray")
	}

	return wtf.Config.UString("wtf.mods.gcal.colors.description", "white")
}

func (widget *Widget) eventSummary(calEvent *CalEvent, conflict bool) string {
	summary := calEvent.event.Summary

	if calEvent.Now() {
		summary = fmt.Sprintf(
			"%s %s",
			wtf.Config.UString("wtf.mods.gcal.currentIcon", "🔸"),
			summary,
		)
	}

	if conflict {
		return fmt.Sprintf("%s %s", wtf.Config.UString("wtf.mods.gcal.conflictIcon", "🚨"), summary)
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
	color := wtf.Config.UString("wtf.mods.gcal.colors.title", "white")

	for _, untypedArr := range wtf.Config.UList("wtf.mods.gcal.colors.highlights") {
		highlightElements := wtf.ToStrs(untypedArr.([]interface{}))

		match, _ := regexp.MatchString(
			strings.ToLower(highlightElements[0]),
			strings.ToLower(calEvent.event.Summary),
		)

		if match == true {
			color = highlightElements[1]
		}
	}

	if calEvent.Past() {
		color = wtf.Config.UString("wtf.mods.gcal.colors.past", "gray")
	}

	return color
}

func (widget *Widget) location(calEvent *CalEvent) string {
	if wtf.Config.UBool("wtf.mods.gcal.displayLocation", true) == false {
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
	if false == wtf.Config.UBool("wtf.mods.gcal.displayResponseStatus", true) {
		return ""
	}

	icon := "[gray]"

	switch calEvent.ResponseFor(wtf.Config.UString("wtf.mods.gcal.email")) {
	case "accepted":
		return icon + "✔︎"
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

func removeDeclined(events []*CalEvent) []*CalEvent {
	var ret []*CalEvent
	for _, e := range events {
		if e.ResponseFor(wtf.Config.UString("wtf.mods.gcal.email")) != "declined" {
			ret = append(ret, e)
		}
	}
	return ret
}
