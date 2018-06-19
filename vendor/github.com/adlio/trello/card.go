// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Card struct {
	client *Client

	// Key metadata
	ID               string     `json:"id"`
	IDShort          int        `json:"idShort"`
	Name             string     `json:"name"`
	Pos              float64    `json:"pos"`
	Email            string     `json:"email"`
	ShortLink        string     `json:"shortLink"`
	ShortUrl         string     `json:"shortUrl"`
	Url              string     `json:"url"`
	Desc             string     `json:"desc"`
	Due              *time.Time `json:"due"`
	DueComplete      bool       `json:"dueComplete"`
	Closed           bool       `json:"closed"`
	Subscribed       bool       `json:"subscribed"`
	DateLastActivity *time.Time `json:"dateLastActivity"`

	// Board
	Board   *Board
	IDBoard string `json:"idBoard"`

	// List
	List   *List
	IDList string `json:"idList"`

	// Badges
	Badges struct {
		Votes              int        `json:"votes"`
		ViewingMemberVoted bool       `json:"viewingMemberVoted"`
		Subscribed         bool       `json:"subscribed"`
		Fogbugz            string     `json:"fogbugz,omitempty"`
		CheckItems         int        `json:"checkItems"`
		CheckItemsChecked  int        `json:"checkItemsChecked"`
		Comments           int        `json:"comments"`
		Attachments        int        `json:"attachments"`
		Description        bool       `json:"description"`
		Due                *time.Time `json:"due,omitempty"`
	} `json:"badges"`

	// Actions
	Actions ActionCollection `json:"actions,omitempty"`

	// Checklists
	IDCheckLists    []string          `json:"idCheckLists"`
	Checklists      []*Checklist      `json:"checklists,omitempty"`
	CheckItemStates []*CheckItemState `json:"checkItemStates,omitempty"`

	// Members
	IDMembers      []string  `json:"idMembers,omitempty"`
	IDMembersVoted []string  `json:"idMembersVoted,omitempty"`
	Members        []*Member `json:"members,omitempty"`

	// Attachments
	IDAttachmentCover     string        `json:"idAttachmentCover"`
	ManualCoverAttachment bool          `json:"manualCoverAttachment"`
	Attachments           []*Attachment `json:"attachments,omitempty"`

	// Labels
	IDLabels []string `json:"idLabels,omitempty"`
	Labels   []*Label `json:"labels,omitempty"`
}

func (c *Card) CreatedAt() time.Time {
	t, _ := IDToTime(c.ID)
	return t
}

func (c *Card) MoveToList(listID string, args Arguments) error {
	path := fmt.Sprintf("cards/%s", c.ID)
	args["idList"] = listID
	return c.client.Put(path, args, &c)
}

func (c *Card) SetPos(newPos float64) error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, Arguments{"pos": fmt.Sprintf("%f", newPos)}, c)
}

func (c *Card) RemoveMember(memberID string) error {
	path := fmt.Sprintf("cards/%s/idMembers/%s", c.ID, memberID)
	return c.client.Delete(path, Defaults(), nil)
}

func (c *Card) AddMemberID(memberID string) (member []*Member, err error) {
	path := fmt.Sprintf("cards/%s/idMembers", c.ID)
	err = c.client.Post(path, Arguments{"value": memberID}, &member)
	return member, err
}

func (c *Card) MoveToTopOfList() error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, Arguments{"pos": "top"}, c)
}

func (c *Card) MoveToBottomOfList() error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, Arguments{"pos": "bottom"}, c)
}

func (c *Card) Update(args Arguments) error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, args, c)
}

func (c *Client) CreateCard(card *Card, extraArgs Arguments) error {
	path := "cards"
	args := Arguments{
		"name":      card.Name,
		"desc":      card.Desc,
		"pos":       strconv.FormatFloat(card.Pos, 'g', -1, 64),
		"idList":    card.IDList,
		"idMembers": strings.Join(card.IDMembers, ","),
		"idLabels":  strings.Join(card.IDLabels, ","),
	}
	if card.Due != nil {
		args["due"] = card.Due.Format(time.RFC3339)
	}
	// Allow overriding the creation position with 'top' or 'botttom'
	if pos, ok := extraArgs["pos"]; ok {
		args["pos"] = pos
	}
	err := c.Post(path, args, &card)
	if err == nil {
		card.client = c
	}
	return err
}

func (l *List) AddCard(card *Card, extraArgs Arguments) error {
	path := fmt.Sprintf("lists/%s/cards", l.ID)
	args := Arguments{
		"name":      card.Name,
		"desc":      card.Desc,
		"idMembers": strings.Join(card.IDMembers, ","),
		"idLabels":  strings.Join(card.IDLabels, ","),
	}
	if card.Due != nil {
		args["due"] = card.Due.Format(time.RFC3339)
	}
	// Allow overwriting the creation position with 'top' or 'bottom'
	if pos, ok := extraArgs["pos"]; ok {
		args["pos"] = pos
	}
	err := l.client.Post(path, args, &card)
	if err == nil {
		card.client = l.client
	} else {
		err = errors.Wrapf(err, "Error adding card to list %s", l.ID)
	}
	return err
}

// Try these Arguments
//
//	Arguments["keepFromSource"] = "all"
//  Arguments["keepFromSource"] = "none"
//	Arguments["keepFromSource"] = "attachments,checklists,comments"
//
func (c *Card) CopyToList(listID string, args Arguments) (*Card, error) {
	path := "cards"
	args["idList"] = listID
	args["idCardSource"] = c.ID
	newCard := Card{}
	err := c.client.Post(path, args, &newCard)
	if err == nil {
		newCard.client = c.client
	} else {
		err = errors.Wrapf(err, "Error copying card '%s' to list '%s'.", c.ID, listID)
	}
	return &newCard, err
}

func (c *Card) AddComment(comment string, args Arguments) (*Action, error) {
	path := fmt.Sprintf("cards/%s/actions/comments", c.ID)
	args["text"] = comment
	action := Action{}
	err := c.client.Post(path, args, &action)
	if err != nil {
		err = errors.Wrapf(err, "Error commenting on card %s", c.ID)
	}
	return &action, err
}

// If this Card was created from a copy of another Card, this func retrieves
// the originating Card. Returns an error only when a low-level failure occurred.
// If this Card has no parent, a nil card and nil error are returned. In other words, the
// non-existence of a parent is not treated as an error.
//
func (c *Card) GetParentCard(args Arguments) (*Card, error) {

	// Hopefully the card came pre-loaded with Actions including the card creation
	action := c.Actions.FirstCardCreateAction()

	if action == nil {
		// No luck. Go get copyCard actions for this card.
		c.client.log("Creation action wasn't supplied before GetParentCard() on '%s'. Getting copyCard actions.", c.ID)
		actions, err := c.GetActions(Arguments{"filter": "copyCard"})
		if err != nil {
			err = errors.Wrapf(err, "GetParentCard() failed to GetActions() for card '%s'", c.ID)
			return nil, err
		}
		action = actions.FirstCardCreateAction()
	}

	if action != nil && action.Data != nil && action.Data.CardSource != nil {
		card, err := c.client.GetCard(action.Data.CardSource.ID, args)
		return card, err
	}

	return nil, nil
}

func (c *Card) GetAncestorCards(args Arguments) (ancestors []*Card, err error) {

	// Get the first parent
	parent, err := c.GetParentCard(args)
	if IsNotFound(err) || IsPermissionDenied(err) {
		c.client.log("[trello] Can't get details about the parent of card '%s' due to lack of permissions or card deleted.", c.ID)
		return ancestors, nil
	}

	for parent != nil {
		ancestors = append(ancestors, parent)
		parent, err = parent.GetParentCard(args)
		if IsNotFound(err) || IsPermissionDenied(err) {
			c.client.log("[trello] Can't get details about the parent of card '%s' due to lack of permissions or card deleted.", c.ID)
			return ancestors, nil
		} else if err != nil {
			return ancestors, err
		}
	}

	return ancestors, err
}

func (c *Card) GetOriginatingCard(args Arguments) (*Card, error) {
	ancestors, err := c.GetAncestorCards(args)
	if err != nil {
		return c, err
	}
	if len(ancestors) > 0 {
		return ancestors[len(ancestors)-1], nil
	} else {
		return c, nil
	}
}

func (c *Card) CreatorMember() (*Member, error) {
	var actions ActionCollection
	var err error

	if len(c.Actions) == 0 {
		c.Actions, err = c.GetActions(Arguments{"filter": "all", "limit": "1000", "memberCreator_fields": "all"})
		if err != nil {
			err = errors.Wrapf(err, "GetActions() call failed.")
			return nil, err
		}
	}
	actions = c.Actions.FilterToCardCreationActions()

	if len(actions) > 0 {
		return actions[0].MemberCreator, nil
	}
	return nil, errors.Errorf("No card creation actions on Card %s with a .MemberCreator", c.ID)
}

func (c *Card) CreatorMemberID() (string, error) {

	var actions ActionCollection
	var err error

	if len(c.Actions) == 0 {
		c.client.log("[trello] CreatorMemberID() called on card '%s' without any Card.Actions. Fetching fresh.", c.ID)
		c.Actions, err = c.GetActions(Defaults())
		if err != nil {
			err = errors.Wrapf(err, "GetActions() call failed.")
		}
	}
	actions = c.Actions.FilterToCardCreationActions()

	if len(actions) > 0 {
		if actions[0].IDMemberCreator != "" {
			return actions[0].IDMemberCreator, err
		}
	}

	return "", errors.Wrapf(err, "No Actions on card '%s' could be used to find its creator.", c.ID)
}

func (b *Board) ContainsCopyOfCard(cardID string, args Arguments) (bool, error) {
	args["filter"] = "copyCard"
	actions, err := b.GetActions(args)
	if err != nil {
		err := errors.Wrapf(err, "GetCards() failed inside ContainsCopyOf() for board '%s' and card '%s'.", b.ID, cardID)
		return false, err
	}
	for _, action := range actions {
		if action.Data != nil && action.Data.CardSource != nil && action.Data.CardSource.ID == cardID {
			return true, nil
		}
	}
	return false, nil
}

func (c *Client) GetCard(cardID string, args Arguments) (card *Card, err error) {
	path := fmt.Sprintf("cards/%s", cardID)
	err = c.Get(path, args, &card)
	if card != nil {
		card.client = c
	}
	return card, err
}

/**
 * Retrieves all Cards on a Board
 *
 * If before
 */
func (b *Board) GetCards(args Arguments) (cards []*Card, err error) {
	path := fmt.Sprintf("boards/%s/cards", b.ID)

	err = b.client.Get(path, args, &cards)

	// Naive implementation would return here. To make sure we get all
	// cards, we begin
	if len(cards) > 0 {
		moreCards := true
		for moreCards == true {
			nextCardBatch := make([]*Card, 0)
			args["before"] = EarliestCardID(cards)
			err = b.client.Get(path, args, &nextCardBatch)
			if len(nextCardBatch) > 0 {
				cards = append(cards, nextCardBatch...)
			} else {
				moreCards = false
			}
		}
	}

	for i := range cards {
		cards[i].client = b.client
	}

	return
}

/**
 * Retrieves all Cards in a List
 */
func (l *List) GetCards(args Arguments) (cards []*Card, err error) {
	path := fmt.Sprintf("lists/%s/cards", l.ID)
	err = l.client.Get(path, args, &cards)
	for i := range cards {
		cards[i].client = l.client
	}
	return
}

func EarliestCardID(cards []*Card) string {
	if len(cards) == 0 {
		return ""
	}
	earliest := cards[0].ID
	for _, card := range cards {
		if card.ID < earliest {
			earliest = card.ID
		}
	}
	return earliest
}
