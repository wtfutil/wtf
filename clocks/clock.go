package clocks

import (
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
