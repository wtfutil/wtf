package bamboohr

import (
	"encoding/xml"
	"fmt"
)

// A Client represents the data required to connect to the BambooHR API
type Client struct {
	apiBase   string
	apiKey    string
	subdomain string
}

// NewClient creates and returns a new BambooHR client
func NewClient(url string, apiKey string, subdomain string) *Client {
	client := Client{
		apiBase:   url,
		apiKey:    apiKey,
		subdomain: subdomain,
	}

	return &client
}

/* -------------------- Public Functions -------------------- */

// Away returns a string representation of the people who are out of the office during the defined period
func (client *Client) Away(itemType, startDate, endDate string) []Item {
	calendar, err := client.getWhoIsAway(startDate, endDate)
	if err != nil {
		return []Item{}
	}

	items := calendar.ItemsByType(itemType)

	return items
}

/* -------------------- Private Functions -------------------- */

// getWhoIsAway is the private interface for retrieving structural data about who will be out of the office
// This method does the actual communication with BambooHR and returns the raw Go
// data structures used by the public interface
func (client *Client) getWhoIsAway(startDate, endDate string) (cal Calendar, err error) {
	apiURL := fmt.Sprintf(
		"%s/%s/v1/time_off/whos_out?start=%s&end=%s",
		client.apiBase,
		client.subdomain,
		startDate,
		endDate,
	)

	data, err := Request(client.apiKey, apiURL)
	if err != nil {
		return cal, err
	}
	err = xml.Unmarshal(data, &cal)

	return
}
