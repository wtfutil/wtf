package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

/* NOTE: Currently single application ONLY
* bearer tokens are only supported for applications, not single-users
 */

// Client represents the data required to connect to the Twitter API
type Client struct {
	apiBase    string
	count      int
	screenName string
	httpClient *http.Client
}

// NewClient creates and returns a new Twitter client
func NewClient(settings *Settings) *Client {
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
		apiBase:    "https://api.twitter.com/1.1/",
		count:      settings.count,
		screenName: "",
		httpClient: httpClient,
	}

	return &client
}

/* -------------------- Public Functions -------------------- */

// Tweets returns a list of tweets of a user
func (client *Client) Tweets() []Tweet {
	tweets, err := client.getTweets()
	if err != nil {
		return []Tweet{}
	}

	return tweets
}

/* -------------------- Private Functions -------------------- */

// tweets is the private interface for retrieving the list of user tweets
func (client *Client) getTweets() (tweets []Tweet, err error) {
	apiURL := fmt.Sprintf(
		"%s/statuses/user_timeline.json?screen_name=%s&count=%s",
		client.apiBase,
		client.screenName,
		strconv.Itoa(client.count),
	)

	data, err := Request(client.httpClient, apiURL)
	if err != nil {
		return tweets, err
	}
	err = json.Unmarshal(data, &tweets)

	return
}
