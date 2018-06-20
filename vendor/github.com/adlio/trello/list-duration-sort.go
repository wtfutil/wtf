package trello

type ByFirstEntered []*ListDuration

func (durs ByFirstEntered) Len() int { return len(durs) }
func (durs ByFirstEntered) Less(i, j int) bool {
	return durs[i].FirstEntered.Before(durs[j].FirstEntered)
}
func (durs ByFirstEntered) Swap(i, j int) { durs[i], durs[j] = durs[j], durs[i] }
