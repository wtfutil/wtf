// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

// Checklist represents Trello card's checklists.
// A card can have one zero or more checklists.
// https://developers.trello.com/reference/#checklist-object
type Checklist struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	IDBoard    string      `json:"idBoard,omitempty"`
	IDCard     string      `json:"idCard,omitempty"`
	Pos        float64     `json:"pos,omitempty"`
	CheckItems []CheckItem `json:"checkItems,omitempty"`
	client *Client
}

// CheckItem is a nested resource representing an item in Checklist.
type CheckItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"state"`
	IDChecklist string  `json:"idChecklist,omitempty"`
	Pos         float64 `json:"pos,omitempty"`
}

// CheckItemState represents a CheckItem when it appears in CheckItemStates on a Card.
type CheckItemState struct {
	IDCheckItem string `json:"idCheckItem"`
	State       string `json:"state"`
}

// CreateChecklist creates a checklist.
// Attribute currently supported as extra argument: pos.
// Attributes currently known to be unsupported: idChecklistSource.
//
// API Docs: https://developers.trello.com/reference#cardsidchecklists-1
func (c *Client) CreateChecklist(card *Card, name string, extraArgs Arguments) (checklist *Checklist, err error) {
	path := "cards/" + card.ID + "/checklists"
	args := Arguments{
		"name": name,
		"pos": "bottom",
	}

	if pos, ok := extraArgs["pos"]; ok{
		args["pos"] = pos
	}

	checklist = &Checklist{}
	err = c.Post(path, args, &checklist)
	if err == nil {
		checklist.client = c
		checklist.IDCard = card.ID
		card.Checklists = append(card.Checklists, checklist)
	}
	return
}

// CreateCheckItem creates a checkitem inside the checklist.
// Attribute currently supported as extra argument: pos.
// Attributes currently known to be unsupported: checked.
//
// API Docs: https://developers.trello.com/reference#checklistsidcheckitems
func (cl *Checklist) CreateCheckItem(name string, extraArgs Arguments) (item *CheckItem, err error) {
	return cl.client.CreateCheckItem(cl, name, extraArgs)
}

// CreateCheckItem creates a checkitem inside the given checklist.
// Attribute currently supported as extra argument: pos.
// Attributes currently known to be unsupported: checked.
//
// API Docs: https://developers.trello.com/reference#checklistsidcheckitems
func (c *Client) CreateCheckItem(checklist *Checklist, name string, extraArgs Arguments) (item *CheckItem, err error) {
	path := "checklists/" + checklist.ID + "/checkItems"
	args := Arguments {
		"name": name,
		"pos": "bottom",
		"checked": "false",
	}

	if pos, ok := extraArgs["pos"]; ok{
		args["pos"] = pos
	}
	if checked, ok := extraArgs["checked"]; ok {
		args["checked"] = checked
	}

	item = &CheckItem{}
	err = c.Post(path, args, item)
	if err == nil {
		checklist.CheckItems = append(checklist.CheckItems, *item)
	}
	return
}
