package helix

import (
	"strings"
	"time"
)

const (
	requestDateTimeFormat = "2006-01-02 15:04:05 -0700 MST"
)

var (
	datetimeFields = []string{"started_at", "ended_at"}
)

// Time is our custom time struct.
type Time struct {
	time.Time
}

// UnmarshalJSON is our custom datetime unmarshaller. Twitch sometimes
// returns datetimes as empty strings, which casuses issues with the native time
// UnmarshalJSON method when decoding the JSON string. Here we hanlde that scenario,
// by returning a zero time value for any JSON time field that is either an
// empty string or "null".
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	timeStr := strings.Trim(string(b), "\"")

	if timeStr == "" || timeStr == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(time.RFC3339, timeStr)

	return
}

func isDatetimeTagField(tag string) bool {
	for _, tagField := range datetimeFields {
		if tagField == tag {
			return true
		}
	}

	return false
}
