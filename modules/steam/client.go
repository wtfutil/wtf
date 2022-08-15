package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Steam struct {
	client  *http.Client
	baseUrl string
}

type ClientOpts struct {
	key     string
	baseUrl string
}

func NewClient(opts *ClientOpts) *Steam {
	baseUrl := opts.baseUrl

	if opts.baseUrl == "" {
		baseUrl = "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key="
	}

	baseUrl += opts.key + "&steamids="

	return &Steam{
		client:  &http.Client{},
		baseUrl: baseUrl,
	}
}

type Player struct {
	Personaname   string `json:"personaname"`
	ProfileUrl    string `json:"profileurl"`
	Personastate  int    `json:"personastate"`
	Gameextrainfo string `json:"gameextrainfo"`
}

type SteamResponse struct {
	Response struct {
		Players []Player `json:"players"`
	} `json:"response"`
}

func (s *Steam) Status(steamID string) (*Player, error) {
	resp, err := s.fetch(steamID)
	if err != nil {
		return nil, err
	}

	var response SteamResponse

	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return &response.Response.Players[0], nil
}

func (s *Steam) fetch(id string) ([]byte, error) {
	resp, err := http.Get(s.baseUrl + id)

	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching %s steam status: %v, status: %d", id, err, resp.StatusCode)
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
