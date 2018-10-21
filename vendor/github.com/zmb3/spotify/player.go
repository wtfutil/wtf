package spotify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// PlayerDevice contains information about a device that a user can play music on
type PlayerDevice struct {
	// ID of the device. This may be empty.
	ID ID `json:"id"`
	// Active If this device is the currently active device.
	Active bool `json:"is_active"`
	// Restricted Whether controlling this device is restricted. At present if
	// this is "true" then no Web API commands will be accepted by this device.
	Restricted bool `json:"is_restricted"`
	// Name The name of the device.
	Name string `json:"name"`
	// Type of device, such as "Computer", "Smartphone" or "Speaker".
	Type string `json:"type"`
	// Volume The current volume in percent.
	Volume int `json:"volume_percent"`
}

// PlayerState contains information about the current playback.
type PlayerState struct {
	CurrentlyPlaying
	// Device The device that is currently active
	Device PlayerDevice `json:"device"`
	// ShuffleState Shuffle is on or off
	ShuffleState bool `json:"shuffle_state"`
	// RepeatState off, track, context
	RepeatState string `json:"repeat_state"`
}

// PlaybackContext is the playback context
type PlaybackContext struct {
	// ExternalURLs of the context, or null if not available.
	ExternalURLs map[string]string `json:"external_urls"`
	// Endpoint of the context, or null if not available.
	Endpoint string `json:"href"`
	// Type of the item's context. Can be one of album, artist or playlist.
	Type string `json:"type"`
	// URI is the Spotify URI for the context.
	URI URI `json:"uri"`
}

// CurrentlyPlaying contains the information about currently playing items
type CurrentlyPlaying struct {
	// Timestamp when data was fetched
	Timestamp int64 `json:"timestamp"`
	// PlaybackContext current context
	PlaybackContext PlaybackContext `json:"context"`
	// Progress into the currently playing track.
	Progress int `json:"progress_ms"`
	// Playing If something is currently playing.
	Playing bool `json:"is_playing"`
	// The currently playing track. Can be null.
	Item *FullTrack `json:"Item"`
}

type RecentlyPlayedItem struct {
	// Track is the track information
	Track SimpleTrack `json:"track"`

	// PlayedAt is the time that this song was played
	PlayedAt time.Time `json:"played_at"`

	// PlaybackContext is the current playback context
	PlaybackContext PlaybackContext `json:"context"`
}

type RecentlyPlayedResult struct {
	Items []RecentlyPlayedItem `json:"items"`
}

// PlaybackOffset can be specified either by track URI OR Position. If both are present the
// request will return 400 BAD REQUEST. If incorrect values are provided for position or uri,
// the request may be accepted but with an unpredictable resulting action on playback.
type PlaybackOffset struct {
	// Position is zero based and can’t be negative.
	Position int `json:"position,omitempty"`
	// URI is a string representing the uri of the item to start at.
	URI URI `json:"uri,omitempty"`
}

type PlayOptions struct {
	// DeviceID The id of the device this command is targeting. If not
	// supplied, the user's currently active device is the target.
	DeviceID *ID `json:"-"`
	// PlaybackContext Spotify URI of the context to play.
	// Valid contexts are albums, artists & playlists.
	PlaybackContext *URI `json:"context_uri,omitempty"`
	// URIs Array of the Spotify track URIs to play
	URIs []URI `json:"uris,omitempty"`
	// PlaybackOffset Indicates from where in the context playback should start.
	// Only available when context corresponds to an album or playlist
	// object, or when the URIs parameter is used.
	PlaybackOffset *PlaybackOffset `json:"offset,omitempty"`
	// PositionMs Indicates from what position to start playback.
	// Must be a positive number. Passing in a position that is greater
	// than the length of the track will cause the player to start playing the next song.
	// Defaults to 0, starting a track from the beginning.
	PositionMs int `json:"position_ms,omitempty"`
}

// RecentlyPlayedOptions describes options for the recently-played request. All
// fields are optional. Only one of `AfterEpochMs` and `BeforeEpochMs` may be
// given. Note that it seems as if Spotify only remembers the fifty most-recent
// tracks as of right now.
type RecentlyPlayedOptions struct {
	// Limit is the maximum number of items to return. Must be no greater than
	// fifty.
	Limit int

	// AfterEpochMs is a Unix epoch in milliseconds that describes a time after
	// which to return songs.
	AfterEpochMs int64

	// BeforeEpochMs is a Unix epoch in milliseconds that describes a time
	// before which to return songs.
	BeforeEpochMs int64
}

// PlayerDevices information about available devices for the current user.
//
// Requires the ScopeUserReadPlaybackState scope in order to read information
func (c *Client) PlayerDevices() ([]PlayerDevice, error) {
	var result struct {
		PlayerDevices []PlayerDevice `json:"devices"`
	}

	err := c.get(c.baseURL+"me/player/devices", &result)
	if err != nil {
		return nil, err
	}

	return result.PlayerDevices, nil
}

// PlayerState gets information about the playing state for the current user
//
// Requires the ScopeUserReadPlaybackState scope in order to read information
func (c *Client) PlayerState() (*PlayerState, error) {
	return c.PlayerStateOpt(nil)
}

// PlayerStateOpt is like PlayerState, but it accepts additional
// options for sorting and filtering the results.
func (c *Client) PlayerStateOpt(opt *Options) (*PlayerState, error) {
	spotifyURL := c.baseURL + "me/player"
	if opt != nil {
		v := url.Values{}
		if opt.Country != nil {
			v.Set("market", *opt.Country)
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	var result PlayerState

	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// PlayerCurrentlyPlaying gets information about the currently playing status
// for the current user.
//
// Requires the ScopeUserReadCurrentlyPlaying scope or the ScopeUserReadPlaybackState
// scope in order to read information
func (c *Client) PlayerCurrentlyPlaying() (*CurrentlyPlaying, error) {
	return c.PlayerCurrentlyPlayingOpt(nil)
}

// PlayerCurrentlyPlayingOpt is like PlayerCurrentlyPlaying, but it accepts
// additional options for sorting and filtering the results.
func (c *Client) PlayerCurrentlyPlayingOpt(opt *Options) (*CurrentlyPlaying, error) {
	spotifyURL := c.baseURL + "me/player/currently-playing"
	if opt != nil {
		v := url.Values{}
		if opt.Country != nil {
			v.Set("market", *opt.Country)
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	req, err := http.NewRequest("GET", spotifyURL, nil)
	if err != nil {
		return nil, err
	}

	var result CurrentlyPlaying
	err = c.execute(req, &result, http.StatusNoContent)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// PlayerRecentlyPlayed gets a list of recently-played tracks for the current
// user. This call requires ScopeUserReadRecentlyPlayed.
func (c *Client) PlayerRecentlyPlayed() ([]RecentlyPlayedItem, error) {
	return c.PlayerRecentlyPlayedOpt(nil)
}

// PlayerRecentlyPlayedOpt is like PlayerRecentlyPlayed, but it accepts
// additional options for sorting and filtering the results.
func (c *Client) PlayerRecentlyPlayedOpt(opt *RecentlyPlayedOptions) ([]RecentlyPlayedItem, error) {
	spotifyURL := c.baseURL + "me/player/recently-played"
	if opt != nil {
		v := url.Values{}
		if opt.Limit != 0 {
			v.Set("limit", strconv.FormatInt(int64(opt.Limit), 10))
		}
		if opt.BeforeEpochMs != 0 {
			v.Set("before", strconv.FormatInt(int64(opt.BeforeEpochMs), 10))
		}
		if opt.AfterEpochMs != 0 {
			v.Set("after", strconv.FormatInt(int64(opt.AfterEpochMs), 10))
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	result := RecentlyPlayedResult{}
	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

// TransferPlayback transfers playback to a new device and determine if
// it should start playing.
//
// Note that a value of false for the play parameter when also transferring
// to another device_id will not pause playback. To ensure that playback is
// paused on the new device you should send a pause command to the currently
// active device before transferring to the new device_id.
//
// Requires the ScopeUserModifyPlaybackState in order to modify the player state
func (c *Client) TransferPlayback(deviceID ID, play bool) error {
	reqData := struct {
		DeviceID []ID `json:"device_ids"`
		Play     bool `json:"play"`
	}{
		DeviceID: []ID{deviceID},
		Play:     play,
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(reqData)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, c.baseURL+"me/player", buf)
	if err != nil {
		return err
	}
	err = c.execute(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}

	return nil
}

// Play Start a new context or resume current playback on the user's active
// device. This call requires ScopeUserModifyPlaybackState in order to modify the player state.
func (c *Client) Play() error {
	return c.PlayOpt(nil)
}

// PlayOpt is like Play but with more options
func (c *Client) PlayOpt(opt *PlayOptions) error {
	spotifyURL := c.baseURL + "me/player/play"
	buf := new(bytes.Buffer)

	if opt != nil {
		v := url.Values{}
		if opt.DeviceID != nil {
			v.Set("device_id", opt.DeviceID.String())
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}

		err := json.NewEncoder(buf).Encode(opt)
		if err != nil {
			return err
		}
	}
	req, err := http.NewRequest(http.MethodPut, spotifyURL, buf)
	if err != nil {
		return err
	}
	err = c.execute(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// Pause Playback on the user's currently active device.
//
// Requires the ScopeUserModifyPlaybackState in order to modify the player state
func (c *Client) Pause() error {
	return c.PauseOpt(nil)
}

// PauseOpt is like Pause but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored
func (c *Client) PauseOpt(opt *PlayOptions) error {
	spotifyURL := c.baseURL + "me/player/pause"

	if opt != nil {
		v := url.Values{}
		if opt.DeviceID != nil {
			v.Set("device_id", opt.DeviceID.String())
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}
	req, err := http.NewRequest(http.MethodPut, spotifyURL, nil)
	if err != nil {
		return err
	}
	err = c.execute(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// Next skips to the next track in the user's queue in the user's
// currently active device. This call requires ScopeUserModifyPlaybackState
// in order to modify the player state
func (c *Client) Next() error {
	return c.NextOpt(nil)
}

// NextOpt is like Next but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored
func (c *Client) NextOpt(opt *PlayOptions) error {
	spotifyURL := c.baseURL + "me/player/next"

	if opt != nil {
		v := url.Values{}
		if opt.DeviceID != nil {
			v.Set("device_id", opt.DeviceID.String())
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}
	req, err := http.NewRequest(http.MethodPost, spotifyURL, nil)
	if err != nil {
		return err
	}
	err = c.execute(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// Previous skips to the Previous track in the user's queue in the user's
// currently active device. This call requires ScopeUserModifyPlaybackState
// in order to modify the player state
func (c *Client) Previous() error {
	return c.PreviousOpt(nil)
}

// PreviousOpt is like Previous but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored
func (c *Client) PreviousOpt(opt *PlayOptions) error {
	spotifyURL := c.baseURL + "me/player/previous"

	if opt != nil {
		v := url.Values{}
		if opt.DeviceID != nil {
			v.Set("device_id", opt.DeviceID.String())
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}
	req, err := http.NewRequest(http.MethodPost, spotifyURL, nil)
	if err != nil {
		return err
	}
	err = c.execute(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// Seek to the given position in the user’s currently playing track.
//
// The position in milliseconds to seek to. Must be a positive number.
// Passing in a position that is greater than the length of the track
// will cause the player to start playing the next song.
//
// Requires the ScopeUserModifyPlaybackState in order to modify the player state
func (c *Client) Seek(position int) error {
	return c.SeekOpt(position, nil)
}

// SeekOpt is like Seek but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored
func (c *Client) SeekOpt(position int, opt *PlayOptions) error {
	return c.playerFuncWithOpt(
		"me/player/seek",
		url.Values{
			"position_ms": []string{strconv.FormatInt(int64(position), 10)},
		},
		opt,
	)
}

// Repeat Set the repeat mode for the user's playback.
//
// Options are repeat-track, repeat-context, and off.
//
// Requires the ScopeUserModifyPlaybackState in order to modify the player state.
func (c *Client) Repeat(state string) error {
	return c.RepeatOpt(state, nil)
}

// RepeatOpt is like Repeat but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored.
func (c *Client) RepeatOpt(state string, opt *PlayOptions) error {
	return c.playerFuncWithOpt(
		"me/player/repeat",
		url.Values{
			"state": []string{state},
		},
		opt,
	)
}

// Volume set the volume for the user's current playback device.
//
// Percent is must be a value from 0 to 100 inclusive.
//
// Requires the ScopeUserModifyPlaybackState in order to modify the player state
func (c *Client) Volume(percent int) error {
	return c.VolumeOpt(percent, nil)
}

// VolumeOpt is like Volume but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored
func (c *Client) VolumeOpt(percent int, opt *PlayOptions) error {
	return c.playerFuncWithOpt(
		"me/player/volume",
		url.Values{
			"volume_percent": []string{strconv.FormatInt(int64(percent), 10)},
		},
		opt,
	)
}

// Shuffle switches shuffle on or off for user's playback.
//
// Requires the ScopeUserModifyPlaybackState in order to modify the player state
func (c *Client) Shuffle(shuffle bool) error {
	return c.ShuffleOpt(shuffle, nil)
}

// ShuffleOpt is like Shuffle but with more options
//
// Only expects PlayOptions.DeviceID, all other options will be ignored
func (c *Client) ShuffleOpt(shuffle bool, opt *PlayOptions) error {
	return c.playerFuncWithOpt(
		"me/player/shuffle",
		url.Values{
			"state": []string{strconv.FormatBool(shuffle)},
		},
		opt,
	)
}

func (c *Client) playerFuncWithOpt(urlSuffix string, values url.Values, opt *PlayOptions) error {
	spotifyURL := c.baseURL + urlSuffix

	if opt != nil {
		if opt.DeviceID != nil {
			values.Set("device_id", opt.DeviceID.String())
		}
	}

	if params := values.Encode(); params != "" {
		spotifyURL += "?" + params
	}

	req, err := http.NewRequest(http.MethodPut, spotifyURL, nil)
	if err != nil {
		return err
	}
	err = c.execute(req, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}
