package spotify

import (
	"fmt"
	"net/url"
	"strconv"
)

// Category is used by Spotify to tag items in.  For example, on the Spotify
// player's "Browse" tab.
type Category struct {
	// A link to the Web API endpoint returning full details of the category
	Endpoint string `json:"href"`
	// The category icon, in various sizes
	Icons []Image `json:"icons"`
	// The Spotify category ID.  This isn't a base-62 Spotify ID, its just
	// a short string that describes and identifies the category (ie "party").
	ID string `json:"id"`
	// The name of the category
	Name string `json:"name"`
}

// GetCategoryOpt is like GetCategory, but it accepts optional arguments.
// The country parameter is an ISO 3166-1 alpha-2 country code.  It can be
// used to ensure that the category exists for a particular country.  The
// locale argument is an ISO 639 language code and an ISO 3166-1 alpha-2
// country code, separated by an underscore.  It can be used to get the
// category strings in a particular language (for example: "es_MX" means
// get categories in Mexico, returned in Spanish).
//
// This call requires authorization.
func (c *Client) GetCategoryOpt(id, country, locale string) (Category, error) {
	cat := Category{}
	spotifyURL := fmt.Sprintf("%sbrowse/categories/%s", c.baseURL, id)
	values := url.Values{}
	if country != "" {
		values.Set("country", country)
	}
	if locale != "" {
		values.Set("locale", locale)
	}
	if query := values.Encode(); query != "" {
		spotifyURL += "?" + query
	}

	err := c.get(spotifyURL, &cat)
	if err != nil {
		return cat, err
	}

	return cat, err
}

// GetCategory gets a single category used to tag items in Spotify
// (on, for example, the Spotify player's Browse tab).
func (c *Client) GetCategory(id string) (Category, error) {
	return c.GetCategoryOpt(id, "", "")
}

// GetCategoryPlaylists gets a list of Spotify playlists tagged with a particular category.
func (c *Client) GetCategoryPlaylists(catID string) (*SimplePlaylistPage, error) {
	return c.GetCategoryPlaylistsOpt(catID, nil)
}

// GetCategoryPlaylistsOpt is like GetCategoryPlaylists, but it accepts optional
// arguments.
func (c *Client) GetCategoryPlaylistsOpt(catID string, opt *Options) (*SimplePlaylistPage, error) {
	spotifyURL := fmt.Sprintf("%sbrowse/categories/%s/playlists", c.baseURL, catID)
	if opt != nil {
		values := url.Values{}
		if opt.Country != nil {
			values.Set("country", *opt.Country)
		}
		if opt.Limit != nil {
			values.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Offset != nil {
			values.Set("offset", strconv.Itoa(*opt.Offset))
		}
		if query := values.Encode(); query != "" {
			spotifyURL += "?" + query
		}
	}

	wrapper := struct {
		Playlists SimplePlaylistPage `json:"playlists"`
	}{}

	err := c.get(spotifyURL, &wrapper)
	if err != nil {
		return nil, err
	}

	return &wrapper.Playlists, nil
}

// GetCategories gets a list of categories used to tag items in Spotify
// (on, for example, the Spotify player's "Browse" tab).
func (c *Client) GetCategories() (*CategoryPage, error) {
	return c.GetCategoriesOpt(nil, "")
}

// GetCategoriesOpt is like GetCategories, but it accepts optional parameters.
//
// The locale option can be used to get the results in a particular language.
// It consists of an ISO 639 language code and an ISO 3166-1 alpha-2 country
// code, separated by an underscore.  Specify the empty string to have results
// returned in the Spotify default language (American English).
func (c *Client) GetCategoriesOpt(opt *Options, locale string) (*CategoryPage, error) {
	spotifyURL := c.baseURL + "browse/categories"
	values := url.Values{}
	if locale != "" {
		values.Set("locale", locale)
	}
	if opt != nil {
		if opt.Country != nil {
			values.Set("country", *opt.Country)
		}
		if opt.Limit != nil {
			values.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Offset != nil {
			values.Set("offset", strconv.Itoa(*opt.Offset))
		}
	}
	if query := values.Encode(); query != "" {
		spotifyURL += "?" + query
	}

	wrapper := struct {
		Categories CategoryPage `json:"categories"`
	}{}

	err := c.get(spotifyURL, &wrapper)
	if err != nil {
		return nil, err
	}

	return &wrapper.Categories, nil
}
