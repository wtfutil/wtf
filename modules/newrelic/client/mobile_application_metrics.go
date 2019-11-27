package newrelic

import (
	"fmt"
)

// GetMobileApplicationMetrics will return a slice of Metric items for a
// particular MobileAplication ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetMobileApplicationMetrics(id int, options *MetricsOptions) ([]Metric, error) {
	mc := NewMetricClient(c)

	return mc.GetMetrics(
		fmt.Sprintf(
			"mobile_applications/%d/metrics.json",
			id,
		),
		options,
	)
}

// GetMobileApplicationMetricData will return all metric data for a particular
// MobileAplication and slice of metric names, optionally filtered by
// MetricDataOptions.
func (c *Client) GetMobileApplicationMetricData(id int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	mc := NewMetricClient(c)

	return mc.GetMetricData(
		fmt.Sprintf(
			"mobile_applications/%d/metrics/data.json",
			id,
		),
		names,
		options,
	)
}
