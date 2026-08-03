package tennis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultBaseURL = "https://api.livetennisapi.com/api/public/v1"

	// FreeKeyURL is where users can sign up for a free API key.
	FreeKeyURL = "https://livetennisapi.com/subscribe/free"
)

// Sentinel errors so the widget can render specific help text per failure mode.
var (
	errUnauthorized = errors.New("unauthorized (401): invalid or missing API key")
	errRateLimited  = errors.New("rate limited (429): too many requests")
)

// Client fetches matches from the Live Tennis API.
type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a Client. Pass nil for httpClient to use http.DefaultClient.
// baseURL overrides the API endpoint (useful for testing); pass "" for the default.
func NewClient(apiKey string, httpClient *http.Client, baseURL string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{apiKey: apiKey, httpClient: httpClient, baseURL: baseURL}
}

// FetchMatches retrieves matches filtered by status (live|upcoming|completed),
// tour (optional, e.g. atp/wta) and limit (0 = API default).
func (c *Client) FetchMatches(ctx context.Context, status, tour string, limit int) ([]Match, error) {
	u, err := url.Parse(c.baseURL + "/matches")
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	q := u.Query()
	if status != "" {
		q.Set("status", status)
	}
	if tour != "" {
		q.Set("tour", tour)
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	switch resp.StatusCode {
	case http.StatusOK:
		// fall through to parsing
	case http.StatusUnauthorized:
		return nil, errUnauthorized
	case http.StatusTooManyRequests:
		return nil, errRateLimited
	default:
		return nil, fmt.Errorf("unexpected status %d from Live Tennis API", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var envelope matchesResponse
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, fmt.Errorf("parsing Live Tennis API response: %w", err)
	}

	return envelope.Data, nil
}
