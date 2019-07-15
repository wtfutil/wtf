package gofeed

import (
	"encoding/json"
	"time"

	"github.com/mmcdole/gofeed/extensions"
)

// Feed is the universal Feed type that atom.Feed
// and rss.Feed gets translated to. It represents
// a web feed.
type Feed struct {
	Title           string                   `json:"title,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Link            string                   `json:"link,omitempty"`
	FeedLink        string                   `json:"feedLink,omitempty"`
	Updated         string                   `json:"updated,omitempty"`
	UpdatedParsed   *time.Time               `json:"updatedParsed,omitempty"`
	Published       string                   `json:"published,omitempty"`
	PublishedParsed *time.Time               `json:"publishedParsed,omitempty"`
	Author          *Person                  `json:"author,omitempty"`
	Language        string                   `json:"language,omitempty"`
	Image           *Image                   `json:"image,omitempty"`
	Copyright       string                   `json:"copyright,omitempty"`
	Generator       string                   `json:"generator,omitempty"`
	Categories      []string                 `json:"categories,omitempty"`
	DublinCoreExt   *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	ITunesExt       *ext.ITunesFeedExtension `json:"itunesExt,omitempty"`
	Extensions      ext.Extensions           `json:"extensions,omitempty"`
	Custom          map[string]string        `json:"custom,omitempty"`
	Items           []*Item                  `json:"items"`
	FeedType        string                   `json:"feedType"`
	FeedVersion     string                   `json:"feedVersion"`
}

func (f Feed) String() string {
	json, _ := json.MarshalIndent(f, "", "    ")
	return string(json)
}

// Item is the universal Item type that atom.Entry
// and rss.Item gets translated to.  It represents
// a single entry in a given feed.
type Item struct {
	Title           string                   `json:"title,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Content         string                   `json:"content,omitempty"`
	Link            string                   `json:"link,omitempty"`
	Updated         string                   `json:"updated,omitempty"`
	UpdatedParsed   *time.Time               `json:"updatedParsed,omitempty"`
	Published       string                   `json:"published,omitempty"`
	PublishedParsed *time.Time               `json:"publishedParsed,omitempty"`
	Author          *Person                  `json:"author,omitempty"`
	GUID            string                   `json:"guid,omitempty"`
	Image           *Image                   `json:"image,omitempty"`
	Categories      []string                 `json:"categories,omitempty"`
	Enclosures      []*Enclosure             `json:"enclosures,omitempty"`
	DublinCoreExt   *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	ITunesExt       *ext.ITunesItemExtension `json:"itunesExt,omitempty"`
	Extensions      ext.Extensions           `json:"extensions,omitempty"`
	Custom          map[string]string        `json:"custom,omitempty"`
}

// Person is an individual specified in a feed
// (e.g. an author)
type Person struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// Image is an image that is the artwork for a given
// feed or item.
type Image struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

// Enclosure is a file associated with a given Item.
type Enclosure struct {
	URL    string `json:"url,omitempty"`
	Length string `json:"length,omitempty"`
	Type   string `json:"type,omitempty"`
}
