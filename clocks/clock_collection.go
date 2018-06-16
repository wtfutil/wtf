package clocks

import (
	"sort"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
)

type ClockCollection struct {
	Clocks []Clock
}

func (clocks *ClockCollection) Sorted() []Clock {
	if "chronological" == wtf.Config.UString("wtf.mods.clocks.sort", "alphabetical") {
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
	now := time.Now()
	sort.Slice(clocks.Clocks, func(i, j int) bool {
		clock := clocks.Clocks[i]
		other := clocks.Clocks[j]

		return clock.ToLocal(now).String() < other.ToLocal(now).String()
	})
}
