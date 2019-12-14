package pagerduty

import (
	"github.com/google/go-querystring/query"
)

// OnCall represents a contiguous unit of time for which a user will be on call for a given escalation policy and escalation rule.
type OnCall struct {
	User             User             `json:"user,omitempty"`
	Schedule         Schedule         `json:"schedule,omitempty"`
	EscalationPolicy EscalationPolicy `json:"escalation_policy,omitempty"`
	EscalationLevel  uint             `json:"escalation_level,omitempty"`
	Start            string           `json:"start,omitempty"`
	End              string           `json:"end,omitempty"`
}

// ListOnCallsResponse is the data structure returned from calling the ListOnCalls API endpoint.
type ListOnCallsResponse struct {
	APIListObject
	OnCalls []OnCall `json:"oncalls"`
}

// ListOnCallOptions is the data structure used when calling the ListOnCalls API endpoint.
type ListOnCallOptions struct {
	APIListObject
	TimeZone            string   `url:"time_zone,omitempty"`
	Includes            []string `url:"include,omitempty,brackets"`
	UserIDs             []string `url:"user_ids,omitempty,brackets"`
	EscalationPolicyIDs []string `url:"escalation_policy_ids,omitempty,brackets"`
	ScheduleIDs         []string `url:"schedule_ids,omitempty,brackets"`
	Since               string   `url:"since,omitempty"`
	Until               string   `url:"until,omitempty"`
	Earliest            bool     `url:"earliest,omitempty"`
}

// ListOnCalls list the on-call entries during a given time range.
func (c *Client) ListOnCalls(o ListOnCallOptions) (*ListOnCallsResponse, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}
	resp, err := c.get("/oncalls?" + v.Encode())
	if err != nil {
		return nil, err
	}
	var result ListOnCallsResponse
	return &result, c.decodeJSON(resp, &result)
}
