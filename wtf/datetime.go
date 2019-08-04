package wtf

import (
	"fmt"
	"time"
)

const (
	// DateFormat defines the format we expect to receive dates from BambooHR in
	DateFormat = "2006-01-02"

	// TimeFormat defines the format we expect to receive times from BambooHR in
	TimeFormat = "15:04"
)

// IsToday returns TRUE if the date is today, FALSE if the date is not today
func IsToday(date time.Time) bool {
	now := time.Now().Local()

	return (date.Year() == now.Year()) &&
		(date.Month() == now.Month()) &&
		(date.Day() == now.Day())
}

// PrettyDate takes a programmer-style date string and converts it
// in a friendlier-to-read format
func PrettyDate(dateStr string) string {
	newTime, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return dateStr
	}

	return fmt.Sprint(newTime.Format("Jan 2, 2006"))
}

// UnixTime takes a Unix epoch time (in seconds) and returns a
// time.Time instance
func UnixTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
