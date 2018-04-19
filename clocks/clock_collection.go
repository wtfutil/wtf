package clocks

import (
	"sort"
)

type ClockCollection struct {
	Clocks []Clock
}

func (clockColl *ClockCollection) Sorted() []Clock {
	if "chronological" == Config.UString("wtf.mods.clocks.sort", "alphabetical") {
		clockColl.SortedChronologically()
	} else {
		clockColl.SortedAlphabetically()
	}

	return clockColl.Clocks
}

func (clockColl *ClockCollection) SortedAlphabetically() {
	sort.Slice(clockColl.Clocks, func(i, j int) bool {
		clock := clockColl.Clocks[i]
		other := clockColl.Clocks[j]

		return clock.Label < other.Label
	})
}

func (clockColl *ClockCollection) SortedChronologically() {
	sort.Slice(clockColl.Clocks, func(i, j int) bool {
		clock := clockColl.Clocks[i]
		other := clockColl.Clocks[j]

		return clock.LocalTime.Before(other.LocalTime)
	})
}
