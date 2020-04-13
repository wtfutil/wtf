package pihole

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Status struct {
	DomainsBeingBlocked string `json:"domains_being_blocked"`
	DNSQueriesToday     string `json:"dns_queries_today"`
	AdsBlockedToday     string `json:"ads_blocked_today"`
	AdsPercentageToday  string `json:"ads_percentage_today"`
	UniqueDomains       string `json:"unique_domains"`
	QueriesForwarded    string `json:"queries_forwarded"`
	QueriesCached       string `json:"queries_cached"`
	Status              string `json:"status"`
	GravityLastUpdated  struct {
		Relative struct {
			Days    string `json:"days"`
			Hours   string `json:"hours"`
			Minutes string `json:"minutes"`
		}
	} `json:"gravity_last_updated"`
}

func getStatus(c http.Client, apiURL string) (status Status, err error) {
	var req *http.Request

	var url *url2.URL

	if url, err = url2.Parse(apiURL); err != nil {
		return status, fmt.Errorf(" failed to parse API URL\n %s\n", parseError(err))
	}

	var query url2.Values

	if query, err = url2.ParseQuery(url.RawQuery); err != nil {
		return status, fmt.Errorf(" failed to parse query\n %s\n", parseError(err))
	}

	query.Add("summary", "")

	url.RawQuery = query.Encode()
	if req, err = http.NewRequest("GET", url.String(), nil); err != nil {
		return status, fmt.Errorf(" failed to create request\n %s\n", parseError(err))
	}

	var resp *http.Response

	if resp, err = c.Do(req); err != nil || resp == nil {
		return status, fmt.Errorf(" failed to connect to Pi-hole server\n %s\n", parseError(err))
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return status, fmt.Errorf(" failed to retrieve version from Pi-hole server\n http status code: %d",
			resp.StatusCode)
	}

	var rBody []byte

	if rBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return status, fmt.Errorf(" failed to read status response\n")
	}

	if err = json.Unmarshal(rBody, &status); err != nil {
		return status, fmt.Errorf(" failed to retrieve top items: check provided api URL and token\n %s\n\n",
			parseError(err))
	}

	return status, err
}

type TopItems struct {
	TopQueries map[string]int `json:"top_queries"`
	TopAds     map[string]int `json:"top_ads"`
}

func getTopItems(c http.Client, settings *Settings) (ti TopItems, err error) {
	var req *http.Request

	var url *url2.URL

	if url, err = url2.Parse(settings.apiUrl); err != nil {
		return ti, fmt.Errorf(" failed to parse API URL\n %s\n", parseError(err))
	}

	var query url2.Values

	if query, err = url2.ParseQuery(url.RawQuery); err != nil {
		return ti, fmt.Errorf(" failed to parse query\n %s\n", parseError(err))
	}

	query.Add("auth", settings.token)
	query.Add("topItems", strconv.Itoa(settings.showTopItems))

	url.RawQuery = query.Encode()

	req, err = http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return ti, fmt.Errorf(" failed to create request\n %s\n", parseError(err))
	}

	var resp *http.Response

	if resp, err = c.Do(req); err != nil || resp == nil {
		return ti, fmt.Errorf(" failed to connect to Pi-hole server\n %s\n", parseError(err))
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return ti, fmt.Errorf(" failed to retrieve version from Pi-hole server\n http status code: %d",
			resp.StatusCode)
	}

	var rBody []byte

	rBody, err = ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(rBody, &ti); err != nil {
		return ti, fmt.Errorf(" failed to retrieve top items: check provided api URL and token\n %s\n\n",
			parseError(err))
	}

	return ti, err
}

type TopClients struct {
	TopSources map[string]int `json:"top_sources"`
}

// parseError removes any token from output and ensures a non-nil response
func parseError(err error) string {
	if err == nil {
		return "unknown error"
	}

	var re = regexp.MustCompile(`auth=[a-zA-Z0-9]*`)

	return re.ReplaceAllString(err.Error(), "auth=<token>")
}

func getTopClients(c http.Client, settings *Settings) (tc TopClients, err error) {
	var req *http.Request

	var url *url2.URL

	if url, err = url2.Parse(settings.apiUrl); err != nil {
		return tc, fmt.Errorf(" failed to parse API URL\n %s\n", parseError(err))
	}

	var query url2.Values

	if query, err = url2.ParseQuery(url.RawQuery); err != nil {
		return tc, fmt.Errorf(" failed to parse query\n %s\n", parseError(err))
	}

	query.Add("topClients", strconv.Itoa(settings.showTopClients))
	query.Add("auth", settings.token)
	url.RawQuery = query.Encode()

	if req, err = http.NewRequest("GET", url.String(), nil); err != nil {
		return tc, fmt.Errorf(" failed to create request\n %s\n", parseError(err))
	}

	var resp *http.Response

	if resp, err = c.Do(req); err != nil || resp == nil {
		return tc, fmt.Errorf(" failed to connect to Pi-hole server\n %s\n", parseError(err))
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return tc, fmt.Errorf(" failed to retrieve version from Pi-hole server\n http status code: %d",
			resp.StatusCode)
	}

	var rBody []byte

	if rBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return tc, fmt.Errorf(" failed to read top clients response\n %s\n", parseError(err))
	}

	if err = json.Unmarshal(rBody, &tc); err != nil {
		return tc, fmt.Errorf(" failed to retrieve top clients: check provided api URL and token\n %s\n\n",
			parseError(err))
	}

	return tc, err
}

type QueryTypes struct {
	QueryTypes map[string]float32 `json:"querytypes"`
}

func getQueryTypes(c http.Client, settings *Settings) (qt QueryTypes, err error) {
	var req *http.Request

	var url *url2.URL

	if url, err = url2.Parse(settings.apiUrl); err != nil {
		return qt, fmt.Errorf(" failed to parse API URL\n %s\n", parseError(err))
	}

	var query url2.Values

	if query, err = url2.ParseQuery(url.RawQuery); err != nil {
		return qt, fmt.Errorf(" failed to parse query\n %s\n", parseError(err))
	}

	query.Add("getQueryTypes", strconv.Itoa(settings.showTopClients))
	query.Add("auth", settings.token)

	url.RawQuery = query.Encode()

	if req, err = http.NewRequest("GET", url.String(), nil); err != nil {
		return qt, fmt.Errorf(" failed to create request\n %s\n", parseError(err))
	}

	var resp *http.Response

	if resp, err = c.Do(req); err != nil || resp == nil {
		return qt, fmt.Errorf(" failed to connect to Pi-hole server\n %s\n", parseError(err))
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return qt, fmt.Errorf(" failed to retrieve version from Pi-hole server\n http status code: %d",
			resp.StatusCode)
	}

	var rBody []byte

	if rBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return qt, fmt.Errorf(" failed to read top clients response\n %s\n", parseError(err))
	}

	if err = json.Unmarshal(rBody, &qt); err != nil {
		return qt, fmt.Errorf(" failed to parse query types response\n %s\n", parseError(err))
	}

	return qt, err
}

func checkServer(c http.Client, apiURL string) error {
	var err error

	var req *http.Request

	var url *url2.URL

	if url, err = url2.Parse(apiURL); err != nil {
		return fmt.Errorf(" failed to parse API URL\n %s\n", parseError(err))
	}

	if url.Host == "" {
		return fmt.Errorf(" please specify 'apiUrl' in Pi-hole settings, e.g.\n apiUrl: http://<server>:<port>/admin/api.php")
	}

	if req, err = http.NewRequest("GET", fmt.Sprintf("%s?version",
		url.String()), nil); err != nil {
		return fmt.Errorf("invalid request: %s\n", parseError(err))
	}

	var resp *http.Response

	if resp, err = c.Do(req); err != nil {
		return fmt.Errorf(" failed to connect to Pi-hole server\n %s\n", parseError(err))
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf(" failed to retrieve version from Pi-hole server\n http status code: %d",
			resp.StatusCode)
	}

	var vResp struct {
		Version int `json:"version"`
	}

	var rBody []byte

	if rBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return fmt.Errorf(" Pi-hole server failed to respond\n %s\n", parseError(err))
	}

	if err = json.Unmarshal(rBody, &vResp); err != nil {
		return fmt.Errorf(" invalid response returned from Pi-hole Server\n %s\n", parseError(err))
	}

	if vResp.Version != 3 {
		return fmt.Errorf(" only Pi-hole API version 3 is supported\n version %d was detected", vResp.Version)
	}

	return err
}

func (widget *Widget) adblockSwitch(action string) {
	var req *http.Request

	var url *url2.URL
	url, _ = url2.Parse(widget.settings.apiUrl)

	var query url2.Values
	query, _ = url2.ParseQuery(url.RawQuery)

	query.Add(strings.ToLower(action), "")
	query.Add("auth", widget.settings.token)

	url.RawQuery = query.Encode()

	req, _ = http.NewRequest("GET", url.String(), nil)

	c := getClient()
	resp, _ := c.Do(req)

	defer func() {
		_ = resp.Body.Close()
	}()

	widget.Refresh()
}

func getClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout:   5 * time.Second,
			DisableKeepAlives:     false,
			DisableCompression:    false,
			ResponseHeaderTimeout: 20 * time.Second,
		},
		Timeout: 21 * time.Second,
	}
}
