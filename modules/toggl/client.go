package toggl

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

func (client *Client) me() (myself togglPerson, err error) {
	me := myself

	resp, err := client.togglRequest("me")
	if err != nil {
		return me, err
	}

	err = utils.ParseJSON(&me, resp.Body)
	if err != nil {
		return me, err
	}

	return me, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	togglAPIURL = &url.URL{Scheme: "https", Host: "www.toggl.com", Path: "/api/v8/"}
)

func (client *Client) togglRequest(path string) (*http.Response, error) {
	params := url.Values{}
	params.Set("with_related_data", "true")

	url := togglAPIURL.ResolveReference(&url.URL{Path: path, RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(client.apiKey, "api_token")

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
