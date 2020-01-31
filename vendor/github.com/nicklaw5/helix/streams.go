package helix

import "time"

// Stream ...
type Stream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	GameID       string    `json:"game_id"`
	TagIDs       []string  `json:"tag_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}

// ManyStreams ...
type ManyStreams struct {
	Streams    []Stream   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// StreamsResponse ...
type StreamsResponse struct {
	ResponseCommon
	Data ManyStreams
}

// StreamsParams ...
type StreamsParams struct {
	After      string   `query:"after"`
	Before     string   `query:"before"`
	First      int      `query:"first,20"`   // Limit 100
	GameIDs    []string `query:"game_id"`    // Limit 100
	Language   []string `query:"language"`   // Limit 100
	Type       string   `query:"type,all"`   // "all" (default), "live" and "vodcast"
	UserIDs    []string `query:"user_id"`    // limit 100
	UserLogins []string `query:"user_login"` // limit 100
}

// GetStreams ...
func (c *Client) GetStreams(params *StreamsParams) (*StreamsResponse, error) {
	resp, err := c.get("/streams", &ManyStreams{}, params)
	if err != nil {
		return nil, err
	}

	streams := &StreamsResponse{}
	streams.StatusCode = resp.StatusCode
	streams.Header = resp.Header
	streams.Error = resp.Error
	streams.ErrorStatus = resp.ErrorStatus
	streams.ErrorMessage = resp.ErrorMessage
	streams.Data.Streams = resp.Data.(*ManyStreams).Streams
	streams.Data.Pagination = resp.Data.(*ManyStreams).Pagination

	return streams, nil
}

// HearthstoneHero ...
type HearthstoneHero struct {
	Class string `json:"class"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

// HearthstonePlayerData ...
type HearthstonePlayerData struct {
	Hero HearthstoneHero `json:"hero"`
}

// HearthstoneMetadata ...
type HearthstoneMetadata struct {
	Broadcaster HearthstonePlayerData `json:"broadcaster"`
	Opponent    HearthstonePlayerData `json:"opponent"`
}

// OverwatchHero ...
type OverwatchHero struct {
	Ability string `json:"ability"`
	Name    string `json:"name"`
	Role    string `json:"role"`
}

// OverwatchBroadcaster ...
type OverwatchBroadcaster struct {
	Hero OverwatchHero `json:"hero"`
}

// OverwatchMetadata ...
type OverwatchMetadata struct {
	Broadcaster OverwatchBroadcaster `json:"broadcaster"`
}

// StreamMetadata ...
type StreamMetadata struct {
	UserID      string              `json:"user_id"`
	UserName    string              `json:"user_name"`
	GameID      string              `json:"game_id"`
	Hearthstone HearthstoneMetadata `json:"hearthstone"`
	Overwatch   OverwatchMetadata   `json:"overwatch"`
}

// ManyStreamsMetadata ...
type ManyStreamsMetadata struct {
	Streams    []StreamMetadata `json:"data"`
	Pagination Pagination       `json:"pagination"`
}

// StreamsMetadataResponse ...
type StreamsMetadataResponse struct {
	ResponseCommon
	Data ManyStreamsMetadata
}

// GetStreamsMetadataRateLimit returns the "Ratelimit-Helixstreamsmetadata-Limit"
// header as an int.
func (sr *StreamsMetadataResponse) GetStreamsMetadataRateLimit() int {
	return sr.convertHeaderToInt(sr.Header.Get("Ratelimit-Helixstreamsmetadata-Limit"))
}

// GetStreamsMetadataRateLimitRemaining returns the "Ratelimit-Helixstreamsmetadata-Remaining"
// header as an int.
func (sr *StreamsMetadataResponse) GetStreamsMetadataRateLimitRemaining() int {
	return sr.convertHeaderToInt(sr.Header.Get("Ratelimit-Helixstreamsmetadata-Remaining"))
}

// StreamsMetadataParams ...
type StreamsMetadataParams StreamsParams

// GetStreamsMetadata ...
func (c *Client) GetStreamsMetadata(params *StreamsMetadataParams) (*StreamsMetadataResponse, error) {
	resp, err := c.get("/streams/metadata", &ManyStreamsMetadata{}, params)
	if err != nil {
		return nil, err
	}

	streams := &StreamsMetadataResponse{}
	streams.StatusCode = resp.StatusCode
	streams.Header = resp.Header
	streams.Error = resp.Error
	streams.ErrorStatus = resp.ErrorStatus
	streams.ErrorMessage = resp.ErrorMessage
	streams.Data.Streams = resp.Data.(*ManyStreamsMetadata).Streams
	streams.Data.Pagination = resp.Data.(*ManyStreamsMetadata).Pagination

	return streams, nil
}
