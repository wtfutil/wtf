package git

import ()

type Client struct {
	Repository string
}

func NewClient() *Client {
	client := Client{
		Repository: "./Documents/Lendesk/core-api",
	}

	return &client
}
