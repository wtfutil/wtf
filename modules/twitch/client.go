package twitch

import (
	"fmt"

	"github.com/nicklaw5/helix"
)

type Twitch struct {
	client *helix.Client
}

func NewClient(clientId string) *Twitch {
	client, err := helix.NewClient(&helix.Options{
		ClientID: clientId,
	})
	if err != nil {
		fmt.Println(err)
	}
	return &Twitch{client: client}
}

func (t *Twitch) TopStreams(params *helix.StreamsParams) (*helix.StreamsResponse, error) {
	if params == nil {
		params = &helix.StreamsParams{}
	}
	return t.client.GetStreams(params)
}
