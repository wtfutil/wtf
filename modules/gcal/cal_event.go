package gcal

import (
	"fmt"
	"time"

	"github.com/wtfutil/wtf/utils"
	"google.golang.org/api/calendar/v3"
)

type CalEvent struct {
	event *calendar.Event
}

func NewCalEvent(event *calendar.Event) *CalEvent {
	calEvent := CalEvent{
		event: event,
	}

	return &calEvent
}

/* -------------------- Exported Functions -------------------- */

func (calEvent *CalEvent) AllDay() bool {
	return len(calEvent.event.Start.Date) > 0
}

func (calEvent *CalEvent) ConflictsWith(otherEvents []*CalEvent) bool {
	hasConflict := false

	for _, otherEvent := range otherEvents {
		if calEvent.event == otherEvent.event {
			continue
		}

		if calEvent.Start().Before(otherEvent.End()) && calEvent.End().After(otherEvent.Start()) {
			hasConflict = true
			break
		}
	}

	return hasConflict
}

func (calEvent *CalEvent) Now() bool {
	return time.Now().After(calEvent.Start()) && time.Now().Before(calEvent.End())
}

func (calEvent *CalEvent) Past() bool {
	if calEvent.AllDay() {
		// FIXME: This should calculate properly
		return false
	}

	return !calEvent.Now() && calEvent.Start().Before(time.Now())
}

func (calEvent *CalEvent) ResponseFor(email string) string {
	for _, attendee := range calEvent.event.Attendees {
		if attendee.Email == email {
			return attendee.ResponseStatus
		}
	}

	return ""
}

/* -------------------- DateTimes -------------------- */

func (calEvent *CalEvent) End() time.Time {
	var calcTime string
	var end time.Time

	if calEvent.AllDay() {
		calcTime = calEvent.event.End.Date
		end, _ = time.ParseInLocation("2006-01-02", calcTime, time.Local)
	} else {
		calcTime = calEvent.event.End.DateTime
		end, _ = time.Parse(time.RFC3339, calcTime)
	}

	return end
}

func (calEvent *CalEvent) Start() time.Time {
	var calcTime string
	var start time.Time

	if calEvent.AllDay() {
		calcTime = calEvent.event.Start.Date
		start, _ = time.ParseInLocation("2006-01-02", calcTime, time.Local)
	} else {
		calcTime = calEvent.event.Start.DateTime
		start, _ = time.Parse(time.RFC3339, calcTime)
	}

	return start
}

func (calEvent *CalEvent) Timestamp(hourFormat string, showEndTime bool) string {
	if calEvent.AllDay() {
		startTime, _ := time.ParseInLocation("2006-01-02", calEvent.event.Start.Date, time.Local)
		return startTime.Format(utils.FriendlyDateFormat)
	}

	startTime, _ := time.Parse(time.RFC3339, calEvent.event.Start.DateTime)
	endTime, _ := time.Parse(time.RFC3339, calEvent.event.End.DateTime)

	timeFormat := utils.MinimumTimeFormat24
	if hourFormat == "12" {
		timeFormat = utils.MinimumTimeFormat12
	}

	if showEndTime {
		return fmt.Sprintf("%s-%s", startTime.Format(timeFormat), endTime.Format(timeFormat))
	}

	return startTime.Format(timeFormat)
}
