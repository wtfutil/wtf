package twitterstats

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Client contains state that allows stats to be fetched about a list of Twitter users
type Client struct {
	httpClient  *http.Client
	screenNames []string
}

// TwitterStats Represents a stats snapshot for a single Twitter user at a point in time
type TwitterStats struct {
	FollowerCount int64 `json:"followers_count"`
	TweetCount    int64 `json:"statuses_count"`
}

const (
	userTimelineURL = "https://api.twitter.com/1.1/users/show.json"
)

// NewClient creates a new twitterstats client that contains an OAuth2 HTTP client which can be used
func NewClient(settings *Settings) *Client {
	usernames := make([]string, len(settings.screenNames))
	for i, username := range settings.screenNames {
		var ok bool
		if usernames[i], ok = username.(string); !ok {
			log.Fatalf("All `screenName`s in twitterstats config must be of type string")
		}
	}

	var httpClient *http.Client
	// If a bearer token is supplied, use that directly.  Otherwise, let the Oauth client fetch a token
	// using the consumer key and secret.
	if settings.bearerToken == "" {
		conf := &clientcredentials.Config{
			ClientID:     settings.consumerKey,
			ClientSecret: settings.consumerSecret,
			TokenURL:     "https://api.twitter.com/oauth2/token",
		}
		httpClient = conf.Client(context.Background())
	} else {
		ctx := context.Background()
		httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: settings.bearerToken,
			TokenType:   "Bearer",
		}))
	}

	client := Client{
		httpClient:  httpClient,
		screenNames: usernames,
	}

	return &client
}

// GetStatsForUser Fetches stats for a single user.  If there is an error fetching or parsing the response
// from the Twitter API, an empty stats struct will be returned.
func (client *Client) GetStatsForUser(username string) TwitterStats {
	stats := TwitterStats{
		FollowerCount: 0,
		TweetCount:    0,
	}

	url := fmt.Sprintf("%s?screen_name=%s", userTimelineURL, username)
	resp, err := client.httpClient.Get(url)
	if err != nil {
		return stats
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return stats
	}

	// If there is an error while parsing, just discard the error and return the empty stats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return stats
	}

	return stats
}

// GetStats Returns a slice of `TwitterStats` structs for each username in `client.screenNames` in the same
// order of `client.screenNames`
func (client *Client) GetStats() []TwitterStats {
	stats := make([]TwitterStats, len(client.screenNames))

	for i, username := range client.screenNames {
		stats[i] = client.GetStatsForUser(username)
	}

	return stats
}
