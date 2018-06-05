/*
 * NewRelic API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package newrelic

import (
	"net/http"
	"net/url"
	"time"
)

const (
	// defaultAPIURL is the default base URL for New Relic's latest API.
	defaultAPIURL = "https://api.newrelic.com/v2/"
	// defaultTimeout is the default timeout for the http.Client used.
	defaultTimeout = 5 * time.Second
)

// Client provides a set of methods to interact with the New Relic API.
type Client struct {
	apiKey     string
	httpClient *http.Client
	url        *url.URL
}

// NewWithHTTPClient returns a new Client object for interfacing with the New
// Relic API, allowing for override of the http.Client object.
func NewWithHTTPClient(apiKey string, client *http.Client) *Client {
	u, err := url.Parse(defaultAPIURL)
	if err != nil {
		panic(err)
	}
	return &Client{
		apiKey:     apiKey,
		httpClient: client,
		url:        u,
	}
}

// NewClient returns a new Client object for interfacing with the New Relic API.
func NewClient(apiKey string) *Client {
	return NewWithHTTPClient(apiKey, &http.Client{Timeout: defaultTimeout})
}
