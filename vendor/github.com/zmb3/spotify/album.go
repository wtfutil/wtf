package spotify

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// SimpleAlbum contains basic data about an album.
type SimpleAlbum struct {
	// The name of the album.
	Name string `json:"name"`
	// A slice of SimpleArtists
	Artists []SimpleArtist `json:"artists"`
	// The field is present when getting an artist’s
	// albums. Possible values are “album”, “single”,
	// “compilation”, “appears_on”. Compare to album_type
	// this field represents relationship between the artist
	// and the album.
	AlbumGroup string `json:"album_group"`
	// The type of the album: one of "album",
	// "single", or "compilation".
	AlbumType string `json:"album_type"`
	// The SpotifyID for the album.
	ID ID `json:"id"`
	// The SpotifyURI for the album.
	URI URI `json:"uri"`
	// The markets in which the album is available,
	// identified using ISO 3166-1 alpha-2 country
	// codes.  Note that al album is considered
	// available in a market when at least 1 of its
	// tracks is available in that market.
	AvailableMarkets []string `json:"available_markets"`
	// A link to the Web API enpoint providing full
	// details of the album.
	Endpoint string `json:"href"`
	// The cover art for the album in various sizes,
	// widest first.
	Images []Image `json:"images"`
	// Known external URLs for this album.
	ExternalURLs map[string]string `json:"external_urls"`
	// The date the album was first released.  For example, "1981-12-15".
	// Depending on the ReleaseDatePrecision, it might be shown as
	// "1981" or "1981-12". You can use ReleaseDateTime to convert this
	// to a time.Time value.
	ReleaseDate string `json:"release_date"`
	// The precision with which ReleaseDate value is known: "year", "month", or "day"
	ReleaseDatePrecision string `json:"release_date_precision"`
}

// Copyright contains the copyright statement associated with an album.
type Copyright struct {
	// The copyright text for the album.
	Text string `json:"text"`
	// The type of copyright.
	Type string `json:"type"`
}

// FullAlbum provides extra album data in addition to the data provided by SimpleAlbum.
type FullAlbum struct {
	SimpleAlbum
	Artists    []SimpleArtist `json:"artists"`
	Copyrights []Copyright    `json:"copyrights"`
	Genres     []string       `json:"genres"`
	// The popularity of the album, represented as an integer between 0 and 100,
	// with 100 being the most popular.  Popularity of an album is calculated
	// from the popularify of the album's individual tracks.
	Popularity int `json:"popularity"`
	// The date the album was first released.  For example, "1981-12-15".
	// Depending on the ReleaseDatePrecision, it might be shown as
	// "1981" or "1981-12". You can use ReleaseDateTime to convert this
	// to a time.Time value.
	ReleaseDate string `json:"release_date"`
	// The precision with which ReleaseDate value is known: "year", "month", or "day"
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Tracks               SimpleTrackPage   `json:"tracks"`
	ExternalIDs          map[string]string `json:"external_ids"`
}

// SavedAlbum provides info about an album saved to an user's account.
type SavedAlbum struct {
	// The date and time the track was saved, represented as an ISO
	// 8601 UTC timestamp with a zero offset (YYYY-MM-DDTHH:MM:SSZ).
	// You can use the TimestampLayout constant to convert this to
	// a time.Time value.
	AddedAt   string `json:"added_at"`
	FullAlbum `json:"album"`
}

// ReleaseDateTime converts the album's ReleaseDate to a time.TimeValue.
// All of the fields in the result may not be valid.  For example, if
// f.ReleaseDatePrecision is "month", then only the month and year
// (but not the day) of the result are valid.
func (f *FullAlbum) ReleaseDateTime() time.Time {
	if f.ReleaseDatePrecision == "day" {
		result, _ := time.Parse(DateLayout, f.ReleaseDate)
		return result
	}
	if f.ReleaseDatePrecision == "month" {
		ym := strings.Split(f.ReleaseDate, "-")
		year, _ := strconv.Atoi(ym[0])
		month, _ := strconv.Atoi(ym[1])
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}
	year, _ := strconv.Atoi(f.ReleaseDate)
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
}

// GetAlbum gets Spotify catalog information for a single album, given its Spotify ID.
func (c *Client) GetAlbum(id ID) (*FullAlbum, error) {
	spotifyURL := fmt.Sprintf("%salbums/%s", c.baseURL, id)

	var a FullAlbum

	err := c.get(spotifyURL, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func toStringSlice(ids []ID) []string {
	result := make([]string, len(ids))
	for i, str := range ids {
		result[i] = str.String()
	}
	return result
}

// GetAlbums gets Spotify Catalog information for multiple albums, given their
// Spotify IDs.  It supports up to 20 IDs in a single call.  Albums are returned
// in the order requested.  If an album is not found, that position in the
// result slice will be nil.
func (c *Client) GetAlbums(ids ...ID) ([]*FullAlbum, error) {
	if len(ids) > 20 {
		return nil, errors.New("spotify: exceeded maximum number of albums")
	}
	spotifyURL := fmt.Sprintf("%salbums?ids=%s", c.baseURL, strings.Join(toStringSlice(ids), ","))

	var a struct {
		Albums []*FullAlbum `json:"albums"`
	}

	err := c.get(spotifyURL, &a)
	if err != nil {
		return nil, err
	}

	return a.Albums, nil
}

// AlbumType represents the type of an album. It can be used to filter
// results when searching for albums.
type AlbumType int

// AlbumType values that can be used to filter which types of albums are
// searched for.  These are flags that can be bitwise OR'd together
// to search for multiple types of albums simultaneously.
const (
	AlbumTypeAlbum       AlbumType = 1 << iota
	AlbumTypeSingle                = 1 << iota
	AlbummTypeAppearsOn            = 1 << iota
	AlbumTypeCompilation           = 1 << iota
)

func (at AlbumType) encode() string {
	types := []string{}
	if at&AlbumTypeAlbum != 0 {
		types = append(types, "album")
	}
	if at&AlbumTypeSingle != 0 {
		types = append(types, "single")
	}
	if at&AlbummTypeAppearsOn != 0 {
		types = append(types, "appears_on")
	}
	if at&AlbumTypeCompilation != 0 {
		types = append(types, "compilation")
	}
	return strings.Join(types, ",")
}

// GetAlbumTracks gets the tracks for a particular album.
// If you only care about the tracks, this call is more efficient
// than GetAlbum.
func (c *Client) GetAlbumTracks(id ID) (*SimpleTrackPage, error) {
	return c.GetAlbumTracksOpt(id, -1, -1)
}

// GetAlbumTracksOpt behaves like GetAlbumTracks, with the exception that it
// allows you to specify extra parameters that limit the number of results returned.
// The maximum number of results to return is specified by limit.
// The offset argument can be used to specify the index of the first track to return.
// It can be used along with limit to reqeust the next set of results.
func (c *Client) GetAlbumTracksOpt(id ID, limit, offset int) (*SimpleTrackPage, error) {
	spotifyURL := fmt.Sprintf("%salbums/%s/tracks", c.baseURL, id)
	v := url.Values{}
	if limit != -1 {
		v.Set("limit", strconv.Itoa(limit))
	}
	if offset != -1 {
		v.Set("offset", strconv.Itoa(offset))
	}
	optional := v.Encode()
	if optional != "" {
		spotifyURL = spotifyURL + "?" + optional
	}

	var result SimpleTrackPage
	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
