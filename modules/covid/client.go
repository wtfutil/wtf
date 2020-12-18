package covid

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/wtfutil/wtf/utils"
)

// LatestCases queries the /latest endpoint
func LatestCases() (*Latest, error) {
	// TODO: figure out the pointer logic here
	latest := &Latest{}

	resp, err := covidAPIRequest("latest")
	if err != nil {
		// TODO: figure out the pointer logic here
		return latest, err
	}

	err = utils.ParseJSON(&latest, resp.Body)
	if err != nil {
		// TODO: figure out the pointer logic here
		return latest, err
	}
	// TODO: figure out the pointer logic here
	return latest, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	covidTrackerAPIURL = &url.URL{Scheme: "https", Host: "coronavirus-tracker-api.herokuapp.com", Path: "/v2/"}
)

func covidAPIRequest(path string) (*http.Response, error) {
	uri := covidTrackerAPIURL.ResolveReference(&url.URL{Path: path})

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
