package newrelic

import (
	"time"
)

// Usage describes usage over a single time period.
type Usage struct {
	From  time.Time `json"from,omitempty"`
	To    time.Time `json:"to,omitempty"`
	Usage int       `json:"usage,omitempty"`
}

// UsageData represents usage data for a product over a time frame, including
// a slice of Usages.
type UsageData struct {
	Product string    `json:"product,omitempty"`
	From    time.Time `json:"from,omitempty"`
	To      time.Time `json:"to,omitempty"`
	Unit    string    `json:"unit,omitempty"`
	Usages  []Usage   `json:"usages,omitempty"`
}

type usageParams struct {
	Start             time.Time
	End               time.Time
	IncludeSubaccount bool
}

func (o *usageParams) String() string {
	return encodeGetParams(map[string]interface{}{
		"start_date":          o.Start.Format("2006-01-02"),
		"end_date":            o.End.Format("2006-01-02"),
		"include_subaccounts": o.IncludeSubaccount,
	})
}

// GetUsages will return usage for a product in a given time frame.
func (c *Client) GetUsages(product string, start, end time.Time, includeSubaccounts bool) (*UsageData, error) {
	resp := &struct {
		UsageData *UsageData `json:"usage_data,omitempty"`
	}{}
	options := &usageParams{start, end, includeSubaccounts}
	err := c.doGet("usages/"+product+".json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.UsageData, nil
}
