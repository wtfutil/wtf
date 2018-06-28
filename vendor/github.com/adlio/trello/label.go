// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import "fmt"

type Label struct {
	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}

func (c *Client) GetLabel(labelID string, args Arguments) (label *Label, err error) {
	path := fmt.Sprintf("labels/%s", labelID)
	err = c.Get(path, args, &label)
	return
}

func (b *Board) GetLabels(args Arguments) (labels []*Label, err error) {
	path := fmt.Sprintf("boards/%s/labels", b.ID)
	err = b.client.Get(path, args, &labels)
	return
}