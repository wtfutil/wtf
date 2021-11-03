package clocks

import (
	"strings"
	"time"
)

type Clock struct {
	Label    string
	Location *time.Location
}

func NewClock(label string, timeLoc *time.Location) Clock {
	clock := Clock{
		Label:    label,
		Location: timeLoc,
	}

	return clock
}

func BuildClock(label string, location string) (clock Clock, err error) {
	timeLoc, err := time.LoadLocation(sanitizeLocation(location))
	if err != nil {
		return Clock{}, err
	}
	return NewClock(label, timeLoc), nil
}

func (clock *Clock) Date(dateFormat string) string {
	return clock.LocalTime().Format(dateFormat)
}

func (clock *Clock) LocalTime() time.Time {
	return clock.ToLocal(time.Now())
}

func (clock *Clock) ToLocal(t time.Time) time.Time {
	return t.In(clock.Location)
}

func (clock *Clock) Time(timeFormat string) string {
	return clock.LocalTime().Format(timeFormat)
}

func sanitizeLocation(locStr string) string {
	return strings.ReplaceAll(locStr, " ", "_")
}
