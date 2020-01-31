package helix

// Game ...
type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}

// ManyGames ...
type ManyGames struct {
	Games []Game `json:"data"`
}

// GamesResponse ...
type GamesResponse struct {
	ResponseCommon
	Data ManyGames
}

// GamesParams ...
type GamesParams struct {
	IDs   []string `query:"id"`   // Limit 100
	Names []string `query:"name"` // Limit 100
}

// GetGames ...
func (c *Client) GetGames(params *GamesParams) (*GamesResponse, error) {
	resp, err := c.get("/games", &ManyGames{}, params)
	if err != nil {
		return nil, err
	}

	games := &GamesResponse{}
	games.StatusCode = resp.StatusCode
	games.Header = resp.Header
	games.Error = resp.Error
	games.ErrorStatus = resp.ErrorStatus
	games.ErrorMessage = resp.ErrorMessage
	games.Data.Games = resp.Data.(*ManyGames).Games

	return games, nil
}

// ManyGamesWithPagination ...
type ManyGamesWithPagination struct {
	ManyGames
	Pagination Pagination `json:"pagination"`
}

// TopGamesParams ...
type TopGamesParams struct {
	After  string `query:"after"`
	Before string `query:"before"`
	First  int    `query:"first,20"` // Limit 100
}

// TopGamesResponse ...
type TopGamesResponse struct {
	ResponseCommon
	Data ManyGamesWithPagination
}

// GetTopGames ...
func (c *Client) GetTopGames(params *TopGamesParams) (*TopGamesResponse, error) {
	resp, err := c.get("/games/top", &ManyGamesWithPagination{}, params)
	if err != nil {
		return nil, err
	}

	games := &TopGamesResponse{}
	games.StatusCode = resp.StatusCode
	games.Header = resp.Header
	games.Error = resp.Error
	games.ErrorStatus = resp.ErrorStatus
	games.ErrorMessage = resp.ErrorMessage
	games.Data.Games = resp.Data.(*ManyGamesWithPagination).Games
	games.Data.Pagination = resp.Data.(*ManyGamesWithPagination).Pagination

	return games, nil
}
