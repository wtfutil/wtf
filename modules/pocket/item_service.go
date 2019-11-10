package pocket

import "sort"

type sortByTimeAdded []Item

func (a sortByTimeAdded) Len() int           { return len(a) }
func (a sortByTimeAdded) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByTimeAdded) Less(i, j int) bool { return a[i].TimeAdded > a[j].TimeAdded }

func orderItemResponseByKey(response ItemLists) []Item {

	var items sortByTimeAdded
	for _, v := range response.List {
		items = append(items, v)
	}
	sort.Sort(items)
	return items
}
