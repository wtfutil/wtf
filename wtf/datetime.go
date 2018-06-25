package wtf

import (
	"fmt"
	"time"
)

// DateFormat defines the format we expect to receive dates from BambooHR in
const DateFormat = "2006-01-02"
const TimeFormat = "15:04"

func IsToday(date time.Time) bool {
	now := Now()

	return (date.Year() == now.Year()) &&
		(date.Month() == now.Month()) &&
		(date.Day() == now.Day())
}

func Now() time.Time {
	return time.Now().Local()
}

func PrettyDate(dateStr string) string {
	newTime, _ := time.Parse(DateFormat, dateStr)
	return fmt.Sprint(newTime.Format("Jan 2, 2006"))
}

func Tomorrow() time.Time {
	return Now().AddDate(0, 0, 1)
}

func UnixTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
