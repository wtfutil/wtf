package trello

import (
	"sort"
)

// ActionCollection is an alias of []*Action, which sorts by the Action's ID.
// Which is the same as sorting by the Action's time of occurrence
type ActionCollection []*Action

func (c ActionCollection) Len() int           { return len(c) }
func (c ActionCollection) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ActionCollection) Less(i, j int) bool { return c[i].ID < c[j].ID }

// FirstCardCreateAction returns first card-create action
func (c ActionCollection) FirstCardCreateAction() *Action {
	sort.Sort(c)
	for _, action := range c {
		if action.DidCreateCard() {
			return action
		}
	}
	return nil
}

// ContainsCardCreation returns true if collection contains a card-create action
func (c ActionCollection) ContainsCardCreation() bool {
	return c.FirstCardCreateAction() != nil
}

// FilterToCardCreationActions returns this collection's card-create actions
func (c ActionCollection) FilterToCardCreationActions() ActionCollection {
	newSlice := make(ActionCollection, 0, len(c))
	for _, action := range c {
		if action.DidCreateCard() {
			newSlice = append(newSlice, action)
		}
	}
	return newSlice
}

// FilterToListChangeActions returns card-change-list actions
func (c ActionCollection) FilterToListChangeActions() ActionCollection {
	newSlice := make(ActionCollection, 0, len(c))
	for _, action := range c {
		if action.DidChangeListForCard() {
			newSlice = append(newSlice, action)
		}
	}
	return newSlice
}

// FilterToCardMembershipChangeActions returns the collection's card-change, archive and unarchive actions
func (c ActionCollection) FilterToCardMembershipChangeActions() ActionCollection {
	newSlice := make(ActionCollection, 0, len(c))
	for _, action := range c {
		if action.DidChangeCardMembership() || action.DidArchiveCard() || action.DidUnarchiveCard() {
			newSlice = append(newSlice, action)
		}
	}
	return newSlice
}
