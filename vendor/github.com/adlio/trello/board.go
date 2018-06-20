// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

type Board struct {
	client         *Client
	ID             string `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Closed         bool   `json:"closed"`
	IdOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Url            string `json:"url"`
	ShortUrl       string `json:"shortUrl"`
	Prefs          struct {
		PermissionLevel       string            `json:"permissionLevel"`
		Voting                string            `json:"voting"`
		Comments              string            `json:"comments"`
		Invitations           string            `json:"invitations"`
		SelfJoin              bool              `json:"selfjoin"`
		CardCovers            bool              `json:"cardCovers"`
		CardAging             string            `json:"cardAging"`
		CalendarFeedEnabled   bool              `json:"calendarFeedEnabled"`
		Background            string            `json:"background"`
		BackgroundColor       string            `json:"backgroundColor"`
		BackgroundImage       string            `json:"backgroundImage"`
		BackgroundImageScaled []BackgroundImage `json:"backgroundImageScaled"`
		BackgroundTile        bool              `json:"backgroundTile"`
		BackgroundBrightness  string            `json:"backgroundBrightness"`
		CanBePublic           bool              `json:"canBePublic"`
		CanBeOrg              bool              `json:"canBeOrg"`
		CanBePrivate          bool              `json:"canBePrivate"`
		CanInvite             bool              `json:"canInvite"`
	} `json:"prefs"`
	LabelNames struct {
		Black  string `json:"black,omitempty"`
		Blue   string `json:"blue,omitempty"`
		Green  string `json:"green,omitempty"`
		Lime   string `json:"lime,omitempty"`
		Orange string `json:"orange,omitempty"`
		Pink   string `json:"pink,omitempty"`
		Purple string `json:"purple,omitempty"`
		Red    string `json:"red,omitempty"`
		Sky    string `json:"sky,omitempty"`
		Yellow string `json:"yellow,omitempty"`
	} `json:"labelNames"`
	Lists   []*List   `json:"lists"`
	Actions []*Action `json:"actions"`
}

type BackgroundImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

func (b *Board) CreatedAt() time.Time {
	t, _ := IDToTime(b.ID)
	return t
}

/**
 * Board retrieves a Trello board by its ID.
 */
func (c *Client) GetBoard(boardID string, args Arguments) (board *Board, err error) {
	path := fmt.Sprintf("boards/%s", boardID)
	err = c.Get(path, args, &board)
	if board != nil {
		board.client = c
	}
	return
}

func (m *Member) GetBoards(args Arguments) (boards []*Board, err error) {
	path := fmt.Sprintf("members/%s/boards", m.ID)
	err = m.client.Get(path, args, &boards)
	for i := range boards {
		boards[i].client = m.client
	}
	return
}
