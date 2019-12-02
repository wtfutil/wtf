package newrelic

import (
	"fmt"
)

// GetApplicationMetrics will return a slice of Metric items for a
// particular Application ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetApplicationMetrics(id int, options *MetricsOptions) ([]Metric, error) {
	mc := NewMetricClient(c)

	return mc.GetMetrics(
		fmt.Sprintf(
			"applications/%d/metrics.json",
			id,
		),
		options,
	)
}

// GetApplicationMetricData will return all metric data for a particular
// application and slice of metric names, optionally filtered by
// MetricDataOptions.
func (c *Client) GetApplicationMetricData(id int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	mc := NewMetricClient(c)

	return mc.GetMetricData(
		fmt.Sprintf(
			"applications/%d/metrics/data.json",
			id,
		),
		names,
		options,
	)
}
