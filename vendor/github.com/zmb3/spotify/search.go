package spotify

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	// MarketFromToken can be used in place of the Market parameter
	// if the Client has a valid access token.  In this case, the
	// results will be limited to content that is playable in the
	// country associated with the user's account.  The user must have
	// granted access to the user-read-private scope when the access
	// token was issued.
	MarketFromToken = "from_token"
)

// SearchType represents the type of a query used in the Search function.
type SearchType int

// Search type values that can be passed to the Search function.  These are flags
// that can be bitwise OR'd together to search for multiple types of content simultaneously.
const (
	SearchTypeAlbum    SearchType = 1 << iota
	SearchTypeArtist              = 1 << iota
	SearchTypePlaylist            = 1 << iota
	SearchTypeTrack               = 1 << iota
)

func (st SearchType) encode() string {
	types := []string{}
	if st&SearchTypeAlbum != 0 {
		types = append(types, "album")
	}
	if st&SearchTypeArtist != 0 {
		types = append(types, "artist")
	}
	if st&SearchTypePlaylist != 0 {
		types = append(types, "playlist")
	}
	if st&SearchTypeTrack != 0 {
		types = append(types, "track")
	}
	return strings.Join(types, ",")
}

// SearchResult contains the results of a call to Search.
// Fields that weren't searched for will be nil pointers.
type SearchResult struct {
	Artists   *FullArtistPage     `json:"artists"`
	Albums    *SimpleAlbumPage    `json:"albums"`
	Playlists *SimplePlaylistPage `json:"playlists"`
	Tracks    *FullTrackPage      `json:"tracks"`
}

// Search gets Spotify catalog information about artists, albums, tracks,
// or playlists that match a keyword string.  t is a mask containing one or more
// search types.  For example, `Search(query, SearchTypeArtist|SearchTypeAlbum)`
// will search for artists or albums matching the specified keywords.
//
// Matching
//
// Matching of search keywords is NOT case sensitive.  Keywords are matched in
// any order unless surrounded by double quotes. Searching for playlists will
// return results where the query keyword(s) match any part of the playlist's
// name or description. Only popular public playlists are returned.
//
// Operators
//
// The operator NOT can be used to exclude results.  For example,
// query = "roadhouse NOT blues" returns items that match "roadhouse" but exludes
// those that also contain the keyword "blues".  Similarly, the OR operator can
// be used to broaden the search.  query = "roadhouse OR blues" returns all results
// that include either of the terms.  Only one OR operator can be used in a query.
//
// Operators should be specified in uppercase.
//
// Wildcards
//
// The asterisk (*) character can, with some limitations, be used as a wildcard
// (maximum of 2 per query).  It will match a variable number of non-white-space
// characters.  It cannot be used in a quoted phrase, in a field filter, or as
// the first character of a keyword string.
//
// Field filters
//
// By default, results are returned when a match is found in any field of the
// target object type.  Searches can be made more specific by specifying an album,
// artist, or track field filter.  For example, "album:gold artist:abba type:album"
// will only return results with the text "gold" in the album name and the text
// "abba" in the artist's name.
//
// The field filter "year" can be used with album, artist, and track searches to
// limit the results to a particular year. For example "bob year:2014" or
// "bob year:1980-2020".
//
// The field filter "tag:new" can be used in album searches to retrieve only
// albums released in the last two weeks. The field filter "tag:hipster" can be
// used in album searches to retrieve only albums with the lowest 10% popularity.
//
// Other possible field filters, depending on object types being searched,
// include "genre", "upc", and "isrc".  For example "damian genre:reggae-pop".
func (c *Client) Search(query string, t SearchType) (*SearchResult, error) {
	return c.SearchOpt(query, t, nil)
}

// SearchOpt works just like Search, but it accepts additional
// parameters for filtering the output.  See the documentation for Search more
// more information.
//
// If the Country field is specified in the options, then the results will only
// contain artists, albums, and tracks playable in the specified country
// (playlist results are not affected by the Country option).  Additionally,
// the constant MarketFromToken can be used with authenticated clients.
// If the client has a valid access token, then the results will only include
// content playable in the user's country.
func (c *Client) SearchOpt(query string, t SearchType, opt *Options) (*SearchResult, error) {
	v := url.Values{}
	v.Set("q", query)
	v.Set("type", t.encode())
	if opt != nil {
		if opt.Limit != nil {
			v.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Country != nil {
			v.Set("market", *opt.Country)
		}
		if opt.Offset != nil {
			v.Set("offset", strconv.Itoa(*opt.Offset))
		}
	}

	spotifyURL := c.baseURL + "search?" + v.Encode()

	var result SearchResult

	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

// NextArtistResults loads the next page of artists into the specified search result.
func (c *Client) NextArtistResults(s *SearchResult) error {
	if s.Artists == nil || s.Artists.Next == "" {
		return ErrNoMorePages
	}
	return c.get(s.Artists.Next, s)
}

// PreviousArtistResults loads the previous page of artists into the specified search result.
func (c *Client) PreviousArtistResults(s *SearchResult) error {
	if s.Artists == nil || s.Artists.Previous == "" {
		return ErrNoMorePages
	}
	return c.get(s.Artists.Previous, s)
}

// NextAlbumResults loads the next page of albums into the specified search result.
func (c *Client) NextAlbumResults(s *SearchResult) error {
	if s.Albums == nil || s.Albums.Next == "" {
		return ErrNoMorePages
	}
	return c.get(s.Albums.Next, s)
}

// PreviousAlbumResults loads the previous page of albums into the specified search result.
func (c *Client) PreviousAlbumResults(s *SearchResult) error {
	if s.Albums == nil || s.Albums.Previous == "" {
		return ErrNoMorePages
	}
	return c.get(s.Albums.Previous, s)
}

// NextPlaylistResults loads the next page of playlists into the specified search result.
func (c *Client) NextPlaylistResults(s *SearchResult) error {
	if s.Playlists == nil || s.Playlists.Next == "" {
		return ErrNoMorePages
	}
	return c.get(s.Playlists.Next, s)
}

// PreviousPlaylistResults loads the previous page of playlists into the specified search result.
func (c *Client) PreviousPlaylistResults(s *SearchResult) error {
	if s.Playlists == nil || s.Playlists.Previous == "" {
		return ErrNoMorePages
	}
	return c.get(s.Playlists.Previous, s)
}

// PreviousTrackResults loads the previous page of tracks into the specified search result.
func (c *Client) PreviousTrackResults(s *SearchResult) error {
	if s.Tracks == nil || s.Tracks.Previous == "" {
		return ErrNoMorePages
	}
	return c.get(s.Tracks.Previous, s)
}

// NextTrackResults loads the next page of tracks into the specified search result.
func (c *Client) NextTrackResults(s *SearchResult) error {
	if s.Tracks == nil || s.Tracks.Next == "" {
		return ErrNoMorePages
	}
	return c.get(s.Tracks.Next, s)
}
