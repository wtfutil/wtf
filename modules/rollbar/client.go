package rollbar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func CurrentActiveItems(accessToken, assignedToName string, activeOnly bool) (*ActiveItems, error) {
	items := &ActiveItems{}

	rollbarAPIURL.Host = "api.rollbar.com"
	rollbarAPIURL.Path = "/api/1/items"
	resp, err := rollbarItemRequest(accessToken, assignedToName, activeOnly)
	if err != nil {
		return items, err
	}

	parseJSON(&items, resp.Body)

	return items, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	rollbarAPIURL = &url.URL{Scheme: "https"}
)

func rollbarItemRequest(accessToken, assignedToName string, activeOnly bool) (*http.Response, error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	params.Add("assigned_user", assignedToName)
	if activeOnly {
		params.Add("status", "active")
	}

	requestURL := rollbarAPIURL.ResolveReference(&url.URL{RawQuery: params.Encode()})
	req, err := http.NewRequest("GET", requestURL.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

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

func parseJSON(obj interface{}, text io.Reader) {
	jsonStream, err := ioutil.ReadAll(text)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonStream))

	for {
		if err := decoder.Decode(obj); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
