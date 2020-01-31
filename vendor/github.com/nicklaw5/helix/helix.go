package helix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// APIBaseURL is the base URL for composing API requests.
	APIBaseURL = "https://api.twitch.tv/helix"

	// AuthBaseURL is the base URL for composing authentication requests.
	AuthBaseURL = "https://id.twitch.tv/oauth2"
)

// HTTPClient ...
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client ...
type Client struct {
	mu           sync.RWMutex
	opts         *Options
	baseURL      string
	lastResponse *Response
}

// Options ...
type Options struct {
	ClientID        string
	ClientSecret    string
	AppAccessToken  string
	UserAccessToken string
	UserAgent       string
	RedirectURI     string
	Scopes          []string
	HTTPClient      HTTPClient
	RateLimitFunc   RateLimitFunc
}

// DateRange is a generic struct used by various responses.
type DateRange struct {
	StartedAt Time `json:"started_at"`
	EndedAt   Time `json:"ended_at"`
}

// RateLimitFunc ...
type RateLimitFunc func(*Response) error

// ResponseCommon ...
type ResponseCommon struct {
	StatusCode   int
	Header       http.Header
	Error        string `json:"error"`
	ErrorStatus  int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func (rc *ResponseCommon) convertHeaderToInt(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

// GetRateLimit returns the "RateLimit-Limit" header as an int.
func (rc *ResponseCommon) GetRateLimit() int {
	return rc.convertHeaderToInt(rc.Header.Get("RateLimit-Limit"))
}

// GetRateLimitRemaining returns the "RateLimit-Remaining" header as an int.
func (rc *ResponseCommon) GetRateLimitRemaining() int {
	return rc.convertHeaderToInt(rc.Header.Get("RateLimit-Remaining"))
}

// GetRateLimitReset returns the "RateLimit-Reset" header as an int.
func (rc *ResponseCommon) GetRateLimitReset() int {
	return rc.convertHeaderToInt(rc.Header.Get("RateLimit-Reset"))
}

// Response ...
type Response struct {
	ResponseCommon
	Data interface{}
}

// Pagination ...
type Pagination struct {
	Cursor string `json:"cursor"`
}

// NewClient returns a new Twicth Helix API client. It returns an
// if clientID is an empty string. It is concurrecy safe.
func NewClient(options *Options) (*Client, error) {
	if options.ClientID == "" {
		return nil, errors.New("A client ID was not provided but is required")
	}

	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}

	client := &Client{
		opts:    options,
		baseURL: APIBaseURL,
	}

	return client, nil
}

func (c *Client) get(path string, respData, reqData interface{}) (*Response, error) {
	return c.sendRequest(http.MethodGet, path, respData, reqData)
}

func (c *Client) post(path string, respData, reqData interface{}) (*Response, error) {
	return c.sendRequest(http.MethodPost, path, respData, reqData)
}

func (c *Client) put(path string, respData, reqData interface{}) (*Response, error) {
	return c.sendRequest(http.MethodPut, path, respData, reqData)
}

func (c *Client) sendRequest(method, path string, respData, reqData interface{}) (*Response, error) {
	resp := &Response{}
	if respData != nil {
		resp.Data = respData
	}

	req, err := c.newRequest(method, path, reqData)
	if err != nil {
		return nil, err
	}

	err = c.doRequest(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func buildQueryString(req *http.Request, v interface{}) (string, error) {
	isNil, err := isZero(v)
	if err != nil {
		return "", err
	}

	if isNil {
		return "", nil
	}

	query := req.URL.Query()
	vType := reflect.TypeOf(v).Elem()
	vValue := reflect.ValueOf(v).Elem()

	for i := 0; i < vType.NumField(); i++ {
		var defaultValue string

		field := vType.Field(i)
		tag := field.Tag.Get("query")

		// Get the default value from the struct tag
		if strings.Contains(tag, ",") {
			tagSlice := strings.Split(tag, ",")

			tag = tagSlice[0]
			defaultValue = tagSlice[1]
		}

		if field.Type.Kind() == reflect.Slice {
			// Attach any slices as query params
			fieldVal := vValue.Field(i)
			for j := 0; j < fieldVal.Len(); j++ {
				query.Add(tag, fmt.Sprintf("%v", fieldVal.Index(j)))
			}
		} else if isDatetimeTagField(tag) {
			// Get and correctly format datetime fields, and attach them query params
			dateStr := fmt.Sprintf("%v", vValue.Field(i))

			if strings.Contains(dateStr, " m=") {
				datetimeSplit := strings.Split(dateStr, " m=")
				dateStr = datetimeSplit[0]
			}

			date, err := time.Parse(requestDateTimeFormat, dateStr)
			if err != nil {
				return "", err
			}

			// Determine if the date has been set. If it has we'll add it to the query.
			if !date.IsZero() {
				query.Add(tag, date.Format(time.RFC3339))
			}
		} else {
			// Add any scalar values as query params
			fieldVal := fmt.Sprintf("%v", vValue.Field(i))

			// If no value was set by the user, use the default
			// value specified in the struct tag.
			if fieldVal == "" || fieldVal == "0" {
				if defaultValue == "" {
					continue
				}

				fieldVal = defaultValue
			}

			query.Add(tag, fieldVal)
		}
	}

	return query.Encode(), nil
}

func isZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}

func (c *Client) newRequest(method, path string, data interface{}) (*http.Request, error) {
	url := c.getBaseURL(path) + path

	// We want to send Webhook POST request data in JSON form.
	// So here we determine if data is of type `WebhookSubscriptionPayload`.
	_, ok := data.(*WebhookSubscriptionPayload)
	if ok {
		return c.newJSONRequest(method, url, data)
	}

	return c.newStandardRequest(method, url, data)
}

func (c *Client) newStandardRequest(method, url string, data interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return req, nil
	}

	query, err := buildQueryString(req, data)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query

	return req, nil
}

func (c *Client) newJSONRequest(method, url string, data interface{}) (*http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) getBaseURL(path string) string {
	for _, authPath := range authPaths {
		if strings.Contains(path, authPath) {
			return AuthBaseURL
		}
	}

	return APIBaseURL
}

func (c *Client) doRequest(req *http.Request, resp *Response) error {
	c.setRequestHeaders(req)

	rateLimitFunc := c.opts.RateLimitFunc

	for {
		if c.lastResponse != nil && rateLimitFunc != nil {
			err := rateLimitFunc(c.lastResponse)
			if err != nil {
				return err
			}
		}

		response, err := c.opts.HTTPClient.Do(req)
		if err != nil {
			return fmt.Errorf("Failed to execute API request: %s", err.Error())
		}
		defer response.Body.Close()

		resp.Header = response.Header

		setResponseStatusCode(resp, "StatusCode", response.StatusCode)

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		// Only attempt to decode the response if we have a response we can handle
		if len(bodyBytes) > 0 && resp.StatusCode < http.StatusInternalServerError {
			if resp.Data != nil && resp.StatusCode < http.StatusBadRequest {
				// Successful request
				err = json.Unmarshal(bodyBytes, &resp.Data)
			} else {
				// Failed request
				err = json.Unmarshal(bodyBytes, &resp)
			}

			if err != nil {
				return fmt.Errorf("Failed to decode API response: %s", err.Error())
			}
		}

		if rateLimitFunc == nil {
			break
		} else {
			c.mu.Lock()
			c.lastResponse = resp
			c.mu.Unlock()

			if rateLimitFunc != nil &&
				c.lastResponse.StatusCode == http.StatusTooManyRequests {
				// Rate limit exceeded, retry to send request after
				// applying rate limiter callback
				continue
			}

			break
		}
	}

	return nil
}

func (c *Client) setRequestHeaders(req *http.Request) {
	opts := c.opts

	req.Header.Set("Client-ID", opts.ClientID)

	if opts.UserAgent != "" {
		req.Header.Set("User-Agent", opts.UserAgent)
	}

	var bearerToken string
	if opts.AppAccessToken != "" {
		bearerToken = opts.AppAccessToken
	}
	if opts.UserAccessToken != "" {
		bearerToken = opts.UserAccessToken
	}

	if bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+bearerToken)
	}
}

func setResponseStatusCode(v interface{}, fieldName string, code int) {
	s := reflect.ValueOf(v).Elem()
	field := s.FieldByName(fieldName)
	field.SetInt(int64(code))
}

// SetAppAccessToken ...
func (c *Client) SetAppAccessToken(accessToken string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opts.AppAccessToken = accessToken
}

// SetUserAccessToken ...
func (c *Client) SetUserAccessToken(accessToken string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opts.UserAccessToken = accessToken
}

// SetUserAgent ...
func (c *Client) SetUserAgent(userAgent string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opts.UserAgent = userAgent
}

// SetScopes ...
func (c *Client) SetScopes(scopes []string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opts.Scopes = scopes
}

// SetRedirectURI ...
func (c *Client) SetRedirectURI(uri string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opts.RedirectURI = uri
}
