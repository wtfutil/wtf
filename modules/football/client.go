package football

import (
	"fmt"
	"net/http"
)

var (
	footballAPIUrl = "https://api.football-data.org/v2"
)

type leagueInfo struct {
	id      int
	caption string
}

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	client := Client{
		apiKey: apiKey,
	}

	return &client
}

func (client *Client) footballRequest(path string, id int) (*http.Response, error) {

	url := fmt.Sprintf("%s/competitions/%d/%s", footballAPIUrl, id, path)
	req, err := http.NewRequest("GET", url, http.NoBody)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", client.apiKey)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	return resp, nil
}
