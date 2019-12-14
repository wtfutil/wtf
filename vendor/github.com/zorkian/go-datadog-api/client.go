/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Client is the object that handles talking to the Datadog API. This maintains
// state information for a particular application connection.
type Client struct {
	apiKey, appKey, baseUrl string

	//The Http Client that is used to make requests
	HttpClient   *http.Client
	RetryTimeout time.Duration

	//Option to specify extra headers like User-Agent
	ExtraHeader map[string]string
}

// valid is the struct to unmarshal validation endpoint responses into.
type valid struct {
	Errors  []string `json:"errors"`
	IsValid bool     `json:"valid"`
}

// NewClient returns a new datadog.Client which can be used to access the API
// methods. The expected argument is the API key.
func NewClient(apiKey, appKey string) *Client {
	baseUrl := os.Getenv("DATADOG_HOST")
	if baseUrl == "" {
		baseUrl = "https://api.datadoghq.com"
	}

	return &Client{
		apiKey:       apiKey,
		appKey:       appKey,
		baseUrl:      baseUrl,
		HttpClient:   http.DefaultClient,
		RetryTimeout: time.Duration(60 * time.Second),
	}
}

// SetKeys changes the value of apiKey and appKey.
func (c *Client) SetKeys(apiKey, appKey string) {
	c.apiKey = apiKey
	c.appKey = appKey
}

// SetBaseUrl changes the value of baseUrl.
func (c *Client) SetBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

// GetBaseUrl returns the baseUrl.
func (c *Client) GetBaseUrl() string {
	return c.baseUrl
}

// Validate checks if the API key (not the APP key) is valid.
func (client *Client) Validate() (bool, error) {
	var out valid
	var resp *http.Response

	uri, err := client.uriForAPI("/v1/validate")
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("DD-API-KEY", client.apiKey)
	req.Header.Set("DD-APPLICATION-KEY", client.appKey)

	resp, err = client.doRequestWithRetries(req, client.RetryTimeout)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return false, err
	}

	return out.IsValid, nil
}
