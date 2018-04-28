package todo

import (
	"fmt"
)

type List struct {
	Items []*Item

	selected int
}

/* -------------------- Exported Functions -------------------- */

func (list *List) Add(text string) {
	item := Item{
		Checked: false,
		Text:    text,
	}

	list.Items = append([]*Item{&item}, list.Items...)
	fmt.Println("added")
}

func (list *List) Delete() {
	list.Items = append(list.Items[:list.selected], list.Items[list.selected+1:]...)
}

func (list *List) Demote() {
	j := list.selected + 1
	if j >= len(list.Items) {
		j = 0
	}

	list.Swap(list.selected, j)
	list.selected = j
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

func (list *List) Promote() {
	j := list.selected - 1
	if j < 0 {
		j = len(list.Items) - 1
	}

	list.Swap(list.selected, j)
	list.selected = j
}

func (list *List) Selected() *Item {
	if list.selected < 0 || list.selected >= len(list.Items) {
		return nil
	}

	return list.Items[list.selected]
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
	return list.Items[i].Text < list.Items[j].Text
}

func (list *List) Swap(i, j int) {
	list.Items[i], list.Items[j] = list.Items[j], list.Items[i]
}
