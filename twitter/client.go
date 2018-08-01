package twitter

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/senorprogrammer/wtf/wtf"
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
func NewClient(url string) *Client {
	client := Client{
		apiBase:    url,
		screenName: "wtfutil",
		count:      5,
	}

	client.loadAPICredentials()

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

func (client *Client) loadAPICredentials() {
	client.bearerToken = wtf.Config.UString(
		"wtf.mods.twitter.bearerToken",
		os.Getenv("WTF_TWITTER_BEARER_TOKEN"),
	)
}
