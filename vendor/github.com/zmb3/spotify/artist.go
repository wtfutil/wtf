package spotify

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// SimpleArtist contains basic info about an artist.
type SimpleArtist struct {
	Name string `json:"name"`
	ID   ID     `json:"id"`
	// The Spotify URI for the artist.
	URI URI `json:"uri"`
	// A link to the Web API enpoint providing full details of the artist.
	Endpoint     string            `json:"href"`
	ExternalURLs map[string]string `json:"external_urls"`
}

// FullArtist provides extra artist data in addition to what is provided by SimpleArtist.
type FullArtist struct {
	SimpleArtist
	// The popularity of the artist, expressed as an integer between 0 and 100.
	// The artist's popularity is calculated from the popularity of the artist's tracks.
	Popularity int `json:"popularity"`
	// A list of genres the artist is associated with.  For example, "Prog Rock"
	// or "Post-Grunge".  If not yet classified, the slice is empty.
	Genres    []string `json:"genres"`
	Followers Followers
	// Images of the artist in various sizes, widest first.
	Images []Image `json:"images"`
}

// GetArtist gets Spotify catalog information for a single artist, given its Spotify ID.
func (c *Client) GetArtist(id ID) (*FullArtist, error) {
	spotifyURL := fmt.Sprintf("%sartists/%s", c.baseURL, id)

	var a FullArtist
	err := c.get(spotifyURL, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// GetArtists gets spotify catalog information for several artists based on their
// Spotify IDs.  It supports up to 50 artists in a single call.  Artists are
// returned in the order requested.  If an artist is not found, that position
// in the result will be nil.  Duplicate IDs will result in duplicate artists
// in the result.
func (c *Client) GetArtists(ids ...ID) ([]*FullArtist, error) {
	spotifyURL := fmt.Sprintf("%sartists?ids=%s", c.baseURL, strings.Join(toStringSlice(ids), ","))

	var a struct {
		Artists []*FullArtist
	}

	err := c.get(spotifyURL, &a)
	if err != nil {
		return nil, err
	}

	return a.Artists, nil
}

// GetArtistsTopTracks gets Spotify catalog information about an artist's top
// tracks in a particular country.  It returns a maximum of 10 tracks.  The
// country is specified as an ISO 3166-1 alpha-2 country code.
func (c *Client) GetArtistsTopTracks(artistID ID, country string) ([]FullTrack, error) {
	spotifyURL := fmt.Sprintf("%sartists/%s/top-tracks?country=%s", c.baseURL, artistID, country)

	var t struct {
		Tracks []FullTrack `json:"tracks"`
	}

	err := c.get(spotifyURL, &t)
	if err != nil {
		return nil, err
	}

	return t.Tracks, nil
}

// GetRelatedArtists gets Spotify catalog information about artists similar to a
// given artist.  Similarity is based on analysis of the Spotify community's
// listening history.  This function returns up to 20 artists that are considered
// related to the specified artist.
func (c *Client) GetRelatedArtists(id ID) ([]FullArtist, error) {
	spotifyURL := fmt.Sprintf("%sartists/%s/related-artists", c.baseURL, id)

	var a struct {
		Artists []FullArtist `json:"artists"`
	}

	err := c.get(spotifyURL, &a)
	if err != nil {
		return nil, err
	}

	return a.Artists, nil
}

// GetArtistAlbums gets Spotify catalog information about an artist's albums.
// It is equivalent to GetArtistAlbumsOpt(artistID, nil).
func (c *Client) GetArtistAlbums(artistID ID) (*SimpleAlbumPage, error) {
	return c.GetArtistAlbumsOpt(artistID, nil, nil)
}

// GetArtistAlbumsOpt is just like GetArtistAlbums, but it accepts optional
// parameters used to filter and sort the result.
//
// The AlbumType argument can be used to find a particular type of album.  Search
// for multiple types by OR-ing the types together.
func (c *Client) GetArtistAlbumsOpt(artistID ID, options *Options, t *AlbumType) (*SimpleAlbumPage, error) {
	spotifyURL := fmt.Sprintf("%sartists/%s/albums", c.baseURL, artistID)
	// add optional query string if options were specified
	values := url.Values{}
	if t != nil {
		values.Set("album_type", t.encode())
	}
	if options != nil {
		if options.Country != nil {
			values.Set("market", *options.Country)
		} else {
			// if the market is not specified, Spotify will likely return a lot
			// of duplicates (one for each market in which the album is available)
			// - prevent this behavior by falling back to the US by default
			// TODO: would this ever be the desired behavior?
			values.Set("market", CountryUSA)
		}
		if options.Limit != nil {
			values.Set("limit", strconv.Itoa(*options.Limit))
		}
		if options.Offset != nil {
			values.Set("offset", strconv.Itoa(*options.Offset))
		}
	}
	if query := values.Encode(); query != "" {
		spotifyURL += "?" + query
	}

	var p SimpleAlbumPage

	err := c.get(spotifyURL, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
