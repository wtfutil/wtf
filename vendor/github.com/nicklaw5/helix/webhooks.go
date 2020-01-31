package helix

import (
	"net/http"
	"regexp"
)

// WebhookSubscription ...
type WebhookSubscription struct {
	Topic     string `json:"topic"`
	Callback  string `json:"callback"`
	ExpiresAt Time   `json:"expires_at"`
}

// ManyWebhookSubscriptions ...
type ManyWebhookSubscriptions struct {
	Total                int                   `json:"total"`
	WebhookSubscriptions []WebhookSubscription `json:"data"`
	Pagination           Pagination            `json:"pagination"`
}

// WebhookSubscriptionsResponse ...
type WebhookSubscriptionsResponse struct {
	ResponseCommon
	Data ManyWebhookSubscriptions
}

// WebhookSubscriptionsParams ...
type WebhookSubscriptionsParams struct {
	After string `query:"after"`
	First int    `query:"first,20"` // Limit 100
}

// GetWebhookSubscriptions gets webhook subscriptions, in order of expiration.
// Requires an app access token.
func (c *Client) GetWebhookSubscriptions(params *WebhookSubscriptionsParams) (*WebhookSubscriptionsResponse, error) {
	resp, err := c.get("/webhooks/subscriptions", &ManyWebhookSubscriptions{}, params)
	if err != nil {
		return nil, err
	}

	webhooks := &WebhookSubscriptionsResponse{}
	webhooks.StatusCode = resp.StatusCode
	webhooks.Header = resp.Header
	webhooks.Error = resp.Error
	webhooks.ErrorStatus = resp.ErrorStatus
	webhooks.ErrorMessage = resp.ErrorMessage
	webhooks.Data.Total = resp.Data.(*ManyWebhookSubscriptions).Total
	webhooks.Data.WebhookSubscriptions = resp.Data.(*ManyWebhookSubscriptions).WebhookSubscriptions
	webhooks.Data.Pagination = resp.Data.(*ManyWebhookSubscriptions).Pagination

	return webhooks, nil
}

// WebhookSubscriptionResponse ...
type WebhookSubscriptionResponse struct {
	ResponseCommon
}

// WebhookSubscriptionPayload ...
type WebhookSubscriptionPayload struct {
	Mode         string `json:"hub.mode"`
	Topic        string `json:"hub.topic"`
	Callback     string `json:"hub.callback"`
	LeaseSeconds int    `json:"hub.lease_seconds,omitempty"`
	Secret       string `json:"hub.secret,omitempty"`
}

// PostWebhookSubscription ...
func (c *Client) PostWebhookSubscription(payload *WebhookSubscriptionPayload) (*WebhookSubscriptionResponse, error) {
	resp, err := c.post("/webhooks/hub", nil, payload)
	if err != nil {
		return nil, err
	}

	webhook := &WebhookSubscriptionResponse{}
	webhook.StatusCode = resp.StatusCode
	webhook.Header = resp.Header
	webhook.Error = resp.Error
	webhook.ErrorStatus = resp.ErrorStatus
	webhook.ErrorMessage = resp.ErrorMessage

	return webhook, nil
}

// Regular expressions used for parsing webhook link headers
var (
	UserFollowsRegexp        = regexp.MustCompile("helix/users/follows\\?first=1(&from_id=(?P<from_id>\\d+))?(&to_id=(?P<to_id>\\d+))?>")
	StreamChangedRegexp      = regexp.MustCompile("helix/streams\\?user_id=(?P<user_id>\\d+)>")
	UserChangedRegexp        = regexp.MustCompile("helix/users\\?id=(?P<id>\\d+)>")
	GameAnalyticsRegexp      = regexp.MustCompile("helix/analytics\\?game_id=(?P<game_id>\\w+)>")
	ExtensionAnalyticsRegexp = regexp.MustCompile("helix/analytics\\?extension_id=(?P<extension_id>\\w+)>")
)

// WebhookTopic is a topic that relates to a specific webhook event.
type WebhookTopic int

// Enumerated webhook topics
const (
	UserFollowsTopic WebhookTopic = iota
	StreamChangedTopic
	UserChangedTopic
	GameAnalyticsTopic
	ExtensionAnalyticsTopic
)

// GetWebhookTopicFromRequest inspects the "Link" request header to
// determine if it matches against any recognised webhooks topics.
// The matched topic is returned. Otherwise -1 is returned.
func GetWebhookTopicFromRequest(req *http.Request) WebhookTopic {
	header := getLinkHeaderFromWebhookRequest(req)

	if UserFollowsRegexp.MatchString(header) {
		return UserFollowsTopic
	}
	if StreamChangedRegexp.MatchString(header) {
		return StreamChangedTopic
	}
	if UserChangedRegexp.MatchString(header) {
		return UserChangedTopic
	}
	if GameAnalyticsRegexp.MatchString(header) {
		return GameAnalyticsTopic
	}
	if ExtensionAnalyticsRegexp.MatchString(header) {
		return ExtensionAnalyticsTopic
	}

	return -1
}

// GetWebhookTopicValuesFromRequest inspects the "Link" request header to
// determine if it matches against any recognised webhooks topics and
// returns the unique values specified in the header.
//
// For example, say we receive a "User Follows" webhook event from Twitch.
// Its "Link" header value look likes the following:
//
// 		<https://api.twitch.tv/helix/webhooks/hub>; rel="hub", <https://api.twitch.tv/helix/users/follows?first=1&from_id=111116&to_id=22222>; rel="self"
//
// From which GetWebhookTopicValuesFromRequest will return a map with the
// values of from_id and to_id:
//
// 		map[from_id:111116 to_id:22222]
//
// This is particularly useful for webhooks events that do not have a distinguishable
// JSON payload, such as the "Stream Changed" down event.
//
// Additionally, if topic is not known you can pass -1 as its value and
func GetWebhookTopicValuesFromRequest(req *http.Request, topic WebhookTopic) map[string]string {
	values := make(map[string]string)
	webhookTopic := topic
	header := getLinkHeaderFromWebhookRequest(req)

	// Webhook topic may not be known, so let's attempt to
	// determine its value based on the request.
	if webhookTopic < 0 && header != "" {
		webhookTopic = GetWebhookTopicFromRequest(req)
	}

	switch webhookTopic {
	case UserFollowsTopic:
		values = findStringSubmatchMap(UserFollowsRegexp, header)
	case StreamChangedTopic:
		values = findStringSubmatchMap(StreamChangedRegexp, header)
	case UserChangedTopic:
		values = findStringSubmatchMap(UserChangedRegexp, header)
	case GameAnalyticsTopic:
		values = findStringSubmatchMap(GameAnalyticsRegexp, header)
	case ExtensionAnalyticsTopic:
		values = findStringSubmatchMap(ExtensionAnalyticsRegexp, header)
	}

	return values
}

func getLinkHeaderFromWebhookRequest(req *http.Request) string {
	return req.Header.Get("link")
}

func findStringSubmatchMap(r *regexp.Regexp, s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}
		captures[name] = match[i]

	}
	return captures
}
