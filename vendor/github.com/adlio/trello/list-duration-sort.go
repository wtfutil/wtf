package trello

// ByFirstEntered is a slice of ListDurations
type ByFirstEntered []*ListDuration

// ByFirstEntered returns the length of the receiver.
func (durs ByFirstEntered) Len() int { return len(durs) }

// Less takes two indexes i and j and returns true exactly if the ListDuration
// at i was entered before j.
func (durs ByFirstEntered) Less(i, j int) bool {
	return durs[i].FirstEntered.Before(durs[j].FirstEntered)
}

// Swap takes two indexes i and j and swaps the ListDurations at the indexes.
func (durs ByFirstEntered) Swap(i, j int) { durs[i], durs[j] = durs[j], durs[i] }
