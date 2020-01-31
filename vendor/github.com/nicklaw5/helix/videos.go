package helix

// Video ...
type Video struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
	PublishedAt  string `json:"published_at"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Viewable     string `json:"viewable"`
	ViewCount    int    `json:"view_count"`
	Language     string `json:"language"`
	Type         string `json:"type"`
	Duration     string `json:"duration"`
}

// ManyVideos ...
type ManyVideos struct {
	Videos     []Video    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// VideosParams ...
type VideosParams struct {
	IDs    []string `query:"id"`      // Limit 100
	UserID string   `query:"user_id"` // Limit 1
	GameID string   `query:"game_id"` // Limit 1

	// Optional
	After    string `query:"after"`
	Before   string `query:"before"`
	First    int    `query:"first,20"`   // Limit 100
	Language string `query:"language"`   // Limit 1
	Period   string `query:"period,all"` // "all" (default), "day", "month", and "week"
	Sort     string `query:"sort,time"`  // "time" (default), "trending", and "views"
	Type     string `query:"type,all"`   // "all" (default), "upload", "archive", and "highlight"
}

// VideosResponse ...
type VideosResponse struct {
	ResponseCommon
	Data ManyVideos
}

// GetVideos gets video information by video ID (one or more), user ID (one only),
// or game ID (one only).
func (c *Client) GetVideos(params *VideosParams) (*VideosResponse, error) {
	resp, err := c.get("/videos", &ManyVideos{}, params)
	if err != nil {
		return nil, err
	}

	videos := &VideosResponse{}
	videos.StatusCode = resp.StatusCode
	videos.Header = resp.Header
	videos.Error = resp.Error
	videos.ErrorStatus = resp.ErrorStatus
	videos.ErrorMessage = resp.ErrorMessage
	videos.Data.Videos = resp.Data.(*ManyVideos).Videos
	videos.Data.Pagination = resp.Data.(*ManyVideos).Pagination

	return videos, nil
}
