// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

type SearchResult struct {
	Options SearchOptions `json:"options"`
	Actions []*Action     `json:"actions,omitempty"`
	Cards   []*Card       `json:"cards,omitempty"`
	Boards  []*Board      `json:"boards,omitempty"`
	Members []*Member     `json:"members,omitempty"`
}

type SearchOptions struct {
	Terms      []SearchTerm     `json:"terms"`
	Modifiers  []SearchModifier `json:"modifiers,omitempty"`
	ModelTypes []string         `json:"modelTypes,omitempty"`
	Partial    bool             `json:"partial"`
}

type SearchModifier struct {
	Text string `json:"text"`
}

type SearchTerm struct {
	Text    string `json:"text"`
	Negated bool   `json:"negated,omitempty"`
}

func (c *Client) SearchCards(query string, args Arguments) (cards []*Card, err error) {
	args["query"] = query
	args["modelTypes"] = "cards"
	res := SearchResult{}
	err = c.Get("search", args, &res)
	cards = res.Cards
	return
}

func (c *Client) SearchBoards(query string, args Arguments) (boards []*Board, err error) {
	args["query"] = query
	args["modelTypes"] = "boards"
	res := SearchResult{}
	err = c.Get("search", args, &res)
	boards = res.Boards
	for _, board := range boards {
		board.client = c
	}
	return
}

func (c *Client) SearchMembers(query string, args Arguments) (members []*Member, err error) {
	args["query"] = query
	err = c.Get("search/members", args, &members)
	return
}
