package todo

import ("fmt")

type List struct {
	Items []*Item

	selected int
}

func (list *List) Delete() {
	fmt.Println("del")
	list.Items = append(list.Items[:list.selected], list.Items[list.selected+1:]...)
}

func (list *List) Next() {
	list.selected = list.selected + 1
	if list.selected >= len(list.Items) {
		list.selected = 0
	}
}

func (list *List) Prev() {
	list.selected = list.selected - 1
	if list.selected < 0 {
		list.selected = len(list.Items) - 1
	}
}

// Toggle switches the checked state of the selected item
func (list *List) Toggle() {
	list.Items[list.selected].Toggle()
}

func (list *List) Unselect() {
	list.selected = -1
}

/* -------------------- Sort Interface -------------------- */

func (list *List) Len() int {
	return len(list.Items)
}

func (list *List) Less(i, j int) bool {
	return list.Items[i].Index < list.Items[j].Index
}

func (list *List) Swap(i, j int) {
	list.Items[i], list.Items[j] = list.Items[j], list.Items[i]
}
