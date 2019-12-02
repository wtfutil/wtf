package newrelic

import (
	"strconv"
)

// ApplicationHostSummary describes an Application's host.
type ApplicationHostSummary struct {
	ApdexScore    float64 `json:"apdex_score,omitempty"`
	ErrorRate     float64 `json:"error_rate,omitempty"`
	InstanceCount int     `json:"instance_count,omitempty"`
	ResponseTime  float64 `json:"response_time,omitempty"`
	Throughput    float64 `json:"throughput,omitempty"`
}

// ApplicationHostEndUserSummary describes the end user summary component of
// an ApplicationHost.
type ApplicationHostEndUserSummary struct {
	ResponseTime float64 `json:"response_time,omitempty"`
	Throughput   float64 `json:"throughput,omitempty"`
	ApdexScore   float64 `json:"apdex_score,omitempty"`
}

// ApplicationHostLinks list IDs associated with an ApplicationHost.
type ApplicationHostLinks struct {
	Application          int   `json:"application,omitempty"`
	ApplicationInstances []int `json:"application_instances,omitempty"`
	Server               int   `json:"server,omitempty"`
}

// ApplicationHost describes a New Relic Application Host.
type ApplicationHost struct {
	ApplicationName    string                        `json:"application_name,omitempty"`
	ApplicationSummary ApplicationHostSummary        `json:"application_summary,omitempty"`
	HealthStatus       string                        `json:"health_status,omitempty"`
	Host               string                        `json:"host,omitempty"`
	ID                 int                           `json:"idomitempty"`
	Language           string                        `json:"language,omitempty"`
	Links              ApplicationHostLinks          `json:"links,omitempty"`
	EndUserSummary     ApplicationHostEndUserSummary `json:"end_user_summary,omitempty"`
}

// ApplicationHostsFilter provides a means to filter requests through
// ApplicationHostsOptions when calling GetApplicationHosts.
type ApplicationHostsFilter struct {
	Hostname string
	IDs      []int
}

// ApplicationHostsOptions provide a means to filter results when calling
// GetApplicationHosts.
type ApplicationHostsOptions struct {
	Filter ApplicationHostsFilter
	Page   int
}

// GetApplicationHosts returns a slice of New Relic Application Hosts,
// optionally filtering by ApplicationHostOptions.
func (c *Client) GetApplicationHosts(id int, options *ApplicationHostsOptions) ([]ApplicationHost, error) {
	resp := &struct {
		ApplicationHosts []ApplicationHost `json:"application_hosts,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(id) + "/hosts.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.ApplicationHosts, nil
}

// GetApplicationHost returns a single Application Host associated with the
// given application host ID and host ID.
func (c *Client) GetApplicationHost(appID, hostID int) (*ApplicationHost, error) {
	resp := &struct {
		ApplicationHost ApplicationHost `json:"application_host,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(appID) + "/hosts/" + strconv.Itoa(hostID) + ".json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return &resp.ApplicationHost, nil
}

func (o *ApplicationHostsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[hostname]": o.Filter.Hostname,
		"filter[ids]":      o.Filter.IDs,
		"page":             o.Page,
	})
}
