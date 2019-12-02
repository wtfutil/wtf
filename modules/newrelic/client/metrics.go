package newrelic

import (
	"time"
)

// Metric describes a New Relic metric.
type Metric struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

// MetricsOptions options allow filtering when getting lists of metric names
// associated with an entity.
type MetricsOptions struct {
	Name string
	Page int
}

// MetricTimeslice describes the period to which a Metric pertains.
type MetricTimeslice struct {
	From   time.Time          `json:"from,omitempty"`
	To     time.Time          `json:"to,omitempty"`
	Values map[string]float64 `json:"values,omitempty"`
}

// MetricData describes the data for a particular metric.
type MetricData struct {
	Name       string            `json:"name,omitempty"`
	Timeslices []MetricTimeslice `json:"timeslices,omitempty"`
}

// MetricDataOptions allow filtering when getting data about a particular set
// of New Relic metrics.
type MetricDataOptions struct {
	Names     Array
	Values    Array
	From      time.Time
	To        time.Time
	Period    int
	Summarize bool
	Raw       bool
}

// MetricDataResponse is the response received from New Relic for any request
// for metric data.
type MetricDataResponse struct {
	From            time.Time    `json:"from,omitempty"`
	To              time.Time    `json:"to,omitempty"`
	MetricsNotFound []string     `json:"metrics_not_found,omitempty"`
	MetricsFound    []string     `json:"metrics_found,omitempty"`
	Metrics         []MetricData `json:"metrics,omitempty"`
}

func (o *MetricsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"name": o.Name,
		"page": o.Page,
	})
}

func (o *MetricDataOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"names[]":   o.Names,
		"values[]":  o.Values,
		"from":      o.From,
		"to":        o.To,
		"period":    o.Period,
		"summarize": o.Summarize,
		"raw":       o.Raw,
	})
}

// MetricClient implements a generic New Relic metrics client.
// This is used as a general client for fetching metric names and data.
type MetricClient struct {
	newRelicClient *Client
}

// NewMetricClient creates and returns a new MetricClient.
func NewMetricClient(newRelicClient *Client) *MetricClient {
	return &MetricClient{
		newRelicClient: newRelicClient,
	}
}

// GetMetrics is a generic function for fetching a list of available metrics
// from different parts of New Relic.
// Example: Application metrics, Component metrics, etc.
func (mc *MetricClient) GetMetrics(path string, options *MetricsOptions) ([]Metric, error) {
	resp := &struct {
		Metrics []Metric `json:"metrics,omitempty"`
	}{}

	err := mc.newRelicClient.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}

	return resp.Metrics, nil
}

// GetMetricData is a generic function for fetching data for a specific metric.
// from different parts of New Relic.
// Example: Application metric data, Component metric data, etc.
func (mc *MetricClient) GetMetricData(path string, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	resp := &struct {
		MetricData MetricDataResponse `json:"metric_data,omitempty"`
	}{}

	if options == nil {
		options = &MetricDataOptions{}
	}

	options.Names = Array{names}
	err := mc.newRelicClient.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}

	return &resp.MetricData, nil
}
