package newrelic

import (
	"strconv"
	"time"
)

// ApplicationSummary describes the brief summary component of an Application.
type ApplicationSummary struct {
	ResponseTime            float64 `json:"response_time,omitempty"`
	Throughput              float64 `json:"throughput,omitempty"`
	ErrorRate               float64 `json:"error_rate,omitempty"`
	ApdexTarget             float64 `json:"apdex_target,omitempty"`
	ApdexScore              float64 `json:"apdex_score,omitempty"`
	HostCount               int     `json:"host_count,omitempty"`
	InstanceCount           int     `json:"instance_count,omitempty"`
	ConcurrentInstanceCount int     `json:"concurrent_instance_count,omitempty"`
}

// EndUserSummary describes the end user summary component of an Application.
type EndUserSummary struct {
	ResponseTime float64 `json:"response_time,omitempty"`
	Throughput   float64 `json:"throughput,omitempty"`
	ApdexTarget  float64 `json:"apdex_target,omitempty"`
	ApdexScore   float64 `json:"apdex_score,omitempty"`
}

// Settings describe settings for an Application.
type Settings struct {
	AppApdexThreshold        float64 `json:"app_apdex_threshold,omitempty"`
	EndUserApdexThreshold    float64 `json:"end_user_apdex_threshold,omitempty"`
	EnableRealUserMonitoring bool    `json:"enable_real_user_monitoring,omitempty"`
	UseServerSideConfig      bool    `json:"use_server_side_config,omitempty"`
}

// Links list IDs associated with an Application.
type Links struct {
	Servers              []int `json:"servers,omitempty"`
	ApplicationHosts     []int `json:"application_hosts,omitempty"`
	ApplicationInstances []int `json:"application_instances,omitempty"`
	AlertPolicy          int   `json:"alert_policy,omitempty"`
}

// Application describes a New Relic Application.
type Application struct {
	ID                 int                `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Language           string             `json:"language,omitempty"`
	HealthStatus       string             `json:"health_status,omitempty"`
	Reporting          bool               `json:"reporting,omitempty"`
	LastReportedAt     time.Time          `json:"last_reported_at,omitempty"`
	ApplicationSummary ApplicationSummary `json:"application_summary,omitempty"`
	EndUserSummary     EndUserSummary     `json:"end_user_summary,omitempty"`
	Settings           Settings           `json:"settings,omitempty"`
	Links              Links              `json:"links,omitempty"`
}

// ApplicationFilter provides a means to filter requests through
// ApplicaitonOptions when calling GetApplications.
type ApplicationFilter struct {
	Name     string
	Host     string
	IDs      []int
	Language string
}

// ApplicationOptions provides a means to filter results when calling
// GetApplicaitons.
type ApplicationOptions struct {
	Filter ApplicationFilter
	Page   int
}

func (o *ApplicationOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[name]":     o.Filter.Name,
		"filter[host]":     o.Filter.Host,
		"filter[ids]":      o.Filter.IDs,
		"filter[language]": o.Filter.Language,
		"page":             o.Page,
	})
}

// GetApplications returns a slice of New Relic Applications, optionally
// filtering by ApplicationOptions.
func (c *Client) GetApplications(options *ApplicationOptions) ([]Application, error) {
	resp := &struct {
		Applications []Application `json:"applications,omitempty"`
	}{}
	err := c.doGet("applications.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Applications, nil
}

// GetApplication returns a single Application associated with a given ID.
func (c *Client) GetApplication(id int) (*Application, error) {
	resp := &struct {
		Application Application `json:"application,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(id) + ".json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Application, nil
}
