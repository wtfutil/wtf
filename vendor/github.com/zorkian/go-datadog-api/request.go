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

// uriForAPI is to be called with something like "/v1/events" and it will give
// the proper request URI to be posted to.
func (client *Client) uriForAPI(api string) (string, error) {
	apiBase, err := url.Parse(client.baseUrl + "/api" + api)
	if err != nil {
		return "", err
	}
	q := apiBase.Query()
	q.Add("api_key", client.apiKey)
	q.Add("application_key", client.appKey)
	apiBase.RawQuery = q.Encode()
	return apiBase.String(), nil
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

	// If they don't care about the body, then we don't care to give them one,
	// so bail out because we're done.
	if out == nil {
		// read the response body so http conn can be reused immediately
		io.Copy(ioutil.Discard, resp.Body)
		return nil
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
	if bodyReader != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}
