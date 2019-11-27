package newrelic

// BrowserApplicationsFilter is the filtering component of
// BrowserApplicationsOptions
type BrowserApplicationsFilter struct {
	Name string
	IDs  []int
}

// BrowserApplicationsOptions provides a filtering mechanism for
// GetBrowserApplications.
type BrowserApplicationsOptions struct {
	Filter BrowserApplicationsFilter
	Page   int
}

// BrowserApplication describes a New Relic Browser Application.
type BrowserApplication struct {
	ID                   int    `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	BrowserMonitoringKey string `json:"browser_monitoring_key,omitempty"`
	LoaderScript         string `json:"loader_script,omitempty"`
}

// GetBrowserApplications will return a slice of New Relic Browser
// Applications, optionally filtered by BrowserApplicationsOptions.
func (c *Client) GetBrowserApplications(opt *BrowserApplicationsOptions) ([]BrowserApplication, error) {
	resp := &struct {
		BrowserApplications []BrowserApplication `json:"browser_applications,omitempty"`
	}{}
	path := "browser_applications.json"
	err := c.doGet(path, opt, resp)
	if err != nil {
		return nil, err
	}
	return resp.BrowserApplications, nil
}

func (o *BrowserApplicationsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[name]": o.Filter.Name,
		"filter[ids]":  o.Filter.IDs,
		"page":         o.Page,
	})
}
