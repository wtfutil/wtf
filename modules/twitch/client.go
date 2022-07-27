package twitch

import (
	helix "github.com/nicklaw5/helix/v2"
)

type Twitch struct {
	client *helix.Client
}

type ClientOpts struct {
	ClientID        string
	ClientSecret    string
	AppAccessToken  string
	UserAccessToken string
}

func NewClient(opts *ClientOpts) (*Twitch, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        opts.ClientID,
		ClientSecret:    opts.ClientSecret,
		AppAccessToken:  opts.AppAccessToken,
		UserAccessToken: opts.UserAccessToken,
	})
	if err != nil {
		return nil, err
	}

	t := &Twitch{client: client}

	if opts.AppAccessToken == "" && opts.ClientSecret != "" {
		if err := t.RefreshOAuthToken(); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *Twitch) RefreshOAuthToken() error {
	resp, err := t.client.RequestAppAccessToken([]string{})
	if err != nil {
		return err
	}
	t.client.SetAppAccessToken(resp.Data.AccessToken)
	return nil
}

func (t *Twitch) TopStreams(params *helix.StreamsParams) (*helix.StreamsResponse, error) {
	if params == nil {
		params = &helix.StreamsParams{}
	}
	return t.client.GetStreams(params)
}
