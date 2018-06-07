package newrelic

// AlertEvent describes a triggered event.
type AlertEvent struct {
	ID            int    `json:"id,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	Product       string `json:"product,omitempty"`
	EntityType    string `json:"entity_type,omitempty"`
	EntityGroupID int    `json:"entity_group_id,omitempty"`
	EntityID      int    `json:"entity_id,omitempty"`
	Priority      string `json:"priority,omitempty"`
	Description   string `json:"description,omitempty"`
	Timestamp     int    `json:"timestamp,omitempty"`
	IncidentID    int    `json:"incident_id"`
}

// AlertEventFilter provides filters for AlertEventOptions when calling
// GetAlertEvents.
type AlertEventFilter struct {
	// TODO: New relic restricts these options
	Product       string
	EntityType    string
	EntityGroupID int
	EntityID      int
	EventType     string
}

// AlertEventOptions is an optional means of filtering AlertEvents when
// calling GetAlertEvents.
type AlertEventOptions struct {
	Filter AlertEventFilter
	Page   int
}

func (o *AlertEventOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[product]":         o.Filter.Product,
		"filter[entity_type]":     o.Filter.EntityType,
		"filter[entity_group_id]": o.Filter.EntityGroupID,
		"filter[entity_id]":       o.Filter.EntityID,
		"filter[event_type]":      o.Filter.EventType,
		"page":                    o.Page,
	})
}

// GetAlertEvents will return a slice of recent AlertEvent items triggered,
// optionally filtering by AlertEventOptions.
func (c *Client) GetAlertEvents(options *AlertEventOptions) ([]AlertEvent, error) {
	resp := &struct {
		RecentEvents []AlertEvent `json:"recent_events,omitempty"`
	}{}
	err := c.doGet("alerts_events.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.RecentEvents, nil
}
