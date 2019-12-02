package newrelic

import (
	"strconv"
)

// ApplicationInstanceSummary describes an Application's instance.
type ApplicationInstanceSummary struct {
	ResponseTime  float64 `json:"response_time,omitempty"`
	Throughput    float64 `json:"throughput,omitempty"`
	ErrorRate     float64 `json:"error_rate,omitempty"`
	ApdexScore    float64 `json:"apdex_score,omitempty"`
	InstanceCount int     `json:"instance_count,omitempty"`
}

// ApplicationInstanceEndUserSummary describes the end user summary component
// of an ApplicationInstance.
type ApplicationInstanceEndUserSummary struct {
	ResponseTime float64 `json:"response_time,omitempty"`
	Throughput   float64 `json:"throughput,omitempty"`
	ApdexScore   float64 `json:"apdex_score,omitempty"`
}

// ApplicationInstanceLinks lists IDs associated with an ApplicationInstances.
type ApplicationInstanceLinks struct {
	Application     int `json:"application,omitempty"`
	ApplicationHost int `json:"application_host,omitempty"`
	Server          int `json:"server,omitempty"`
}

// ApplicationInstance describes a New Relic Application instance.
type ApplicationInstance struct {
	ID                 int                               `json:"id,omitempty"`
	ApplicationName    string                            `json:"application_name,omitempty"`
	Host               string                            `json:"host,omitempty"`
	Port               int                               `json:"port,omitempty"`
	Language           string                            `json:"language,omitempty"`
	HealthStatus       string                            `json:"health_status,omitempty"`
	ApplicationSummary ApplicationInstanceSummary        `json:"application_summary,omitempty"`
	EndUserSummary     ApplicationInstanceEndUserSummary `json:"end_user_summary,omitempty"`
	Links              ApplicationInstanceLinks          `json:"links,omitempty"`
}

// ApplicationInstancesFilter provides a means to filter requests through
// ApplicationInstancesOptions when calling GetApplicationInstances.
type ApplicationInstancesFilter struct {
	Hostname string
	IDs      []int
}

// ApplicationInstancesOptions provides a means to filter results when calling
// GetApplicationInstances.
type ApplicationInstancesOptions struct {
	Filter ApplicationInstancesFilter
	Page   int
}

// GetApplicationInstances returns a slice of New Relic Application Instances,
// optionall filtering by ApplicationInstancesOptions.
func (c *Client) GetApplicationInstances(appID int, options *ApplicationInstancesOptions) ([]ApplicationInstance, error) {
	resp := &struct {
		ApplicationInstances []ApplicationInstance `json:"application_instances,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(appID) + "/instances.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.ApplicationInstances, nil
}

// GetApplicationInstance returns a single Application Instance associated
// with the given application ID and instance ID
func (c *Client) GetApplicationInstance(appID, instanceID int) (*ApplicationInstance, error) {
	resp := &struct {
		ApplicationInstance ApplicationInstance `json:"application_instance,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(appID) + "/instances/" + strconv.Itoa(instanceID) + ".json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return &resp.ApplicationInstance, nil
}
func (o *ApplicationInstancesOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[hostname]": o.Filter.Hostname,
		"filter[ids]":      o.Filter.IDs,
		"page":             o.Page,
	})
}
