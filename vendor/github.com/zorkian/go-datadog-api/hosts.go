package datadog

type HostActionResp struct {
	Action   string `json:"action"`
	Hostname string `json:"hostname"`
	Message  string `json:"message,omitempty"`
}

type HostActionMute struct {
	Message  *string `json:"message,omitempty"`
	EndTime  *string `json:"end,omitempty"`
	Override *bool   `json:"override,omitempty"`
}

// MuteHost mutes all monitors for the given host
func (client *Client) MuteHost(host string, action *HostActionMute) (*HostActionResp, error) {
	var out HostActionResp
	uri := "/v1/host/" + host + "/mute"
	if err := client.doJsonRequest("POST", uri, action, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UnmuteHost unmutes all monitors for the given host
func (client *Client) UnmuteHost(host string) (*HostActionResp, error) {
	var out HostActionResp
	uri := "/v1/host/" + host + "/unmute"
	if err := client.doJsonRequest("POST", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// HostTotalsResp defines response to GET /v1/hosts/totals.
type HostTotalsResp struct {
	TotalUp     *int `json:"total_up"`
	TotalActive *int `json:"total_active"`
}

// GetHostTotals returns number of total active hosts and total up hosts.
// Active means the host has reported in the past hour, and up means it has reported in the past two hours.
func (client *Client) GetHostTotals() (*HostTotalsResp, error) {
	var out HostTotalsResp
	uri := "/v1/hosts/totals"
	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
