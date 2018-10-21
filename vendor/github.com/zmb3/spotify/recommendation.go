package spotify

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Seeds contains IDs of artists, genres and/or tracks
// to be used as seeds for recommendations
type Seeds struct {
	Artists []ID
	Tracks  []ID
	Genres  []string
}

// count returns the total number of seeds contained in s
func (s Seeds) count() int {
	return len(s.Artists) + len(s.Tracks) + len(s.Genres)
}

// Recommendations contains a list of recommended tracks based on seeds
type Recommendations struct {
	Seeds  []RecommendationSeed `json:"seeds"`
	Tracks []SimpleTrack        `json:"tracks"`
}

// RecommendationSeed represents a recommendation seed after
// being processed by the Spotify API
type RecommendationSeed struct {
	AfterFilteringSize int    `json:"afterFilteringSize"`
	AfterRelinkingSize int    `json:"afterRelinkingSize"`
	Endpoint           string `json:"href"`
	ID                 ID     `json:"id"`
	InitialPoolSize    int    `json:"initialPoolSize"`
	Type               string `json:"type"`
}

// MaxNumberOfSeeds allowed by Spotify for a recommendation request
const MaxNumberOfSeeds = 5

// setSeedValues sets url values into v for each seed in seeds
func setSeedValues(seeds Seeds, v url.Values) {
	if len(seeds.Artists) != 0 {
		v.Set("seed_artists", strings.Join(toStringSlice(seeds.Artists), ","))
	}
	if len(seeds.Tracks) != 0 {
		v.Set("seed_tracks", strings.Join(toStringSlice(seeds.Tracks), ","))
	}
	if len(seeds.Genres) != 0 {
		v.Set("seed_genres", strings.Join(seeds.Genres, ","))
	}
}

// setTrackAttributesValues sets track attributes values to the given url values
func setTrackAttributesValues(trackAttributes *TrackAttributes, values url.Values) {
	if trackAttributes == nil {
		return
	}
	for attr, val := range trackAttributes.intAttributes {
		values.Set(attr, strconv.Itoa(val))
	}
	for attr, val := range trackAttributes.floatAttributes {
		values.Set(attr, strconv.FormatFloat(val, 'f', -1, 64))
	}
}

// GetRecommendations returns a list of recommended tracks based on the given seeds.
// Recommendations are generated based on the available information for a given seed entity
// and matched against similar artists and tracks. If there is sufficient information
// about the provided seeds, a list of tracks will be returned together with pool size details.
// For artists and tracks that are very new or obscure
// there might not be enough data to generate a list of tracks.
func (c *Client) GetRecommendations(seeds Seeds, trackAttributes *TrackAttributes, opt *Options) (*Recommendations, error) {
	v := url.Values{}

	if seeds.count() == 0 {
		return nil, fmt.Errorf("spotify: at least one seed is required")
	}
	if seeds.count() > MaxNumberOfSeeds {
		return nil, fmt.Errorf("spotify: exceeded maximum of %d seeds", MaxNumberOfSeeds)
	}

	setSeedValues(seeds, v)
	setTrackAttributesValues(trackAttributes, v)

	if opt != nil {
		if opt.Limit != nil {
			v.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Country != nil {
			v.Set("market", *opt.Country)
		}
	}

	spotifyURL := c.baseURL + "recommendations?" + v.Encode()

	var recommendations Recommendations
	err := c.get(spotifyURL, &recommendations)
	if err != nil {
		return nil, err
	}

	return &recommendations, err
}

// GetAvailableGenreSeeds retrieves a list of available genres seed parameter values for
// recommendations.
func (c *Client) GetAvailableGenreSeeds() ([]string, error) {
	spotifyURL := c.baseURL + "recommendations/available-genre-seeds"

	genreSeeds := make(map[string][]string)

	err := c.get(spotifyURL, &genreSeeds)
	if err != nil {
		return nil, err
	}

	return genreSeeds["genres"], nil
}
