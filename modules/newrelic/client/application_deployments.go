package newrelic

import (
	"strconv"
	"time"
)

// ApplicationDeploymentLinks represents links that apply to an
// ApplicationDeployment.
type ApplicationDeploymentLinks struct {
	Application int `json:"application,omitempty"`
}

// ApplicationDeploymentOptions provide a means to filter when calling
// GetApplicationDeployments.
type ApplicationDeploymentOptions struct {
	Page int
}

// ApplicationDeployment contains information about a New Relic Application
// Deployment.
type ApplicationDeployment struct {
	ID          int                        `json:"id,omitempty"`
	Revision    string                     `json:"revision,omitempty"`
	Changelog   string                     `json:"changelog,omitempty"`
	Description string                     `json:"description,omitempty"`
	User        string                     `json:"user,omitemtpy"`
	Timestamp   time.Time                  `json:"timestamp,omitempty"`
	Links       ApplicationDeploymentLinks `json:"links,omitempty"`
}

// GetApplicationDeployments returns a slice of New Relic Application
// Deployments.
func (c *Client) GetApplicationDeployments(id int, opt *ApplicationDeploymentOptions) ([]ApplicationDeployment, error) {
	resp := &struct {
		Deployments []ApplicationDeployment `json:"deployments,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(id) + "/deployments.json"
	err := c.doGet(path, opt, resp)
	if err != nil {
		return nil, err
	}
	return resp.Deployments, nil
}

func (o *ApplicationDeploymentOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"page": o.Page,
	})
}
