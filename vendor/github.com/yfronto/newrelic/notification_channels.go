package newrelic

import (
	"strconv"
)

// NotificationChannelLinks describes object links for notification channels.
type NotificationChannelLinks struct {
	NotificationChannels []int `json:"notification_channels,omitempty"`
	User                 int   `json:"user,omitempty"`
}

// NotificationChannel describes a New Relic notification channel.
type NotificationChannel struct {
	ID           int                      `json:"id,omitempty"`
	Type         string                   `json:"type,omitempty"`
	DowntimeOnly bool                     `json:"downtime_only,omitempty"`
	URL          string                   `json:"url,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Description  string                   `json:"description,omitempty"`
	Email        string                   `json:"email,omitempty"`
	Subdomain    string                   `json:"subdomain,omitempty"`
	Service      string                   `json:"service,omitempty"`
	MobileAlerts bool                     `json:"mobile_alerts,omitempty"`
	EmailAlerts  bool                     `json:"email_alerts,omitempty"`
	Room         string                   `json:"room,omitempty"`
	Links        NotificationChannelLinks `json:"links,omitempty"`
}

// NotificationChannelsFilter provides filters for
// NotificationChannelsOptions.
type NotificationChannelsFilter struct {
	Type []string
	IDs  []int
}

// NotificationChannelsOptions is an optional means of filtering when calling
// GetNotificationChannels.
type NotificationChannelsOptions struct {
	Filter NotificationChannelsFilter
	Page   int
}

func (o *NotificationChannelsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[type]": o.Filter.Type,
		"filter[ids]":  o.Filter.IDs,
		"page":         o.Page,
	})
}

// GetNotificationChannel will return the NotificationChannel with  particular ID.
func (c *Client) GetNotificationChannel(id int) (*NotificationChannel, error) {
	resp := &struct {
		NotificationChannel *NotificationChannel `json:"notification_channel,omitempty"`
	}{}
	err := c.doGet("notification_channels/"+strconv.Itoa(id)+".json", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.NotificationChannel, nil
}

// GetNotificationChannels will return a slice of NotificationChannel items,
// optionally filtering by NotificationChannelsOptions.
func (c *Client) GetNotificationChannels(options *NotificationChannelsOptions) ([]NotificationChannel, error) {
	resp := &struct {
		NotificationChannels []NotificationChannel `json:"notification_channels,omitempty"`
	}{}
	err := c.doGet("notification_channels.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.NotificationChannels, nil
}
