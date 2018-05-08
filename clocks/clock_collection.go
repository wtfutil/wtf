package clocks

import (
	"sort"
)

type ClockCollection struct {
	Clocks []Clock
}

func (clocks *ClockCollection) Sorted() []Clock {
	if "chronological" == Config.UString("wtf.mods.clocks.sort", "alphabetical") {
		clocks.SortedChronologically()
	} else {
		clocks.SortedAlphabetically()
	}

	return clocks.Clocks
}

func (clocks *ClockCollection) SortedAlphabetically() {
	sort.Slice(clocks.Clocks, func(i, j int) bool {
		clock := clocks.Clocks[i]
		other := clocks.Clocks[j]

		return clock.Label < other.Label
	})
}

func (clocks *ClockCollection) SortedChronologically() {
	sort.Slice(clocks.Clocks, func(i, j int) bool {
		clock := clocks.Clocks[i]
		other := clocks.Clocks[j]

		return clock.LocalTime.Before(other.LocalTime)
	})
}
