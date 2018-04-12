package status

import (
	"time"
)

func Timezones(zones []string) []time.Time {
	times := []time.Time{}

	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			continue
		}

		times = append(times, time.Now().In(loc))
	}

	return times
}
