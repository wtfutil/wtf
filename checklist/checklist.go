package checklist

// Checklist is a module for creating generic checklist implementations
// See 'Todo' for an implementation example
type Checklist struct {
	Selected int

	Items []*ChecklistItem
}

func NewChecklist() Checklist {
	list := Checklist{
		Selected: -1,
	}

	return list
}

/* -------------------- Exported Functions -------------------- */

// Add creates a new item in the checklist
func (list *Checklist) Add(checked bool, text string) {
	item := ChecklistItem{
		Checked: checked,
		Text:    text,
	}

	list.Items = append([]*ChecklistItem{&item}, list.Items...)
}

// CheckedItems returns a slice of all the checked items
func (list *Checklist) CheckedItems() []*ChecklistItem {
	items := []*ChecklistItem{}

	for _, item := range list.Items {
		if item.Checked {
			items = append(items, item)
		}
	}

	return items
}

// Delete removes the selected item from the checklist
func (list *Checklist) Delete() {
	list.Items = append(list.Items[:list.Selected], list.Items[list.Selected+1:]...)
	list.Prev()
}

// Demote moves the selected item down in the checklist
func (list *Checklist) Demote() {
	if list.IsUnselectable() {
		return
	}

	j := list.Selected + 1
	if j >= len(list.Items) {
		j = 0
	}

	list.Swap(list.Selected, j)
	list.Selected = j
}

// IsSelectable returns true if the checklist has selectable items, false if it does not
func (list *Checklist) IsSelectable() bool {
	return list.Selected >= 0 && list.Selected < len(list.Items)
}

// IsUnselectable returns true if the checklist has no selectable items, false if it does
func (list *Checklist) IsUnselectable() bool {
	return !list.IsSelectable()
}

// Next selects the next item in the checklist
func (list *Checklist) Next() {
	list.Selected = list.Selected + 1
	if list.Selected >= len(list.Items) {
		list.Selected = 0
	}
}

// LongestLine returns the length of the longest checklist item's text
func (list *Checklist) LongestLine() int {
	maxLen := 0

	for _, item := range list.Items {
		if len(item.Text) > maxLen {
			maxLen = len(item.Text)
		}
	}

	return maxLen
}

// Prev selects the previous item in the checklist
func (list *Checklist) Prev() {
	list.Selected = list.Selected - 1
	if list.Selected < 0 {
		list.Selected = len(list.Items) - 1
	}
}

// Promote moves the selected item up in the checklist
func (list *Checklist) Promote() {
	if list.IsUnselectable() {
		return
	}

	j := list.Selected - 1
	if j < 0 {
		j = len(list.Items) - 1
	}

	list.Swap(list.Selected, j)
	list.Selected = j
}

// SelectedItem returns the currently-selected checklist item or nil if no item is selected
func (list *Checklist) SelectedItem() *ChecklistItem {
	if list.IsUnselectable() {
		return nil
	}

	return list.Items[list.Selected]
}

func (list *Checklist) SetSelectedByItem(selectableItem *ChecklistItem) {
	for idx, item := range list.Items {
		if item == selectableItem {
			list.Selected = idx
			break
		}
	}
}

// Toggle switches the checked state of the currently-selected item
func (list *Checklist) Toggle() {
	if list.IsUnselectable() {
		return
	}

	list.SelectedItem().Toggle()
}

// UncheckedItems returns a slice of all the unchecked items
func (list *Checklist) UncheckedItems() []*ChecklistItem {
	items := []*ChecklistItem{}

	for _, item := range list.Items {
		if !item.Checked {
			items = append(items, item)
		}
	}

	return items
}

// Unselect removes the current select such that no item is selected
func (list *Checklist) Unselect() {
	list.Selected = -1
}

// Update sets the text of the currently-selected item to the provided text
func (list *Checklist) Update(text string) {
	item := list.SelectedItem()

	if item == nil {
		return
	}

	item.Text = text
}

/* -------------------- Sort Interface -------------------- */

func (list *Checklist) Len() int {
	return len(list.Items)
}

func (list *Checklist) Less(i, j int) bool {
	return list.Items[i].Text < list.Items[j].Text
}

func (list *Checklist) Swap(i, j int) {
	list.Items[i], list.Items[j] = list.Items[j], list.Items[i]
}
