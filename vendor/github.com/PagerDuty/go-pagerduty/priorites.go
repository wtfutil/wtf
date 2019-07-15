package pagerduty

import (
	"encoding/json"
)

// PriorityProperty is a single priorty object returned from the Priorities endpoint
type PriorityProperty struct {
	APIObject
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Priorities struct {
	APIListObject
	Priorities []PriorityProperty `json:"priorities"`
}

// ListPriorities lists existing priorities
func (c *Client) ListPriorities() (*Priorities, error) {
	resp, e := c.get("/priorities")
	if e != nil {
		return nil, e
	}

	var p Priorities
	e = json.NewDecoder(resp.Body).Decode(&p)
	if e != nil {
		return nil, e
	}

	return &p, nil
}
