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

// Card represents the card resource.
// https://developers.trello.com/reference/#card-object
type Card struct {
	client *Client

	// Key metadata
	ID               string     `json:"id"`
	IDShort          int        `json:"idShort"`
	Name             string     `json:"name"`
	Pos              float64    `json:"pos"`
	Email            string     `json:"email"`
	ShortLink        string     `json:"shortLink"`
	ShortURL         string     `json:"shortUrl"`
	URL              string     `json:"url"`
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

	// Custom Fields
	CustomFieldItems []*CustomFieldItem `json:"customFieldItems,omitempty"`

	customFieldMap *map[string]interface{}
}

// CreatedAt returns the receiver card's created-at attribute as time.Time.
func (c *Card) CreatedAt() time.Time {
	t, _ := IDToTime(c.ID)
	return t
}

// CustomFields returns the card's custom fields.
func (c *Card) CustomFields(boardCustomFields []*CustomField) map[string]interface{} {

	cfm := c.customFieldMap

	if cfm == nil {
		cfm = &(map[string]interface{}{})

		// bcfOptionNames[CustomField ID] = Custom Field Name
		bcfOptionNames := map[string]string{}

		// bcfOptionsMap[CustomField ID][ID of the option] = Value of the option
		bcfOptionsMap := map[string]map[string]interface{}{}

		for _, bcf := range boardCustomFields {
			bcfOptionNames[bcf.ID] = bcf.Name
			for _, cf := range bcf.Options {
				// create 2nd level map when not available yet
				map2, ok := bcfOptionsMap[cf.IDCustomField]
				if !ok {
					map2 = map[string]interface{}{}
					bcfOptionsMap[bcf.ID] = map2
				}

				bcfOptionsMap[bcf.ID][cf.ID] = cf.Value.Text
			}
		}

		for _, cf := range c.CustomFieldItems {
			name := bcfOptionNames[cf.IDCustomField]

			// create 2nd level map when not available yet
			map2, ok := bcfOptionsMap[cf.IDCustomField]
			if !ok {
				continue
			}
			value, ok := map2[cf.IDValue]

			if ok {
				(*cfm)[name] = value
			}
		}
		c.customFieldMap = cfm
	}
	return *cfm
}

// MoveToList moves a card to a list given by listID.
func (c *Card) MoveToList(listID string, args Arguments) error {
	path := fmt.Sprintf("cards/%s", c.ID)
	args["idList"] = listID
	return c.client.Put(path, args, &c)
}

// SetPos sets a card's new position.
func (c *Card) SetPos(newPos float64) error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, Arguments{"pos": fmt.Sprintf("%f", newPos)}, c)
}

// RemoveMember receives the id of a member and removes the corresponding member from the card.
func (c *Card) RemoveMember(memberID string) error {
	path := fmt.Sprintf("cards/%s/idMembers/%s", c.ID, memberID)
	return c.client.Delete(path, Defaults(), nil)
}

// AddMemberID receives a member id and adds the corresponding member to the card.
// Returns a list of the card's members or an error.
func (c *Card) AddMemberID(memberID string) (member []*Member, err error) {
	path := fmt.Sprintf("cards/%s/idMembers", c.ID)
	err = c.client.Post(path, Arguments{"value": memberID}, &member)
	return member, err
}

// RemoveIDLabel removes a label id from the card.
func (c *Card) RemoveIDLabel(labelID string, label *Label) error {
	path := fmt.Sprintf("cards/%s/idLabels/%s", c.ID, labelID)
	return c.client.Delete(path, Defaults(), label)

}

// AddIDLabel receives a label id and adds the corresponding label or returns an error.
func (c *Card) AddIDLabel(labelID string) error {
	path := fmt.Sprintf("cards/%s/idLabels", c.ID)
	err := c.client.Post(path, Arguments{"value": labelID}, &c.IDLabels)
	return err
}

// MoveToTopOfList moves the card to the top of it's list.
func (c *Card) MoveToTopOfList() error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, Arguments{"pos": "top"}, c)
}

// MoveToBottomOfList moves the card to the bottom of its list.
func (c *Card) MoveToBottomOfList() error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, Arguments{"pos": "bottom"}, c)
}

// Update UPDATEs the card's attributes.
func (c *Card) Update(args Arguments) error {
	path := fmt.Sprintf("cards/%s", c.ID)
	return c.client.Put(path, args, c)
}

// CreateCard takes a Card and Arguments and POSTs the card.
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

// AddCard takes a Card and Arguments and adds the card to the receiver list.
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

// CopyToList takes a list id and Arguments and returns the matching Card.
// The following Arguments are supported.
//
//	Arguments["keepFromSource"] = "all"
//  Arguments["keepFromSource"] = "none"
//	Arguments["keepFromSource"] = "attachments,checklists,comments"
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

// AddComment takes a comment string and Arguments and adds the comment to the card.
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

// AddURLAttachment takes an Attachment and adds it to the card.
func (c *Card) AddURLAttachment(attachment *Attachment) error {
	path := fmt.Sprintf("cards/%s/attachments", c.ID)
	args := Arguments{
		"url":  attachment.URL,
		"name": attachment.Name,
	}
	err := c.client.Post(path, args, &attachment)
	if err != nil {
		err = errors.Wrapf(err, "Error adding attachment to card %s", c.ID)
	}
	return err

}

// GetParentCard retrieves the originating Card if the Card was created
// from a copy of another Card. Returns an error only when a low-level failure occurred.
// If this Card has no parent, a nil card and nil error are returned. In other words, the
// non-existence of a parent is not treated as an error.
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

// GetAncestorCards takes Arguments, GETs the card's ancestors and returns them as a slice.
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

// GetOriginatingCard takes Arguments, GETs ancestors and returns most recent ancestor card of the Card.
func (c *Card) GetOriginatingCard(args Arguments) (*Card, error) {
	ancestors, err := c.GetAncestorCards(args)
	if err != nil {
		return c, err
	}
	if len(ancestors) > 0 {
		return ancestors[len(ancestors)-1], nil
	}

	return c, nil
}

// CreatorMember returns the member of the card who created it or and error.
// The creator is the member who is associated with the card's first action.
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

// CreatorMemberID returns as string the id of the member who created the card or an error.
// The creator is the member who is associated with the card's first action.
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

// ContainsCopyOfCard accepts a card id and Arguments and returns true
// if the receiver Board contains a Card with the id.
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

// GetCard receives a card id and Arguments and returns the card if found
// with the credentials given for the receiver Client. Returns an error
// otherwise.
func (c *Client) GetCard(cardID string, args Arguments) (card *Card, err error) {
	path := fmt.Sprintf("cards/%s", cardID)
	err = c.Get(path, args, &card)
	if card != nil {
		card.client = c
	}
	return card, err
}

// GetCards takes Arguments and retrieves all Cards on a Board as slice or returns error.
func (b *Board) GetCards(args Arguments) (cards []*Card, err error) {
	path := fmt.Sprintf("boards/%s/cards", b.ID)

	err = b.client.Get(path, args, &cards)

	// Naive implementation would return here. To make sure we get all
	// cards, we begin
	if len(cards) > 0 {
		moreCards := true
		for moreCards == true {
			nextCardBatch := make([]*Card, 0)
			args["before"] = earliestCardID(cards)
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

// GetCards retrieves all Cards in a List or an error if something goes wrong.
func (l *List) GetCards(args Arguments) (cards []*Card, err error) {
	path := fmt.Sprintf("lists/%s/cards", l.ID)
	err = l.client.Get(path, args, &cards)
	for i := range cards {
		cards[i].client = l.client
	}
	return
}

func earliestCardID(cards []*Card) string {
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
