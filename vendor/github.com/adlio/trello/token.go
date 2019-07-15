// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

// Token represents Trello tokens. Tokens can be used for setting up Webhooks among other things.
// https://developers.trello.com/reference/#tokens
type Token struct {
	client      *Client
	ID          string       `json:"id"`
	DateCreated time.Time    `json:"dateCreated"`
	DateExpires *time.Time   `json:"dateExpires"`
	IDMember    string       `json:"idMember"`
	Identifier  string       `json:"identifier"`
	Permissions []Permission `json:"permissions"`
}

// Permission represent a Token's permissions.
type Permission struct {
	IDModel   string `json:"idModel"`
	ModelType string `json:"modelType"`
	Read      bool   `json:"read"`
	Write     bool   `json:"write"`
}

// GetToken takes a token id and Arguments and GETs and returns the Token or an error.
func (c *Client) GetToken(tokenID string, args Arguments) (token *Token, err error) {
	path := fmt.Sprintf("tokens/%s", tokenID)
	err = c.Get(path, args, &token)
	if token != nil {
		token.client = c
	}
	return
}
