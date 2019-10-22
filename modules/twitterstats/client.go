package twitterstats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	followerCount int64
	tweetCount    int64
}

const (
	userTimelineURL = "https://api.twitter.com/1.1/users/show.json"
)

// NewClient creates a new twitterstats client that contains an OAuth2 HTTP client which can be used
func NewClient(settings *Settings) *Client {
	usernames := make([]string, len(settings.screenNames))
	for i, username := range settings.screenNames {
		switch username.(type) {
		default:
			{
				log.Fatalf("All `screenName`s in twitterstats config must be of type string")
			}
		case string:
			usernames[i] = username.(string)
		}

	}

	conf := &clientcredentials.Config{
		ClientID:     settings.consumerKey,
		ClientSecret: settings.consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := conf.Client(oauth2.NoContext)

	client := Client{
		httpClient:  httpClient,
		screenNames: usernames,
	}

	return &client
}

// GetStats Returns a slice of `TwitterStats` structs for each username in `client.screenNames` in the same
// order of `client.screenNames`
func (client *Client) GetStats() []TwitterStats {
	stats := make([]TwitterStats, len(client.screenNames))

	for i, username := range client.screenNames {
		stats[i] = TwitterStats{
			followerCount: 0,
			tweetCount:    0,
		}

		res, err := client.httpClient.Get(fmt.Sprintf("%s?screen_name=%s", userTimelineURL, username))
		if err != nil {
			continue
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
		}

		var parsed map[string]interface{}
		err = json.Unmarshal(body, &parsed)
		if err != nil {
			continue
		}

		stats[i] = TwitterStats{
			followerCount: int64(parsed["followers_count"].(float64)),
			tweetCount:    int64(parsed["statuses_count"].(float64)),
		}
	}

	return stats
}
