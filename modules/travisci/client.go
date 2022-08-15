package travisci

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/wtfutil/wtf/utils"
)

var TRAVIS_HOSTS = map[bool]string{
	false: "travis-ci.org",
	true:  "travis-ci.com",
}

func BuildsFor(settings *Settings) (*Builds, error) {
	builds := &Builds{}

	travisAPIURL.Host = "api." + TRAVIS_HOSTS[settings.pro]
	if settings.baseURL != "" {
		travisAPIURL.Host = settings.baseURL
	}

	resp, err := travisBuildRequest(settings)
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
	travisAPIURL = &url.URL{Scheme: "https", Path: "/"}
)

func travisBuildRequest(settings *Settings) (*http.Response, error) {
	var path string = "builds"
	if settings.baseURL != "" {
		travisAPIURL.Path = "/api/"
	}
	params := url.Values{}
	params.Add("limit", settings.limit)
	params.Add("sort_by", settings.sort_by)

	requestUrl := travisAPIURL.ResolveReference(&url.URL{Path: path, RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", requestUrl.String(), http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Travis-API-Version", "3")

	bearer := fmt.Sprintf("token %s", settings.apiKey)
	req.Header.Add("Authorization", bearer)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
