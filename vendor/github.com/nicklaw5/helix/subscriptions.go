package helix

// Subscription ...
type Subscription struct {
	BroadcasterID   string `json:"broadcaster_id"`
	BroadcasterName string `json:"broadcaster_name"`
	IsGift          bool   `json:"is_gift"`
	Tier            string `json:"tier"`
	PlanName        string `json:"plan_name"`
	UserID          string `json:"user_id"`
	UserName        string `json:"user_name"`
}

// ManySubscriptions ...
type ManySubscriptions struct {
	Subscriptions []Subscription `json:"data"`
	Pagination    Pagination     `json:"pagination"`
}

// SubscriptionsResponse ...
type SubscriptionsResponse struct {
	ResponseCommon
	Data ManySubscriptions
}

// SubscriptionsParams ...
type SubscriptionsParams struct {
	BroadcasterID string   `query:"broadcaster_id"` // Limit 1
	UserID        []string `query:"user_id"`        // Limit 100
}

// GetSubscriptions gets subscriptions about one Twitch broadcaster.
// Broadcasters can only request their own subscriptions.
//
// Required scope: channel:read:subscriptions
func (c *Client) GetSubscriptions(params *SubscriptionsParams) (*SubscriptionsResponse, error) {
	resp, err := c.get("/subscriptions", &ManySubscriptions{}, params)
	if err != nil {
		return nil, err
	}

	subscriptions := &SubscriptionsResponse{}
	subscriptions.StatusCode = resp.StatusCode
	subscriptions.Header = resp.Header
	subscriptions.Error = resp.Error
	subscriptions.ErrorStatus = resp.ErrorStatus
	subscriptions.ErrorMessage = resp.ErrorMessage
	subscriptions.Data.Subscriptions = resp.Data.(*ManySubscriptions).Subscriptions
	subscriptions.Data.Pagination = resp.Data.(*ManySubscriptions).Pagination

	return subscriptions, nil
}
