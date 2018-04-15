package status

import (
	"time"
)

func Timezones(timezones map[string]interface{}) map[string]time.Time {
	times := make(map[string]time.Time)

	for label, timezone := range timezones {
		tzloc, err := time.LoadLocation(timezone.(string))

		if err != nil {
			continue
		}

		times[label] = time.Now().In(tzloc)
	}

	return times
}
