package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/* NOTE: Currently single application ONLY
* bearer tokens are only supported for applications, not single-users
 */

// Client represents the data required to connect to the Twitter API
type Client struct {
	apiBase     string
	bearerToken string
	count       int
	screenName  string
}

// NewClient creates and returns a new Twitter client
func NewClient(settings *Settings) *Client {
	client := Client{
		apiBase:     "https://api.twitter.com/1.1/",
		count:       settings.count,
		screenName:  "",
		bearerToken: settings.bearerToken,
	}

	return &client
}

/* -------------------- Public Functions -------------------- */

// Tweets returns a list of tweets of a user
func (client *Client) Tweets() []Tweet {
	tweets, err := client.tweets()
	if err != nil {
		return []Tweet{}
	}

	return tweets
}

/* -------------------- Private Functions -------------------- */

// tweets is the private interface for retrieving the list of user tweets
func (client *Client) tweets() (tweets []Tweet, err error) {
	apiURL := fmt.Sprintf(
		"%s/statuses/user_timeline.json?screen_name=%s&count=%s",
		client.apiBase,
		client.screenName,
		strconv.Itoa(client.count),
	)

	data, err := Request(client.bearerToken, apiURL)
	if err != nil {
		return tweets, err
	}
	err = json.Unmarshal(data, &tweets)

	return
}
