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
