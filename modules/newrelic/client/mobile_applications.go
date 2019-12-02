package newrelic

import (
	"strconv"
)

// MobileApplicationSummary describes an Application's host.
type MobileApplicationSummary struct {
	ActiveUsers     int     `json:"active_users,omitempty"`
	LaunchCount     int     `json:"launch_count,omitempty"`
	Throughput      float64 `json:"throughput,omitempty"`
	ResponseTime    float64 `json:"response_time,omitempty"`
	CallsPerSession float64 `json:"calls_per_session,omitempty"`
	InteractionTime float64 `json:"interaction_time,omitempty"`
	FailedCallRate  float64 `json:"failed_call_rate,omitempty"`
	RemoteErrorRate float64 `json:"remote_error_rate"`
}

// MobileApplicationCrashSummary describes a MobileApplication's crash data.
type MobileApplicationCrashSummary struct {
	SupportsCrashData    bool    `json:"supports_crash_data,omitempty"`
	UnresolvedCrashCount int     `json:"unresolved_crash_count,omitempty"`
	CrashCount           int     `json:"crash_count,omitempty"`
	CrashRate            float64 `json:"crash_rate,omitempty"`
}

// MobileApplication describes a New Relic Application Host.
type MobileApplication struct {
	ID            int                           `json:"id,omitempty"`
	Name          string                        `json:"name,omitempty"`
	HealthStatus  string                        `json:"health_status,omitempty"`
	Reporting     bool                          `json:"reporting,omitempty"`
	MobileSummary MobileApplicationSummary      `json:"mobile_summary,omitempty"`
	CrashSummary  MobileApplicationCrashSummary `json:"crash_summary,omitempty"`
}

// GetMobileApplications returns a slice of New Relic Mobile Applications.
func (c *Client) GetMobileApplications() ([]MobileApplication, error) {
	resp := &struct {
		Applications []MobileApplication `json:"applications,omitempty"`
	}{}
	path := "mobile_applications.json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.Applications, nil
}

// GetMobileApplication returns a single Mobile Application with the id.
func (c *Client) GetMobileApplication(id int) (*MobileApplication, error) {
	resp := &struct {
		Application MobileApplication `json:"application,omitempty"`
	}{}
	path := "mobile_applications/" + strconv.Itoa(id) + ".json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Application, nil
}
