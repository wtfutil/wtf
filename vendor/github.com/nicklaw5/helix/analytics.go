package helix

// ExtensionAnalytic ...
type ExtensionAnalytic struct {
	ExtensionID string    `json:"extension_id"`
	URL         string    `json:"URL"`
	Type        string    `json:"type"`
	DateRange   DateRange `json:"date_range"`
}

// ManyExtensionAnalytics ...
type ManyExtensionAnalytics struct {
	ExtensionAnalytics []ExtensionAnalytic `json:"data"`
	Pagination         Pagination          `json:"pagination"`
}

// ExtensionAnalyticsResponse ...
type ExtensionAnalyticsResponse struct {
	ResponseCommon
	Data ManyExtensionAnalytics
}

// ExtensionAnalyticsParams ...
type ExtensionAnalyticsParams struct {
	ExtensionID string `query:"extension_id"`
	First       int    `query:"first,20"`
	After       string `query:"after"`
	StartedAt   Time   `query:"started_at"`
	EndedAt     Time   `query:"ended_at"`
	Type        string `query:"type"`
}

// GetExtensionAnalytics returns a URL to the downloadable CSV file
// containing analytics data. Valid for 5 minutes.
func (c *Client) GetExtensionAnalytics(params *ExtensionAnalyticsParams) (*ExtensionAnalyticsResponse, error) {
	resp, err := c.get("/analytics/extensions", &ManyExtensionAnalytics{}, params)
	if err != nil {
		return nil, err
	}

	users := &ExtensionAnalyticsResponse{}
	users.StatusCode = resp.StatusCode
	users.Header = resp.Header
	users.Error = resp.Error
	users.ErrorStatus = resp.ErrorStatus
	users.ErrorMessage = resp.ErrorMessage
	users.Data.ExtensionAnalytics = resp.Data.(*ManyExtensionAnalytics).ExtensionAnalytics
	users.Data.Pagination = resp.Data.(*ManyExtensionAnalytics).Pagination
	return users, nil
}

// GameAnalytic ...
type GameAnalytic struct {
	GameID    string    `json:"game_id"`
	URL       string    `json:"URL"`
	Type      string    `json:"type"`
	DateRange DateRange `json:"date_range"`
}

// ManyGameAnalytics ...
type ManyGameAnalytics struct {
	GameAnalytics []GameAnalytic `json:"data"`
	Pagination    Pagination     `json:"pagination"`
}

// GameAnalyticsResponse ...
type GameAnalyticsResponse struct {
	ResponseCommon
	Data ManyGameAnalytics
}

// GameAnalyticsParams ...
type GameAnalyticsParams struct {
	GameID    string `query:"game_id"`
	First     int    `query:"first,20"`
	After     string `query:"after"`
	StartedAt Time   `query:"started_at"`
	EndedAt   Time   `query:"ended_at"`
	Type      string `query:"type"`
}

// GetGameAnalytics returns a URL to the downloadable CSV file
// containing analytics data for the specified game. Valid for 5 minutes.
func (c *Client) GetGameAnalytics(params *GameAnalyticsParams) (*GameAnalyticsResponse, error) {

	resp, err := c.get("/analytics/games", &ManyGameAnalytics{}, params)
	if err != nil {
		return nil, err
	}

	users := &GameAnalyticsResponse{}
	users.StatusCode = resp.StatusCode
	users.Header = resp.Header
	users.Error = resp.Error
	users.ErrorStatus = resp.ErrorStatus
	users.ErrorMessage = resp.ErrorMessage
	users.Data.GameAnalytics = resp.Data.(*ManyGameAnalytics).GameAnalytics
	users.Data.Pagination = resp.Data.(*ManyGameAnalytics).Pagination

	return users, nil
}
