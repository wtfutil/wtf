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

func (actions ActionCollection) FirstCardCreateAction() *Action {
	sort.Sort(actions)
	for _, action := range actions {
		if action.DidCreateCard() {
			return action
		}
	}
	return nil
}

func (actions ActionCollection) ContainsCardCreation() bool {
	return actions.FirstCardCreateAction() != nil
}

func (c ActionCollection) FilterToCardCreationActions() ActionCollection {
	newSlice := make(ActionCollection, 0, len(c))
	for _, action := range c {
		if action.DidCreateCard() {
			newSlice = append(newSlice, action)
		}
	}
	return newSlice
}

func (c ActionCollection) FilterToListChangeActions() ActionCollection {
	newSlice := make(ActionCollection, 0, len(c))
	for _, action := range c {
		if action.DidChangeListForCard() {
			newSlice = append(newSlice, action)
		}
	}
	return newSlice
}

func (c ActionCollection) FilterToCardMembershipChangeActions() ActionCollection {
	newSlice := make(ActionCollection, 0, len(c))
	for _, action := range c {
		if action.DidChangeCardMembership() || action.DidArchiveCard() || action.DidUnarchiveCard() {
			newSlice = append(newSlice, action)
		}
	}
	return newSlice
}
