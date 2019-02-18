package rollbar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/wtfutil/wtf/wtf"
)

func CurrentActiveItems() (*ActiveItems, error) {
	items := &ActiveItems{}

	accessToken := wtf.Config.UString("wtf.mods.rollbar.accessToken", "")
	rollbarAPIURL.Host = "api.rollbar.com"
	rollbarAPIURL.Path = "/api/1/items"
	resp, err := rollbarItemRequest(accessToken)
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

func rollbarItemRequest(accessToken string) (*http.Response, error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	userName := wtf.Config.UString("wtf.mods.rollbar.assignedToName", "")
	params.Add("assigned_user", userName)
	active := wtf.Config.UBool("wtf.mods.rollbar.activeOnly", false)
	if active {
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
