package clocks

import (
	"sort"
	"time"
)

type ClockCollection struct {
	Clocks []Clock
}

func (clocks *ClockCollection) Sorted(sortOrder string) []Clock {
	if sortOrder == "natural" {
		//no-op
	} else if sortOrder == "chronological" {
		clocks.SortedChronologically()
	} else if sortOrder == "reversechronological" {
		clocks.SortedReverseChronologically()
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
	now := time.Now()
	sort.Slice(clocks.Clocks, func(i, j int) bool {
		clock := clocks.Clocks[i]
		other := clocks.Clocks[j]

		return clock.ToLocal(now).String() < other.ToLocal(now).String()
	})
}

func (clocks *ClockCollection) SortedReverseChronologically() {
	now := time.Now()
	sort.Slice(clocks.Clocks, func(i, j int) bool {
		clock := clocks.Clocks[i]
		other := clocks.Clocks[j]

		return clock.ToLocal(now).String() > other.ToLocal(now).String()
	})
}
