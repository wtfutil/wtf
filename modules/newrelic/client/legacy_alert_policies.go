package newrelic

import (
	"strconv"
)

// LegacyAlertPolicyLinks describes object links for Alert Policies.
type LegacyAlertPolicyLinks struct {
	NotificationChannels []int `json:"notification_channels,omitempty"`
	Servers              []int `json:"servers,omitempty"`
}

// LegacyAlertPolicyCondition describes conditions that trigger an LegacyAlertPolicy.
type LegacyAlertPolicyCondition struct {
	ID             int     `json:"id,omitempty"`
	Enabled        bool    `json:"enabled,omitempty"`
	Severity       string  `json:"severity,omitempty"`
	Threshold      float64 `json:"threshold,omitempty"`
	TriggerMinutes int     `json:"trigger_minutes,omitempty"`
	Type           string  `json:"type,omitempty"`
}

// LegacyAlertPolicy describes a New Relic alert policy.
type LegacyAlertPolicy struct {
	Conditions         []LegacyAlertPolicyCondition `json:"conditions,omitempty"`
	Enabled            bool                         `json:"enabled,omitempty"`
	ID                 int                          `json:"id,omitempty"`
	Links              LegacyAlertPolicyLinks       `json:"links,omitempty"`
	IncidentPreference string                       `json:"incident_preference,omitempty"`
	Name               string                       `json:"name,omitempty"`
}

// LegacyAlertPolicyFilter provides filters for LegacyAlertPolicyOptions.
type LegacyAlertPolicyFilter struct {
	Name string
}

// LegacyAlertPolicyOptions is an optional means of filtering when calling
// GetLegacyAlertPolicies.
type LegacyAlertPolicyOptions struct {
	Filter LegacyAlertPolicyFilter
	Page   int
}

func (o *LegacyAlertPolicyOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[name]": o.Filter.Name,
		"page":         o.Page,
	})
}

// GetLegacyAlertPolicy will return the LegacyAlertPolicy with  particular ID.
func (c *Client) GetLegacyAlertPolicy(id int) (*LegacyAlertPolicy, error) {
	resp := &struct {
		LegacyAlertPolicy *LegacyAlertPolicy `json:"alert_policy,omitempty"`
	}{}
	err := c.doGet("alert_policies/"+strconv.Itoa(id)+".json", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.LegacyAlertPolicy, nil
}

// GetLegacyAlertPolicies will return a slice of LegacyAlertPolicy items,
// optionally filtering by LegacyAlertPolicyOptions.
func (c *Client) GetLegacyAlertPolicies(options *LegacyAlertPolicyOptions) ([]LegacyAlertPolicy, error) {
	resp := &struct {
		LegacyAlertPolicies []LegacyAlertPolicy `json:"alert_policies,omitempty"`
	}{}
	err := c.doGet("alert_policies.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.LegacyAlertPolicies, nil
}
