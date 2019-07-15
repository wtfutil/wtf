// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// DefaultBaseURL is the default API base url used by Client to send requests to Trello.
const DefaultBaseURL = "https://api.trello.com/1"

// Client is the central object for making API calls. It wraps a http client,
// context, logger and identity configuration (Key and Token) of the Trello member.
type Client struct {
	Client   *http.Client
	Logger   logger
	BaseURL  string
	Key      string
	Token    string
	throttle <-chan time.Time
	testMode bool
	ctx      context.Context
}

type logger interface {
	Debugf(string, ...interface{})
}

// NewClient is a constructor for the Client. It takes the key and token credentials
// of a Trello member to authenticate and authorise requests with.
func NewClient(key, token string) *Client {
	return &Client{
		Client:   http.DefaultClient,
		BaseURL:  DefaultBaseURL,
		Key:      key,
		Token:    token,
		throttle: time.Tick(time.Second / 8), // Actually 10/second, but we're extra cautious
		testMode: false,
		ctx:      context.Background(),
	}
}

// WithContext takes a context.Context, sets it as context on the client and returns
// a Client pointer.
func (c *Client) WithContext(ctx context.Context) *Client {
	newC := *c
	newC.ctx = ctx
	return &newC
}

// Throttle starts receiving throttles from throttle channel each ticker period.
func (c *Client) Throttle() {
	if !c.testMode {
		<-c.throttle
	}
}

// Get takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a GET request on the Trello API endpoint and the path and uses the
// Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Get(path string, args Arguments, target interface{}) error {

	// Trello prohibits more than 10 seconds/second per token
	c.Throttle()

	params := args.ToURLValues()
	c.log("[trello] GET %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}
	req = req.WithContext(c.ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return makeHTTPClientError(url, resp)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s", url)
	}

	return nil
}

// Put takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a PUT request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Put(path string, args Arguments, target interface{}) error {

	// Trello prohibits more than 10 seconds/second per token
	c.Throttle()

	params := args.ToURLValues()
	c.log("[trello] PUT %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("PUT", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid PUT request %s", url)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return makeHTTPClientError(url, resp)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s", url)
	}

	return nil
}

// Post takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a POST request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Post(path string, args Arguments, target interface{}) error {

	// Trello prohibits more than 10 seconds/second per token
	c.Throttle()

	params := args.ToURLValues()
	c.log("[trello] POST %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("POST", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid POST request %s", url)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "HTTP Read error on response for %s", url)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(b))
	err = decoder.Decode(target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s:\n%s", url, string(b))
	}

	return nil
}

// Delete takes a path, Arguments, and a target interface (e.g. Board or Card).
// It runs a DELETE request on the Trello API endpoint with the path and uses
// the Arguments as URL parameters. Then it returns either the target interface
// updated from the response or an error.
func (c *Client) Delete(path string, args Arguments, target interface{}) error {

	c.Throttle()

	params := args.ToURLValues()
	c.log("[trello] DELETE %s?%s", path, params.Encode())

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("DELETE", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid DELETE request %s", url)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "HTTP Read error on response for %s", url)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(b))
	err = decoder.Decode(target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s:\n%s", url, string(b))
	}

	return nil
}

func (c *Client) log(format string, args ...interface{}) {
	if c.Logger != nil {
		c.Logger.Debugf(format, args...)
	}
}
