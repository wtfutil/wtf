// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

// Board represents a Trello Board.
// https://developers.trello.com/reference/#boardsid
type Board struct {
	client         *Client
	ID             string `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Closed         bool   `json:"closed"`
	IDOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	URL            string `json:"url"`
	ShortURL       string `json:"shortUrl"`
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
	Lists        []*List      `json:"lists"`
	Actions      []*Action    `json:"actions"`
	Organization Organization `json:"organization"`
}

// NewBoard is a constructor that sets the default values
// for Prefs.SelfJoin and Prefs.CardCovers also set by the API.
func NewBoard(name string) Board {
	b := Board{Name: name}

	// default values in line with API POST
	b.Prefs.SelfJoin = true
	b.Prefs.CardCovers = true

	return b
}

// BackgroundImage is a nested resource of Board.
type BackgroundImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// CreatedAt returns a board's created-at attribute as time.Time.
func (b *Board) CreatedAt() time.Time {
	t, _ := IDToTime(b.ID)
	return t
}

// CreateBoard creates a board remote.
// Attribute currently supported as exra argument: powerUps.
// Attributes currently known to be unsupported: idBoardSource, keepFromSource.
//
// API Docs: https://developers.trello.com/reference/#boardsid
func (c *Client) CreateBoard(board *Board, extraArgs Arguments) error {
	path := "boards"
	args := Arguments{
		"desc":             board.Desc,
		"name":             board.Name,
		"prefs_selfJoin":   fmt.Sprintf("%t", board.Prefs.SelfJoin),
		"prefs_cardCovers": fmt.Sprintf("%t", board.Prefs.CardCovers),
		"idOrganization":   board.IDOrganization,
	}

	if board.Prefs.Voting != "" {
		args["prefs_voting"] = board.Prefs.Voting
	}
	if board.Prefs.PermissionLevel != "" {
		args["prefs_permissionLevel"] = board.Prefs.PermissionLevel
	}
	if board.Prefs.Comments != "" {
		args["prefs_comments"] = board.Prefs.Comments
	}
	if board.Prefs.Invitations != "" {
		args["prefs_invitations"] = board.Prefs.Invitations
	}
	if board.Prefs.Background != "" {
		args["prefs_background"] = board.Prefs.Background
	}
	if board.Prefs.CardAging != "" {
		args["prefs_cardAging"] = board.Prefs.CardAging
	}

	// Expects one of "all", "calendar", "cardAging", "recap", or "voting".
	if powerUps, ok := extraArgs["powerUps"]; ok {
		args["powerUps"] = powerUps
	}

	err := c.Post(path, args, &board)
	if err == nil {
		board.client = c
	}
	return err
}

// Delete makes a DELETE call for the receiver Board.
func (b *Board) Delete(extraArgs Arguments) error {
	path := fmt.Sprintf("boards/%s", b.ID)
	return b.client.Delete(path, Arguments{}, b)
}

// GetBoard retrieves a Trello board by its ID.
func (c *Client) GetBoard(boardID string, args Arguments) (board *Board, err error) {
	path := fmt.Sprintf("boards/%s", boardID)
	err = c.Get(path, args, &board)
	if board != nil {
		board.client = c
	}
	return
}

// GetMyBoards returns a slice of all boards associated with the credentials set on the client.
func (c *Client) GetMyBoards(args Arguments) (boards []*Board, err error) {
	path := "members/me/boards"
	err = c.Get(path, args, &boards)
	for i := range boards {
		boards[i].client = c
	}
	return
}

// GetBoards returns a slice of all public boards of the receiver Member.
func (m *Member) GetBoards(args Arguments) (boards []*Board, err error) {
	path := fmt.Sprintf("members/%s/boards", m.ID)
	err = m.client.Get(path, args, &boards)
	for i := range boards {
		boards[i].client = m.client
	}
	return
}
