package twitterstats

import (
	"log"
)

type Client struct {
	apiBase           string
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
	screenNames       []string
}

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

	client := Client{
		apiBase:           "https://api.twitter.com/1.1/",
		consumerKey:       settings.consumerKey,
		consumerSecret:    settings.consumerSecret,
		accessToken:       settings.accessToken,
		accessTokenSecret: settings.accessTokenSecret,
		screenNames:       usernames,
	}

	return &client
}

func (client *Client) GetFollowerCounts() []int64 {
	return []int64{0, 0, 0, 0, 0} // TODO
}

func (client *Client) GetTweetCounts() []int64 {
	return []int64{0, 0, 0, 0, 0} // TODO
}
