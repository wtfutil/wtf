package clocks

import (
	"time"
)

func Timezones(locations map[string]interface{}) map[string]time.Time {
	times := make(map[string]time.Time)

	for label, location := range locations {
		tzloc, err := time.LoadLocation(location.(string))

		if err != nil {
			continue
		}

		times[label] = time.Now().In(tzloc)
	}

	return times
}
