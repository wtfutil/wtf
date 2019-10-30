package circleci

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/wtfutil/wtf/utils"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	client := Client{
		apiKey: apiKey,
	}

	return &client
}

func (client *Client) BuildsFor() ([]*Build, error) {
	builds := []*Build{}

	resp, err := client.circleRequest("recent-builds")
	if err != nil {
		return builds, err
	}

	err = utils.ParseJSON(&builds, resp.Body)
	if err != nil {
		return builds, err
	}

	return builds, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	circleAPIURL = &url.URL{Scheme: "https", Host: "circleci.com", Path: "/api/v1/"}
)

func (client *Client) circleRequest(path string) (*http.Response, error) {
	params := url.Values{}
	params.Add("circle-token", client.apiKey)

	url := circleAPIURL.ResolveReference(&url.URL{Path: path, RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
