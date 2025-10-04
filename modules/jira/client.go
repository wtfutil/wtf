package jira

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/wtfutil/wtf/utils"
)

// UserIDCache represent a cached username to account ID mapping
type UserIDCache struct {
	AccountID string
	ExpiresAt time.Time
}

// UserIDCacheMap holds the cache with thread safety
type UserIDCacheMap struct {
	cache map[string]UserIDCache
	mutex sync.RWMutex
}

// Global cache instance
var userIDCache = &UserIDCacheMap{
	cache: make(map[string]UserIDCache),
}

// JQLConversionRequest represents the request body for the JQL conversion API
type JQLConversionRequest struct {
	QueryStrings []string `json:"queryStrings"`
}

// JQLConversionResponse represents the response from the JQL conversion API
type JQLConversionResponse struct {
	QueryStrings []ConvertedQuery `json:"queryStrings"`
}

// ConvertedQuery represents a single converted JQL query
type ConvertedQuery struct {
	Query          string        `json:"query"`
	ConvertedQuery string        `json:"convertedQuery"`
	UserMessages   []UserMessage `json:"userMessages"`
}

// UserMessage represents messages about the conversion
type UserMessage struct {
	MessageKey  string            `json:"messageKey"`
	MessageArgs map[string]string `json:"messageArgs"`
}

// Get retrieves a cache account ID for a username
func (c *UserIDCacheMap) Get(username string) (string, bool) {
	c.mutex.RLock()
	entry, exists := c.cache[username]
	if !exists {
		c.mutex.RUnlock()
		return "", false
	}

	// Check if cache entry has expired
	if time.Now().After(entry.ExpiresAt) {
		c.mutex.RUnlock()
		// Remove expired entry - upgrade to write lock
		c.mutex.Lock()
		delete(c.cache, username)
		c.mutex.Unlock()
		return "", false
	}

	accountID := entry.AccountID
	c.mutex.RUnlock()
	return accountID, true
}

// Set stores a username to account ID mapping with expiration
func (c *UserIDCacheMap) Set(username, accountID string, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[username] = UserIDCache{
		AccountID: accountID,
		ExpiresAt: time.Now().Add(duration),
	}
}

// Clear removes all expired entries from the cache
func (c *UserIDCacheMap) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for username, entry := range c.cache {
		if now.After(entry.ExpiresAt) {
			delete(c.cache, username)
		}
	}
}

// ConvertJQLWithUsername converts a JQL query containing username to account ID
func (widget *Widget) ConvertJQLWithUsername(username string) (string, error) {
	// Check cache first
	if accountID, found := userIDCache.Get(username); found {
		return fmt.Sprintf("assignee = \"%s\"", accountID), nil
	}

	// Create a JQL query with the username that needs conversion
	originalJQL := fmt.Sprintf("assignee = \"%s\"", username)

	// Prepare the request body
	requestBody := JQLConversionRequest{
		QueryStrings: []string{originalJQL},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// Make the POST request to the JQL conversion API
	resp, err := widget.jiraPostRequest("/rest/api/3/jql/pdcleaner", jsonData)
	if err != nil {
		return "", err
	}

	var conversionResult JQLConversionResponse
	err = utils.ParseJSON(&conversionResult, bytes.NewReader(resp))
	if err != nil {
		return "", err
	}

	if len(conversionResult.QueryStrings) == 0 {
		return "", fmt.Errorf("no conversion result for username: %s", username)
	}

	// Return the converted JQL query part (just the assignee part)
	convertedQuery := conversionResult.QueryStrings[0].ConvertedQuery

	// Extract account ID properly
	accountID := extractAccountIDFromJQL(convertedQuery)
	if accountID == "" {
		return "", fmt.Errorf("failed to extract account ID from converted query: %s", convertedQuery)
	}

	// Cache the result for 10 minutes
	userIDCache.Set(username, accountID, 10*time.Minute)

	return convertedQuery, nil
}

// extractAccountIDFromJQL extracts the account ID from a converted JQL query
func extractAccountIDFromJQL(jql string) string {
	// Example: "assignee = \"account:5b10ac8d82e05b22cc7d4ef5\""
	// We want to extract: "account:5b10ac8d82e05b22cc7d4ef5"

	start := strings.Index(jql, "\"")
	if start == -1 {
		return ""
	}

	end := strings.LastIndex(jql, "\"")
	if end == -1 || end <= start {
		return ""
	}

	return jql[start+1 : end]
}

// IssuesFor returns a collection of issues for a given collection of projects.
// If username is provided, it scopes the issues to that person
func (widget *Widget) IssuesFor(username string, projects []string, jql string) (*SearchResult, error) {
	query := []string{}

	var projQuery = getProjectQuery(projects)
	if projQuery != "" {
		query = append(query, projQuery)
	}

	if username != "" {
		// Convert JQL with username to account ID
		convertedJQL, err := widget.ConvertJQLWithUsername(username)
		if err != nil {
			return &SearchResult{}, fmt.Errorf("failed to convert username %s to account ID: %v", username, err)
		}
		query = append(query, convertedJQL)
	}

	if jql != "" {
		query = append(query, jql)
	}

	v := url.Values{}

	v.Set("jql", strings.Join(query, " AND "))

	url := fmt.Sprintf("/rest/api/2/search?%s", v.Encode())

	resp, err := widget.jiraRequest(url)
	if err != nil {
		return &SearchResult{}, err
	}

	searchResult := &SearchResult{}
	err = utils.ParseJSON(searchResult, bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func buildJql(key string, value string) string {
	return fmt.Sprintf("%s = \"%s\"", key, value)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) jiraRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", widget.settings.domain, path)

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	if widget.settings.personalAccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+widget.settings.personalAccessToken)
	} else {
		req.SetBasicAuth(widget.settings.email, widget.settings.apiKey)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !widget.settings.verifyServerCertificate,
			},
			Proxy: http.ProxyFromEnvironment,
		},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (widget *Widget) jiraPostRequest(path string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s%s", widget.settings.domain, path)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if widget.settings.personalAccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+widget.settings.personalAccessToken)
	} else {
		req.SetBasicAuth(widget.settings.email, widget.settings.apiKey)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !widget.settings.verifyServerCertificate,
			},
			Proxy: http.ProxyFromEnvironment,
		},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getProjectQuery(projects []string) string {
	singleEmptyProject := len(projects) == 1 && projects[0] == ""
	if len(projects) == 0 || singleEmptyProject {
		return ""
	} else if len(projects) == 1 {
		return buildJql("project", projects[0])
	}

	quoted := make([]string, len(projects))
	for i := range projects {
		quoted[i] = fmt.Sprintf("\"%s\"", projects[i])
	}
	return fmt.Sprintf("project in (%s)", strings.Join(quoted, ", "))
}
