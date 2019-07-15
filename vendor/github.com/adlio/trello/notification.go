package trello

import (
	"time"
)

// Notification represents a Trello Notification.
// https://developers.trello.com/reference/#notifications
type Notification struct {
	client *Client

	ID              string           `json:"id"`
	IDAction        string           `json:"idAction"`
	Unread          bool             `json:"unread"`
	Type            string           `json:"type"`
	IDMemberCreator string           `json:"idMemberCreator"`
	Date            time.Time        `json:"date"`
	DateRead        time.Time        `json:"dataRead"`
	Data            NotificationData `json:"data,omitempty"`
	MemberCreator   *Member          `json:"memberCreator,omitempty"`
}

// NotificationData represents the 'notificaiton.data'
type NotificationData struct {
	Text  string                 `json:"text"`
	Card  *NotificationDataCard  `json:"card,omitempty"`
	Board *NotificationDataBoard `json:"board,omitempty"`
}

// NotificationDataBoard represents the 'notification.data.board'
type NotificationDataBoard struct {
	ID        string `json:"id"`
	ShortLink string `json:"shortLink"`
	Name      string `json:"name"`
}

// NotificationDataCard represents the 'notification.data.card'
type NotificationDataCard struct {
	ID        string `json:"id"`
	IDShort   int    `json:"idShort"`
	Name      string `json:"name"`
	ShortLink string `json:"shortLink"`
}

// GetMyNotifications returns the notifications of the authenticated user
func (c *Client) GetMyNotifications(args Arguments) (notifications []*Notification, err error) {
	path := "members/me/notifications"
	err = c.Get(path, args, &notifications)
	for i := range notifications {
		notifications[i].client = c
	}
	return
}
