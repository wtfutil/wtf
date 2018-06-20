// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

type Checklist struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	IDBoard    string      `json:"idBoard,omitempty"`
	IDCard     string      `json:"idCard,omitempty"`
	Pos        float64     `json:"pos,omitempty"`
	CheckItems []CheckItem `json:"checkItems,omitempty"`
}

type CheckItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"state"`
	IDChecklist string  `json:"idChecklist,omitempty"`
	Pos         float64 `json:"pos,omitempty"`
}

// Manifestation of CheckItem when it appears in CheckItemStates
// on a Card.
type CheckItemState struct {
	IDCheckItem string `json:"idCheckItem"`
	State       string `json:"state"`
}
