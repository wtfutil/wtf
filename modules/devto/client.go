package devto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://dev.to/api/articles"

// Article represents a DEV.to article from the public API.
type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	User  struct {
		Username string `json:"username"`
	} `json:"user"`
}

// Client fetches articles from the DEV.to API.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a Client. Pass nil for httpClient to use http.DefaultClient.
// baseURL overrides the API endpoint (useful for testing); pass "" for the default.
func NewClient(httpClient *http.Client, baseURL string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{httpClient: httpClient, baseURL: baseURL}
}

// FetchArticles retrieves articles matching the given filters.
func (c *Client) FetchArticles(ctx context.Context, tag, username, state string, perPage int) ([]Article, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	q := u.Query()
	if tag != "" {
		q.Set("tag", tag)
	}
	if username != "" {
		q.Set("username", username)
	}
	if state != "" {
		q.Set("state", state)
	}
	if perPage > 0 {
		q.Set("per_page", fmt.Sprintf("%d", perPage))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from DEV.to API", resp.StatusCode)
	}

	var articles []Article
	if err := json.NewDecoder(resp.Body).Decode(&articles); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return articles, nil
}
