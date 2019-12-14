/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
)

// Response contains common fields that might be present in any API response.
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func (client *Client) apiAcceptsKeysInHeaders(api string) bool {
	for _, prefix := range []string{"/v1/series", "/v1/check_run", "/v1/events", "/v1/screen"} {
		if strings.HasPrefix(api, prefix) {
			return false
		}
	}
	return true
}

// uriForAPI is to be called with either an API resource like "/v1/events"
// or a full URL like the IP Ranges one
// and it will give the proper request URI to be posted to.
func (client *Client) uriForAPI(api string) (string, error) {
	var err error
	// If api is a URI such as /v1/hosts/, /v2/dashboards... add credentials and return a properly formatted URL
	if !(strings.HasPrefix(api, "https://") || strings.HasPrefix(api, "http://")) {
		apiBase, err := url.Parse(client.baseUrl + "/api" + api)
		if err != nil {
			return "", err
		}
		q := apiBase.Query()
		if !client.apiAcceptsKeysInHeaders(api) {
			q.Add("api_key", client.apiKey)
			q.Add("application_key", client.appKey)
		}
		apiBase.RawQuery = q.Encode()
		return apiBase.String(), nil
	}
	// if api is a generic URL we simply return it
	apiBase, err := url.Parse(api)
	if err != nil {
		return "", err
	}
	return apiBase.String(), nil
}

// URLIPRanges returns the IP Ranges URL used to whitelist IP addresses in use to send data to Datadog
// agents, api, apm, logs, process, synthetics, webhooks
func (client *Client) URLIPRanges() (string, error) {
	baseURL := client.GetBaseUrl()
	// Get the domain from the URL: eu, com...
	domain := strings.Split(baseURL, ".")[2]
	var urlIPRanges string
	switch domain {
	case "eu":
		urlIPRanges = "https://ip-ranges.datadoghq.eu"
	case "com":
		urlIPRanges = "https://ip-ranges.datadoghq.com"
	}
	return urlIPRanges, nil
}

// redactError removes api and application keys from error strings
func (client *Client) redactError(err error) error {
	if err == nil {
		return nil
	}
	errString := err.Error()

	if len(client.apiKey) > 0 {
		errString = strings.Replace(errString, client.apiKey, "redacted", -1)
	}
	if len(client.appKey) > 0 {
		errString = strings.Replace(errString, client.appKey, "redacted", -1)
	}

	// Return original error if no replacements were made to keep the original,
	// probably more useful error type information.
	if errString == err.Error() {
		return err
	}
	return fmt.Errorf("%s", errString)
}

// doJsonRequest is the simplest type of request: a method on a URI that
// returns some JSON result which we unmarshal into the passed interface. It
// wraps doJsonRequestUnredacted to redact api and application keys from
// errors.
func (client *Client) doJsonRequest(method, api string,
	reqbody, out interface{}) error {
	if err := client.doJsonRequestUnredacted(method, api, reqbody, out); err != nil {
		return client.redactError(err)
	}
	return nil
}

// doJsonRequestUnredacted is the simplest type of request: a method on a URI that returns
// some JSON result which we unmarshal into the passed interface.
func (client *Client) doJsonRequestUnredacted(method, api string,
	reqbody, out interface{}) error {
	req, err := client.createRequest(method, api, reqbody)
	if err != nil {
		return err
	}

	// Perform the request and retry it if it's not a POST or PUT request
	var resp *http.Response
	if method == "POST" || method == "PUT" {
		resp, err = client.HttpClient.Do(req)
	} else {
		resp, err = client.doRequestWithRetries(req, client.RetryTimeout)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("API error %s: %s", resp.Status, body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// If we got no body, by default let's just make an empty JSON dict. This
	// saves us some work in other parts of the code.
	if len(body) == 0 {
		body = []byte{'{', '}'}
	}

	// Try to parse common response fields to check whether there's an error reported in a response.
	var common *Response
	err = json.Unmarshal(body, &common)
	if err != nil {
		// UnmarshalTypeError errors are ignored, because in some cases API response is an array that cannot be
		// unmarshalled into a struct.
		_, ok := err.(*json.UnmarshalTypeError)
		if !ok {
			return err
		}
	}
	if common != nil && common.Status == "error" {
		return fmt.Errorf("API returned error: %s", common.Error)
	}

	// If they don't care about the body, then we don't care to give them one,
	// so bail out because we're done.
	if out == nil {
		return nil
	}

	return json.Unmarshal(body, &out)
}

// doRequestWithRetries performs an HTTP request repeatedly for maxTime or until
// no error and no acceptable HTTP response code was returned.
func (client *Client) doRequestWithRetries(req *http.Request, maxTime time.Duration) (*http.Response, error) {
	var (
		err  error
		resp *http.Response
		bo   = backoff.NewExponentialBackOff()
		body []byte
	)

	bo.MaxElapsedTime = maxTime

	// Save the body for retries
	if req.Body != nil {
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return resp, err
		}
	}

	operation := func() error {
		if body != nil {
			r := bytes.NewReader(body)
			req.Body = ioutil.NopCloser(r)
		}

		resp, err = client.HttpClient.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// 2xx all done
			return nil
		} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			// 4xx are not retryable
			return nil
		}

		return fmt.Errorf("Received HTTP status code %d", resp.StatusCode)
	}

	err = backoff.Retry(operation, bo)

	return resp, err
}

func (client *Client) createRequest(method, api string, reqbody interface{}) (*http.Request, error) {
	// Handle the body if they gave us one.
	var bodyReader io.Reader
	if method != "GET" && reqbody != nil {
		bjson, err := json.Marshal(reqbody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bjson)
	}

	apiUrlStr, err := client.uriForAPI(api)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, apiUrlStr, bodyReader)
	if err != nil {
		return nil, err
	}
	if client.apiAcceptsKeysInHeaders(api) {
		req.Header.Set("DD-API-KEY", client.apiKey)
		req.Header.Set("DD-APPLICATION-KEY", client.appKey)
	}
	if bodyReader != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	for k, v := range client.ExtraHeader {
		req.Header.Add(k, v)
	}

	return req, nil
}
