package checklist

// Checklist is a module for creating generic checklist implementations
// See 'Todo' for an implementation example
type Checklist struct {
	Items []*ChecklistItem

	checkedIcon   string
	selected      int
	uncheckedIcon string
}

func NewChecklist(checkedIcon, uncheckedIcon string) Checklist {
	list := Checklist{
		checkedIcon:   checkedIcon,
		selected:      -1,
		uncheckedIcon: uncheckedIcon,
	}

	return list
}

/* -------------------- Exported Functions -------------------- */

// Add creates a new item in the checklist
func (list *Checklist) Add(checked bool, text string) {
	item := NewChecklistItem(
		checked,
		text,
		list.checkedIcon,
		list.uncheckedIcon,
	)

	list.Items = append([]*ChecklistItem{item}, list.Items...)
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
	list.Items = append(list.Items[:list.selected], list.Items[list.selected+1:]...)
	list.Prev()
}

// IsSelectable returns true if the checklist has selectable items, false if it does not
func (list *Checklist) IsSelectable() bool {
	return list.selected >= 0 && list.selected < len(list.Items)
}

// IsUnselectable returns true if the checklist has no selectable items, false if it does
func (list *Checklist) IsUnselectable() bool {
	return !list.IsSelectable()
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

func (list *Checklist) Selected() int {
	return list.selected
}

// SelectedItem returns the currently-selected checklist item or nil if no item is selected
func (list *Checklist) SelectedItem() *ChecklistItem {
	if list.IsUnselectable() {
		return nil
	}

	return list.Items[list.selected]
}

func (list *Checklist) SetSelectedByItem(selectableItem *ChecklistItem) {
	for idx, item := range list.Items {
		if item == selectableItem {
			list.selected = idx
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
	list.selected = -1
}

// Update sets the text of the currently-selected item to the provided text
func (list *Checklist) Update(text string) {
	item := list.SelectedItem()

	if item == nil {
		return
	}

	item.Text = text
}

/* -------------------- Item Movement -------------------- */

// Prev selects the previous item UP in the checklist
func (list *Checklist) Prev() {
	list.selected--
	if list.selected < 0 {
		list.selected = len(list.Items) - 1
	}
}

// Next selects the next item DOWN in the checklist
func (list *Checklist) Next() {
	list.selected++
	if list.selected >= len(list.Items) {
		list.selected = 0
	}
}

// Promote moves the selected item UP in the checklist
func (list *Checklist) Promote() {
	if list.IsUnselectable() {
		return
	}

	k := list.selected - 1
	if k < 0 {
		k = len(list.Items) - 1
	}

	list.Swap(list.selected, k)
	list.selected = k
}

// Demote moves the selected item DOWN in the checklist
func (list *Checklist) Demote() {
	if list.IsUnselectable() {
		return
	}

	j := list.selected + 1
	if j >= len(list.Items) {
		j = 0
	}

	list.Swap(list.selected, j)
	list.selected = j
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
