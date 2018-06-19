// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

type Action struct {
	ID              string      `json:"id"`
	IDMemberCreator string      `json:"idMemberCreator"`
	Type            string      `json:"type"`
	Date            time.Time   `json:"date"`
	Data            *ActionData `json:"data,omitempty"`
	MemberCreator   *Member     `json:"memberCreator,omitempty"`
	Member          *Member     `json:"member,omitempty"`
}

type ActionData struct {
	Text           string    `json:"text,omitempty"`
	List           *List     `json:"list,omitempty"`
	Card           *Card     `json:"card,omitempty"`
	CardSource     *Card     `json:"cardSource,omitempty"`
	Board          *Board    `json:"board,omitempty"`
	Old            *Card     `json:"old,omitempty"`
	ListBefore     *List     `json:"listBefore,omitempty"`
	ListAfter      *List     `json:"listAfter,omitempty"`
	DateLastEdited time.Time `json:"dateLastEdited"`

	CheckItem *CheckItem `json:"checkItem"`
	Checklist *Checklist `json:"checklist"`
}

func (b *Board) GetActions(args Arguments) (actions ActionCollection, err error) {
	path := fmt.Sprintf("boards/%s/actions", b.ID)
	err = b.client.Get(path, args, &actions)
	return
}

func (l *List) GetActions(args Arguments) (actions ActionCollection, err error) {
	path := fmt.Sprintf("lists/%s/actions", l.ID)
	err = l.client.Get(path, args, &actions)
	return
}

func (c *Card) GetActions(args Arguments) (actions ActionCollection, err error) {
	path := fmt.Sprintf("cards/%s/actions", c.ID)
	err = c.client.Get(path, args, &actions)
	return
}

// GetListChangeActions retrieves a slice of Actions which resulted in changes
// to the card's active List. This includes the createCard and copyCard action (which
// place the card in its first list, and the updateCard:closed action, which remove it
// from its last list.
//
// This function is just an alias for:
//   card.GetActions(Arguments{"filter": "createCard,copyCard,updateCard:idList,updateCard:closed", "limit": "1000"})
//
func (c *Card) GetListChangeActions() (actions ActionCollection, err error) {
	return c.GetActions(Arguments{"filter": "createCard,copyCard,updateCard:idList,updateCard:closed"})
}

func (c *Card) GetMembershipChangeActions() (actions ActionCollection, err error) {
	// We include updateCard:closed as if the member is implicitly removed from the card when it's closed.
	// This allows us to "close out" the duration length.
	return c.GetActions(Arguments{"filter": "addMemberToCard,removeMemberFromCard,updateCard:closed"})
}

// DidCreateCard() returns true if this action created a card, false otherwise.
func (a *Action) DidCreateCard() bool {
	switch a.Type {
	case "createCard", "emailCard", "copyCard", "convertToCardFromCheckItem":
		return true
	case "moveCardToBoard":
		return true // Unsure about this one
	default:
		return false
	}
}

func (a *Action) DidArchiveCard() bool {
	return (a.Type == "updateCard") && a.Data != nil && a.Data.Card != nil && a.Data.Card.Closed
}

func (a *Action) DidUnarchiveCard() bool {
	return (a.Type == "updateCard") && a.Data != nil && a.Data.Old != nil && a.Data.Old.Closed
}

// Returns true if this action created the card (in which case it caused it to enter its
// first list), archived the card (in which case it caused it to leave its last List),
// or was an updateCard action involving a change to the list. This is supporting
// functionality for ListDuration.
//
func (a *Action) DidChangeListForCard() bool {
	if a.DidCreateCard() {
		return true
	}
	if a.DidArchiveCard() {
		return true
	}
	if a.DidUnarchiveCard() {
		return true
	}
	if a.Type == "updateCard" {
		if a.Data != nil && a.Data.ListAfter != nil {
			return true
		}
	}
	return false
}

func (a *Action) DidChangeCardMembership() bool {
	switch a.Type {
	case "addMemberToCard":
		return true
	case "removeMemberFromCard":
		return true
	default:
		return false
	}
}

// ListAfterAction calculates which List the card ended up in after this action
// completed. Returns nil when the action resulted in the card being archived (in
// which case we consider it to not be in a list anymore), or when the action isn't
// related to a list at all (in which case this is a nonsensical question to ask).
//
func ListAfterAction(a *Action) *List {
	switch a.Type {
	case "createCard", "copyCard", "emailCard", "convertToCardFromCheckItem":
		return a.Data.List
	case "updateCard":
		if a.DidArchiveCard() {
			return nil
		} else if a.DidUnarchiveCard() {
			return a.Data.List
		}
		if a.Data.ListAfter != nil {
			return a.Data.ListAfter
		}
	}
	return nil
}
