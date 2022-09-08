package twitch

import (
	helix "github.com/nicklaw5/helix/v2"
)

type Twitch struct {
	client           *helix.Client
	UserRefreshToken string
	UserID           string
	Streams          string
}

type ClientOpts struct {
	ClientID         string
	ClientSecret     string
	AppAccessToken   string
	UserAccessToken  string
	UserRefreshToken string
	RedirectURI      string
	Streams          string
	UserID           string
}

func NewClient(opts *ClientOpts) (*Twitch, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        opts.ClientID,
		ClientSecret:    opts.ClientSecret,
		AppAccessToken:  opts.AppAccessToken,
		UserAccessToken: opts.UserAccessToken,
		RedirectURI:     opts.RedirectURI,
	})
	if err != nil {
		return nil, err
	}

	t := &Twitch{client: client}
	t.UserRefreshToken = opts.UserRefreshToken
	t.UserID = opts.UserID
	t.Streams = opts.Streams
	if opts.AppAccessToken == "" && opts.ClientSecret != "" {
		if err := t.RefreshOAuthToken(); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *Twitch) RefreshOAuthToken() error {
	if t.Streams == "followed" {
		userResp, err := t.client.RefreshUserAccessToken(t.UserRefreshToken)
		if err != nil {
			return err
		}
		t.client.SetUserAccessToken(userResp.Data.AccessToken)
		t.UserRefreshToken = userResp.Data.RefreshToken
	} else if t.Streams == "top" {
		appResp, err := t.client.RequestAppAccessToken([]string{})
		if err != nil {
			return err
		}
		t.client.SetAppAccessToken(appResp.Data.AccessToken)
	}

	return nil
}

func (t *Twitch) TopStreams(params *helix.StreamsParams) (*helix.StreamsResponse, error) {
	if params == nil {
		params = &helix.StreamsParams{}
	}
	return t.client.GetStreams(params)
}

func (t *Twitch) FollowedStreams(params *helix.FollowedStreamsParams) (*helix.StreamsResponse, error) {
	if params == nil {
		params = &helix.FollowedStreamsParams{}
	}
	return t.client.GetFollowedStream(params)
}
