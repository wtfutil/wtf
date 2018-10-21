// Package spotify provides utilties for interfacing
// with Spotify's Web API.
package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Version is the version of this library.
const Version = "1.0.0"

const (
	// DateLayout can be used with time.Parse to create time.Time values
	// from Spotify date strings.  For example, PrivateUser.Birthdate
	// uses this format.
	DateLayout = "2006-01-02"
	// TimestampLayout can be used with time.Parse to create time.Time
	// values from SpotifyTimestamp strings.  It is an ISO 8601 UTC timestamp
	// with a zero offset.  For example, PlaylistTrack's AddedAt field uses
	// this format.
	TimestampLayout = "2006-01-02T15:04:05Z"

	// defaultRetryDurationS helps us fix an apparent server bug whereby we will
	// be told to retry but not be given a wait-interval.
	defaultRetryDuration = time.Second * 5

	// rateLimitExceededStatusCode is the code that the server returns when our
	// request frequency is too high.
	rateLimitExceededStatusCode = 429
)

const baseAddress = "https://api.spotify.com/v1/"

// Client is a client for working with the Spotify Web API.
// To create an authenticated client, use the `Authenticator.NewClient` method.
type Client struct {
	http    *http.Client
	baseURL string

	AutoRetry bool
}

// URI identifies an artist, album, track, or category.  For example,
// spotify:track:6rqhFgbbKwnb9MLmUQDhG6
type URI string

// ID is a base-62 identifier for an artist, track, album, etc.
// It can be found at the end of a spotify.URI.
type ID string

func (id *ID) String() string {
	return string(*id)
}

// Followers contains information about the number of people following a
// particular artist or playlist.
type Followers struct {
	// The total number of followers.
	Count uint `json:"total"`
	// A link to the Web API endpoint providing full details of the followers,
	// or the empty string if this data is not available.
	Endpoint string `json:"href"`
}

// Image identifies an image associated with an item.
type Image struct {
	// The image height, in pixels.
	Height int `json:"height"`
	// The image width, in pixels.
	Width int `json:"width"`
	// The source URL of the image.
	URL string `json:"url"`
}

// Download downloads the image and writes its data to the specified io.Writer.
func (i Image) Download(dst io.Writer) error {
	resp, err := http.Get(i.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// TODO: get Content-Type from header?
	if resp.StatusCode != http.StatusOK {
		return errors.New("Couldn't download image - HTTP" + strconv.Itoa(resp.StatusCode))
	}
	_, err = io.Copy(dst, resp.Body)
	return err
}

// Error represents an error returned by the Spotify Web API.
type Error struct {
	// A short description of the error.
	Message string `json:"message"`
	// The HTTP status code.
	Status int `json:"status"`
}

func (e Error) Error() string {
	return e.Message
}

// decodeError decodes an Error from an io.Reader.
func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("spotify: HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e struct {
		E Error `json:"error"`
	}
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("spotify: couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	if e.E.Message == "" {
		// Some errors will result in there being a useful status-code but an
		// empty message, which will confuse the user (who only has access to
		// the message and not the code). An example of this is when we send
		// some of the arguments directly in the HTTP query and the URL ends-up
		// being too long.

		e.E.Message = fmt.Sprintf("spotify: unexpected HTTP %d: %s (empty error)",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return e.E
}

// shouldRetry determines whether the status code indicates that the
// previous operation should be retried at a later time
func shouldRetry(status int) bool {
	return status == http.StatusAccepted || status == http.StatusTooManyRequests
}

// isFailure determines whether the code indicates failure
func isFailure(code int, validCodes []int) bool {
	for _, item := range validCodes {
		if item == code {
			return false
		}
	}
	return true
}

// execute executes a non-GET request. `needsStatus` describes other HTTP status codes
// that can represent success. Note that in all current usages of this function,
// we need to still allow a 200 even if we'd also like to check for additional
// success codes.
func (c *Client) execute(req *http.Request, result interface{}, needsStatus ...int) error {
	for {
		resp, err := c.http.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if c.AutoRetry && shouldRetry(resp.StatusCode) {
			time.Sleep(retryDuration(resp))
			continue
		}
		if resp.StatusCode != http.StatusOK && isFailure(resp.StatusCode, needsStatus) {
			return c.decodeError(resp)
		}

		if result != nil {
			if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
				return err
			}
		}
		break
	}
	return nil
}

func retryDuration(resp *http.Response) time.Duration {
	raw := resp.Header.Get("Retry-After")
	if raw == "" {
		return defaultRetryDuration
	}
	seconds, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		return defaultRetryDuration
	}
	return time.Duration(seconds) * time.Second
}

func (c *Client) get(url string, result interface{}) error {
	for {
		resp, err := c.http.Get(url)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode == rateLimitExceededStatusCode && c.AutoRetry {
			time.Sleep(retryDuration(resp))
			continue
		}
		if resp.StatusCode != http.StatusOK {
			return c.decodeError(resp)
		}

		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}

		break
	}

	return nil
}

// Options contains optional parameters that can be provided
// to various API calls.  Only the non-nil fields are used
// in queries.
type Options struct {
	// Country is an ISO 3166-1 alpha-2 country code.  Provide
	// this parameter if you want the list of returned items to
	// be relevant to a particular country.  If omitted, the
	// results will be relevant to all countries.
	Country *string
	// Limit is the maximum number of items to return.
	Limit *int
	// Offset is the index of the first item to return.  Use it
	// with Limit to get the next set of items.
	Offset *int
	// Timerange is the period of time from which to return results
	// in certain API calls. The three options are the following string
	// literals: "short", "medium", and "long"
	Timerange *string
}

// NewReleasesOpt is like NewReleases, but it accepts optional parameters
// for filtering the results.
func (c *Client) NewReleasesOpt(opt *Options) (albums *SimpleAlbumPage, err error) {
	spotifyURL := c.baseURL + "browse/new-releases"
	if opt != nil {
		v := url.Values{}
		if opt.Country != nil {
			v.Set("country", *opt.Country)
		}
		if opt.Limit != nil {
			v.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Offset != nil {
			v.Set("offset", strconv.Itoa(*opt.Offset))
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	var objmap map[string]*json.RawMessage
	err = c.get(spotifyURL, &objmap)
	if err != nil {
		return nil, err
	}

	var result SimpleAlbumPage
	err = json.Unmarshal(*objmap["albums"], &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NewReleases gets a list of new album releases featured in Spotify.
// This call requires bearer authorization.
func (c *Client) NewReleases() (albums *SimpleAlbumPage, err error) {
	return c.NewReleasesOpt(nil)
}
