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

type Client struct {
	apiBase     string
	httpClient  *http.Client
	screenNames []string
}

type TwitterStats struct {
	followerCount int64
	tweetCount    int64
}

const (
	userTimelineUrl = "https://api.twitter.com/1.1/users/show.json"
)

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

	// token, err := conf.Token(oauth2.NoContext)
	httpClient := conf.Client(oauth2.NoContext)

	client := Client{
		apiBase:     "https://api.twitter.com/1.1/",
		httpClient:  httpClient,
		screenNames: usernames,
	}

	return &client
}

func (client *Client) GetStats() []TwitterStats {
	stats := make([]TwitterStats, len(client.screenNames))

	for i, username := range client.screenNames {
		stats[i] = TwitterStats{
			followerCount: 0,
			tweetCount:    0,
		}

		res, err := client.httpClient.Get(fmt.Sprintf("%s?screen_name=%s", userTimelineUrl, username))
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
